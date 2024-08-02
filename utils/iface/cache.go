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

package iface

import (
	"unsafe"
)

// Cache 因为Golang原生的接口转换性能较差，所以在某些性能要求较高的场景下，需要尽量较少接口转换。
// 如果必须转换接口，那么目前可用的优化方案是，在编码时已知接口类型，可以将接口转换为[2]unsafe.Pointer，使用时再转换回接口。
// Cache套件就是使用此优化方案，注意不安全，在明确了解此方案时再使用。
type Cache [2]unsafe.Pointer

// NilCache nil cache
var NilCache Cache

// Cache2Iface Cache转换为接口
func Cache2Iface[T any](c Cache) T {
	return *(*T)(unsafe.Pointer(&c))
}

// Iface2Cache 接口转换为Cache
func Iface2Cache[T any](i T) Cache {
	return *(*Cache)(unsafe.Pointer(&i))
}
