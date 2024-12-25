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

package pt

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/generic"
)

// Deprecated: UnsafeEntityLib 访问实体原型库的内部方法
func UnsafeEntityLib(entityLib EntityLib) _UnsafeEntityLib {
	return _UnsafeEntityLib{
		EntityLib: entityLib,
	}
}

type _UnsafeEntityLib struct {
	EntityLib
}

// SetCallback 设置回调
func (u _UnsafeEntityLib) SetCallback(declareCB, redeclareCB, undeclareCB generic.Action1[ec.EntityPT]) {
	u.setCallback(declareCB, redeclareCB, undeclareCB)
}
