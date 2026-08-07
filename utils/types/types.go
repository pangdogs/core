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

package types

import (
	"reflect"
	"strings"
)

// Zero 返回 T 的零值。
func Zero[T any]() T {
	var zero T
	return zero
}

// New 返回指向新 T 零值的指针。
func New[T any]() *T {
	var zero T
	return &zero
}

// Pointer 返回指向 src 副本的指针。
func Pointer[T any](src T) *T {
	return &src
}

// Name 返回值或 reflect.Type 的具名类型名；参数没有动态类型时 panic。
func Name(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return t.Name()
}

// NameRT 返回反射类型的具名类型名；t 为 nil 时 panic。
func NameRT(t reflect.Type) string {
	return t.Name()
}

// NameT 返回 T 的具名类型名。
func NameT[T any]() string {
	return reflect.TypeFor[T]().Name()
}

// FullName 返回值或 reflect.Type 的“包导入路径.类型名”；参数没有动态类型时 panic。
func FullName(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return FullNameRT(t)
}

// FullNameRT 返回 t 的“包导入路径.类型名”；无包路径时仅返回类型名。
func FullNameRT(t reflect.Type) string {
	pkgPath := t.PkgPath()
	name := t.Name()
	if pkgPath == "" {
		return name
	}
	return t.PkgPath() + "." + t.Name()
}

// FullNameT 返回 T 的“包导入路径.类型名”。
func FullNameT[T any]() string {
	return FullNameRT(reflect.TypeFor[T]())
}

// WriteFullName 将值或 reflect.Type 的完整名称写入 sb。
func WriteFullName(sb *strings.Builder, i any) {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	WriteFullNameRT(sb, t)
}

// WriteFullNameRT 将 t 的完整名称写入 sb。
func WriteFullNameRT(sb *strings.Builder, t reflect.Type) {
	pkgPath := t.PkgPath()
	name := t.Name()
	if pkgPath == "" {
		sb.WriteString(name)
		return
	}
	sb.WriteString(pkgPath)
	sb.WriteString(".")
	sb.WriteString(name)
}

// WriteFullNameT 将 T 的完整名称写入 sb。
func WriteFullNameT[T any](sb *strings.Builder) {
	WriteFullNameRT(sb, reflect.TypeFor[T]())
}
