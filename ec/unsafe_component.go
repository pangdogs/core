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
	"context"
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
func (u _UnsafeComponent) Init(name string, entity Entity, instance Component) {
	u.init(name, entity, instance)
}

// WithContext 传递上下文
func (u _UnsafeComponent) WithContext(ctx context.Context) {
	u.withContext(ctx)
}

// SetId 设置Id
func (u _UnsafeComponent) SetId(id uid.Id) {
	u.setId(id)
}

// SetBuiltin 设置实体原型中的组件信息
func (u _UnsafeComponent) SetBuiltin(builtin *BuiltinComponent) {
	u.setBuiltin(builtin)
}

// SetState 设置状态
func (u _UnsafeComponent) SetState(state ComponentState) {
	u.setState(state)
}

// SetReflected 设置反射值
func (u _UnsafeComponent) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// SetNonRemovable 设置是否不可删除
func (u _UnsafeComponent) SetNonRemovable(b bool) {
	u.setNonRemovable(b)
}

// CleanManagedHooks 清理所有的托管hook
func (u _UnsafeComponent) CleanManagedHooks() {
	u.cleanManagedHooks()
}
