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
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// Run 运行
func (rt *RuntimeBehavior) Run() async.AsyncRet {
	ctx := rt.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrRuntime, ctx.Err())
	case <-ctx.Terminated():
		exception.Panicf("%w: terminated", ErrRuntime)
	default:
	}

	if !rt.isRunning.CompareAndSwap(false, true) {
		exception.Panicf("%w: already running", ErrRuntime)
	}

	if parentCtx, ok := ctx.GetParentContext().(corectx.Context); ok {
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

	rt.emitEventRunningEvent(runtime.RunningEvent_Starting)

	handles := rt.loopStart()

	rt.emitEventRunningEvent(runtime.RunningEvent_Started)

	rt.mainLoop()

	rt.emitEventRunningEvent(runtime.RunningEvent_Terminating)

	rt.loopStop(handles)
	ctx.GetWaitGroup().Wait()

	rt.emitEventRunningEvent(runtime.RunningEvent_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(corectx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	corectx.UnsafeContext(ctx).ReturnTerminated()
}

func (rt *RuntimeBehavior) emitEventRunningEvent(runningEvent runtime.RunningEvent, args ...any) {
	runtime.UnsafeContext(rt.ctx).EmitEventRunningEvent(runningEvent, args...)
}

func (rt *RuntimeBehavior) onBeforeContextRunningEvent(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
	switch runningEvent {
	case runtime.RunningEvent_Birth:
		if rt.options.AutoRun {
			rt.getInstance().Run()
		}
	case runtime.RunningEvent_Starting:
		rt.initAddIn()
	case runtime.RunningEvent_FrameLoopBegin:
		runtime.UnsafeFrame(rt.options.Frame).LoopBegin()
	case runtime.RunningEvent_FrameUpdateBegin:
		runtime.UnsafeFrame(rt.options.Frame).UpdateBegin()
	case runtime.RunningEvent_FrameUpdateEnd:
		runtime.UnsafeFrame(rt.options.Frame).UpdateEnd()
	case runtime.RunningEvent_FrameLoopEnd:
		runtime.UnsafeFrame(rt.options.Frame).LoopEnd()
	}
}

func (rt *RuntimeBehavior) onAfterContextRunningEvent(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
	switch runningEvent {
	case runtime.RunningEvent_Terminated:
		rt.shutAddIn()
	}
}

func (rt *RuntimeBehavior) initAddIn() {
	addInManager := runtime.UnsafeContext(rt.ctx).GetAddInManager()

	rt.managedAddInManagerHandles[0] = extension.BindEventRuntimeInstallAddIn(addInManager, extension.HandleEventRuntimeInstallAddIn(rt.activateAddIn))
	rt.managedAddInManagerHandles[1] = extension.BindEventRuntimeUninstallAddIn(addInManager, extension.HandleEventRuntimeUninstallAddIn(rt.deactivateAddIn))

	addInStatusList := addInManager.List()
	for i := range addInStatusList {
		rt.activateAddIn(addInStatusList[i])
	}
}

func (rt *RuntimeBehavior) shutAddIn() {
	addInManager := runtime.UnsafeContext(rt.ctx).GetAddInManager()

	rt.managedAddInManagerHandles[0].Unbind()

	addInStatusList := addInManager.List()
	for i := len(addInStatusList) - 1; i >= 0; i-- {
		addInStatusList[i].Uninstall()
	}

	rt.managedAddInManagerHandles[1].Unbind()
}

func (rt *RuntimeBehavior) activateAddIn(status extension.AddInStatus) {
	if status.State() != extension.AddInState_Loaded {
		return
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivating, status)

	if status.State() != extension.AddInState_Loaded {
		rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivatingAborted, status)
		return
	}

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInInit); ok {
		generic.CastAction2(cb.Init).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), service.Current(rt), rt.ctx)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleRuntimeAddInInit); ok {
		generic.CastAction1(cb.Init).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}

	if status.State() != extension.AddInState_Loaded {
		rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivatingAborted, status)
		return
	}

	addInStatus := status.(extension.RuntimeAddInStatus)
	extension.UnsafeRuntimeAddInStatus(addInStatus).SetState(extension.AddInState_Running)

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivatingDone, status)

	if status.State() != extension.AddInState_Running {
		return
	}

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInOnRuntimeRunningEvent); ok {
		extension.UnsafeRuntimeAddInStatus(addInStatus).ManagedRuntimeRunningEventHandle(
			runtime.BindEventContextRunningEvent(rt.ctx, runtime.HandleEventContextRunningEvent(cb.OnContextRunningEvent)),
		)
	}
}

func (rt *RuntimeBehavior) deactivateAddIn(status extension.AddInStatus) {
	if status.State() != extension.AddInState_Running {
		return
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInDeactivating, status)

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), service.Current(rt), rt.ctx)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleRuntimeAddInShut); ok {
		generic.CastAction1(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), rt.ctx)
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInDeactivatingDone, status)
}

func (rt *RuntimeBehavior) loopStart() []event.Handle {
	ctx := rt.ctx
	frame := rt.options.Frame

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningBegin()
	}

	return []event.Handle{
		runtime.BindEventEntityManagerAddEntity(ctx.GetEntityManager(), rt.handleEventEntityManagerAddEntity),
		runtime.BindEventEntityManagerRemoveEntity(ctx.GetEntityManager(), rt.handleEventEntityManagerRemoveEntity),
		runtime.BindEventEntityManagerEntityAddComponents(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityAddComponents),
		runtime.BindEventEntityManagerEntityRemoveComponent(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityRemoveComponent),
		runtime.BindEventEntityManagerEntityComponentEnableChanged(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityComponentEnableChanged),
		runtime.BindEventEntityManagerEntityFirstTouchComponent(ctx.GetEntityManager(), rt.handleEventEntityManagerEntityFirstTouchComponent),
	}
}

func (rt *RuntimeBehavior) loopStop(handles []event.Handle) {
	frame := rt.options.Frame

	event.UnbindHandles(handles)

	if frame != nil {
		runtime.UnsafeFrame(frame).RunningEnd()
	}
}

func (rt *RuntimeBehavior) mainLoop() {
	frame := rt.options.Frame

	if frame == nil {
		rt.loopingNoFrame()
	} else {
		rt.loopingRealTime()
	}
}

func (rt *RuntimeBehavior) runTask(task _Task) {
	switch task.typ {
	case _TaskType_Call:
		rt.emitEventRunningEvent(runtime.RunningEvent_RunCallBegin)
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		rt.emitEventRunningEvent(runtime.RunningEvent_RunCallEnd)
	case _TaskType_Frame:
		task.run(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}
}

func (rt *RuntimeBehavior) runGC() {
	rt.emitEventRunningEvent(runtime.RunningEvent_RunGCBegin)
	rt.gc()
	rt.emitEventRunningEvent(runtime.RunningEvent_RunGCEnd)
}
