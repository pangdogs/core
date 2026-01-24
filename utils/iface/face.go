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
	"reflect"

	"git.golaxy.org/core/utils/exception"
)

// FaceAny face with any
type FaceAny = Face[any]

// MakeFaceAny 创建 FaceAny
func MakeFaceAny[C any](cache C) FaceAny {
	return Face[any]{
		Iface: cache,
		Cache: Iface2Cache[C](cache),
	}
}

// Face 面，用于存储接口与接口存储器，接口用于断言转换类型，接口存储器用于重解释接口
type Face[T any] struct {
	Iface T     // 接口
	Cache Cache // 接口存储器
}

// IsNil 是否为空
func (f *Face[T]) IsNil() bool {
	return Iface2Cache[T](f.Iface) == NilCache || f.Cache == NilCache
}

// MakeFaceT 创建面（Face），接口存储器重解释接口与接口类型相同
func MakeFaceT[T any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// MakeFaceTC 创建面（Face），接口存储器重解释接口与接口类型可以不同
func MakeFaceTC[T, C any](iface T, cache C) Face[T] {
	if Iface2Cache(iface)[1] != Iface2Cache(cache)[1] {
		exception.Panicf("%w: incorrect face pointer", exception.ErrCore)
	}
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](cache),
	}
}

// MakeFaceTReflectC 创建面（Face），自动反射获取接口存储器
func MakeFaceTReflectC[T, C any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](reflect.ValueOf(iface).Interface().(C)),
	}
}
