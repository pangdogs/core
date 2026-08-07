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

package option

import (
	"git.golaxy.org/core/utils/generic"
)

// Setting 是按地址修改选项值的函数式设置项；nil Setting 是空操作。
type Setting[T any] generic.Action1[*T]

// Apply 将设置项应用到 options。
func (s Setting[T]) Apply(options *T) {
	generic.CastAction1(s).UnsafeCall(options)
}

// New 从 T 的零值开始，先应用 defaults，再按给定顺序应用 settings。
func New[T any](defaults Setting[T], settings ...Setting[T]) (options T) {
	defaults.Apply(&options)

	for i := range settings {
		settings[i].Apply(&options)
	}

	return
}

// Append 在 options 的浅拷贝上按顺序应用 settings，并返回结果。
func Append[T any](options T, settings ...Setting[T]) T {
	for i := range settings {
		settings[i].Apply(&options)
	}
	return options
}

// Change 在 options 的浅拷贝上按顺序应用 settings，并返回结果。
//
// Change 当前与 Append 语义相同，用于在调用处表达“修改现有选项”的意图。
func Change[T any](options T, settings ...Setting[T]) T {
	for i := range settings {
		settings[i].Apply(&options)
	}
	return options
}
