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

// EntityAtti 实体原型属性
type EntityAtti struct {
	Prototype          string    // 实体原型名称（必填）
	Instance           any       // 实体实例
	Scope              *ec.Scope // 可访问作用域
	AwakeOnFirstAccess *bool     // 设置开启组件被首次访问时，检测并调用Awake()
}

// EntityWith 创建实体原型属性，用于注册实体原型时自定义相关属性
func EntityWith(prototype string, inst any, scope *ec.Scope, awakeOnFirstAccess *bool) EntityAtti {
	return EntityAtti{
		Prototype:          prototype,
		Instance:           inst,
		Scope:              scope,
		AwakeOnFirstAccess: awakeOnFirstAccess,
	}
}

// CompAtti 组件原型属性
type CompAtti struct {
	Instance     any    // 组件实例（必填）
	Name         string // 组件名称
	NonRemovable bool   // 是否不可删除
}

// CompWith 创建组件原型属性，用于注册实体原型时自定义相关属性
func CompWith(inst any, name string, nonRemovable bool) CompAtti {
	return CompAtti{
		Instance:     inst,
		Name:         name,
		NonRemovable: nonRemovable,
	}
}

// CompInterfaceWith 创建组件原型属性，用于注册实体原型时使用接口名作为组件名称，并自定义相关属性
func CompInterfaceWith[FACE any](inst any, nonRemovable bool) CompAtti {
	return CompAtti{
		Instance:     inst,
		Name:         types.FullNameT[FACE](),
		NonRemovable: nonRemovable,
	}
}
