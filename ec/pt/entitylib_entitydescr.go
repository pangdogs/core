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
	"git.golaxy.org/core/utils/meta"
)

// NewEntityDescriptor 创建实体原型描述；prototype 为空时 panic。
func NewEntityDescriptor(prototype string) *EntityDescriptor {
	if prototype == "" {
		exception.Panicf("%w: %w: prototype is empty", ErrPt, exception.ErrArgs)
	}
	return &EntityDescriptor{
		Prototype:                  prototype,
		Instance:                   nil,
		Scope:                      ec.Scope_Global,
		ComponentAwakeOnFirstTouch: false,
		ComponentUniqueID:          false,
		Meta:                       nil,
	}
}

// EntityDescriptor 描述一个可注册的实体原型。
type EntityDescriptor struct {
	Prototype                  string    // Prototype 是实体原型名，不能为空。
	Instance                   any       // Instance 是自定义实体值或反射类型；nil 表示使用默认实体实现。
	Scope                      ec.Scope  // Scope 是构造实体时使用的默认作用域。
	ComponentAwakeOnFirstTouch bool      // ComponentAwakeOnFirstTouch 指示正常激活期间被访问的组件是否优先执行 Awake。
	ComponentUniqueID          bool      // ComponentUniqueID 指示是否为每个组件分配唯一 ID。
	Meta                       meta.Meta // Meta 是实体原型元数据。
}

// SetInstance 设置自定义实体实例类型并返回 descr，以便链式调用。
func (descr *EntityDescriptor) SetInstance(instance any) *EntityDescriptor {
	descr.Instance = instance
	return descr
}

// SetScope 设置默认实体作用域并返回 descr。
func (descr *EntityDescriptor) SetScope(scope ec.Scope) *EntityDescriptor {
	descr.Scope = scope
	return descr
}

// SetComponentAwakeOnFirstTouch 设置正常激活期间被访问的组件是否优先执行 Awake。
func (descr *EntityDescriptor) SetComponentAwakeOnFirstTouch(b bool) *EntityDescriptor {
	descr.ComponentAwakeOnFirstTouch = b
	return descr
}

// SetComponentUniqueID 设置是否为每个组件分配唯一 ID。
func (descr *EntityDescriptor) SetComponentUniqueID(b bool) *EntityDescriptor {
	descr.ComponentUniqueID = b
	return descr
}

// SetMeta 使用 dict 的副本替换元数据并返回 descr。
func (descr *EntityDescriptor) SetMeta(dict map[string]any) *EntityDescriptor {
	descr.Meta = meta.New(dict)
	return descr
}

// MergeMeta 合并 dict；同名键会覆盖原值。
func (descr *EntityDescriptor) MergeMeta(dict map[string]any) *EntityDescriptor {
	for k, v := range dict {
		descr.Meta.Add(k, v)
	}
	return descr
}

// MergeIfAbsent 合并 dict，但保留已有的同名键。
func (descr *EntityDescriptor) MergeIfAbsent(dict map[string]any) *EntityDescriptor {
	for k, v := range dict {
		descr.Meta.TryAdd(k, v)
	}
	return descr
}

// AssignMeta 直接绑定 m 并返回 descr；m 不会被复制。
func (descr *EntityDescriptor) AssignMeta(m meta.Meta) *EntityDescriptor {
	descr.Meta = m
	return descr
}
