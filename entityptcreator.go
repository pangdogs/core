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
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/service"
)

// CreateEntityPT 创建实体原型
func CreateEntityPT(svcCtx service.Context, prototype string) EntityPTCreator {
	if svcCtx == nil {
		panic(fmt.Errorf("%w: %w: svcCtx is nil", ErrCore, ErrArgs))
	}
	return EntityPTCreator{
		svcCtx:    svcCtx,
		prototype: prototype,
	}
}

// EntityPTCreator 实体原型构建器
type EntityPTCreator struct {
	svcCtx    service.Context
	prototype string
	atti      pt.Atti
	comps     []any
}

// Instance 设置实例，用于扩展实体能力
func (c EntityPTCreator) Instance(instance any) EntityPTCreator {
	c.atti.Instance = instance
	return c
}

// Scope 设置实体的可访问作用域
func (c EntityPTCreator) Scope(scope ec.Scope) EntityPTCreator {
	c.atti.Scope = &scope
	return c
}

// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
func (c EntityPTCreator) AwakeOnFirstAccess(b bool) EntityPTCreator {
	c.atti.AwakeOnFirstAccess = &b
	return c
}

// AddComponent 添加组件
func (c EntityPTCreator) AddComponent(comp any, alias ...string) EntityPTCreator {
	if len(alias) > 0 {
		c.comps = append(c.comps, pt.CompAlias(comp, true, alias[0]))
	} else {
		c.comps = append(c.comps, comp)
	}
	return c
}

// AddMutableComponent 添加不固定的组件
func (c EntityPTCreator) AddMutableComponent(comp any, alias ...string) EntityPTCreator {
	if len(alias) > 0 {
		c.comps = append(c.comps, pt.CompAlias(comp, false, alias[0]))
	} else {
		c.comps = append(c.comps, pt.CompAlias(comp, false, ""))
	}
	return c
}

// Declare 声明实体原型
func (c EntityPTCreator) Declare() {
	if c.svcCtx == nil {
		panic(fmt.Errorf("%w: setting svcCtx is nil", ErrCore))
	}
	c.svcCtx.GetEntityLib().Declare(c.prototype, c.atti, c.comps...)
}
