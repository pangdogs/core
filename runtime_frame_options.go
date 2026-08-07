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

// FrameOptions 定义运行时帧循环的选项。
type FrameOptions struct {
	Enabled     bool    // 是否启用帧循环。
	TargetFPS   float64 // 目标 FPS；设置时会四舍五入为整数值。
	TotalFrames int64   // 最大运行帧数；0 表示不限制。
}

type _FrameOption struct{}

// Default 返回帧选项的默认设置。
func (_FrameOption) Default() option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		With.Frame.Enabled(true).Apply(options)
		With.Frame.TargetFPS(30).Apply(options)
		With.Frame.TotalFrames(0).Apply(options)
	}
}

// Enabled 设置是否启用帧循环。
func (_FrameOption) Enabled(b bool) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		options.Enabled = b
	}
}

// TargetFPS 设置目标 FPS；fps 必须大于 0，并会被四舍五入为整数值。
func (_FrameOption) TargetFPS(fps float64) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		if fps <= 0 {
			exception.Panicf("%w: %w: TargetFPS must be greater than 0", runtime.ErrFrame, exception.ErrArgs)
		}
		options.TargetFPS = math.Round(fps)
	}
}

// TotalFrames 设置最大运行帧数；0 表示不限制，负值会导致 panic。
func (_FrameOption) TotalFrames(v int64) option.Setting[FrameOptions] {
	return func(options *FrameOptions) {
		if v < 0 {
			exception.Panicf("%w: %w: TotalFrames must be greater than or equal to 0", runtime.ErrFrame, exception.ErrArgs)
		}
		options.TotalFrames = v
	}
}
