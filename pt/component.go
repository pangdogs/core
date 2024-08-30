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

// ComponentPT 组件原型
type ComponentPT struct {
	Name       string       // 组件名称
	InstanceRT reflect.Type // 实例反射类型
}

// Construct 创建组件
func (pt ComponentPT) Construct() ec.Component {
	vfComp := reflect.New(pt.InstanceRT)

	comp := vfComp.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflected(vfComp)

	return comp
}
