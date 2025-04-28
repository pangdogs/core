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

// BuildEntityAttribute 创建实体原型属性，用于注册实体原型时自定义相关属性
func BuildEntityAttribute(prototype string) *EntityAttributeCreator {
	if prototype == "" {
		exception.Panicf("%w: %w: prototype is empty", ErrPt, exception.ErrArgs)
	}
	return &EntityAttributeCreator{
		atti: EntityAttribute{
			Prototype: prototype,
		},
	}
}

type EntityAttributeCreator struct {
	atti EntityAttribute
}

func (c *EntityAttributeCreator) SetInstance(instance any) *EntityAttributeCreator {
	c.atti.Instance = instance
	return c
}

func (c *EntityAttributeCreator) SetScope(scope ec.Scope) *EntityAttributeCreator {
	c.atti.Scope = &scope
	return c
}

func (c *EntityAttributeCreator) SetComponentNameIndexing(b bool) *EntityAttributeCreator {
	c.atti.ComponentNameIndexing = &b
	return c
}

func (c *EntityAttributeCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityAttributeCreator {
	c.atti.ComponentAwakeOnFirstTouch = &b
	return c
}

func (c *EntityAttributeCreator) SetComponentUniqueID(b bool) *EntityAttributeCreator {
	c.atti.ComponentUniqueID = &b
	return c
}

func (c *EntityAttributeCreator) SetExtra(dict map[string]any) *EntityAttributeCreator {
	c.atti.Extra = meta.M(dict)
	return c
}

func (c *EntityAttributeCreator) MergeExtra(dict map[string]any) *EntityAttributeCreator {
	for k, v := range dict {
		c.atti.Extra.Add(k, v)
	}
	return c
}

func (c *EntityAttributeCreator) MergeIfAbsent(dict map[string]any) *EntityAttributeCreator {
	for k, v := range dict {
		c.atti.Extra.TryAdd(k, v)
	}
	return c
}

func (c *EntityAttributeCreator) AssignExtra(m meta.Meta) *EntityAttributeCreator {
	c.atti.Extra = m
	return c
}

func (c *EntityAttributeCreator) Get() *EntityAttribute {
	return &c.atti
}
