package golaxy

import (
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
	"time"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (_service *ServiceBehavior) Run() <-chan struct{} {
	if !service.UnsafeContext(_service.ctx).MarkRunning(true) {
		panic("service already running")
	}

	shutChan := make(chan struct{}, 1)

	if parentCtx, ok := _service.ctx.GetParentContext().(internal.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go _service.running(shutChan)

	return shutChan
}

// Terminate 停止
func (_service *ServiceBehavior) Terminate() {
	_service.ctx.GetCancelFunc()()
}

func (_service *ServiceBehavior) running(shutChan chan struct{}) {
	ctx := _service.ctx

	_service.changeRunningState(service.RunningState_Starting)

	_service.initPlugin()

	_service.changeRunningState(service.RunningState_Started)

	defer func() {
		_service.changeRunningState(service.RunningState_Terminating)

		ctx.GetWaitGroup().Wait()
		_service.shutPlugin()

		_service.changeRunningState(service.RunningState_Terminated)

		if parentCtx, ok := ctx.GetParentContext().(internal.Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		service.UnsafeContext(ctx).MarkRunning(false)
		shutChan <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (_service *ServiceBehavior) changeRunningState(state service.RunningState) {
	if handler := service.UnsafeContext(_service.ctx).GetOptions().RunningHandler; handler != nil {
		internal.CallOuterVoid(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
			handler(_service.ctx, state)
		})
	}
}

func (_service *ServiceBehavior) initPlugin() {
	if pluginBundle := service.UnsafeContext(_service.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginInit, ok := pluginFace.Iface.(LifecycleServicePluginInit); ok {
				internal.CallOuterVoid(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
					pluginInit.InitSP(_service.ctx)
				})
			}
			return true
		})
	}
}

func (_service *ServiceBehavior) shutPlugin() {
	if pluginBundle := service.UnsafeContext(_service.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.ReverseRange(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginShut, ok := pluginFace.Iface.(LifecycleServicePluginShut); ok {
				internal.CallOuterVoid(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
					pluginShut.ShutSP(_service.ctx)
				})
			}
			return true
		})
	}
}
