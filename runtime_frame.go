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
	"time"
)

type _Frame struct {
	targetFPS            float64
	totalFrames          int64
	curFPS               float64
	curFrames            int64
	runningBeginTime     time.Time
	runningElapseTime    time.Duration
	loopBeginTime        time.Time
	lastLoopElapseTime   time.Duration
	updateBeginTime      time.Time
	lastUpdateElapseTime time.Duration
	statFPSBeginTime     time.Time
	statFPSFrames        int64
}

// TargetFPS 返回目标 FPS。
func (frame *_Frame) TargetFPS() float64 {
	return frame.targetFPS
}

// CurFPS 返回最近一个统计周期内的实际 FPS。
func (frame *_Frame) CurFPS() float64 {
	return frame.curFPS
}

// TotalFrames 返回最大运行帧数；0 表示不限制。
func (frame *_Frame) TotalFrames() int64 {
	return frame.totalFrames
}

// CurFrames 返回已经开始执行的帧数。
func (frame *_Frame) CurFrames() int64 {
	return frame.curFrames
}

// RunningBeginTime 返回帧循环的启动时间。
func (frame *_Frame) RunningBeginTime() time.Time {
	return frame.runningBeginTime
}

// RunningElapseTime 返回已累计的帧循环运行时长。
func (frame *_Frame) RunningElapseTime() time.Duration {
	return frame.runningElapseTime
}

// LoopBeginTime 返回当前帧循环的开始时间；循环包含异步调用处理。
func (frame *_Frame) LoopBeginTime() time.Time {
	return frame.loopBeginTime
}

// LastLoopElapseTime 返回上一帧完整循环的耗时，包含异步调用处理。
func (frame *_Frame) LastLoopElapseTime() time.Duration {
	return frame.lastLoopElapseTime
}

// UpdateBeginTime 返回当前帧更新阶段的开始时间。
func (frame *_Frame) UpdateBeginTime() time.Time {
	return frame.updateBeginTime
}

// LastUpdateElapseTime 返回上一帧更新阶段的耗时。
func (frame *_Frame) LastUpdateElapseTime() time.Duration {
	return frame.lastUpdateElapseTime
}

func (frame *_Frame) init(targetFPS float64, totalFrames int64) {
	frame.targetFPS = targetFPS
	frame.totalFrames = totalFrames
}

func (frame *_Frame) setCurFrames(v int64) {
	frame.curFrames = v
}

func (frame *_Frame) runningBegin() {
	now := time.Now()

	frame.curFPS = 0
	frame.curFrames = 0

	frame.statFPSBeginTime = now
	frame.statFPSFrames = 0

	frame.runningBeginTime = now
	frame.runningElapseTime = 0

	frame.loopBeginTime = now
	frame.lastLoopElapseTime = 0

	frame.updateBeginTime = now
	frame.lastUpdateElapseTime = 0
}

func (frame *_Frame) runningEnd() {
}

func (frame *_Frame) loopBegin() {
	now := time.Now()

	frame.loopBeginTime = now

	statInterval := now.Sub(frame.statFPSBeginTime).Seconds()
	if statInterval >= 1 {
		frame.curFPS = float64(frame.statFPSFrames) / statInterval
		frame.statFPSBeginTime = now
		frame.statFPSFrames = 0
	}
}

func (frame *_Frame) loopEnd() {
	frame.lastLoopElapseTime = time.Now().Sub(frame.loopBeginTime)
	frame.runningElapseTime += frame.lastLoopElapseTime
	frame.statFPSFrames++
}

func (frame *_Frame) updateBegin() {
	frame.updateBeginTime = time.Now()
}

func (frame *_Frame) updateEnd() {
	frame.lastUpdateElapseTime = time.Now().Sub(frame.updateBeginTime)
}
