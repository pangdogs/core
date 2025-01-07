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
	"github.com/elliotchance/pie/v2"
)

// EntityAttribute 实体原型属性
type EntityAttribute struct {
	Prototype                  string                        // 实体原型名称（必填）
	Instance                   any                           // 实体实例
	Scope                      *ec.Scope                     // 可访问作用域
	ComponentAwakeOnFirstTouch *bool                         // 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentUniqueID          *bool                         // 是否为实体组件分配唯一Id
	Extra                      generic.SliceMap[string, any] // 自定义属性
}

// EntityWith 创建实体原型属性，用于注册实体原型时自定义相关属性
func EntityWith(prototype string, inst any, scope *ec.Scope, componentAwakeOnFirstTouch, componentUniqueID *bool, extra ...map[string]any) EntityAttribute {
	return EntityAttribute{
		Prototype:                  prototype,
		Instance:                   inst,
		Scope:                      scope,
		ComponentAwakeOnFirstTouch: componentAwakeOnFirstTouch,
		ComponentUniqueID:          componentUniqueID,
		Extra:                      generic.MakeSliceMapFromGoMap(pie.First(extra)),
	}
}

// ComponentAttribute 组件原型属性
type ComponentAttribute struct {
	Instance  any                           // 组件实例（必填）
	Name      string                        // 组件名称
	removable bool                          // 是否可以删除
	Extra     generic.SliceMap[string, any] // 自定义属性
}

// ComponentWith 创建组件原型属性，用于注册实体原型时自定义相关属性
func ComponentWith(inst any, name string, removable bool, extra ...map[string]any) ComponentAttribute {
	return ComponentAttribute{
		Instance:  inst,
		Name:      name,
		removable: removable,
		Extra:     generic.MakeSliceMapFromGoMap(pie.First(extra)),
	}
}
