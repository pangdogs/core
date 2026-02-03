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

var With _EntityOption

type _EntityOption struct{}

// Default 默认值
func (_EntityOption) Default() option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		With.InstanceFace(iface.Face[Entity]{}).Apply(options)
		With.Scope(Scope_Global).Apply(options)
		With.PersistId(uid.Nil).Apply(options)
		With.ComponentAwakeOnFirstTouch(false).Apply(options)
		With.ComponentUniqueID(false).Apply(options)
		With.Meta(nil).Apply(options)
	}
}

// InstanceFace 实例，用于扩展实体能力
func (_EntityOption) InstanceFace(face iface.Face[Entity]) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.InstanceFace = face
	}
}

// Scope 可访问作用域
func (_EntityOption) Scope(scope Scope) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.Scope = scope
	}
}

// PersistId 实体持久化Id
func (_EntityOption) PersistId(id uid.Id) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.PersistId = id
	}
}

// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (_EntityOption) ComponentAwakeOnFirstTouch(b bool) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.ComponentAwakeOnFirstTouch = b
	}
}

// ComponentUniqueID 是否为实体组件分配唯一Id
func (_EntityOption) ComponentUniqueID(b bool) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.ComponentUniqueID = b
	}
}

// Meta Meta信息
func (_EntityOption) Meta(m meta.Meta) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.Meta = m
	}
}
