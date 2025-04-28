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

// BuildComponentAttribute 创建组件原型属性，用于注册实体原型时自定义相关属性
func BuildComponentAttribute(instance any) *ComponentAttribute {
	if instance == nil {
		exception.Panicf("%w: %w: instance is nil", ErrPt, exception.ErrArgs)
	}
	return &ComponentAttribute{
		Instance: instance,
	}
}

// ComponentAttribute 组件原型属性
type ComponentAttribute struct {
	Instance  any       // 组件实例（必填）
	Name      string    // 组件名称
	Removable bool      // 是否可以删除
	Extra     meta.Meta // 自定义属性
}

func (atti *ComponentAttribute) SetName(name string) *ComponentAttribute {
	atti.Name = name
	return atti
}

func (atti *ComponentAttribute) SetRemovable(b bool) *ComponentAttribute {
	atti.Removable = b
	return atti
}

func (atti *ComponentAttribute) SetExtra(dict map[string]any) *ComponentAttribute {
	atti.Extra = meta.M(dict)
	return atti
}

func (atti *ComponentAttribute) MergeExtra(dict map[string]any) *ComponentAttribute {
	for k, v := range dict {
		atti.Extra.Add(k, v)
	}
	return atti
}

func (atti *ComponentAttribute) MergeExtraIfAbsent(dict map[string]any) *ComponentAttribute {
	for k, v := range dict {
		atti.Extra.TryAdd(k, v)
	}
	return atti
}

func (atti *ComponentAttribute) AssignExtra(m meta.Meta) *ComponentAttribute {
	atti.Extra = m
	return atti
}
