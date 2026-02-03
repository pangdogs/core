/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core

import (
	"sync"
	"time"

	"git.golaxy.org/core/runtime"
)

func (rt *RuntimeBehavior) loopingRealTime() {
	gcTicker := time.NewTicker(rt.options.GCInterval)
	defer gcTicker.Stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go rt.scheduleFrameTasks(&wg, rt.frame.GetCurFrames()+1, rt.frame.GetTotalFrames(), rt.frame.GetTargetFPS())

	taskOut := rt.taskQueue.out()

loop:
	for rt.frameLoopBegin(); ; {
		select {
		case task := <-taskOut:
			rt.runTask(task)

		case <-gcTicker.C:
			rt.runGC()

		case <-rt.ctx.Done():
			break loop
		}
	}

	wg.Wait()
	rt.taskQueue.close()

loopEnding:
	for {
		select {
		case task, ok := <-taskOut:
			if !ok {
				break loopEnding
			}
			rt.runTask(task)

		default:
			break loopEnding
		}
	}

	rt.runGC()
	rt.frameLoopEnd()
}

func (rt *RuntimeBehavior) scheduleFrameTasks(wg *sync.WaitGroup, curFrames, totalFrames int64, targetFPS float64) {
	defer wg.Done()

	updateTicker := time.NewTicker(time.Duration(float64(time.Second) / targetFPS))
	defer updateTicker.Stop()

	done := make(chan struct{}, 1)

	for {
		if totalFrames > 0 && curFrames >= totalFrames {
			rt.Terminate()
			return
		}

		select {
		case <-updateTicker.C:
			if rt.taskQueue.pushFrame(rt.ctx, rt.frameLoop, done) {
				curFrames++
			}
		case <-rt.ctx.Done():
			return
		}
	}
}

func (rt *RuntimeBehavior) frameLoop(...any) {
	rt.frameLoopEnd()
	rt.frameLoopBegin()
}

func (rt *RuntimeBehavior) frameLoopBegin() {
	rt.emitEventRunningEvent(runtime.RunningEvent_FrameLoopBegin)
	rt.emitEventRunningEvent(runtime.RunningEvent_FrameUpdateBegin)

	_EmitEventUpdate(&rt.runtimeEventTab)
	_EmitEventLateUpdate(&rt.runtimeEventTab)

	rt.emitEventRunningEvent(runtime.RunningEvent_FrameUpdateEnd)
}

func (rt *RuntimeBehavior) frameLoopEnd() {
	rt.emitEventRunningEvent(runtime.RunningEvent_FrameLoopEnd)
	rt.frame.setCurFrames(rt.frame.GetCurFrames() + 1)
}
