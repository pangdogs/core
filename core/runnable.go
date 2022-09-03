package core

import (
	"sync/atomic"
)

// _Runnable 可运行接口
type _Runnable interface {
	// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
	Run() <-chan struct{}

	// Stop 停止
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
