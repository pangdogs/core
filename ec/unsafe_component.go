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
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/uid"
	"reflect"
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

// Init 初始化
func (uc _UnsafeComponent) Init(name string, entity Entity, composite Component) {
	uc.init(name, entity, composite)
}

// SetId 设置Id
func (uc _UnsafeComponent) SetId(id uid.Id) {
	uc.setId(id)
}

// SetState 设置状态
func (uc _UnsafeComponent) SetState(state ComponentState) {
	uc.setState(state)
}

// SetReflected 设置反射值
func (uc _UnsafeComponent) SetReflected(v reflect.Value) {
	uc.setReflected(v)
}

// SetFixed 设置是否固定
func (uc _UnsafeComponent) SetFixed(b bool) {
	uc.setFixed(b)
}

// GetComposite 获取扩展者
func (uc _UnsafeComponent) GetComposite() Component {
	return uc.getComposite()
}

// EventComponentDestroySelf 事件：组件销毁自身
func (uc _UnsafeComponent) EventComponentDestroySelf() event.IEvent {
	return uc.eventComponentDestroySelf()
}

// CleanManagedHooks 清理所有的托管hook
func (uc _UnsafeComponent) CleanManagedHooks() {
	uc.cleanManagedHooks()
}
