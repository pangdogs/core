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

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (serv *ServiceBehavior) Run() <-chan struct{} {
	ctx := serv.ctx

	select {
	case <-ctx.Done():
		panic(fmt.Errorf("%w: %w", ErrService, context.Canceled))
	default:
	}

	if !serv.started.CompareAndSwap(false, true) {
		panic(fmt.Errorf("%w: running already started", ErrService))
	}

	if parentCtx, ok := serv.ctx.GetParentContext().(concurrent.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	shutChan := make(chan struct{}, 1)
	go serv.running(shutChan)

	return shutChan
}

// Terminate 停止
func (serv *ServiceBehavior) Terminate() {
	serv.ctx.GetCancelFunc()()
}

func (serv *ServiceBehavior) running(shutChan chan struct{}) {
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

	shutChan <- struct{}{}
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

	pluginBundle.ReversedRange(func(pluginInfo plugin.PluginInfo) bool {
		serv.deactivatePlugin(pluginInfo)
		return true
	})
}

func (serv *ServiceBehavior) activatePlugin(pluginInfo plugin.PluginInfo) {
	if pluginInit, ok := pluginInfo.Face.Iface.(LifecycleServicePluginInit); ok {
		generic.MakeAction1(pluginInit.InitSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
	plugin.UnsafePluginBundle(service.UnsafeContext(serv.ctx).GetOptions().PluginBundle).SetActive(pluginInfo.Name, true)
}

func (serv *ServiceBehavior) deactivatePlugin(pluginInfo plugin.PluginInfo) {
	plugin.UnsafePluginBundle(service.UnsafeContext(serv.ctx).GetOptions().PluginBundle).SetActive(pluginInfo.Name, false)
	if pluginShut, ok := pluginInfo.Face.Iface.(LifecycleServicePluginShut); ok {
		generic.MakeAction1(pluginShut.ShutSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
	}
}
