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
	"git.golaxy.org/core/runtime"
	"sync"
	"time"
)

func (rt *RuntimeBehavior) loopingRealTime() {
	gcTicker := time.NewTicker(rt.opts.GCInterval)
	defer gcTicker.Stop()

	wg := &sync.WaitGroup{}
	frame := rt.opts.Frame

	wg.Add(1)
	go rt.makeFrameTasks(wg, frame.GetCurFrames()+1, frame.GetTotalFrames(), frame.GetTargetFPS())

loop:
	for rt.frameLoopBegin(); ; {
		select {
		case task := <-rt.processQueue:
			rt.runTask(task)

		case <-gcTicker.C:
			rt.runGC()

		case <-rt.ctx.Done():
			break loop
		}
	}

	wg.Wait()
	close(rt.processQueue)

loopEnding:
	for {
		select {
		case task, ok := <-rt.processQueue:
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

func (rt *RuntimeBehavior) makeFrameTasks(wg *sync.WaitGroup, curFrames, totalFrames int64, targetFPS float32) {
	defer wg.Done()

	updateTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(targetFPS)))
	defer updateTicker.Stop()

	done := make(chan struct{}, 1)

	for {
		if totalFrames > 0 && curFrames >= totalFrames {
			rt.Terminate()
			return
		}

		select {
		case <-updateTicker.C:
			select {
			case rt.processQueue <- _Task{typ: _TaskType_Frame, action: rt.frameLoop, done: done}:
				select {
				case <-done:
					curFrames++
					continue
				case <-rt.ctx.Done():
					return
				}
			case <-rt.ctx.Done():
				return
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
	rt.changeRunningStatus(runtime.RunningStatus_FrameLoopBegin)
	rt.changeRunningStatus(runtime.RunningStatus_FrameUpdateBegin)

	_EmitEventUpdate(&rt.eventUpdate)
	_EmitEventLateUpdate(&rt.eventLateUpdate)

	rt.changeRunningStatus(runtime.RunningStatus_FrameUpdateEnd)
}

func (rt *RuntimeBehavior) frameLoopEnd() {
	rt.changeRunningStatus(runtime.RunningStatus_FrameLoopEnd)

	frame := runtime.UnsafeFrame(rt.opts.Frame)
	frame.SetCurFrames(frame.GetCurFrames() + 1)
}
