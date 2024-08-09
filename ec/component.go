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
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// Component 组件接口
type Component interface {
	iComponent
	ictx.CurrentContextProvider
	fmt.Stringer

	// GetId 获取组件Id
	GetId() uid.Id
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetFixed 是否固定
	GetFixed() bool
	// DestroySelf 销毁自身
	DestroySelf()
}

type iComponent interface {
	init(name string, entity Entity, composite Component)
	setId(id uid.Id)
	setState(state ComponentState)
	setReflected(v reflect.Value)
	setFixed(b bool)
	getComposite() Component
	eventComponentDestroySelf() event.IEvent
	cleanManagedHooks()
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         uid.Id
	name                       string
	entity                     Entity
	composite                  Component
	state                      ComponentState
	reflected                  reflect.Value
	fixed                      bool
	_eventComponentDestroySelf event.Event
	managedHooks               []event.Hook
}

// GetId 获取组件Id
func (comp *ComponentBehavior) GetId() uid.Id {
	return comp.id
}

// GetName 获取组件名称
func (comp *ComponentBehavior) GetName() string {
	return comp.name
}

// GetEntity 获取组件依附的实体
func (comp *ComponentBehavior) GetEntity() Entity {
	return comp.entity
}

// GetState 获取组件状态
func (comp *ComponentBehavior) GetState() ComponentState {
	return comp.state
}

// GetReflected 获取反射值
func (comp *ComponentBehavior) GetReflected() reflect.Value {
	if comp.reflected.IsValid() {
		return comp.reflected
	}
	comp.reflected = reflect.ValueOf(comp.composite)
	return comp.reflected
}

// GetFixed 是否固定
func (comp *ComponentBehavior) GetFixed() bool {
	return comp.fixed
}

// DestroySelf 销毁自身
func (comp *ComponentBehavior) DestroySelf() {
	switch comp.GetState() {
	case ComponentState_Awake, ComponentState_Start, ComponentState_Alive:
		_EmitEventComponentDestroySelf(UnsafeComponent(comp), comp.composite)
	}
}

// GetCurrentContext 获取当前上下文
func (comp *ComponentBehavior) GetCurrentContext() iface.Cache {
	return comp.entity.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (comp *ComponentBehavior) GetConcurrentContext() iface.Cache {
	return comp.entity.GetConcurrentContext()
}

// String implements fmt.Stringer
func (comp *ComponentBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q, "entity_id":%q}`, comp.GetId(), comp.GetName(), comp.GetEntity().GetId())
}

func (comp *ComponentBehavior) init(name string, entity Entity, composite Component) {
	comp.name = name
	comp.entity = entity
	comp.composite = composite
	comp._eventComponentDestroySelf.Init(false, nil, event.EventRecursion_Discard)
}

func (comp *ComponentBehavior) setId(id uid.Id) {
	comp.id = id
}

func (comp *ComponentBehavior) setState(state ComponentState) {
	if state <= comp.state {
		return
	}

	comp.state = state

	switch comp.state {
	case ComponentState_Detach:
		comp._eventComponentDestroySelf.Close()
	}
}

func (comp *ComponentBehavior) setReflected(v reflect.Value) {
	comp.reflected = v
}

func (comp *ComponentBehavior) setFixed(b bool) {
	comp.fixed = b
}

func (comp *ComponentBehavior) getComposite() Component {
	return comp.composite
}

func (comp *ComponentBehavior) eventComponentDestroySelf() event.IEvent {
	return &comp._eventComponentDestroySelf
}
