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
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// Signed 约束所有有符号整数及其自定义底层类型。
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 约束所有无符号整数及其自定义底层类型。
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer 约束所有整数类型。
type Integer interface {
	Signed | Unsigned
}

// Bool2Int 将 false 转为 0、true 转为 1，并返回目标整数类型。
func Bool2Int[T Integer](b bool) T {
	if b {
		return 1
	}
	return 0
}

// Int2Bool 将 0 转为 false、其他整数转为 true。
func Int2Bool[T Integer](v T) bool {
	if v != 0 {
		return true
	}
	return false
}

// String2Bytes 将字符串零拷贝重解释为字节切片。
//
// 返回值与 s 共享存储，绝不能修改；违反约束可能导致崩溃或未定义行为。
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Bytes2String 将字节切片零拷贝重解释为字符串。
//
// 返回值与 b 共享存储；字符串存活期间不得修改或并发写入 b。
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Panic2Err 将 recover 的结果转换为 error；nil 保持为 nil。
func Panic2Err(panicInfo any) error {
	switch info := panicInfo.(type) {
	case nil:
		return nil
	case error:
		return info
	case string:
		return errors.New(info)
	default:
		return fmt.Errorf("%v", panicInfo)
	}
}
