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

// NewEntityDescriptor 创建实体原型描述，用于注册实体原型
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

// EntityDescriptor 实体原型描述
type EntityDescriptor struct {
	Prototype                  string    // 实体原型名称（必填）
	Instance                   any       // 实体实例
	Scope                      ec.Scope  // 可访问作用域
	ComponentAwakeOnFirstTouch bool      // 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentUniqueID          bool      // 是否为实体组件分配唯一Id
	Meta                       meta.Meta // 原型Meta信息
}

func (descr *EntityDescriptor) SetInstance(instance any) *EntityDescriptor {
	descr.Instance = instance
	return descr
}

func (descr *EntityDescriptor) SetScope(scope ec.Scope) *EntityDescriptor {
	descr.Scope = scope
	return descr
}

func (descr *EntityDescriptor) SetComponentAwakeOnFirstTouch(b bool) *EntityDescriptor {
	descr.ComponentAwakeOnFirstTouch = b
	return descr
}

func (descr *EntityDescriptor) SetComponentUniqueID(b bool) *EntityDescriptor {
	descr.ComponentUniqueID = b
	return descr
}

func (descr *EntityDescriptor) SetMeta(dict map[string]any) *EntityDescriptor {
	descr.Meta = meta.M(dict)
	return descr
}

func (descr *EntityDescriptor) MergeMeta(dict map[string]any) *EntityDescriptor {
	for k, v := range dict {
		descr.Meta.Add(k, v)
	}
	return descr
}

func (descr *EntityDescriptor) MergeIfAbsent(dict map[string]any) *EntityDescriptor {
	for k, v := range dict {
		descr.Meta.TryAdd(k, v)
	}
	return descr
}

func (descr *EntityDescriptor) AssignMeta(m meta.Meta) *EntityDescriptor {
	descr.Meta = m
	return descr
}
