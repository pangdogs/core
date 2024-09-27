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
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"reflect"
	"slices"
)

// ComponentDesc 组件描述
type ComponentDesc struct {
	PT    ComponentPT // 原型
	Alias string      // 别名
	Fixed bool        // 固定
}

// EntityPT 实体原型接口
type EntityPT interface {
	// Prototype 实体原型名称
	Prototype() string
	// InstanceRT 实体实例反射类型
	InstanceRT() reflect.Type
	// Scope 可访问作用域
	Scope() *ec.Scope
	// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
	AwakeOnFirstAccess() *bool
	// CountComponents // 组件数量
	CountComponents() int
	// Component 获取组件
	Component(idx int) ComponentDesc
	// Components 获取所有组件
	Components() []ComponentDesc
	// Construct 创建实体
	Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity
}

type _EntityPT struct {
	prototype          string
	instanceRT         reflect.Type
	scope              *ec.Scope
	awakeOnFirstAccess *bool
	components         []ComponentDesc
}

// Prototype 实体原型名称
func (pt *_EntityPT) Prototype() string {
	return pt.prototype
}

// InstanceRT 实体实例反射类型
func (pt *_EntityPT) InstanceRT() reflect.Type {
	return pt.instanceRT
}

// Scope 可访问作用域
func (pt *_EntityPT) Scope() *ec.Scope {
	return pt.scope
}

// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
func (pt *_EntityPT) AwakeOnFirstAccess() *bool {
	return pt.awakeOnFirstAccess
}

// CountComponents // 组件数量
func (pt *_EntityPT) CountComponents() int {
	return len(pt.components)
}

// Component 获取组件
func (pt *_EntityPT) Component(idx int) ComponentDesc {
	if idx < 0 || idx >= len(pt.components) {
		panic(fmt.Errorf("%w: %w: idx out of range", ErrPt, exception.ErrArgs))
	}
	return pt.components[idx]
}

// Components 获取所有组件
func (pt *_EntityPT) Components() []ComponentDesc {
	return slices.Clone(pt.components)
}

// Construct 创建实体
func (pt *_EntityPT) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.Make(ec.With.Default())
	if pt.instanceRT != nil {
		options.InstanceFace = iface.MakeFaceT(reflect.New(pt.instanceRT).Interface().(ec.Entity))
	}
	if pt.scope != nil {
		options.Scope = *pt.scope
	}
	if pt.awakeOnFirstAccess != nil {
		options.AwakeOnFirstAccess = *pt.awakeOnFirstAccess
	}
	options = option.Append(options, settings...)
	options.Prototype = pt.prototype

	return pt.assemble(ec.UnsafeNewEntity(options))
}

func (pt *_EntityPT) assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}

	for i := range pt.components {
		compDesc := &pt.components[i]

		comp := compDesc.PT.Construct()
		ec.UnsafeComponent(comp).SetFixed(compDesc.Fixed)

		if err := entity.AddComponent(compDesc.Alias, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}
	}

	return entity
}
