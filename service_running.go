package galaxy

import (
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/service"
)

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (serv *_ServiceBehavior) Run() <-chan struct{} {
	if !service.UnsafeContext(serv.ctx).MarkRunning() {
		panic("service already running")
	}

	shutChan := make(chan struct{}, 1)

	if parentCtx, ok := serv.ctx.GetParentCtx().(internal.Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	go serv.running(shutChan)

	return shutChan
}

// Stop 停止
func (serv *_ServiceBehavior) Stop() {
	serv.ctx.GetCancelFunc()()
}

func (serv *_ServiceBehavior) running(shutChan chan struct{}) {
	defer func() {
		internal.CallOuterNoRet(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), func() {
			if service.UnsafeContext(serv.ctx).GetOptions().StoppingCallback != nil {
				service.UnsafeContext(serv.ctx).GetOptions().StoppingCallback(serv.ctx)
			}
		})

		serv.ctx.GetWaitGroup().Wait()

		internal.CallOuterNoRet(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), func() {
			if service.UnsafeContext(serv.ctx).GetOptions().StoppedCallback != nil {
				service.UnsafeContext(serv.ctx).GetOptions().StoppedCallback(serv.ctx)
			}
		})

		if parentCtx, ok := serv.ctx.GetParentCtx().(internal.Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		service.UnsafeContext(serv.ctx).MarkShutdown()
		shutChan <- struct{}{}
	}()

	internal.CallOuterNoRet(serv.ctx.GetAutoRecover(), serv.ctx.GetReportError(), func() {
		if service.UnsafeContext(serv.ctx).GetOptions().StartedCallback != nil {
			service.UnsafeContext(serv.ctx).GetOptions().StartedCallback(serv.ctx)
		}
	})

	select {
	case <-serv.ctx.Done():
		return
	}
}
