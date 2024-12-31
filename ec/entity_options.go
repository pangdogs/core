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
	InstanceFace               iface.Face[Entity] // 实例，用于扩展实体能力
	Scope                      Scope              // 可访问作用域
	PersistId                  uid.Id             // 实体持久化Id
	ComponentAwakeOnFirstTouch bool               // 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentUniqueID          bool               // 是否为实体组件分配唯一Id
	Meta                       meta.Meta          // Meta信息
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		With.InstanceFace(iface.Face[Entity]{})(o)
		With.Scope(Scope_Global)(o)
		With.PersistId(uid.Nil)(o)
		With.ComponentAwakeOnFirstTouch(false)(o)
		With.ComponentUniqueID(false)(o)
		With.Meta(nil)(o)
	}
}

// InstanceFace 实例，用于扩展实体能力
func (_Option) InstanceFace(face iface.Face[Entity]) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.InstanceFace = face
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

// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (_Option) ComponentAwakeOnFirstTouch(b bool) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.ComponentAwakeOnFirstTouch = b
	}
}

// ComponentUniqueID 是否为实体组件分配唯一Id
func (_Option) ComponentUniqueID(b bool) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.ComponentUniqueID = b
	}
}

// Meta Meta信息
func (_Option) Meta(m meta.Meta) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Meta = m
	}
}
