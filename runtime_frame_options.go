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
	"math"

	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/option"
)

// FrameOptions 帧的所有选项
type FrameOptions struct {
	Enabled     bool    // 是否启用帧
	TargetFPS   float64 // 目标FPS
	TotalFrames int64   // 运行帧数上限
}

type _FrameOption struct{}

// Default 默认值
func (_FrameOption) Default() option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		With.Frame.Enabled(true).Apply(options)
		With.Frame.TargetFPS(30).Apply(options)
		With.Frame.TotalFrames(0).Apply(options)
	}
}

// Enabled 是否启用帧
func (_FrameOption) Enabled(b bool) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		options.Enabled = b
	}
}

// TargetFPS 目标FPS
func (_FrameOption) TargetFPS(fps float64) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		if fps <= 0 {
			exception.Panicf("%w: %w: TargetFPS less equal 0 is invalid", runtime.ErrFrame, exception.ErrArgs)
		}
		options.TargetFPS = math.Round(fps)
	}
}

// TotalFrames 运行帧数上限
func (_FrameOption) TotalFrames(v int64) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		if v < 0 {
			exception.Panicf("%w: %w: TotalFrames less 0 is invalid", runtime.ErrFrame, exception.ErrArgs)
		}
		options.TotalFrames = v
	}
}
