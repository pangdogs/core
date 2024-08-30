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
)

// ComponentDesc 组件描述
type ComponentDesc struct {
	PT    ComponentPT // 原型
	Alias string      // 别名
	Fixed bool        // 固定
}

// EntityPT 实体原型
type EntityPT struct {
	Prototype          string          // 实体原型名称
	InstanceRT         reflect.Type    // 实例反射类型
	Scope              *ec.Scope       // 可访问作用域
	AwakeOnFirstAccess *bool           // 设置开启组件被首次访问时，检测并调用Awake()
	Components         []ComponentDesc // 组件信息
}

// Construct 创建实体
func (pt EntityPT) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.Make(ec.With.Default())
	if pt.InstanceRT != nil {
		options.InstanceFace = iface.MakeFaceT(reflect.New(pt.InstanceRT).Interface().(ec.Entity))
	}
	if pt.Scope != nil {
		options.Scope = *pt.Scope
	}
	if pt.AwakeOnFirstAccess != nil {
		options.AwakeOnFirstAccess = *pt.AwakeOnFirstAccess
	}
	options = option.Append(options, settings...)
	options.Prototype = pt.Prototype

	return pt.assemble(ec.UnsafeNewEntity(options))
}

func (pt EntityPT) assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}

	for i := range pt.Components {
		compInfo := &pt.Components[i]

		comp := compInfo.PT.Construct()

		ec.UnsafeComponent(comp).SetFixed(compInfo.Fixed)

		if err := entity.AddComponent(compInfo.Alias, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}
	}

	return entity
}
