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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"time"
)

type (
	CustomGC = generic.Action1[Runtime] // 自定义GC函数
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	InstanceFace         iface.Face[Runtime] // 实例，用于扩展运行时能力
	AutoRun              bool                // 是否开启自动运行
	ProcessQueueCapacity int                 // 任务处理流水线大小
	Frame                runtime.Frame       // 帧，设置为nil表示不使用帧更新特性
	GCInterval           time.Duration       // GC间隔时长
	CustomGC             CustomGC            // 自定义GC
}

type _RuntimeOption struct{}

// Default 运行时的默认值
func (_RuntimeOption) Default() option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		With.Runtime.InstanceFace(iface.Face[Runtime]{})(o)
		With.Runtime.AutoRun(false)(o)
		With.Runtime.ProcessQueueCapacity(128)(o)
		With.Runtime.Frame(nil)(o)
		With.Runtime.GCInterval(10 * time.Second)(o)
		With.Runtime.CustomGC(nil)(o)
	}
}

// InstanceFace 实例，用于扩展运行时能力
func (_RuntimeOption) InstanceFace(face iface.Face[Runtime]) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.InstanceFace = face
	}
}

// AutoRun 运行时是否开启自动运行
func (_RuntimeOption) AutoRun(b bool) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.AutoRun = b
	}
}

// ProcessQueueCapacity 任务处理流水线大小
func (_RuntimeOption) ProcessQueueCapacity(cap int) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		if cap <= 0 {
			exception.Panicf("%w: %w: ProcessQueueCapacity less equal 0 is invalid", ErrRuntime, ErrArgs)
		}
		o.ProcessQueueCapacity = cap
	}
}

// Frame 运行时的帧，设置为nil表示不使用帧更新特性
func (_RuntimeOption) Frame(frame runtime.Frame) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.Frame = frame
	}
}

// GCInterval 运行时的GC间隔时长
func (_RuntimeOption) GCInterval(dur time.Duration) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		if dur <= 0 {
			exception.Panicf("%w: %w: GCInterval less equal 0 is invalid", ErrRuntime, ErrArgs)
		}
		o.GCInterval = dur
	}
}

// CustomGC 运行时的自定义GC
func (_RuntimeOption) CustomGC(fn CustomGC) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.CustomGC = fn
	}
}
