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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/option"
)

// TaskQueueOptions 定义运行时任务队列的容量策略。
type TaskQueueOptions struct {
	Unbounded bool // 是否使用无界队列。
	Capacity  int  // 有界队列容量；使用无界队列时忽略。
}

type _TaskQueueOption struct{}

// Default 返回任务队列选项的默认设置。
func (_TaskQueueOption) Default() option.Setting[TaskQueueOptions] {
	return func(options *TaskQueueOptions) {
		With.TaskQueue.Unbounded(true).Apply(options)
		With.TaskQueue.Capacity(128).Apply(options)
	}
}

// Unbounded 设置是否使用无界队列。
func (_TaskQueueOption) Unbounded(b bool) option.Setting[TaskQueueOptions] {
	return func(options *TaskQueueOptions) {
		options.Unbounded = b
	}
}

// Capacity 设置队列容量，cap 必须大于 0。
func (_TaskQueueOption) Capacity(cap int) option.Setting[TaskQueueOptions] {
	return func(options *TaskQueueOptions) {
		if cap <= 0 {
			exception.Panicf("%w: %w: Capacity must be greater than 0", ErrRuntime, exception.ErrArgs)
		}
		options.Capacity = cap
	}
}
