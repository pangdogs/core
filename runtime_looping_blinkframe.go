package core

import "git.golaxy.org/core/runtime"

func (rt *RuntimeBehavior) loopingBlinkFrame() {
	frame := runtime.UnsafeFrame(rt.opts.Frame)

	totalFrames := frame.GetTotalFrames()
	gcFrames := uint64(rt.opts.GCInterval.Seconds() * float64(frame.GetTargetFPS()))

loop:
	for {
		curFrames := frame.GetCurFrames()

		if totalFrames > 0 && curFrames >= totalFrames {
			break loop
		}

		select {
		case <-rt.ctx.Done():
			break loop
		default:
		}

		rt.blinkFrameLoop()

		if curFrames%gcFrames == 0 {
			rt.gc()
		}

		frame.SetCurFrames(curFrames + 1)
	}

	close(rt.processQueue)

	for {
		select {
		case task, ok := <-rt.processQueue:
			if !ok {
				return
			}
			rt.runTask(task)

		default:
			return
		}
	}
}

func (rt *RuntimeBehavior) blinkFrameLoop() {
	rt.changeRunningState(runtime.RunningState_FrameLoopBegin)
	rt.changeRunningState(runtime.RunningState_FrameUpdateBegin)

	emitEventUpdate(&rt.eventUpdate)
	emitEventLateUpdate(&rt.eventLateUpdate)

	rt.changeRunningState(runtime.RunningState_FrameUpdateEnd)
	rt.changeRunningState(runtime.RunningState_FrameLoopEnd)
}
