package internal

func UnsafeRunningMark(runningMark RunningMark) _UnsafeRunningMark {
	return _UnsafeRunningMark{
		RunningMark: runningMark,
	}
}

type _UnsafeRunningMark struct {
	RunningMark
}

func (urm _UnsafeRunningMark) MarkRunning() bool {
	return urm.markRunning()
}

func (urm _UnsafeRunningMark) MarkShutdown() bool {
	return urm.markShutdown()
}
