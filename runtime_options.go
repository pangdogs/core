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

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
)

type (
	CustomGC = generic.Action1[Runtime] // 自定义GC函数
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	InstanceFace                    iface.Face[Runtime] // 实例，用于扩展运行时能力
	AutoRun                         bool                // 是否开启自动运行
	ContinueOnActivatingEntityPanic bool                // 激活实体时发生panic是否继续，不继续将会主动删除实体
	Frame                           FrameOptions        // 帧配置
	TaskQueue                       TaskQueueOptions    // 任务处理流水线配置
	GCInterval                      time.Duration       // GC间隔时长
	CustomGC                        CustomGC            // 自定义GC
}

type _RuntimeOption struct{}

// Default 运行时的默认值
func (_RuntimeOption) Default() option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		With.Runtime.InstanceFace(iface.Face[Runtime]{}).Apply(options)
		With.Runtime.AutoRun(false).Apply(options)
		With.Runtime.ContinueOnActivatingEntityPanic(false).Apply(options)
		With.Runtime.Frame(With.Frame.Default()).Apply(options)
		With.Runtime.TaskQueue(With.TaskQueue.Default()).Apply(options)
		With.Runtime.GCInterval(10 * time.Second).Apply(options)
		With.Runtime.CustomGC(nil).Apply(options)
	}
}

// InstanceFace 实例，用于扩展运行时能力
func (_RuntimeOption) InstanceFace(face iface.Face[Runtime]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.InstanceFace = face
	}
}

// AutoRun 运行时是否开启自动运行
func (_RuntimeOption) AutoRun(b bool) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.AutoRun = b
	}
}

// ContinueOnActivatingEntityPanic 激活实体时发生panic是否继续，不继续将会主动删除实体
func (_RuntimeOption) ContinueOnActivatingEntityPanic(b bool) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.ContinueOnActivatingEntityPanic = b
	}
}

// Frame 帧配置
func (_RuntimeOption) Frame(settings ...option.Setting[FrameOptions]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.Frame = option.Append(options.Frame, settings...)
	}
}

// TaskQueue 任务处理流水线配置
func (_RuntimeOption) TaskQueue(settings ...option.Setting[TaskQueueOptions]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.TaskQueue = option.Append(options.TaskQueue, settings...)
	}
}

// GCInterval 运行时的GC间隔时长
func (_RuntimeOption) GCInterval(dur time.Duration) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		if dur <= 0 {
			exception.Panicf("%w: %w: GCInterval must be greater than 0", ErrRuntime, ErrArgs)
		}
		options.GCInterval = dur
	}
}

// CustomGC 运行时的自定义GC
func (_RuntimeOption) CustomGC(fn CustomGC) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.CustomGC = fn
	}
}
