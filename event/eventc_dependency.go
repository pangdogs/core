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

// Cache 引用iface.Cache
type Cache = iface.Cache

// Cache2Iface 引用iface.Cache2Iface
func Cache2Iface[T any](c Cache) T {
	return iface.Cache2Iface[T](c)
}

// Panicf 引用exception.Panicf
func Panicf(format string, args ...any) {
	exception.Panicf(format, args...)
}
