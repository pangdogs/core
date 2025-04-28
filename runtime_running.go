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
	"git.golaxy.org/core/ec/ictx"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// Run 运行
func (rt *RuntimeBehavior) Run() async.AsyncRet {
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

	return ctx.Terminated()
}

// Terminate 停止
func (rt *RuntimeBehavior) Terminate() async.AsyncRet {
	return rt.ctx.Terminate()
}

// Terminated 已停止
func (rt *RuntimeBehavior) Terminated() async.AsyncRet {
	return rt.ctx.Terminated()
}

func (rt *RuntimeBehavior) running() {
	ctx := rt.ctx

	rt.changeRunningStatus(runtime.RunningStatus_Starting)

	hooks := rt.loopStart()

	rt.changeRunningStatus(runtime.RunningStatus_Started)

	rt.mainLoop()

	rt.changeRunningStatus(runtime.RunningStatus_Terminating)

	rt.loopStop(hooks)
	ctx.GetWaitGroup().Wait()

	rt.changeRunningStatus(runtime.RunningStatus_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	ictx.UnsafeContext(ctx).ReturnTerminated()
}

func (rt *RuntimeBehavior) changeRunningStatus(status runtime.RunningStatus, args ...any) {
	switch status {
	case runtime.RunningStatus_FrameLoopBegin:
		runtime.UnsafeFrame(rt.opts.Frame).LoopBegin()
	case runtime.RunningStatus_FrameUpdateBegin:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateBegin()
	case runtime.RunningStatus_FrameUpdateEnd:
		runtime.UnsafeFrame(rt.opts.Frame).UpdateEnd()
	case runtime.RunningStatus_FrameLoopEnd:
		runtime.UnsafeFrame(rt.opts.Frame).LoopEnd()
	}

	runtime.UnsafeContext(rt.ctx).ChangeRunningStatus(status, args...)

	_EmitEventRuntimeRunningStatusChanged(&rt.eventRuntimeRunningStatusChanged, rt.ctx, status, args...)

	switch status {
	case runtime.RunningStatus_Starting:
		rt.initAddIn()
	case runtime.RunningStatus_Terminated:
		rt.shutAddIn()
	}
}

func (rt *RuntimeBehavior) initAddIn() {
	addInManager := rt.ctx.GetAddInManager()
	if addInManager == nil {
		return
	}

	extension.UnsafeAddInManager(addInManager).SetCallback(rt.activateAddIn, rt.deactivateAddIn)

	addInManager.Range(func(addInStatus extension.AddInStatus) bool {
		rt.activateAddIn(addInStatus)
		return true
	})
}

func (rt *RuntimeBehavior) shutAddIn() {
	addInManager := rt.ctx.GetAddInManager()
	if addInManager == nil {
		return
	}

	extension.UnsafeAddInManager(addInManager).SetCallback(nil, nil)

	addInManager.ReversedRange(func(addInStatus extension.AddInStatus) bool {
		rt.deactivateAddIn(addInStatus)
		return true
	})
}

func (rt *RuntimeBehavior) activateAddIn(addInStatus extension.AddInStatus) {
	if addInStatus.State() != extension.AddInState_Loaded {
		return
	}

	if !func() bool {
		rt.changeRunningStatus(runtime.RunningStatus_ActivatingAddIn, addInStatus)
		defer rt.changeRunningStatus(runtime.RunningStatus_AddInActivated, addInStatus)

		if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInInit); ok {
			generic.CastAction2(cb.Init).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), service.Current(rt), rt.ctx)
		} else if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleRuntimeAddInInit); ok {
			generic.CastAction1(cb.Init).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
		}

		return extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Active, extension.AddInState_Loaded)
	}() {
		return
	}

	if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInOnRuntimeRunningStatusChanged); ok {
		event.Bind[LifecycleAddInOnRuntimeRunningStatusChanged](&rt.eventRuntimeRunningStatusChanged, cb)
	}
}

func (rt *RuntimeBehavior) deactivateAddIn(addInStatus extension.AddInStatus) {
	if addInStatus.State() != extension.AddInState_Active {
		return
	}

	if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInOnRuntimeRunningStatusChanged); ok {
		event.Unbind(&rt.eventRuntimeRunningStatusChanged, cb)
	}

	rt.changeRunningStatus(runtime.RunningStatus_DeactivatingAddIn, addInStatus)
	defer rt.changeRunningStatus(runtime.RunningStatus_AddInDeactivated, addInStatus)

	if !extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Inactive, extension.AddInState_Active) {
		return
	}

	if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), service.Current(rt), rt.ctx)
	} else if cb, ok := addInStatus.InstanceFace().Iface.(LifecycleRuntimeAddInShut); ok {
		generic.CastAction1(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}
}

func (rt *RuntimeBehavior) loopStart() []event.Hook {
	ctx := rt.ctx
	frame := rt.opts.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	return []event.Hook{
		runtime.BindEventEntityManagerAddEntity(ctx.GetEntityManager(), rt.handleEventEntityManagerAddEntity),
		runtime.BindEventEntityManagerRemoveEntity(ctx.GetEntityManager(), rt.handleEventEntityManagerRemoveEntity),
		runtime.BindEventEntityManagerEntityAddComponents(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityAddComponents),
		runtime.BindEventEntityManagerEntityRemoveComponent(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityRemoveComponent),
		runtime.BindEventEntityManagerEntityFirstTouchComponent(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityFirstTouchComponent),
	}
}

func (rt *RuntimeBehavior) loopStop(hooks []event.Hook) {
	frame := rt.opts.Frame

	event.UnbindHooks(hooks)

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
		rt.changeRunningStatus(runtime.RunningStatus_RunCallBegin)
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		rt.changeRunningStatus(runtime.RunningStatus_RunCallEnd)
	case _TaskType_Frame:
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}
}

func (rt *RuntimeBehavior) runGC() {
	rt.changeRunningStatus(runtime.RunningStatus_RunGCBegin)
	rt.gc()
	rt.changeRunningStatus(runtime.RunningStatus_RunGCEnd)
}
