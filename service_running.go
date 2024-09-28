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
	"fmt"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/generic"
	"time"
)

// Run 运行
func (svc *ServiceBehavior) Run() <-chan struct{} {
	ctx := svc.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrService, context.Canceled))
	case <-ctx.Terminated():
		panic(fmt.Errorf("%w: terminated", ErrRuntime))
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
		svc.initPlugin()
	case service.RunningState_Terminated:
		svc.shutPlugin()
	}

	service.UnsafeContext(svc.ctx).ChangeRunningState(state, args...)
}

func (svc *ServiceBehavior) initPlugin() {
	pluginBundle := svc.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(svc.activatePlugin)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(svc.deactivatePlugin)

	pluginBundle.Range(func(pluginStatus plugin.PluginStatus) bool {
		svc.activatePlugin(pluginStatus)
		return true
	})
}

func (svc *ServiceBehavior) shutPlugin() {
	pluginBundle := svc.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(nil)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(nil)

	pluginBundle.ReversedRange(func(pluginStatus plugin.PluginStatus) bool {
		svc.deactivatePlugin(pluginStatus)
		return true
	})
}

func (svc *ServiceBehavior) activatePlugin(pluginStatus plugin.PluginStatus) {
	if pluginStatus.State() != plugin.PluginState_Loaded {
		return
	}

	svc.changeRunningState(service.RunningState_PluginActivating, pluginStatus)
	defer svc.changeRunningState(service.RunningState_PluginActivated, pluginStatus)

	if pluginInit, ok := pluginStatus.InstanceFace().Iface.(LifecycleServicePluginInit); ok {
		generic.MakeAction1(pluginInit.InitSP).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx)
	}

	plugin.UnsafePluginStatus(pluginStatus).SetState(plugin.PluginState_Active, plugin.PluginState_Loaded)
}

func (svc *ServiceBehavior) deactivatePlugin(pluginStatus plugin.PluginStatus) {
	svc.changeRunningState(service.RunningState_PluginDeactivating, pluginStatus)
	defer svc.changeRunningState(service.RunningState_PluginDeactivated, pluginStatus)

	if !plugin.UnsafePluginStatus(pluginStatus).SetState(plugin.PluginState_Inactive, plugin.PluginState_Active) {
		return
	}

	if pluginShut, ok := pluginStatus.InstanceFace().Iface.(LifecycleServicePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutSP).Call(svc.ctx.GetAutoRecover(), svc.ctx.GetReportError(), svc.ctx)
	}
}
