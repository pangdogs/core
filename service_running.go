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
func (serv *ServiceBehavior) Run() <-chan struct{} {
	ctx := serv.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrService, context.Canceled))
	case <-ctx.Terminated():
		panic(fmt.Errorf("%w: terminated", ErrRuntime))
	default:
	}

	if parentCtx, ok := serv.ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go serv.running()

	return ictx.UnsafeContext(ctx).GetTerminatedChan()
}

// Terminate 停止
func (serv *ServiceBehavior) Terminate() <-chan struct{} {
	return serv.ctx.Terminate()
}

// Terminated 已停止
func (serv *ServiceBehavior) Terminated() <-chan struct{} {
	return serv.ctx.Terminated()
}

func (serv *ServiceBehavior) running() {
	ctx := serv.ctx

	serv.changeRunningState(service.RunningState_Starting)
	serv.changeRunningState(service.RunningState_Started)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			time.Sleep(1 * time.Second)
		}
	}

	serv.changeRunningState(service.RunningState_Terminating)

	ctx.GetWaitGroup().Wait()

	serv.changeRunningState(service.RunningState_Terminated)

	if parentCtx, ok := ctx.GetParentContext().(ictx.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	close(ictx.UnsafeContext(ctx).GetTerminatedChan())
}

func (serv *ServiceBehavior) changeRunningState(state service.RunningState) {
	switch state {
	case service.RunningState_Starting:
		serv.initPlugin()
	case service.RunningState_Terminated:
		serv.shutPlugin()
	}

	service.UnsafeContext(serv.ctx).ChangeRunningState(state)
}

func (serv *ServiceBehavior) initPlugin() {
	pluginBundle := serv.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(serv.activatePlugin)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(serv.deactivatePlugin)

	pluginBundle.Range(func(pluginStatus plugin.PluginStatus) bool {
		serv.activatePlugin(pluginStatus)
		return true
	})
}

func (serv *ServiceBehavior) shutPlugin() {
	pluginBundle := serv.ctx.GetPluginBundle()
	if pluginBundle == nil {
		return
	}

	plugin.UnsafePluginBundle(pluginBundle).SetInstallCB(nil)
	plugin.UnsafePluginBundle(pluginBundle).SetUninstallCB(nil)

	pluginBundle.ReversedRange(func(pluginStatus plugin.PluginStatus) bool {
		serv.deactivatePlugin(pluginStatus)
		return true
	})
}

func (serv *ServiceBehavior) activatePlugin(pluginStatus plugin.PluginStatus) {
	if pluginInit, ok := pluginStatus.InstanceFace.Iface.(LifecycleServicePluginInit); ok {
		generic.MakeAction1(pluginInit.InitSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
	plugin.UnsafePluginBundle(serv.ctx.GetPluginBundle()).SetPluginState(pluginStatus.Name, plugin.PluginState_Active)
}

func (serv *ServiceBehavior) deactivatePlugin(pluginStatus plugin.PluginStatus) {
	plugin.UnsafePluginBundle(serv.ctx.GetPluginBundle()).SetPluginState(pluginStatus.Name, plugin.PluginState_Inactive)
	if pluginShut, ok := pluginStatus.InstanceFace.Iface.(LifecycleServicePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
}
