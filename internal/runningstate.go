package internal

import (
	"sync/atomic"
)

// RunningState 运行状态
type RunningState interface {
	markRunning(v bool) bool
}

// RunningStateBehavior 运行状态
type RunningStateBehavior struct {
	running atomic.Bool
}

func (rs *RunningStateBehavior) markRunning(v bool) bool {
	return rs.running.CompareAndSwap(!v, v)
}
