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

package pt

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/meta"
)

// NewComponentDescriptor 创建用于实体原型声明的组件描述；instance 为 nil 时 panic。
func NewComponentDescriptor(instance any) *ComponentDescriptor {
	if instance == nil {
		exception.Panicf("%w: %w: instance is nil", ErrPt, exception.ErrArgs)
	}
	return &ComponentDescriptor{
		Instance:  instance,
		Name:      "",
		Removable: false,
		Meta:      nil,
	}
}

// ComponentDescriptor 描述实体原型中的一个内建组件。
type ComponentDescriptor struct {
	Instance  any       // Instance 是组件值、组件类型或已声明组件原型名。
	Name      string    // Name 是组件加入实体时使用的名称；为空时取组件类型名。
	Removable bool      // Removable 指示组件是否允许动态删除；默认为 false。
	Meta      meta.Meta // Meta 是该内建组件的原型元数据。
}

// SetName 设置组件在实体中的名称并返回 descr，以便链式调用。
func (descr *ComponentDescriptor) SetName(name string) *ComponentDescriptor {
	descr.Name = name
	return descr
}

// SetRemovable 设置组件是否允许动态删除并返回 descr。
func (descr *ComponentDescriptor) SetRemovable(b bool) *ComponentDescriptor {
	descr.Removable = b
	return descr
}

// SetMeta 使用 dict 的副本替换元数据并返回 descr。
func (descr *ComponentDescriptor) SetMeta(dict map[string]any) *ComponentDescriptor {
	descr.Meta = meta.New(dict)
	return descr
}

// MergeMeta 合并 dict；同名键会覆盖原值。
func (descr *ComponentDescriptor) MergeMeta(dict map[string]any) *ComponentDescriptor {
	for k, v := range dict {
		descr.Meta.Add(k, v)
	}
	return descr
}

// MergeMetaIfAbsent 合并 dict，但保留已有的同名键。
func (descr *ComponentDescriptor) MergeMetaIfAbsent(dict map[string]any) *ComponentDescriptor {
	for k, v := range dict {
		descr.Meta.TryAdd(k, v)
	}
	return descr
}

// AssignMeta 直接绑定 m 并返回 descr；m 不会被复制。
func (descr *ComponentDescriptor) AssignMeta(m meta.Meta) *ComponentDescriptor {
	descr.Meta = m
	return descr
}
