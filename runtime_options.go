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
	CustomGC = generic.Action1[Runtime] // 运行时完成内置清理后调用的自定义 GC 函数。
)

// RuntimeOptions 定义创建运行时及其工作循环时使用的选项。
type RuntimeOptions struct {
	InstanceFace                    iface.Face[Runtime] // 自定义运行时实例及其接口缓存。
	AutoRun                         bool                // 是否在 RunningEvent_Birth 后自动启动运行时。
	ContinueOnActivatingEntityPanic bool                // 激活实体发生 panic 后是否继续；为 false 时销毁该实体。
	Frame                           FrameOptions        // 帧循环配置。
	TaskQueue                       TaskQueueOptions    // 任务队列配置。
	GCInterval                      time.Duration       // 两次运行时 GC 之间的最短间隔。
	CustomGC                        CustomGC            // 内置清理完成后执行的自定义 GC。
}

type _RuntimeOption struct{}

// Default 返回运行时选项的默认设置。
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

// InstanceFace 设置用于扩展运行时能力的自定义实例。
func (_RuntimeOption) InstanceFace(face iface.Face[Runtime]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.InstanceFace = face
	}
}

// AutoRun 设置是否在 RunningEvent_Birth 后自动启动运行时。
func (_RuntimeOption) AutoRun(b bool) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.AutoRun = b
	}
}

// ContinueOnActivatingEntityPanic 设置激活实体发生 panic 后是否继续。
// 设置为 false 时，运行时会主动销毁激活失败的实体。
func (_RuntimeOption) ContinueOnActivatingEntityPanic(b bool) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.ContinueOnActivatingEntityPanic = b
	}
}

// Frame 追加帧循环设置。
func (_RuntimeOption) Frame(settings ...option.Setting[FrameOptions]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.Frame = option.Append(options.Frame, settings...)
	}
}

// TaskQueue 追加任务队列设置。
func (_RuntimeOption) TaskQueue(settings ...option.Setting[TaskQueueOptions]) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.TaskQueue = option.Append(options.TaskQueue, settings...)
	}
}

// GCInterval 设置运行时 GC 的最短间隔，dur 必须大于 0。
func (_RuntimeOption) GCInterval(dur time.Duration) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		if dur <= 0 {
			exception.Panicf("%w: %w: GCInterval must be greater than 0", ErrRuntime, ErrArgs)
		}
		options.GCInterval = dur
	}
}

// CustomGC 设置内置清理完成后执行的自定义 GC 函数。
func (_RuntimeOption) CustomGC(fn CustomGC) option.Setting[RuntimeOptions] {
	return func(options *RuntimeOptions) {
		options.CustomGC = fn
	}
}
