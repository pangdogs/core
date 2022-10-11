package galaxy

import (
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (_service *_ServiceBehavior) Run() <-chan struct{} {
	if !service.UnsafeContext(_service.ctx).MarkRunning() {
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
func (_service *_ServiceBehavior) Stop() {
	_service.ctx.GetCancelFunc()()
}

func (_service *_ServiceBehavior) running(shutChan chan struct{}) {
	if pluginLib := service.UnsafeContext(_service.ctx).GetOptions().PluginLib; pluginLib != nil {
		pluginLib.Range(func(pluginName string, pluginFace util.FaceAny) bool {
			if pluginInit, ok := pluginFace.Iface.(_PluginInit); ok {
				internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
					pluginInit.Init()
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

		if pluginLib := service.UnsafeContext(_service.ctx).GetOptions().PluginLib; pluginLib != nil {
			pluginLib.Range(func(pluginName string, pluginFace util.FaceAny) bool {
				if pluginShut, ok := pluginFace.Iface.(_PluginShut); ok {
					internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
						pluginShut.Shut()
					})
				}
				return true
			})
		}

		service.UnsafeContext(_service.ctx).MarkShutdown()
		shutChan <- struct{}{}
	}()

	if callback := service.UnsafeContext(_service.ctx).GetOptions().StartedCallback; callback != nil {
		internal.CallOuterNoRet(_service.ctx.GetAutoRecover(), _service.ctx.GetReportError(), func() {
			callback(_service.ctx)
		})
	}

	select {
	case <-_service.ctx.Done():
		return
	}
}
