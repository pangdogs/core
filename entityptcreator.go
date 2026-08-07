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

// BuildEntityPT 创建绑定到 svcCtx 实体原型库的构建器。
func BuildEntityPT(svcCtx service.Context, prototype string) *EntityPTCreator {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrCore, ErrArgs)
	}
	return &EntityPTCreator{
		svcCtx: svcCtx,
		descr:  pt.NewEntityDescriptor(prototype),
	}
}

// EntityPTCreator 以链式方式配置并声明实体原型。
type EntityPTCreator struct {
	svcCtx service.Context
	descr  *pt.EntityDescriptor
	comps  []any
}

// SetInstance 设置该原型用于构造自定义实体的实例或反射类型。
func (c *EntityPTCreator) SetInstance(instance any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetInstance(instance)
	return c
}

// SetScope 设置由该原型创建的实体可访问作用域。
func (c *EntityPTCreator) SetScope(scope ec.Scope) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetScope(scope)
	return c
}

// SetComponentAwakeOnFirstTouch 设置组件是否延迟到首次访问时进入 Awakened 状态。
func (c *EntityPTCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetComponentAwakeOnFirstTouch(b)
	return c
}

// SetComponentUniqueID 设置是否为每个组件分配独立 ID。
func (c *EntityPTCreator) SetComponentUniqueID(b bool) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetComponentUniqueID(b)
	return c
}

// SetMeta 用 dict 替换原型元数据。
func (c *EntityPTCreator) SetMeta(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.SetMeta(dict)
	return c
}

// MergeMeta 合并原型元数据；同名键会被覆盖。
func (c *EntityPTCreator) MergeMeta(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.MergeMeta(dict)
	return c
}

// MergeMetaIfAbsent 合并原型元数据；已有的同名键保持不变。
func (c *EntityPTCreator) MergeMetaIfAbsent(dict map[string]any) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.MergeIfAbsent(dict)
	return c
}

// AssignMeta 直接采用 m 作为原型元数据。
func (c *EntityPTCreator) AssignMeta(m meta.Meta) *EntityPTCreator {
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.descr.AssignMeta(m)
	return c
}

// AddComponent 向原型追加一个内建组件。
// comp 可以是组件实例或 ComponentDescriptor；未指定名称时使用组件类型名。
func (c *EntityPTCreator) AddComponent(comp any, name ...string) *EntityPTCreator {
	switch v := comp.(type) {
	case pt.ComponentDescriptor, *pt.ComponentDescriptor:
		c.comps = append(c.comps, v)
	default:
		c.comps = append(c.comps, pt.NewComponentDescriptor(comp).SetName(pie.First(name)))
	}
	return c
}

// Declare 将构建结果注册到服务的实体原型库。
func (c *EntityPTCreator) Declare() {
	if c.svcCtx == nil {
		exception.Panicf("%w: svcCtx is nil", ErrCore)
	}
	if c.descr == nil {
		exception.Panicf("%w: descr is nil", ErrCore)
	}
	c.svcCtx.EntityLib().Declare(c.descr, c.comps...)
}
