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
		atti:   pt.NewEntityAttribute(prototype),
	}
}

// EntityPTCreator 实体原型构建器
type EntityPTCreator struct {
	svcCtx service.Context
	atti   *pt.EntityAttribute
	comps  []any
}

// SetInstance 设置实例，用于扩展实体能力
func (c *EntityPTCreator) SetInstance(instance any) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetInstance(instance)
	return c
}

// SetScope 设置实体的可访问作用域
func (c *EntityPTCreator) SetScope(scope ec.Scope) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetScope(scope)
	return c
}

// SetComponentNameIndexing 设置是否开启组件名称索引
func (c *EntityPTCreator) SetComponentNameIndexing(b bool) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetComponentNameIndexing(b)
	return c
}

// SetComponentAwakeOnFirstTouch 设置当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (c *EntityPTCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetComponentAwakeOnFirstTouch(b)
	return c
}

// SetComponentUniqueID 设置是否为实体组件分配唯一Id
func (c *EntityPTCreator) SetComponentUniqueID(b bool) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetComponentUniqueID(b)
	return c
}

// SetExtra 设置自定义属性
func (c *EntityPTCreator) SetExtra(dict map[string]any) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.SetExtra(dict)
	return c
}

// MergeExtra 合并自定义属性，如果存在则覆盖
func (c *EntityPTCreator) MergeExtra(dict map[string]any) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.MergeExtra(dict)
	return c
}

// MergeExtraIfAbsent 合并自定义属性，如果存在则跳过
func (c *EntityPTCreator) MergeExtraIfAbsent(dict map[string]any) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.MergeIfAbsent(dict)
	return c
}

// AssignExtra 赋值自定义属性
func (c *EntityPTCreator) AssignExtra(m meta.Meta) *EntityPTCreator {
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.atti.AssignExtra(m)
	return c
}

// AddComponent 添加组件
func (c *EntityPTCreator) AddComponent(comp any, name ...string) *EntityPTCreator {
	switch v := comp.(type) {
	case pt.ComponentAttribute, *pt.ComponentAttribute:
		c.comps = append(c.comps, v)
	default:
		c.comps = append(c.comps, pt.NewComponentAttribute(comp).SetName(pie.First(name)))
	}
	return c
}

// Declare 声明实体原型
func (c *EntityPTCreator) Declare() {
	if c.svcCtx == nil {
		exception.Panicf("%w: svcCtx is nil", ErrCore)
	}
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.svcCtx.GetEntityLib().Declare(c.atti, c.comps...)
}

// Redeclare 重新声明实体原型
func (c *EntityPTCreator) Redeclare() {
	if c.svcCtx == nil {
		exception.Panicf("%w: svcCtx is nil", ErrCore)
	}
	if c.atti == nil {
		exception.Panicf("%w: atti is nil", ErrCore)
	}
	c.svcCtx.GetEntityLib().Redeclare(c.atti, c.comps...)
}
