package golaxy

import "time"

// PushCall 将代码片段压入接收者的任务处理流水线，串行化的进行调用。
func (_runtime *RuntimeBehavior) PushCall(segment func()) {
	if segment == nil {
		panic("nil segment")
	}

	timeoutTimer := time.NewTimer(_runtime.opts.ProcessQueueTimeout)
	defer timeoutTimer.Stop()

	select {
	case _runtime.processQueue <- segment:
		return
	case <-timeoutTimer.C:
		panic("process queue push segment timeout")
	}
}
