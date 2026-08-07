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

var (
	NilCache Cache // NilCache 是空接口缓存。
)

// Cache 保存 Go 接口值的类型字与数据字，用于跳过高频接口断言。
// Cache 只能保存接口值，不能直接保存具体类型值。
//
// 该表示依赖 Go 运行时内部布局。恢复时的目标必须是缓存对象确实实现的接口类型；
// Cache2Iface 不会进行运行时兼容性检查，错误使用可能导致未定义行为。
type Cache [2]unsafe.Pointer

// Cache2Iface 将 c 直接重解释为 T，不执行类型兼容性检查；T 必须是接口类型。
func Cache2Iface[T any](c Cache) T {
	return *(*T)(unsafe.Pointer(&c))
}

// Iface2Cache 将接口值 i 的内部表示保存为 Cache；T 必须是接口类型，不能是具体类型。
func Iface2Cache[T any](i T) Cache {
	return *(*Cache)(unsafe.Pointer(&i))
}
