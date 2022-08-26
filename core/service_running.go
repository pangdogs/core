package core

// Run ...
func (serv *_ServiceBehavior) Run() <-chan struct{} {
	if !serv.ctx.markRunning() {
		panic("serv already running")
	}

	shutChan := make(chan struct{}, 1)

	go serv.running(shutChan)

	return shutChan
}

// Stop ...
func (serv *_ServiceBehavior) Stop() {
	serv.ctx.GetCancelFunc()()
}

func (serv *_ServiceBehavior) running(shutChan chan struct{}) {
	if parentCtx, ok := serv.ctx.GetParentCtx().(Context); ok {
		parentCtx.GetWaitGroup().Add(1)
	}

	defer func() {
		CallOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
			if serv.ctx.getOptions().StoppingCallback != nil {
				serv.ctx.getOptions().StoppingCallback(serv)
			}
		})

		if parentCtx, ok := serv.ctx.GetParentCtx().(Context); ok {
			parentCtx.GetWaitGroup().Done()
		}

		serv.ctx.GetWaitGroup().Wait()

		CallOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
			if serv.ctx.getOptions().StoppedCallback != nil {
				serv.ctx.getOptions().StoppedCallback(serv)
			}
		})

		serv.ctx.markShutdown()
		shutChan <- struct{}{}
	}()

	CallOuterNoRet(serv.opts.EnableAutoRecover, serv.ctx.GetReportError(), func() {
		if serv.ctx.getOptions().StartedCallback != nil {
			serv.ctx.getOptions().StartedCallback(serv)
		}
	})

	select {
	case <-serv.ctx.Done():
		return
	}
}
