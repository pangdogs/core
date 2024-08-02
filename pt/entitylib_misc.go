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
	"git.golaxy.org/core/utils/types"
)

type _CompAlias struct {
	Comp  any
	Alias string
	Fixed bool
}

// CompAlias 组件与别名，用于注册实体原型时自定义组件别名
func CompAlias(comp any, fixed bool, alias string) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: alias,
		Fixed: fixed,
	}
}

// CompInterface 组件与接口，用于注册实体原型时使用接口名作为别名
func CompInterface[FACE any](comp any, fixed bool) _CompAlias {
	return _CompAlias{
		Comp:  comp,
		Alias: types.FullNameT[FACE](),
		Fixed: fixed,
	}
}

// Attribute 实体原型属性
type Attribute struct {
	Composite          any       // 实体类型
	Scope              *ec.Scope // 可访问作用域
	AwakeOnFirstAccess *bool     // 设置开启组件被首次访问时，检测并调用Awake()
}
