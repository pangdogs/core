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
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// EntityAttribute 实体原型属性
type EntityAttribute struct {
	Prototype                  string                        // 实体原型名称（必填）
	Instance                   any                           // 实体实例
	Scope                      *ec.Scope                     // 可访问作用域
	ComponentAwakeOnFirstTouch *bool                         // 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentUniqueID          *bool                         // 是否为实体组件分配唯一Id
	Extra                      generic.SliceMap[string, any] // 自定义属性
}

func (atti EntityAttribute) SetInstance(instance any) EntityAttribute {
	atti.Instance = instance
	return atti
}

func (atti EntityAttribute) SetScope(scope ec.Scope) EntityAttribute {
	atti.Scope = &scope
	return atti
}

func (atti EntityAttribute) SetComponentAwakeOnFirstTouch(componentAwakeOnFirstTouch bool) EntityAttribute {
	atti.ComponentAwakeOnFirstTouch = &componentAwakeOnFirstTouch
	return atti
}

func (atti EntityAttribute) SetComponentUniqueID(componentUniqueID bool) EntityAttribute {
	atti.ComponentUniqueID = &componentUniqueID
	return atti
}

func (atti EntityAttribute) SetExtra(extra map[string]any) EntityAttribute {
	atti.Extra = generic.MakeSliceMapFromGoMap(extra)
	return atti
}

func (atti EntityAttribute) OverrideExtra(extra map[string]any) EntityAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.Add(k, v)
	}
	return atti
}

func (atti EntityAttribute) CombineExtra(extra map[string]any) EntityAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.TryAdd(k, v)
	}
	return atti
}

// Entity 创建实体原型属性，用于注册实体原型时自定义相关属性
func Entity(prototype string) EntityAttribute {
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

func (atti ComponentAttribute) OverrideExtra(extra map[string]any) ComponentAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.Add(k, v)
	}
	return atti
}

func (atti ComponentAttribute) CombineExtra(extra map[string]any) ComponentAttribute {
	atti.Extra = atti.Extra.Clone()
	for k, v := range extra {
		atti.Extra.TryAdd(k, v)
	}
	return atti
}

// Component 创建组件原型属性，用于注册实体原型时自定义相关属性
func Component(instance any) ComponentAttribute {
	if instance == nil {
		exception.Panicf("%w: %w: instance is nil", ErrPt, exception.ErrArgs)
	}
	return ComponentAttribute{
		Instance: instance,
	}
}
