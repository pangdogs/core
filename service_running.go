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

// Stop 停止
func (_service *ServiceBehavior) Stop() {
	_service.ctx.GetCancelFunc()()
}

func (_service *ServiceBehavior) running(shutChan chan struct{}) {
	if pluginBundle := service.UnsafeContext(_service.ctx).GetOptions().PluginBundle; pluginBundle != nil {
		pluginBundle.Range(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginInit, ok := pluginFace.Iface.(LifecycleServicePluginInit); ok {
				internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
					pluginInit.InitSP(_service.ctx)
				})
			}
			return true
		})
	}

	defer func() {
		if callback := service.UnsafeContext(_service.ctx).GetOptions().StoppingCallback; callback != nil {
			internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
				callback(_service.ctx)
			})
		}

		_service.ctx.GetWaitGroup().Wait()

		if callback := service.UnsafeContext(_service.ctx).GetOptions().StoppedCallback; callback != nil {
			internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
				callback(_service.ctx)
			})
		}

		if parentCtx, ok := _service.ctx.GetParentContext().(internal.Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		if pluginBundle := service.UnsafeContext(_service.ctx).GetOptions().PluginBundle; pluginBundle != nil {
			pluginBundle.ReverseRange(func(pluginName string, pluginFace util.FaceAny) bool {
				if pluginShut, ok := pluginFace.Iface.(LifecycleServicePluginShut); ok {
					internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
						pluginShut.ShutSP(_service.ctx)
					})
				}
				return true
			})
		}

		service.UnsafeContext(_service.ctx).MarkRunning(false)
		shutChan <- struct{}{}
	}()

	if callback := service.UnsafeContext(_service.ctx).GetOptions().StartedCallback; callback != nil {
		internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
			callback(_service.ctx)
		})
	}

	for {
		select {
		case <-_service.ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
