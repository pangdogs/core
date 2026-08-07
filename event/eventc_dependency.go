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

package event

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// Cache 是 iface.Cache 的别名，供 eventc 生成代码使用。
type Cache = iface.Cache

// Cache2Iface 将接口缓存恢复为目标接口，供 eventc 生成代码使用。
func Cache2Iface[T any](c Cache) T {
	return iface.Cache2Iface[T](c)
}

// Panicf 按框架约定抛出格式化 panic，供 eventc 生成代码使用。
func Panicf(format string, args ...any) {
	exception.Panicf(format, args...)
}
