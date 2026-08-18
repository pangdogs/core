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

// EntityOptions 定义实体的构造选项。
type EntityOptions struct {
	InstanceFace               iface.Face[Entity] // InstanceFace 是用于扩展实体行为的实际实例。
	Scope                      Scope              // Scope 是实体的可查询范围。
	PersistID                  uid.ID             // PersistID 是实体的持久化 ID；Nil 表示由框架分配。
	ComponentAwakeOnFirstTouch bool               // ComponentAwakeOnFirstTouch 指示正常激活期间被访问的组件是否优先执行 Awake。
	ComponentUniqueID          bool               // ComponentUniqueID 指示是否为每个组件分配唯一 ID。
	Meta                       meta.Meta          // Meta 是随实体携带的元数据。
}

// With 提供实体选项构造器。
var With _EntityOption

type _EntityOption struct{}

// Default 返回实体选项的默认设置。
func (_EntityOption) Default() option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		With.InstanceFace(iface.Face[Entity]{}).Apply(options)
		With.Scope(Scope_Global).Apply(options)
		With.PersistID(uid.Nil).Apply(options)
		With.ComponentAwakeOnFirstTouch(false).Apply(options)
		With.ComponentUniqueID(false).Apply(options)
		With.Meta(nil).Apply(options)
	}
}

// InstanceFace 设置用于扩展实体行为的实际实例。
func (_EntityOption) InstanceFace(face iface.Face[Entity]) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.InstanceFace = face
	}
}

// Scope 设置实体的可查询范围。
func (_EntityOption) Scope(scope Scope) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.Scope = scope
	}
}

// PersistID 设置实体的持久化 ID。
func (_EntityOption) PersistID(id uid.ID) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.PersistID = id
	}
}

// ComponentAwakeOnFirstTouch 设置正常激活期间被访问的组件是否优先执行 Awake。
func (_EntityOption) ComponentAwakeOnFirstTouch(b bool) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.ComponentAwakeOnFirstTouch = b
	}
}

// ComponentUniqueID 设置是否为实体的每个组件分配唯一 ID。
func (_EntityOption) ComponentUniqueID(b bool) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.ComponentUniqueID = b
	}
}

// Meta 设置实体元数据。
func (_EntityOption) Meta(m meta.Meta) option.Setting[EntityOptions] {
	return func(options *EntityOptions) {
		options.Meta = m
	}
}
