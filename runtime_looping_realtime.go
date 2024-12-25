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
	"time"
)

func (rt *RuntimeBehavior) loopingRealTime() {
	gcTicker := time.NewTicker(rt.opts.GCInterval)
	defer gcTicker.Stop()

	frame := runtime.UnsafeFrame(rt.opts.Frame)
	go rt.makeFrameTasks(frame.GetCurFrames()+1, frame.GetTotalFrames(), frame.GetTargetFPS())

loop:
	for rt.frameLoopBegin(); ; {
		select {
		case task, ok := <-rt.processQueue:
			if !ok {
				break loop
			}
			rt.runTask(task)

		case <-gcTicker.C:
			rt.runGC()

		case <-rt.ctx.Done():
			break loop
		}
	}

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

func (rt *RuntimeBehavior) makeFrameTasks(curFrames, totalFrames int64, targetFPS float32) {
	updateTicker := time.NewTicker(time.Duration(float64(time.Second) / float64(targetFPS)))
	defer updateTicker.Stop()

	for {
		if totalFrames > 0 && curFrames >= totalFrames {
			rt.Terminate()
			return
		}

		select {
		case <-updateTicker.C:
			func() {
				defer func() {
					recover()
				}()
				select {
				case rt.processQueue <- _Task{typ: _TaskType_Frame, action: rt.frameLoop}:
					curFrames++
				case <-rt.ctx.Done():
				}
			}()
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
