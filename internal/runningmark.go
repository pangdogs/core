package internal

import (
	"sync/atomic"
)

// RunningMark 运行标记
type RunningMark interface {
	markRunning() bool
	markShutdown() bool
}

// RunningMarkBehavior 运行标记
type RunningMarkBehavior int32

func (mark *RunningMarkBehavior) markRunning() bool {
	return atomic.CompareAndSwapInt32((*int32)(mark), 0, 1)
}

func (mark *RunningMarkBehavior) markShutdown() bool {
	return atomic.CompareAndSwapInt32((*int32)(mark), 1, 0)
}
