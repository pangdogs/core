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

package async

import "fmt"

var (
	VoidRet = MakeRet(nil, nil) // 空返回值
)

// Ret 调用结果
type Ret = RetT[any]

// MakeRet 创建调用结果
func MakeRet(val any, err error) Ret {
	return MakeRetT(val, err)
}

// MakeRetT 创建调用结果
func MakeRetT[T any](val T, err error) RetT[T] {
	return RetT[T]{
		Value: val,
		Error: err,
	}
}

// CastRetT 转换
func CastRetT[T any](ret Ret) RetT[T] {
	if ret.Value == nil || ret.Error != nil {
		return RetT[T]{
			Error: ret.Error,
		}
	}
	return RetT[T]{
		Value: ret.Value.(T),
		Error: ret.Error,
	}
}

// AsRetT 转换
func AsRetT[T any](ret Ret) (RetT[T], bool) {
	if ret.Value == nil || ret.Error != nil {
		return RetT[T]{
			Error: ret.Error,
		}, true
	}
	v, ok := ret.Value.(T)
	if !ok {
		return RetT[T]{}, false
	}
	return RetT[T]{
		Value: v,
		Error: ret.Error,
	}, true
}

// RetT 调用结果
type RetT[T any] struct {
	Value T     // 返回值
	Error error // error
}

// OK 是否成功
func (ret RetT[T]) OK() bool {
	return ret.Error == nil
}

// String implements fmt.Stringer
func (ret RetT[T]) String() string {
	if ret.Error != nil {
		return ret.Error.Error()
	}
	return fmt.Sprintf("%v", ret.Value)
}
