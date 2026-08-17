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

// BuildEntity 创建绑定到 provider 当前运行时的实体构建器。
// prototype 必须已经在所属服务的实体原型库中声明。
func BuildEntity(provider corectx.CurrentContextProvider, prototype string) *EntityCreator {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrCore, ErrArgs)
	}
	return &EntityCreator{
		rtCtx:     runtime.Current(provider),
		prototype: prototype,
	}
}

// EntityCreator 基于已声明原型配置并创建实体。
type EntityCreator struct {
	rtCtx     runtime.Context
	prototype string
	meta      meta.Meta
	settings  []option.Setting[ec.EntityOptions]
}

// SetInstanceFace 设置用于扩展实体能力的自定义实例及其接口缓存。
func (c *EntityCreator) SetInstanceFace(face iface.Face[ec.Entity]) *EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(face))
	return c
}

// SetInstance 设置用于扩展实体能力的自定义实例。
func (c *EntityCreator) SetInstance(instance ec.Entity) *EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(iface.NewFaceT(instance)))
	return c
}

// SetScope 覆盖原型定义的实体可访问作用域。
func (c *EntityCreator) SetScope(scope ec.Scope) *EntityCreator {
	c.settings = append(c.settings, ec.With.Scope(scope))
	return c
}

// SetPersistId 设置实体的持久化 ID。
func (c *EntityCreator) SetPersistId(id uid.Id) *EntityCreator {
	c.settings = append(c.settings, ec.With.PersistId(id))
	return c
}

// SetComponentAwakeOnFirstTouch 设置正常激活期间被访问的组件是否优先执行 Awake。
func (c *EntityCreator) SetComponentAwakeOnFirstTouch(b bool) *EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentAwakeOnFirstTouch(b))
	return c
}

// SetComponentUniqueID 设置是否为每个组件分配独立 ID。
func (c *EntityCreator) SetComponentUniqueID(b bool) *EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentUniqueID(b))
	return c
}

// SetMeta 用 dict 替换待创建实体的元数据。
func (c *EntityCreator) SetMeta(dict map[string]any) *EntityCreator {
	if c.meta == nil {
		c.settings = append(c.settings, c.withMeta())
	}
	c.meta = meta.New(dict)
	return c
}

// MergeMeta 合并元数据；同名键会被覆盖。
func (c *EntityCreator) MergeMeta(dict map[string]any) *EntityCreator {
	for k, v := range dict {
		if c.meta == nil {
			c.settings = append(c.settings, c.withMeta())
		}
		c.meta.Add(k, v)
	}
	return c
}

// MergeMetaIfAbsent 合并元数据；已有的同名键保持不变。
func (c *EntityCreator) MergeMetaIfAbsent(dict map[string]any) *EntityCreator {
	for k, v := range dict {
		if c.meta == nil {
			c.settings = append(c.settings, c.withMeta())
		}
		c.meta.TryAdd(k, v)
	}
	return c
}

// AssignMeta 直接采用 m 作为待创建实体的元数据。
func (c *EntityCreator) AssignMeta(m meta.Meta) *EntityCreator {
	if m == nil {
		m = meta.New(nil)
	}
	if c.meta == nil {
		c.settings = append(c.settings, c.withMeta())
	}
	c.meta = m
	return c
}

// New 根据原型构造实体，并将其加入绑定运行时的实体管理器。
func (c *EntityCreator) New() (ec.Entity, error) {
	if c.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	entity := pt.For(service.Current(c.rtCtx), c.prototype).Construct(c.settings...)

	if err := c.rtCtx.EntityManager().AddEntity(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (c *EntityCreator) withMeta() option.Setting[ec.EntityOptions] {
	return func(o *ec.EntityOptions) {
		o.Meta = c.meta
	}
}
