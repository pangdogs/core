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

package ec

import (
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	CompositeFace      iface.Face[Entity] // 扩展者，在扩展实体自身能力时使用
	Prototype          string             // 实体原型名称
	Scope              Scope              // 可访问作用域
	PersistId          uid.Id             // 实体持久化Id
	AwakeOnFirstAccess bool               // 开启组件被首次访问时，检测并调用Awake()
	Meta               meta.Meta          // Meta信息
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		With.CompositeFace(iface.Face[Entity]{})(o)
		With.Prototype("")(o)
		With.Scope(Scope_Global)(o)
		With.PersistId(uid.Nil)(o)
		With.AwakeOnFirstAccess(false)(o)
		With.Meta(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (_Option) CompositeFace(face iface.Face[Entity]) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (_Option) Prototype(pt string) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Prototype = pt
	}
}

// Scope 可访问作用域
func (_Option) Scope(scope Scope) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}

// PersistId 实体持久化Id
func (_Option) PersistId(id uid.Id) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.PersistId = id
	}
}

// AwakeOnFirstAccess 开启组件被首次访问时，检测并调用Awake()
func (_Option) AwakeOnFirstAccess(b bool) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.AwakeOnFirstAccess = b
	}
}

// Meta Meta信息
func (_Option) Meta(m meta.Meta) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Meta = m
	}
}
