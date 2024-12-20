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

	svc.changeRunningState(service.RunningState_Starting)
	svc.changeRunningState(service.RunningState_Started)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	svc.changeRunningState(service.RunningState_Terminating)

	ctx.GetWaitGroup().Wait()

	svc.changeRunningState(service.RunningState_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	close(ictx.UnsafeContext(ctx).GetTerminatedChan())
}

func (svc *ServiceBehavior) changeRunningState(state service.RunningState, args ...any) {
	switch state {
	case service.RunningState_Starting:
		svc.initAddIn()
	case service.RunningState_Terminated:
		svc.shutAddIn()
	}

	service.UnsafeContext(svc.ctx).ChangeRunningState(state, args...)
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

	svc.changeRunningState(service.RunningState_AddInActivating, addInStatus)
	defer svc.changeRunningState(service.RunningState_AddInActivated, addInStatus)

	if addInInit, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInInit); ok {
		generic.MakeAction2(addInInit.Init).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
	}

	extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Active, extension.AddInState_Loaded)
}

func (svc *ServiceBehavior) deactivateAddIn(addInStatus extension.AddInStatus) {
	svc.changeRunningState(service.RunningState_AddInDeactivating, addInStatus)
	defer svc.changeRunningState(service.RunningState_AddInDeactivated, addInStatus)

	if !extension.UnsafeAddInStatus(addInStatus).SetState(extension.AddInState_Inactive, extension.AddInState_Active) {
		return
	}

	if addInShut, ok := addInStatus.InstanceFace().Iface.(LifecycleAddInShut); ok {
		generic.MakeAction2(addInShut.Shut).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx, nil)
	}
}
