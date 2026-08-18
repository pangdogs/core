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

package uid

import "github.com/rs/xid"

var (
	// Nil 是空 ID。
	Nil ID = ""

	// New 生成新的 xid 字符串 ID。
	New = func() ID {
		return ID(xid.New().String())
	}

	// From 将字符串直接转换为 ID，不执行格式校验。
	From = func(str string) ID {
		return ID(str)
	}
)

// ID 是框架统一使用的字符串标识类型。
type ID string

// IsNil 报告 ID 是否为空。
func (id ID) IsNil() bool {
	return id == Nil
}

// String 返回 ID 的原始字符串。
func (id ID) String() string {
	return string(id)
}
