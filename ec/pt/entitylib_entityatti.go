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
	"git.golaxy.org/core/utils/generic"
)

// BuildEntityAttribute 创建实体原型属性，用于注册实体原型时自定义相关属性
func BuildEntityAttribute(prototype string) EntityAttribute {
	if prototype == "" {
		exception.Panicf("%w: %w: prototype is empty", ErrPt, exception.ErrArgs)
	}
	return EntityAttribute{
		Prototype: prototype,
	}
}

// ComponentAttribute 组件原型属性
type ComponentAttribute struct {
	Instance  any                           // 组件实例（必填）
	Name      string                        // 组件名称
	Removable bool                          // 是否可以删除
	Extra     generic.SliceMap[string, any] // 自定义属性
}

func (atti ComponentAttribute) SetName(name string) ComponentAttribute {
	atti.Name = name
	return atti
}

func (atti ComponentAttribute) SetRemovable(b bool) ComponentAttribute {
	atti.Removable = b
	return atti
}

func (atti ComponentAttribute) SetExtra(extra map[string]any) ComponentAttribute {
	atti.Extra = generic.MakeSliceMapFromGoMap(extra)
	return atti
}

func (atti ComponentAttribute) OverrideMergeExtra(extra map[string]any) ComponentAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.Add(k, v)
	}
	return atti
}

func (atti ComponentAttribute) IncrementalMergeExtra(extra map[string]any) ComponentAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.TryAdd(k, v)
	}
	return atti
}

// BuildComponentAttribute 创建组件原型属性，用于注册实体原型时自定义相关属性
func BuildComponentAttribute(instance any) ComponentAttribute {
	if instance == nil {
		exception.Panicf("%w: %w: instance is nil", ErrPt, exception.ErrArgs)
	}
	return ComponentAttribute{
		Instance: instance,
	}
}
