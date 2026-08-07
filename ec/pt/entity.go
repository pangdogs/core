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

// Prototype 返回实体原型名。
func (pt *_Entity) Prototype() string {
	return pt.prototype
}

// InstanceRT 返回实体实例的指针类型；使用默认实体实现时返回 nil。
func (pt *_Entity) InstanceRT() reflect.Type {
	if pt.instanceRT == nil {
		return nil
	}
	return reflect.PointerTo(pt.instanceRT)
}

// Scope 返回原型的默认实体作用域。
func (pt *_Entity) Scope() ec.Scope {
	return pt.scope
}

// ComponentAwakeOnFirstTouch 报告组件是否在首次访问时推进至 Awakened。
func (pt *_Entity) ComponentAwakeOnFirstTouch() bool {
	return pt.componentAwakeOnFirstTouch
}

// ComponentUniqueID 报告是否为组件分配唯一 ID。
func (pt *_Entity) ComponentUniqueID() bool {
	return pt.componentUniqueID
}

// Meta 返回实体原型元数据。
func (pt *_Entity) Meta() meta.Meta {
	return pt.meta
}

// CountComponents 返回内建组件数。
func (pt *_Entity) CountComponents() int {
	return len(pt.components)
}

// GetComponent 返回指定位置的内建组件描述；索引越界时 panic。
func (pt *_Entity) GetComponent(idx int) ec.BuiltinComponent {
	if idx < 0 || idx >= len(pt.components) {
		exception.Panicf("%w: %w: idx out of range", ErrPt, exception.ErrArgs)
	}
	return pt.components[idx]
}

// ListComponents 返回全部内建组件描述的副本。
func (pt *_Entity) ListComponents() []ec.BuiltinComponent {
	return slices.Clone(pt.components)
}

// Construct 根据原型创建处于 Born 状态的实体，并应用额外选项。
func (pt *_Entity) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	options := option.New(ec.With.Default())
	if pt.instanceRT != nil {
		options.InstanceFace = iface.NewFaceT(reflect.New(pt.instanceRT).Interface().(ec.Entity))
	}
	options.Scope = pt.scope
	options.ComponentAwakeOnFirstTouch = pt.componentAwakeOnFirstTouch
	options.ComponentUniqueID = pt.componentUniqueID
	options = option.Append(options, settings...)

	return pt.assemble(ec.UnsafeNewEntity(options))
}

// String 返回实体原型的 JSON 文本；编码失败时 panic。
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

// MarshalJSON 将实体原型编码为 JSON。
func (pt *_Entity) MarshalJSON() ([]byte, error) {
	entityStringer := _EntityJSON{
		Prototype:                  pt.prototype,
		Scope:                      pt.scope.String(),
		ComponentAwakeOnFirstTouch: pt.componentAwakeOnFirstTouch,
		ComponentUniqueID:          pt.componentUniqueID,
		Meta:                       pt.meta.ToGoMap(),
		Components:                 pt.components,
	}
	if pt.instanceRT != nil {
		entityStringer.Instance = pt.instanceRT.String()
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
