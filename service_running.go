package core

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
	"time"
)

// Run 运行
func (serv *ServiceBehavior) Run() {
	ctx := serv.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrService, context.Canceled))
	case <-concurrent.UnsafeContext(ctx).GetTerminatedChan():
		panic(fmt.Errorf("%w: terminated", ErrRuntime))
	default:
	}

	if parentCtx, ok := serv.ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go serv.running()
}

// Terminate 停止
func (serv *ServiceBehavior) Terminate() <-chan struct{} {
	return serv.ctx.Terminate()
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

	if parentCtx, ok := ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Done()
	}

	close(concurrent.UnsafeContext(serv.ctx).GetTerminatedChan())
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
