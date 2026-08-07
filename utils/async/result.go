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

// NewResult 将 value 与 err 组合为异步结果。
func NewResult(value any, err error) Result {
	return Result{
		Value: value,
		Error: err,
	}
}

// Result 保存 Future 的一次产出值或错误。
type Result struct {
	Value any   // Value 是本次产出携带的返回值。
	Error error // Error 非 nil 时表示本次产出失败。
}

// OK 报告结果是否不含错误。
func (ret Result) OK() bool {
	return ret.Error == nil
}

// String 返回错误文本；无错误时返回 Value 的默认格式。
func (ret Result) String() string {
	if ret.Error != nil {
		return ret.Error.Error()
	}
	return fmt.Sprintf("%v", ret.Value)
}
