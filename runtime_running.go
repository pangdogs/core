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

// Run 在新 goroutine 中启动运行时循环，并返回运行时终止时完成的 Future。
// 运行时只能启动一次；上下文已经取消或运行时已经启动时会 panic。
func (rt *RuntimeBehavior) Run() async.Future {
	ctx := rt.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrRuntime, ctx.Err())
	default:
	}

	if !rt.isRunning.CompareAndSwap(false, true) {
		exception.Panicf("%w: already running", ErrRuntime)
	}

	if parentCtx, ok := ctx.ParentContext().(corectx.Context); ok {
		if !parentCtx.WaitGroup().Join(1) {
			ctx.Terminate()
			corectx.UnsafeContext(ctx).CloseWaitGroup()
			corectx.UnsafeContext(ctx).ReturnTerminated()
			return ctx.Terminated()
		}
	}

	go rt.running()

	return ctx.Terminated()
}

// Terminate 请求停止运行时，并返回运行时终止时完成的 Future。
func (rt *RuntimeBehavior) Terminate() async.Future {
	return rt.ctx.Terminate()
}

// Terminated 返回运行时终止时完成的 Future。
func (rt *RuntimeBehavior) Terminated() async.Future {
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

	corectx.UnsafeContext(ctx).CloseWaitGroup()
	ctx.WaitGroup().Wait()

	rt.shutAddIn()

	rt.emitEventRunningEvent(runtime.RunningEvent_Terminated)

	if parentCtx, ok := ctx.ParentContext().(corectx.Context); ok {
		parentCtx.WaitGroup().Done()
	}

	corectx.UnsafeContext(ctx).ReturnTerminated()
}

func (rt *RuntimeBehavior) emitEventRunningEvent(runningEvent runtime.RunningEvent, args ...any) {
	runtime.UnsafeContext(rt.ctx).EmitEventRunningEvent(runningEvent, args...)
}

func (rt *RuntimeBehavior) onBeforeContextRunningEvent(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
	switch runningEvent {
	case runtime.RunningEvent_Starting:
		rt.initAddIn()
	case runtime.RunningEvent_FrameLoopBegin:
		rt.frame.loopBegin()
	case runtime.RunningEvent_FrameUpdateBegin:
		rt.frame.updateBegin()
	case runtime.RunningEvent_FrameUpdateEnd:
		rt.frame.updateEnd()
	case runtime.RunningEvent_FrameLoopEnd:
		rt.frame.loopEnd()
	}
}

func (rt *RuntimeBehavior) onAfterContextRunningEvent(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
	switch runningEvent {
	case runtime.RunningEvent_Birth:
		if rt.options.AutoRun {
			rt.getInstance().Run()
		}
	}
}

func (rt *RuntimeBehavior) initAddIn() {
	addInManager := runtime.UnsafeContext(rt.ctx).AddInManager()

	rt.managedAddInManagerHandles[0] = runtime.BindEventInstallAddIn(addInManager, runtime.HandleEventInstallAddIn(rt.activateAddIn))
	rt.managedAddInManagerHandles[1] = runtime.BindEventUninstallAddIn(addInManager, runtime.HandleEventUninstallAddIn(rt.deactivateAddIn))

	statuses := runtime.UnsafeAddInManager(addInManager).ListStatuses()
	for i := range statuses {
		rt.activateAddIn(statuses[i])
	}
}

func (rt *RuntimeBehavior) shutAddIn() {
	addInManager := runtime.UnsafeContext(rt.ctx).AddInManager()

	rt.managedAddInManagerHandles[0].Unbind()

	statuses := runtime.UnsafeAddInManager(addInManager).ListStatuses()
	for i := len(statuses) - 1; i >= 0; i-- {
		addInManager.Uninstall(statuses[i].Name())
	}

	rt.managedAddInManagerHandles[1].Unbind()
}

func (rt *RuntimeBehavior) activateAddIn(status runtime.AddInStatus) {
	if status.State() != extension.AddInState_Loaded {
		return
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivating, status)

	if status.State() != extension.AddInState_Loaded {
		rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivationAborted, status)
		return
	}

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInInit); ok {
		generic.CastAction2(cb.Init).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError(), service.Current(rt), rt.ctx)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleRuntimeAddInInit); ok {
		generic.CastAction1(cb.Init).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError(), rt.ctx)
	}

	if status.State() != extension.AddInState_Loaded {
		rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivationAborted, status)
		return
	}

	runtime.UnsafeAddInStatus(status).Started()

	if status.State() != extension.AddInState_Running {
		rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivationAborted, status)
		return
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInActivated, status)

	if status.State() != extension.AddInState_Running {
		return
	}

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInOnRuntimeRunningEvent); ok {
		runtime.UnsafeAddInStatus(status).ManagedRuntimeRunningEventHandle(
			runtime.BindEventContextRunningEvent(rt.ctx, runtime.HandleEventContextRunningEvent(cb.OnContextRunningEvent)),
		)
	}
}

func (rt *RuntimeBehavior) deactivateAddIn(status runtime.AddInStatus) {
	if status.State() != extension.AddInState_Running {
		return
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInDeactivating, status)

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(cb.Shut).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError(), service.Current(rt), rt.ctx)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleRuntimeAddInShut); ok {
		generic.CastAction1(cb.Shut).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError(), rt.ctx)
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_AddInDeactivated, status)
}

func (rt *RuntimeBehavior) loopStart() []event.Handle {
	ctx := rt.ctx

	if rt.frame != nil {
		rt.frame.runningBegin()
	}

	return []event.Handle{
		runtime.BindEventEntityManagerAddEntity(ctx.EntityManager(), rt.handleEventEntityManagerAddEntity),
		runtime.BindEventEntityManagerRemoveEntity(ctx.EntityManager(), rt.handleEventEntityManagerRemoveEntity),
		runtime.BindEventEntityManagerEntityAddComponents(ctx.EntityManager(), rt.handleEventEntityManagerEntityAddComponents),
		runtime.BindEventEntityManagerEntityRemoveComponent(ctx.EntityManager(), rt.handleEventEntityManagerEntityRemoveComponent),
		runtime.BindEventEntityManagerEntityComponentEnableChanged(ctx.EntityManager(), rt.handleEventEntityManagerEntityComponentEnableChanged),
		runtime.BindEventEntityManagerEntityFirstTouchComponent(ctx.EntityManager(), rt.handleEventEntityManagerEntityFirstTouchComponent),
	}
}

func (rt *RuntimeBehavior) loopStop(handles []event.Handle) {
	event.UnbindHandles(handles)

	if rt.frame != nil {
		rt.frame.runningEnd()
	}
}

func (rt *RuntimeBehavior) mainLoop() {
	if rt.frame == nil {
		rt.loopingNoFrame()
	} else {
		rt.loopingRealTime()
	}
}

func (rt *RuntimeBehavior) runTask(task _Task) {
	switch task.typ {
	case TaskType_Call:
		rt.emitEventRunningEvent(runtime.RunningEvent_RunCallBegin)
		task.run(rt.ctx)
		rt.emitEventRunningEvent(runtime.RunningEvent_RunCallEnd)
	case TaskType_Frame:
		task.run(rt.ctx)
	}
	rt.taskQueue.complete(task.typ)
}

func (rt *RuntimeBehavior) runGC() {
	rt.emitEventRunningEvent(runtime.RunningEvent_RunGCBegin)
	rt.gc()
	rt.emitEventRunningEvent(runtime.RunningEvent_RunGCEnd)
}
