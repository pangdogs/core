package core

import "time"

func (runtime *_RuntimeBehavior) pushCall(segment func()) {
	if segment == nil {
		panic("nil segment")
	}

	timeoutTimer := time.NewTimer(runtime.opts.ProcessQueueTimeout)
	defer timeoutTimer.Stop()

	select {
	case runtime.processQueue <- segment:
		return
	case <-timeoutTimer.C:
		panic("process queue push segment timeout")
	}
}
