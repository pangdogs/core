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

// Frame 帧，在运行时初始化时可以设置帧，用于设置运行时帧更新方式，在逻辑运行过程中可以在运行时上下文中获取帧信息
type Frame interface {
	// TargetFPS 获取目标FPS
	TargetFPS() float64
	// CurFPS 获取当前FPS
	CurFPS() float64
	// TotalFrames 获取运行帧数上限
	TotalFrames() int64
	// CurFrames 获取当前帧数
	CurFrames() int64
	// RunningBeginTime 获取运行开始时间
	RunningBeginTime() time.Time
	// RunningElapseTime 获取运行持续时间
	RunningElapseTime() time.Duration
	// LoopBeginTime 获取当前帧循环开始时间（包含异步调用）
	LoopBeginTime() time.Time
	// LastLoopElapseTime 获取上一帧循环耗时（包含异步调用）
	LastLoopElapseTime() time.Duration
	// UpdateBeginTime 获取当前帧更新开始时间
	UpdateBeginTime() time.Time
	// LastUpdateElapseTime 获取上一次帧更新耗时
	LastUpdateElapseTime() time.Duration
}
