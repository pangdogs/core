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
	if pluginBundle := service.UnsafeContext(serv.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(info plugin.PluginInfo) bool {
			if pluginInit, ok := info.Face.Iface.(LifecycleServicePluginInit); ok {
				generic.CastAction1(pluginInit.InitSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
			}
			plugin.UnsafePluginBundle(pluginBundle).Activate(info.Name, true)
			return true
		})
	}
}

func (serv *ServiceBehavior) shutPlugin() {
	if pluginBundle := service.UnsafeContext(serv.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.ReverseRange(func(info plugin.PluginInfo) bool {
			plugin.UnsafePluginBundle(pluginBundle).Activate(info.Name, false)
			if pluginShut, ok := info.Face.Iface.(LifecycleServicePluginShut); ok {
				generic.CastAction1(pluginShut.ShutSP).Call(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), serv.ctx)
			}
			return true
		})
	}
}
