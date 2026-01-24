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
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

// BuildEntity 创建实体
func BuildEntity(provider corectx.CurrentContextProvider, prototype string) *EntityCreator {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrCore, ErrArgs)
	}
	return &EntityCreator{
		rtCtx:     runtime.Current(provider),
		prototype: prototype,
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx     runtime.Context
	prototype string
	parentId  uid.Id
	meta      meta.Meta
	settings  []option.Setting[ec.EntityOptions]
}

// SetInstanceFace 设置实例，用于扩展实体能力
func (c *EntityCreator) SetInstanceFace(face iface.Face[ec.Entity]) *EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(face))
	return c
}

// SetInstance 设置实例，用于扩展实体能力
func (c *EntityCreator) SetInstance(instance ec.Entity) *EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(iface.MakeFaceT(instance)))
	return c
}

// SetScope 设置实体的可访问作用域
func (c *EntityCreator) SetScope(scope ec.Scope) *EntityCreator {
	c.settings = append(c.settings, ec.With.Scope(scope))
	return c
}

// SetPersistId 设置实体持久化Id
func (c *EntityCreator) SetPersistId(id uid.Id) *EntityCreator {
	c.settings = append(c.settings, ec.With.PersistId(id))
	return c
}

// SetComponentAwakeOnFirstTouch 设置当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (c *EntityCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentAwakeOnFirstTouch(b))
	return c
}

// SetComponentUniqueID 设置是否为实体组件分配唯一Id
func (c *EntityCreator) SetComponentUniqueID(b bool) *EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentUniqueID(b))
	return c
}

// SetMeta 设置Meta信息
func (c *EntityCreator) SetMeta(dict map[string]any) *EntityCreator {
	if c.meta == nil {
		c.settings = append(c.settings, c.withMeta())
	}
	c.meta = meta.M(dict)
	return c
}

// MergeMeta 合并Meta信息，如果存在则覆盖
func (c *EntityCreator) MergeMeta(dict map[string]any) *EntityCreator {
	for k, v := range dict {
		if c.meta == nil {
			c.settings = append(c.settings, c.withMeta())
		}
		c.meta.Add(k, v)
	}
	return c
}

// MergeMetaIfAbsent 合并Meta信息，如果存在则跳过
func (c *EntityCreator) MergeMetaIfAbsent(dict map[string]any) *EntityCreator {
	for k, v := range dict {
		if c.meta == nil {
			c.settings = append(c.settings, c.withMeta())
		}
		c.meta.TryAdd(k, v)
	}
	return c
}

// AssignMeta 赋值Meta信息
func (c *EntityCreator) AssignMeta(m meta.Meta) *EntityCreator {
	if m == nil {
		m = meta.M(nil)
	}
	if c.meta == nil {
		c.settings = append(c.settings, c.withMeta())
	}
	c.meta = m
	return c
}

// SetParentId 设置父实体Id
func (c *EntityCreator) SetParentId(id uid.Id) *EntityCreator {
	c.parentId = id
	return c
}

// New 创建实体
func (c *EntityCreator) New() (ec.Entity, error) {
	if c.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	entity := pt.For(service.Current(c.rtCtx), c.prototype).Construct(c.settings...)

	if err := c.rtCtx.GetEntityManager().AddEntity(entity); err != nil {
		return nil, err
	}

	if !c.parentId.IsNil() {
		if err := c.rtCtx.GetEntityTree().AddChild(c.parentId, entity.GetId()); err != nil {
			entity.Destroy()
			return nil, err
		}
	}

	return entity, nil
}

func (c *EntityCreator) withMeta() option.Setting[ec.EntityOptions] {
	return func(o *ec.EntityOptions) {
		o.Meta = c.meta
	}
}
