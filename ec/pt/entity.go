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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"reflect"
	"slices"
)

type _Entity struct {
	prototype                  string
	instanceRT                 reflect.Type
	scope                      *ec.Scope
	componentAwakeOnFirstTouch *bool
	componentUniqueID          *bool
	extra                      generic.SliceMap[string, any]
	components                 []ec.BuiltinComponent
}

// Prototype 实体原型名称
func (pt *_Entity) Prototype() string {
	return pt.prototype
}

// InstanceRT 实体实例反射类型
func (pt *_Entity) InstanceRT() reflect.Type {
	return reflect.PointerTo(pt.instanceRT)
}

// Scope 可访问作用域
func (pt *_Entity) Scope() *ec.Scope {
	return pt.scope
}

// ComponentAwakeOnFirstTouch 开启组件被首次访问时，检测并调用Awake()
func (pt *_Entity) ComponentAwakeOnFirstTouch() *bool {
	return pt.componentAwakeOnFirstTouch
}

// ComponentUniqueID 开启组件唯一Id
func (pt *_Entity) ComponentUniqueID() *bool {
	return pt.componentUniqueID
}

// CountComponents // 组件数量
func (pt *_Entity) CountComponents() int {
	return len(pt.components)
}

// Extra 自定义原型属性
func (pt *_Entity) Extra() generic.SliceMap[string, any] {
	return pt.extra
}

// Component 获取组件
func (pt *_Entity) Component(idx int) ec.BuiltinComponent {
	if idx < 0 || idx >= len(pt.components) {
		exception.Panicf("%w: %w: idx out of range", ErrPt, exception.ErrArgs)
	}
	return pt.components[idx]
}

// Components 获取所有组件
func (pt *_Entity) Components() []ec.BuiltinComponent {
	return slices.Clone(pt.components)
}

// Construct 创建实体
func (pt *_Entity) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.Make(ec.With.Default())
	if pt.instanceRT != nil {
		options.InstanceFace = iface.MakeFaceT(reflect.New(pt.instanceRT).Interface().(ec.Entity))
	}
	if pt.scope != nil {
		options.Scope = *pt.scope
	}
	if pt.componentAwakeOnFirstTouch != nil {
		options.ComponentAwakeOnFirstTouch = *pt.componentAwakeOnFirstTouch
	}
	if pt.componentUniqueID != nil {
		options.ComponentUniqueID = *pt.componentUniqueID
	}
	options = option.Append(options, settings...)

	return pt.assemble(ec.UnsafeNewEntity(options))
}

func (pt *_Entity) assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrPt, exception.ErrArgs)
	}

	ec.UnsafeEntity(entity).SetPT(pt)

	for i := range pt.components {
		builtin := &pt.components[i]

		comp := builtin.PT.Construct()
		ec.UnsafeComponent(comp).SetBuiltin(builtin)
		ec.UnsafeComponent(comp).SetNonRemovable(builtin.NonRemovable)

		if err := entity.AddComponent(builtin.Name, comp); err != nil {
			exception.Panicf("%w: %w", ErrPt, err)
		}
	}

	return entity
}
