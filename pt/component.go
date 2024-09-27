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
	"reflect"
)

// ComponentPT 组件原型接口
type ComponentPT interface {
	// Name 组件名称
	Name() string
	// InstanceRT 组件实例反射类型
	InstanceRT() reflect.Type
	// Construct 创建组件
	Construct() ec.Component
}

type _ComponentPT struct {
	name       string       // 组件名称
	instanceRT reflect.Type // 实例反射类型
}

// Name 组件名称
func (pt *_ComponentPT) Name() string {
	return pt.name
}

// InstanceRT 组件实例反射类型
func (pt *_ComponentPT) InstanceRT() reflect.Type {
	return pt.instanceRT
}

// Construct 创建组件
func (pt *_ComponentPT) Construct() ec.Component {
	compRV := reflect.New(pt.instanceRT)

	comp := compRV.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflected(compRV)

	return comp
}
