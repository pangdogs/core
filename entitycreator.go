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
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

// CreateEntity 创建实体
func CreateEntity(provider ictx.CurrentContextProvider, prototype string) EntityCreator {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrCore, ErrArgs)
	}
	return EntityCreator{
		rtCtx:     runtime.Current(provider),
		prototype: prototype,
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx     runtime.Context
	prototype string
	parentId  uid.Id
	settings  []option.Setting[ec.EntityOptions]
}

// InstanceFace 设置实例，用于扩展实体能力
func (c EntityCreator) InstanceFace(face iface.Face[ec.Entity]) EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(face))
	return c
}

// Instance 设置实例，用于扩展实体能力
func (c EntityCreator) Instance(instance ec.Entity) EntityCreator {
	c.settings = append(c.settings, ec.With.InstanceFace(iface.MakeFaceT(instance)))
	return c
}

// Scope 设置实体的可访问作用域
func (c EntityCreator) Scope(scope ec.Scope) EntityCreator {
	c.settings = append(c.settings, ec.With.Scope(scope))
	return c
}

// PersistId 设置实体持久化Id
func (c EntityCreator) PersistId(id uid.Id) EntityCreator {
	c.settings = append(c.settings, ec.With.PersistId(id))
	return c
}

// ComponentAwakeOnFirstTouch 开启组件被首次访问时，检测并调用Awake()
func (c EntityCreator) ComponentAwakeOnFirstTouch(b bool) EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentAwakeOnFirstTouch(b))
	return c
}

// ComponentUniqueID 开启组件唯一Id
func (c EntityCreator) ComponentUniqueID(b bool) EntityCreator {
	c.settings = append(c.settings, ec.With.ComponentUniqueID(b))
	return c
}

// Meta 设置Meta信息
func (c EntityCreator) Meta(m meta.Meta) EntityCreator {
	c.settings = append(c.settings, ec.With.Meta(m))
	return c
}

// ParentId 设置父实体Id
func (c EntityCreator) ParentId(id uid.Id) EntityCreator {
	c.parentId = id
	return c
}

// Spawn 创建实体
func (c EntityCreator) Spawn() (ec.Entity, error) {
	if c.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	entity := pt.For(service.Current(c.rtCtx), c.prototype).Construct(c.settings...)

	if c.parentId.IsNil() {
		if err := c.rtCtx.GetEntityManager().AddEntity(entity); err != nil {
			return nil, err
		}
	} else {
		if err := c.rtCtx.GetEntityTree().AddNode(entity, c.parentId); err != nil {
			return nil, err
		}
	}

	return entity, nil
}
