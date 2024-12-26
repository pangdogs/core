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
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"time"
)

// Run 运行
func (svc *ServiceBehavior) Run() <-chan struct{} {
	ctx := svc.ctx

	select {
	case <-ctx.Done():
		exception.Panicf("%w: %w", ErrService, context.Canceled)
	case <-ctx.Terminated():
		exception.Panicf("%w: terminated", ErrRuntime)
	default:
	}

	if parentCtx, ok := svc.ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go svc.running()

	return ictx.UnsafeContext(ctx).GetTerminatedChan()
}

// Terminate 停止
func (svc *ServiceBehavior) Terminate() <-chan struct{} {
	return svc.ctx.Terminate()
}

// Terminated 已停止
func (svc *ServiceBehavior) Terminated() <-chan struct{} {
	return svc.ctx.Terminated()
}

func (svc *ServiceBehavior) running() {
	ctx := svc.ctx

	svc.changeRunningStatus(service.RunningStatus_Starting)
	svc.changeRunningStatus(service.RunningStatus_Started)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	svc.changeRunningStatus(service.RunningStatus_Terminating)

	ctx.GetWaitGroup().Wait()

	svc.changeRunningStatus(service.RunningStatus_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	close(ictx.UnsafeContext(ctx).GetTerminatedChan())
}

func (svc *ServiceBehavior) changeRunningStatus(status service.RunningStatus, args ...any) {
	service.UnsafeContext(svc.ctx).ChangeRunningStatus(status, args...)

	svc.statusChangesCond.L.Lock()
	svc.statusChanges = &_StatusChanges{
		status: status,
		args:   args,
	}
	svc.statusChangesCond.Broadcast()
	svc.statusChangesCond.L.Unlock()

	switch status {
	case service.RunningStatus_Starting:
		svc.initEntityPT()
		svc.initAddIn()
	case service.RunningStatus_Terminated:
		svc.shutAddIn()
		svc.shutEntityPT()
	}
}

func (svc *ServiceBehavior) initEntityPT() {
	entityLib := svc.ctx.GetEntityLib()
	if entityLib == nil {
		return
	}

	pt.UnsafeEntityLib(entityLib).SetCallback(
		func(entityPT ec.EntityPT) {
			svc.changeRunningStatus(service.RunningStatus_EntityPTDeclared, entityPT)
		},
		func(entityPT ec.EntityPT) {
			svc.changeRunningStatus(service.RunningStatus_EntityPTRedeclared, entityPT)
		},
		func(entityPT ec.EntityPT) {
			svc.changeRunningStatus(service.RunningStatus_EntityPTUndeclared, entityPT)
		},
	)

	entityLib.Range(func(entityPT ec.EntityPT) bool {
		svc.changeRunningStatus(service.RunningStatus_EntityPTDeclared, entityPT)
		return true
	})
}

func (svc *ServiceBehavior) shutEntityPT() {
	entityLib := svc.ctx.GetEntityLib()
	if entityLib == nil {
		return
	}

	pt.UnsafeEntityLib(entityLib).SetCallback(nil, nil, nil)
}

func (svc *ServiceBehavior) initAddIn() {
	addInManager := svc.ctx.GetAddInManager()
	if addInManager == nil {
		return
	}

	extension.UnsafeAddInManager(addInManager).SetCallback(svc.activateAddIn, svc.deactivateAddIn)

	addInManager.Range(func(addInStatus extension.AddInStatus) bool {
		svc.activateAddIn(addInStatus)
		return true
	})
}

func (svc *ServiceBehavior) shutAddIn() {
	addInManager := svc.ctx.GetAddInManager()
	if addInManager == nil {
		return
	}

	extension.UnsafeAddInManager(addInManager).SetCallback(nil, nil)

	addInManager.ReversedRange(func(addInStatus extension.AddInStatus) bool {
		svc.deactivateAddIn(addInStatus)
		return true
	})
}

func (svc *ServiceBehavior) activateAddIn(addInStatus extension.AddInStatus) {
	if addInStatus.State() != extension.AddInState_Loaded {
		return
	}

	if !func() bool {
		svc.changeRunningStatus(service.RunningStatus_AddInActivating, addInStatus)
		defer svc.changeRunningStatus(service.RunningStatus_AddInActivated, addInStatus)

		if addInInit, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInInit); ok {
			generic.CastAction2(addInInit.Init).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
		}

		return extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Active, extension.AddInState_Loaded)
	}() {
		return
	}

	addInOnServiceRunningStatusChanged, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInOnServiceRunningStatusChanged)
	if ok {
		go func() {
			for {
				if !func() bool {
					svc.statusChangesCond.L.Lock()
					svc.statusChangesCond.Wait()
					statusChanges := svc.statusChanges
					svc.statusChangesCond.L.Unlock()

					if statusChanges.status == service.RunningStatus_AddInDeactivating && statusChanges.args[0].(extension.AddInStatus) == addInStatus {
						return false
					}
					if addInStatus.State() != extension.AddInState_Active {
						return false
					}

					addInOnServiceRunningStatusChanged.OnServiceRunningStatusChanged(svc.ctx, statusChanges.status, statusChanges.args...)
					return true
				}() {
					return
				}
			}
		}()
	}
}

func (svc *ServiceBehavior) deactivateAddIn(addInStatus extension.AddInStatus) {
	if addInStatus.State() != extension.AddInState_Active {
		return
	}

	svc.changeRunningStatus(service.RunningStatus_AddInDeactivating, addInStatus)
	defer svc.changeRunningStatus(service.RunningStatus_AddInDeactivated, addInStatus)

	if !extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Inactive, extension.AddInState_Active) {
		return
	}

	if addInShut, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.CastAction2(addInShut.Shut).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
	}
}
