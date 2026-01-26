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
	"encoding/json"
	"reflect"
	"slices"
	"sync"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
)

type _Entity struct {
	prototype                  string
	instanceRT                 reflect.Type
	scope                      ec.Scope
	componentAwakeOnFirstTouch bool
	componentUniqueID          bool
	meta                       meta.Meta
	components                 []ec.BuiltinComponent
	stringerOnce               sync.Once
	stringerCache              string
}

// Prototype 实体原型名称
func (pt *_Entity) Prototype() string {
	return pt.prototype
}

// InstanceRT 实体实例反射类型
func (pt *_Entity) InstanceRT() reflect.Type {
	if pt.instanceRT == nil {
		return nil
	}
	return reflect.PointerTo(pt.instanceRT)
}

// Scope 可访问作用域
func (pt *_Entity) Scope() ec.Scope {
	return pt.scope
}

// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (pt *_Entity) ComponentAwakeOnFirstTouch() bool {
	return pt.componentAwakeOnFirstTouch
}

// ComponentUniqueID 是否为实体组件分配唯一Id
func (pt *_Entity) ComponentUniqueID() bool {
	return pt.componentUniqueID
}

// Meta 原型Meta信息
func (pt *_Entity) Meta() meta.Meta {
	return pt.meta
}

// CountComponents 组件数量
func (pt *_Entity) CountComponents() int {
	return len(pt.components)
}

// GetComponent 获取组件
func (pt *_Entity) GetComponent(idx int) ec.BuiltinComponent {
	if idx < 0 || idx >= len(pt.components) {
		exception.Panicf("%w: %w: idx out of range", ErrPt, exception.ErrArgs)
	}
	return pt.components[idx]
}

// ListComponents 获取所有组件
func (pt *_Entity) ListComponents() []ec.BuiltinComponent {
	return slices.Clone(pt.components)
}

// Construct 创建实体
func (pt *_Entity) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.Make(ec.With.Default())
	if pt.instanceRT != nil {
		options.InstanceFace = iface.MakeFaceT(reflect.New(pt.instanceRT).Interface().(ec.Entity))
	}
	options.Scope = pt.scope
	options.ComponentAwakeOnFirstTouch = pt.componentAwakeOnFirstTouch
	options.ComponentUniqueID = pt.componentUniqueID
	options = option.Append(options, settings...)

	return pt.assemble(ec.UnsafeNewEntity(options))
}

// String implements fmt.Stringer
func (pt *_Entity) String() string {
	pt.stringerOnce.Do(func() {
		data, err := json.Marshal(pt)
		if err != nil {
			exception.Panicf("%w: unexpected failure marshaling entity: %s", ErrPt, err)
		}
		pt.stringerCache = string(data)
	})
	return pt.stringerCache
}

type _EntityJSON struct {
	Prototype                  string                `json:"prototype"`
	Instance                   string                `json:"instance"`
	Scope                      string                `json:"scope"`
	ComponentAwakeOnFirstTouch bool                  `json:"component_awake_on_first_touch"`
	ComponentUniqueID          bool                  `json:"component_unique_id"`
	Meta                       map[string]any        `json:"meta"`
	Components                 []ec.BuiltinComponent `json:"components"`
}

// MarshalJSON implements json.Marshaler
func (pt _Entity) MarshalJSON() ([]byte, error) {
	entityStringer := _EntityJSON{
		Prototype:                  pt.prototype,
		Scope:                      pt.scope.String(),
		ComponentAwakeOnFirstTouch: pt.componentAwakeOnFirstTouch,
		ComponentUniqueID:          pt.componentUniqueID,
		Meta:                       pt.meta.ToGoMap(),
		Components:                 pt.components,
	}
	if pt.instanceRT != nil {
		entityStringer.Instance = pt.instanceRT.Name()
	}

	data, err := json.Marshal(entityStringer)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (pt *_Entity) assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrPt, exception.ErrArgs)
	}

	ec.UnsafeEntity(entity).SetPT(pt)

	for i := range pt.components {
		builtin := &pt.components[i]

		comp := builtin.PT.Construct()
		ec.UnsafeComponent(comp).SetBuiltin(builtin)
		ec.UnsafeComponent(comp).SetRemovable(builtin.Removable)

		if err := entity.AddComponent(builtin.Name, comp); err != nil {
			exception.Panicf("%w: %w", ErrPt, err)
		}
	}

	return entity
}
