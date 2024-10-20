/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core

import (
	"context"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// Run 运行
func (rt *RuntimeBehavior) Run() <-chan struct{} {
	ctx := rt.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrRuntime, context.Canceled)
	case <-ctx.Terminated():
		exception.Panicf("%w: terminated", ErrRuntime)
	default:
	}

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go rt.running()

	return ictx.UnsafeContext(ctx).GetTerminatedChan()
}

// Terminate 停止
func (rt *RuntimeBehavior) Terminate() <-chan struct{} {
	return rt.ctx.Terminate()
}

// Terminated 已停止
func (rt *RuntimeBehavior) Terminated() <-chan struct{} {
	return rt.ctx.Terminated()
}

func (rt *RuntimeBehavior) running() {
	ctx := rt.ctx

	rt.changeRunningState(runtime.RunningState_Starting)

	hooks := rt.loopStart()

	rt.changeRunningState(runtime.RunningState_Started)

	rt.mainLoop()

	rt.changeRunningState(runtime.RunningState_Terminating)

	rt.loopStop(hooks)
	ctx.GetWaitGroup().Wait()

	rt.changeRunningState(runtime.RunningState_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	close(ictx.UnsafeContext(ctx).GetTerminatedChan())
}

func (rt *RuntimeBehavior) changeRunningState(state runtime.RunningState, args ...any) {
	switch state {
	case runtime.RunningState_Starting:
		rt.initPlugin()
	case runtime.RunningState_FrameLoopBegin:
		runtime.UnsafeFrame(rt.opts.Frame).LoopBegin()
	case runtime.RunningState_FrameUpdateBegin:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateBegin()
	case runtime.RunningState_FrameUpdateEnd:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateEnd()
	case runtime.RunningState_FrameLoopEnd:
		runtime.UnsafeFrame(rt.opts.Frame).LoopEnd()
	case runtime.RunningState_Terminated:
		rt.shutPlugin()
	}

	runtime.UnsafeContext(rt.ctx).ChangeRunningState(state, args...)
}

func (rt *RuntimeBehavior) initPlugin() {
	pluginBundle := rt.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(rt.activatePlugin)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(rt.deactivatePlugin)

	pluginBundle.Range(func(pluginStatus plugin.PluginStatus) bool {
		rt.activatePlugin(pluginStatus)
		return true
	})
}

func (rt *RuntimeBehavior) shutPlugin() {
	pluginBundle := rt.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(nil)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(nil)

	pluginBundle.ReversedRange(func(pluginStatus plugin.PluginStatus) bool {
		rt.deactivatePlugin(pluginStatus)
		return true
	})
}

func (rt *RuntimeBehavior) activatePlugin(pluginStatus plugin.PluginStatus) {
	if pluginStatus.State() != plugin.PluginState_Loaded {
		return
	}

	rt.changeRunningState(runtime.RunningState_PluginActivating, pluginStatus)
	defer rt.changeRunningState(runtime.RunningState_PluginActivated, pluginStatus)

	if pluginInit, ok := pluginStatus.InstanceFace().Iface.(LifecycleRuntimePluginInit); ok {
		generic.MakeAction1(pluginInit.InitRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}

	plugin.UnsafePluginStatus(pluginStatus).SetState(plugin.PluginState_Active, plugin.PluginState_Loaded)
}

func (rt *RuntimeBehavior) deactivatePlugin(pluginStatus plugin.PluginStatus) {
	rt.changeRunningState(runtime.RunningState_PluginDeactivating, pluginStatus)
	defer rt.changeRunningState(runtime.RunningState_PluginDeactivated, pluginStatus)

	if !plugin.UnsafePluginStatus(pluginStatus).SetState(plugin.PluginState_Inactive, plugin.PluginState_Active) {
		return
	}

	if pluginShut, ok := pluginStatus.InstanceFace().Iface.(LifecycleRuntimePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutRP).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}
}

func (rt *RuntimeBehavior) loopStart() (hooks [5]event.Hook) {
	ctx := rt.ctx
	frame := rt.opts.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	hooks[0] = runtime.BindEventEntityMgrAddEntity(ctx.GetEntityMgr(), rt)
	hooks[1] = runtime.BindEventEntityMgrRemoveEntity(ctx.GetEntityMgr(), rt)
	hooks[2] = runtime.BindEventEntityMgrEntityAddComponents(ctx.GetEntityMgr(), rt)
	hooks[3] = runtime.BindEventEntityMgrEntityRemoveComponent(ctx.GetEntityMgr(), rt)
	hooks[4] = runtime.BindEventEntityMgrEntityFirstAccessComponent(ctx.GetEntityMgr(), rt)

	return
}

func (rt *RuntimeBehavior) loopStop(hooks [5]event.Hook) {
	frame := rt.opts.Frame

	event.Clean(hooks[:])

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningEnd()
	}
}

func (rt *RuntimeBehavior) mainLoop() {
	frame := rt.opts.Frame

	if frame == nil {
		rt.loopingNoFrame()
	} else {
		rt.loopingRealTime()
	}
}

func (rt *RuntimeBehavior) runTask(task _Task) {
	switch task.typ {
	case _TaskType_Call:
		rt.changeRunningState(runtime.RunningState_RunCallBegin)
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		rt.changeRunningState(runtime.RunningState_RunCallEnd)
	case _TaskType_Frame:
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}
}

func (rt *RuntimeBehavior) runGC() {
	rt.changeRunningState(runtime.RunningState_RunGCBegin)
	rt.gc()
	rt.changeRunningState(runtime.RunningState_RunGCEnd)
}
