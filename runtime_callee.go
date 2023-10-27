package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
	"time"
)

// PushCall 将代码片段压入接受者的任务处理流水线，串行化的进行调用。
func (_runtime *RuntimeBehavior) PushCall(segment func()) {
	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrRuntime, errors.ErrArgs))
	}

	if _runtime.opts.ProcessQueueTimeout > 0 {
		timeoutTimer := time.NewTimer(_runtime.opts.ProcessQueueTimeout)
		defer timeoutTimer.Stop()

		select {
		case _runtime.processQueue <- segment:
			return
		case <-timeoutTimer.C:
			panic(fmt.Errorf("%w: process queue is full", ErrRuntime))
		}
	} else {
		select {
		case _runtime.processQueue <- segment:
			return
		default:
			panic(fmt.Errorf("%w: process queue is full", ErrRuntime))
		}
	}
}
