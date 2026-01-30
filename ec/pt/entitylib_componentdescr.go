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

// NewComponentDescriptor 创建组件原型描述，用于注册实体原型的组件
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

// ComponentDescriptor 组件原型描述
type ComponentDescriptor struct {
	Instance  any       // 组件实例（必填）
	Name      string    // 组件名称
	Removable bool      // 是否可以删除
	Meta      meta.Meta // 原型Meta信息
}

func (descr *ComponentDescriptor) SetName(name string) *ComponentDescriptor {
	descr.Name = name
	return descr
}

func (descr *ComponentDescriptor) SetRemovable(b bool) *ComponentDescriptor {
	descr.Removable = b
	return descr
}

func (descr *ComponentDescriptor) SetMeta(dict map[string]any) *ComponentDescriptor {
	descr.Meta = meta.New(dict)
	return descr
}

func (descr *ComponentDescriptor) MergeMeta(dict map[string]any) *ComponentDescriptor {
	for k, v := range dict {
		descr.Meta.Add(k, v)
	}
	return descr
}

func (descr *ComponentDescriptor) MergeMetaIfAbsent(dict map[string]any) *ComponentDescriptor {
	for k, v := range dict {
		descr.Meta.TryAdd(k, v)
	}
	return descr
}

func (descr *ComponentDescriptor) AssignMeta(m meta.Meta) *ComponentDescriptor {
	descr.Meta = m
	return descr
}
