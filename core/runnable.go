package core

import (
	"sync/atomic"
)

type Runnable interface {
	Run() <-chan struct{}
	Stop()
}

type _RunnableMark interface {
	markRunning() bool
	markShutdown() bool
}

type _RunnableMarkBehavior struct {
	runningFlag int32
}

func (r *_RunnableMarkBehavior) markRunning() bool {
	return atomic.CompareAndSwapInt32(&r.runningFlag, 0, 1)
}

func (r *_RunnableMarkBehavior) markShutdown() bool {
	return atomic.CompareAndSwapInt32(&r.runningFlag, 1, 0)
}
