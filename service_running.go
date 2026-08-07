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
	"time"

	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// Run 在新 goroutine 中启动服务循环，并返回服务终止时完成的 Future。
// 服务只能启动一次；上下文已经取消或服务已经启动时会 panic。
func (svc *ServiceBehavior) Run() async.Future {
	ctx := svc.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrService, ctx.Err())
	default:
	}

	if !svc.isRunning.CompareAndSwap(false, true) {
		exception.Panicf("%w: already running", ErrService)
	}

	if parentCtx, ok := ctx.ParentContext().(corectx.Context); ok {
		if !parentCtx.WaitGroup().Join(1) {
			ctx.Terminate()
			corectx.UnsafeContext(ctx).CloseWaitGroup()
			corectx.UnsafeContext(ctx).ReturnTerminated()
			return ctx.Terminated()
		}
	}

	go svc.running()

	return ctx.Terminated()
}

// Terminate 请求停止服务，并返回服务终止时完成的 Future。
func (svc *ServiceBehavior) Terminate() async.Future {
	return svc.ctx.Terminate()
}

// Terminated 返回服务终止时完成的 Future。
func (svc *ServiceBehavior) Terminated() async.Future {
	return svc.ctx.Terminated()
}

func (svc *ServiceBehavior) running() {
	ctx := svc.ctx

	svc.emitEventRunningEvent(service.RunningEvent_Starting)
	svc.emitEventRunningEvent(service.RunningEvent_Started)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticker.C:
			svc.emitEventRunningEvent(service.RunningEvent_Heartbeat)
		}
	}

	svc.emitEventRunningEvent(service.RunningEvent_Terminating)

	corectx.UnsafeContext(ctx).CloseWaitGroup()
	ctx.WaitGroup().Wait()

	svc.shutAddIn()

	svc.emitEventRunningEvent(service.RunningEvent_Terminated)

	if parentCtx, ok := ctx.ParentContext().(corectx.Context); ok {
		parentCtx.WaitGroup().Done()
	}

	corectx.UnsafeContext(ctx).ReturnTerminated()
}

func (svc *ServiceBehavior) emitEventRunningEvent(runningEvent service.RunningEvent, args ...any) {
	svc.onBeforeContextRunningEvent(svc.ctx, runningEvent, args...)
	service.UnsafeContext(svc.ctx).EmitEventRunningEvent(runningEvent, args...)
}

func (svc *ServiceBehavior) onBeforeContextRunningEvent(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
	switch runningEvent {
	case service.RunningEvent_Starting:
		svc.initEntityPT()
		svc.initComponentPT()
		svc.initAddIn()
	}
}

func (svc *ServiceBehavior) initEntityPT() {
	go func() {
		for entityPT := range svc.ctx.EntityLib().Watch(svc.ctx) {
			svc.emitEventRunningEvent(service.RunningEvent_EntityPTDeclared, entityPT)
		}
	}()
}

func (svc *ServiceBehavior) initComponentPT() {
	go func() {
		for compPT := range svc.ctx.EntityLib().ComponentLib().Watch(svc.ctx) {
			svc.emitEventRunningEvent(service.RunningEvent_ComponentPTDeclared, compPT)
		}
	}()
}

func (svc *ServiceBehavior) initAddIn() {
	addInManager := service.UnsafeContext(svc.ctx).AddInManager()
	for _, status := range service.UnsafeAddInManager(addInManager).Freeze() {
		svc.activateAddIn(status)
	}
}

func (svc *ServiceBehavior) shutAddIn() {
	addInManager := service.UnsafeContext(svc.ctx).AddInManager()

	statuses := service.UnsafeAddInManager(addInManager).ListStatuses()
	for i := len(statuses) - 1; i >= 0; i-- {
		svc.deactivateAddIn(statuses[i])
	}
}

func (svc *ServiceBehavior) activateAddIn(status service.AddInStatus) {
	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInInit); ok {
		generic.CastAction2(cb.Init).Call(svc.ctx.AutoRecover(), svc.ctx.ReportError(), svc.ctx, nil)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleServiceAddInInit); ok {
		generic.CastAction1(cb.Init).Call(svc.ctx.AutoRecover(), svc.ctx.ReportError(), svc.ctx)
	}

	service.UnsafeAddInStatus(status).Started()
}

func (svc *ServiceBehavior) deactivateAddIn(status service.AddInStatus) {
	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(cb.Shut).Call(svc.ctx.AutoRecover(), svc.ctx.ReportError(), svc.ctx, nil)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleServiceAddInShut); ok {
		generic.CastAction1(cb.Shut).Call(svc.ctx.AutoRecover(), svc.ctx.ReportError(), svc.ctx)
	}

	service.UnsafeAddInStatus(status).Stopped()
}
