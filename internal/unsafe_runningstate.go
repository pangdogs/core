package internal

func UnsafeRunningState(runningState RunningState) _UnsafeRunningState {
	return _UnsafeRunningState{
		RunningState: runningState,
	}
}

type _UnsafeRunningState struct {
	RunningState
}

func (urm _UnsafeRunningState) MarkRunning(v bool) bool {
	return urm.markRunning(v)
}
