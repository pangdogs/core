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

// BuildEntityAttribute 创建实体原型属性，用于注册实体原型时自定义相关属性
func BuildEntityAttribute(prototype string) *EntityAttribute {
	if prototype == "" {
		exception.Panicf("%w: %w: prototype is empty", ErrPt, exception.ErrArgs)
	}
	return &EntityAttribute{
		Prototype: prototype,
	}
}

// EntityAttribute 实体原型属性
type EntityAttribute struct {
	Prototype                  string    // 实体原型名称（必填）
	Instance                   any       // 实体实例
	Scope                      *ec.Scope // 可访问作用域
	ComponentNameIndexing      *bool     // 是否开启组件名称索引
	ComponentAwakeOnFirstTouch *bool     // 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentUniqueID          *bool     // 是否为实体组件分配唯一Id
	Extra                      meta.Meta // 自定义属性
}

func (atti *EntityAttribute) SetInstance(instance any) *EntityAttribute {
	atti.Instance = instance
	return atti
}

func (atti *EntityAttribute) SetScope(scope ec.Scope) *EntityAttribute {
	atti.Scope = &scope
	return atti
}

func (atti *EntityAttribute) SetComponentNameIndexing(b bool) *EntityAttribute {
	atti.ComponentNameIndexing = &b
	return atti
}

func (atti *EntityAttribute) SetComponentAwakeOnFirstTouch(b bool) *EntityAttribute {
	atti.ComponentAwakeOnFirstTouch = &b
	return atti
}

func (atti *EntityAttribute) SetComponentUniqueID(b bool) *EntityAttribute {
	atti.ComponentUniqueID = &b
	return atti
}

func (atti *EntityAttribute) SetExtra(dict map[string]any) *EntityAttribute {
	atti.Extra = meta.M(dict)
	return atti
}

func (atti *EntityAttribute) MergeExtra(dict map[string]any) *EntityAttribute {
	for k, v := range dict {
		atti.Extra.Add(k, v)
	}
	return atti
}

func (atti *EntityAttribute) MergeIfAbsent(dict map[string]any) *EntityAttribute {
	for k, v := range dict {
		atti.Extra.TryAdd(k, v)
	}
	return atti
}

func (atti *EntityAttribute) AssignExtra(m meta.Meta) *EntityAttribute {
	atti.Extra = m
	return atti
}
