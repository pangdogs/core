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

// FaceAny 是以 any 保存公开接口值的 Face。
type FaceAny = Face[any]

// NewFaceAny 创建 FaceAny，并以 C 的接口视图建立缓存。
func NewFaceAny[C any](cache C) FaceAny {
	return Face[any]{
		Iface: cache,
		Cache: Iface2Cache[C](cache),
	}
}

// Face 同时保存公开接口值与用于快速重解释的接口缓存。
type Face[T any] struct {
	Iface T     // Iface 是安全访问与维持实例生命周期所用的接口值。
	Cache Cache // Cache 是预先选择的接口视图缓存。
}

// IsNil 报告公开接口值或接口缓存是否为空。
func (f *Face[T]) IsNil() bool {
	return Iface2Cache[T](f.Iface) == NilCache || f.Cache == NilCache
}

// NewFaceT 创建公开接口与缓存接口类型相同的 Face。
func NewFaceT[T any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// NewFaceTC 创建公开接口类型为 T、缓存接口类型为 C 的 Face。
// iface 与 cache 必须引用同一实例，否则 panic。
func NewFaceTC[T, C any](iface T, cache C) Face[T] {
	if Iface2Cache(iface)[1] != Iface2Cache(cache)[1] {
		exception.Panicf("%w: incorrect face pointer", exception.ErrCore)
	}
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](cache),
	}
}

// NewFaceTReflectC 通过反射把 iface 断言为 C，并以该视图创建缓存。
// iface 未实现 C 或为无效反射值时 panic。
func NewFaceTReflectC[T, C any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[C](reflect.ValueOf(iface).Interface().(C)),
	}
}
