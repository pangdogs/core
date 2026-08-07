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

package runtime

import (
	"time"
)

// Frame 提供运行时帧循环的配置值和只读统计信息。
type Frame interface {
	// TargetFPS 返回目标 FPS。
	TargetFPS() float64
	// CurFPS 返回最近一个统计周期内的实际 FPS。
	CurFPS() float64
	// TotalFrames 返回最大运行帧数；0 表示不限制。
	TotalFrames() int64
	// CurFrames 返回已经开始执行的帧数。
	CurFrames() int64
	// RunningBeginTime 返回帧循环的启动时间。
	RunningBeginTime() time.Time
	// RunningElapseTime 返回已累计的帧循环运行时长。
	RunningElapseTime() time.Duration
	// LoopBeginTime 返回当前帧循环的开始时间；循环包含异步调用处理。
	LoopBeginTime() time.Time
	// LastLoopElapseTime 返回上一帧完整循环的耗时，包含异步调用处理。
	LastLoopElapseTime() time.Duration
	// UpdateBeginTime 返回当前帧更新阶段的开始时间。
	UpdateBeginTime() time.Time
	// LastUpdateElapseTime 返回上一帧更新阶段的耗时。
	LastUpdateElapseTime() time.Duration
}
