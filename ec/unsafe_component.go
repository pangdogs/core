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

package ec

import (
	"reflect"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
)

// Deprecated: UnsafeComponent 访问组件内部函数
func UnsafeComponent(comp Component) _UnsafeComponent {
	return _UnsafeComponent{
		Component: comp,
	}
}

type _UnsafeComponent struct {
	Component
}

// SetId 设置Id
func (u _UnsafeComponent) SetId(id uid.Id) {
	u.setId(id)
}

// SetState 设置状态
func (u _UnsafeComponent) SetState(state ComponentState) {
	u.setState(state)
}

// SetReflected 设置反射值
func (u _UnsafeComponent) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// SetBuiltin 设置实体原型中的组件信息
func (u _UnsafeComponent) SetBuiltin(builtin *BuiltinComponent) {
	u.setBuiltin(builtin)
}

// SetRemovable 设置是否可以删除
func (u _UnsafeComponent) SetRemovable(b bool) {
	u.setRemovable(b)
}

// GetProcessedStateBits 获取已处理状态标志位
func (u _UnsafeComponent) GetProcessedStateBits() *generic.Bits16 {
	return u.getProcessedStateBits()
}

// GetAttachedHandle 获取加入实体时的句柄
func (u _UnsafeComponent) GetAttachedHandle() (int, int64) {
	return u.getAttachedHandle()
}

// ManagedRuntimeUpdateHandle 托管运行时更新句柄
func (u _UnsafeComponent) ManagedRuntimeUpdateHandle(updateHandle event.Handle) {
	u.managedRuntimeUpdateHandle(updateHandle)
}

// ManagedRuntimeLateUpdateHandle 托管运行时延迟更新句柄
func (u _UnsafeComponent) ManagedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	u.managedRuntimeLateUpdateHandle(lateUpdateHandle)
}

// ManagedUnbindRuntimeHandles 解绑定托管的运行时句柄
func (u _UnsafeComponent) ManagedUnbindRuntimeHandles() {
	u.managedUnbindRuntimeHandles()
}
