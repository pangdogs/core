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

// Signed 有符号整形
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 无符号整形
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer 整形
type Integer interface {
	Signed | Unsigned
}

// Bool2Int bool转int
func Bool2Int[T Integer](b bool) T {
	if b {
		return 1
	}
	return 0
}

// Int2Bool int转bool
func Int2Bool[T Integer](v T) bool {
	if v != 0 {
		return true
	}
	return false
}

// String2Bytes 快速string转bytes
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Bytes2String 快速bytes转string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Panic2Err panic转换为error
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
