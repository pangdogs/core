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
	case <-ctx.TerminatedChan():
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

// TerminatedChan 已停止chan
func (serv *ServiceBehavior) TerminatedChan() <-chan struct{} {
	return serv.ctx.TerminatedChan()
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

	pluginBundle.Range(func(pluginInfo plugin.PluginInfo) bool {
		serv.activatePlugin(pluginInfo)
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

	pluginBundle.ReversedRange(func(pluginInfo plugin.PluginInfo) bool {
		serv.deactivatePlugin(pluginInfo)
		return true
	})
}

func (serv *ServiceBehavior) activatePlugin(pluginInfo plugin.PluginInfo) {
	if pluginInit, ok := pluginInfo.Face.Iface.(LifecycleServicePluginInit); ok {
		generic.MakeAction1(pluginInit.InitSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
	plugin.UnsafePluginBundle(serv.ctx.GetPluginBundle()).SetActive(pluginInfo.Name, true)
}

func (serv *ServiceBehavior) deactivatePlugin(pluginInfo plugin.PluginInfo) {
	plugin.UnsafePluginBundle(serv.ctx.GetPluginBundle()).SetActive(pluginInfo.Name, false)
	if pluginShut, ok := pluginInfo.Face.Iface.(LifecycleServicePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
}
