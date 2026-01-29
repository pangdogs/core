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
	"sync"
	"time"

	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"github.com/elliotchance/pie/v2"
)

// Run 运行
func (svc *ServiceBehavior) Run() async.AsyncRet {
	ctx := svc.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrService, ctx.Err())
	case <-ctx.Terminated():
		exception.Panicf("%w: terminated", ErrService)
	default:
	}

	if !svc.isRunning.CompareAndSwap(false, true) {
		exception.Panicf("%w: already running", ErrService)
	}

	if parentCtx, ok := svc.ctx.GetParentContext().(corectx.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go svc.running()

	return ctx.Terminated()
}

// Terminate 停止
func (svc *ServiceBehavior) Terminate() async.AsyncRet {
	return svc.ctx.Terminate()
}

// Terminated 已停止
func (svc *ServiceBehavior) Terminated() async.AsyncRet {
	return svc.ctx.Terminated()
}

func (svc *ServiceBehavior) running() {
	ctx := svc.ctx

	svc.emitEventRunningEvent(service.RunningEvent_Starting)
	svc.emitEventRunningEvent(service.RunningEvent_Started)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	svc.emitEventRunningEvent(service.RunningEvent_Terminating)

	ctx.GetWaitGroup().Wait()

	svc.emitEventRunningEvent(service.RunningEvent_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(corectx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	corectx.UnsafeContext(ctx).ReturnTerminated()
}

func (svc *ServiceBehavior) emitEventRunningEvent(runningEvent service.RunningEvent, args ...any) {
	svc.onBeforeContextRunningEvent(svc.ctx, runningEvent, args...)
	service.UnsafeContext(svc.ctx).EmitEventRunningEvent(runningEvent, args...)
	svc.onAfterContextRunningEvent(svc.ctx, runningEvent, args...)
}

func (svc *ServiceBehavior) onBeforeContextRunningEvent(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
	switch runningEvent {
	case service.RunningEvent_Starting:
		svc.initEntityPT()
		svc.initComponentPT()
		svc.initAddIn()
	}
}

func (svc *ServiceBehavior) onAfterContextRunningEvent(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
	switch runningEvent {
	case service.RunningEvent_Terminated:
		svc.shutAddIn()
	}
}

func (svc *ServiceBehavior) initEntityPT() {
	go func() {
		for entityPT := range svc.ctx.GetEntityLib().ListAndWatch(svc.ctx) {
			svc.emitEventRunningEvent(service.RunningEvent_EntityPTDeclared, entityPT)
		}
	}()
}

func (svc *ServiceBehavior) initComponentPT() {
	go func() {
		for compPT := range svc.ctx.GetEntityLib().GetComponentLib().ListAndWatch(svc.ctx) {
			svc.emitEventRunningEvent(service.RunningEvent_ComponentPTDeclared, compPT)
		}
	}()
}

func (svc *ServiceBehavior) initAddIn() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for event := range service.UnsafeContext(svc.ctx).GetAddInManager().ListAndWatch(svc.ctx.Terminated().Context(nil)) {
			switch e := event.(type) {
			case *extension.EventServiceAddInSnapshot:
				for _, status := range e.StatusList {
					svc.activateAddIn(status)
				}
				wg.Done()
			case *extension.EventServiceInstallAddIn:
				select {
				case <-svc.ctx.Done():
					continue
				default:
					svc.activateAddIn(e.Status)
				}
			case *extension.EventServiceUninstallAddIn:
				svc.deactivateAddIn(e.Status)
			}
		}
	}()

	wg.Wait()
}

func (svc *ServiceBehavior) shutAddIn() {
	for _, status := range pie.Reverse(service.UnsafeContext(svc.ctx).GetAddInManager().List()) {
		svcAddInStatus := status.(extension.ServiceAddInStatus)
		svcAddInStatus.Uninstall()
		<-svcAddInStatus.WaitState(extension.AddInState_Unloaded)
	}
}

func (svc *ServiceBehavior) activateAddIn(status extension.AddInStatus) {
	svcAddInStatus := status.(extension.ServiceAddInStatus)

	if !extension.UnsafeServiceAddInStatus(svcAddInStatus).DoInstallingOnce() {
		return
	}

	svc.emitEventRunningEvent(service.RunningEvent_AddInActivating, status)

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInInit); ok {
		generic.CastAction2(cb.Init).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleServiceAddInInit); ok {
		generic.CastAction1(cb.Init).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx)
	}

	extension.UnsafeServiceAddInStatus(svcAddInStatus).SetState(extension.AddInState_Loaded, extension.AddInState_Running)

	svc.emitEventRunningEvent(service.RunningEvent_AddInActivatingDone, status)
}

func (svc *ServiceBehavior) deactivateAddIn(status extension.AddInStatus) {
	svcAddInStatus := status.(extension.ServiceAddInStatus)

	if !extension.UnsafeServiceAddInStatus(svcAddInStatus).DoUninstallingOnce() {
		return
	}

	svc.emitEventRunningEvent(service.RunningEvent_AddInDeactivating, status)

	if cb, ok := status.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(cb.Shut).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
	} else if cb, ok := status.InstanceFace().Iface.(LifecycleServiceAddInShut); ok {
		generic.CastAction1(cb.Shut).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx)
	}

	extension.UnsafeServiceAddInStatus(svcAddInStatus).SetState(extension.AddInState_Running, extension.AddInState_Unloaded)

	svc.emitEventRunningEvent(service.RunningEvent_AddInDeactivatingDone, status)
}
