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
	"git.golaxy.org/core/utils/exception"
	"reflect"
)

// Face 用于存储接口与Cache，接口可用于断言转换类型，存储器可用于转换为接口
type Face[T any] struct {
	Iface T
	Cache Cache
}

// IsNil 是否为空
func (f *Face[T]) IsNil() bool {
	return Iface2Cache[T](f.Iface) == NilCache || f.Cache == NilCache
}

// FaceAny any face
type FaceAny = Face[any]

// MakeFaceT 创建Face，不使用Cache
func MakeFaceT[T any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// MakeFaceAny 创建FaceAny，使用Cache，不使用接口
func MakeFaceAny[C any](iface C) FaceAny {
	return Face[any]{
		Iface: iface,
		Cache: Iface2Cache[C](iface),
	}
}

// MakeFaceTC 创建Face，使用Cache，传入接口与Cache
func MakeFaceTC[T, C any](iface T, cache C) Face[T] {
	if Iface2Cache(iface)[1] != Iface2Cache(cache)[1] {
		exception.Panicf("%w: incorrect face pointer", exception.ErrCore)
	}
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](cache),
	}
}

// MakeFaceTReflectC 创建Face，使用Cache，自动反射Cache类型
func MakeFaceTReflectC[T, C any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](reflect.ValueOf(iface).Interface().(C)),
	}
}
