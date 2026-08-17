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

// UnsafeComponent 暴露组件状态机与运行时句柄等框架内部能力。
//
// Deprecated: 仅供框架内部使用。
func UnsafeComponent(comp Component) _UnsafeComponent {
	return _UnsafeComponent{
		Component: comp,
	}
}

type _UnsafeComponent struct {
	Component
}

// Instance 返回实际组件实例。
func (u _UnsafeComponent) Instance() Component {
	return u.getInstance()
}

// SetId 设置组件 ID。
func (u _UnsafeComponent) SetId(id uid.Id) {
	u.setId(id)
}

// SetState 推进组件生命周期状态。
func (u _UnsafeComponent) SetState(state ComponentState) {
	u.setState(state)
}

// SetReflected 缓存实际组件实例的反射值。
func (u _UnsafeComponent) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// SetBuiltin 绑定组件在实体原型中的内建描述。
func (u _UnsafeComponent) SetBuiltin(builtin *BuiltinComponent) {
	u.setBuiltin(builtin)
}

// SetRemovable 设置组件是否允许动态删除。
func (u _UnsafeComponent) SetRemovable(b bool) {
	u.setRemovable(b)
}

// ProcessedStateBits 返回生命周期阶段的已处理标志位。
func (u _UnsafeComponent) ProcessedStateBits() *generic.Bits16 {
	return u.getProcessedStateBits()
}

// AttachedHandle 返回组件在实体组件表中的位置与版本。
func (u _UnsafeComponent) AttachedHandle() (int, int64) {
	return u.getAttachedHandle()
}

// ManagedRuntimeUpdateHandle 替换并托管 Runtime 更新事件句柄。
func (u _UnsafeComponent) ManagedRuntimeUpdateHandle(updateHandle event.Handle) {
	u.managedRuntimeUpdateHandle(updateHandle)
}

// ManagedRuntimeLateUpdateHandle 替换并托管 Runtime 后置更新事件句柄。
func (u _UnsafeComponent) ManagedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	u.managedRuntimeLateUpdateHandle(lateUpdateHandle)
}

// ManagedUnbindRuntimeHandles 解绑全部托管的 Runtime 更新事件句柄。
func (u _UnsafeComponent) ManagedUnbindRuntimeHandles() {
	u.managedUnbindRuntimeHandles()
}
