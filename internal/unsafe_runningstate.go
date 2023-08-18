package internal

// Deprecated: UnsafeRunningState 访问运行状态内部方法
func UnsafeRunningState(runningState RunningState) _UnsafeRunningState {
	return _UnsafeRunningState{
		RunningState: runningState,
	}
}

type _UnsafeRunningState struct {
	RunningState
}

// MarkRunning 标记已经开始运行
func (urs _UnsafeRunningState) MarkRunning(v bool) bool {
	return urs.markRunning(v)
}
