package core

// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
func (serv *_ServiceBehavior) Run() <-chan struct{} {
	if !serv.ctx.markRunning() {
		panic("serv already running")
	}

	shutChan := make(chan struct{}, 1)

	if parentCtx, ok := serv.ctx.GetParentCtx().(_Context); ok {
		parentCtx.getWaitGroup().Add(1)
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
		callOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
			if serv.ctx.getOptions().StoppingCallback != nil {
				serv.ctx.getOptions().StoppingCallback(serv)
			}
		})

		serv.ctx.getWaitGroup().Wait()

		callOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
			if serv.ctx.getOptions().StoppedCallback != nil {
				serv.ctx.getOptions().StoppedCallback(serv)
			}
		})

		if parentCtx, ok := serv.ctx.GetParentCtx().(_Context); ok {
			parentCtx.getWaitGroup().Done()
		}

		serv.ctx.markShutdown()
		shutChan <- struct{}{}
	}()

	callOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
		if serv.ctx.getOptions().StartedCallback != nil {
			serv.ctx.getOptions().StartedCallback(serv)
		}
	})

	select {
	case <-serv.ctx.Done():
		return
	}
}
