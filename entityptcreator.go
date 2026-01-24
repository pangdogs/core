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

package core

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/meta"
	"github.com/elliotchance/pie/v2"
)

// BuildEntityPT 创建实体原型
func BuildEntityPT(svcCtx service.Context, prototype string) *EntityPTCreator {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrCore, ErrArgs)
	}
	return &EntityPTCreator{
		svcCtx: svcCtx,
		descr:  pt.NewEntityDescriptor(prototype),
	}
}

// EntityPTCreator 实体原型构建器
type EntityPTCreator struct {
	svcCtx service.Context
	descr  *pt.EntityDescriptor
	comps  []any
}

// SetInstance 设置实例，用于扩展实体能力
func (c *EntityPTCreator) SetInstance(instance any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetInstance(instance)
	return c
}

// SetScope 设置实体的可访问作用域
func (c *EntityPTCreator) SetScope(scope ec.Scope) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetScope(scope)
	return c
}

// SetComponentAwakeOnFirstTouch 设置当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (c *EntityPTCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetComponentAwakeOnFirstTouch(b)
	return c
}

// SetComponentUniqueID 设置是否为实体组件分配唯一Id
func (c *EntityPTCreator) SetComponentUniqueID(b bool) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetComponentUniqueID(b)
	return c
}

// SetMeta 设置原型Meta信息
func (c *EntityPTCreator) SetMeta(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetMeta(dict)
	return c
}

// MergeMeta 合并原型Meta信息，如果存在则覆盖
func (c *EntityPTCreator) MergeMeta(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.MergeMeta(dict)
	return c
}

// MergeMetaIfAbsent 合并原型Meta信息，如果存在则跳过
func (c *EntityPTCreator) MergeMetaIfAbsent(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.MergeIfAbsent(dict)
	return c
}

// AssignMeta 赋值原型Meta信息
func (c *EntityPTCreator) AssignMeta(m meta.Meta) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.AssignMeta(m)
	return c
}

// AddComponent 添加组件
func (c *EntityPTCreator) AddComponent(comp any, name ...string) *EntityPTCreator {
	switch v := comp.(type) {
	case pt.ComponentDescriptor, *pt.ComponentDescriptor:
		c.comps = append(c.comps, v)
	default:
		c.comps = append(c.comps, pt.NewComponentDescriptor(comp).SetName(pie.First(name)))
	}
	return c
}

// Declare 声明实体原型
func (c *EntityPTCreator) Declare() {
	if c.svcCtx == nil {
		exception.Panicf("%w: svcCtx is nil", ErrCore)
	}
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.svcCtx.GetEntityLib().Declare(c.descr, c.comps...)
}
