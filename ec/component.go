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
	"context"
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
	iContext
	ictx.CurrentContextProvider
	fmt.Stringer

	// GetId 获取组件Id
	GetId() uid.Id
	// GetDesc 获取组件原型信息
	GetDesc() ComponentDesc
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetNonRemovable 是否不可删除
	GetNonRemovable() bool
	// DestroySelf 销毁自身
	DestroySelf()

	IComponentEventTab
}

type iComponent interface {
	init(name string, entity Entity, instance Component)
	withContext(ctx context.Context)
	setId(id uid.Id)
	setDesc(desc *ComponentDesc)
	setState(state ComponentState)
	setReflected(v reflect.Value)
	setNonRemovable(b bool)
	cleanManagedHooks()
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	context.Context
	terminate    context.CancelFunc
	terminated   chan struct{}
	id           uid.Id
	desc         *ComponentDesc
	name         string
	entity       Entity
	instance     Component
	state        ComponentState
	reflected    reflect.Value
	nonRemovable bool
	managedHooks []event.Hook

	componentEventTab componentEventTab
}

// GetId 获取组件Id
func (comp *ComponentBehavior) GetId() uid.Id {
	return comp.id
}

// GetDesc 获取组件原型信息
func (comp *ComponentBehavior) GetDesc() ComponentDesc {
	if comp.desc == nil {
		return *noneComponentDesc
	}
	return *comp.desc
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
	comp.reflected = reflect.ValueOf(comp.instance)
	return comp.reflected
}

// GetNonRemovable 是否不可删除
func (comp *ComponentBehavior) GetNonRemovable() bool {
	return comp.nonRemovable
}

// DestroySelf 销毁自身
func (comp *ComponentBehavior) DestroySelf() {
	switch comp.GetState() {
	case ComponentState_Awake, ComponentState_Start, ComponentState_Alive:
		_EmitEventComponentDestroySelf(comp, comp.instance)
	}
}

// EventComponentDestroySelf 事件：组件销毁自身
func (comp *ComponentBehavior) EventComponentDestroySelf() event.IEvent {
	return comp.componentEventTab.EventComponentDestroySelf()
}

// Terminated 已停止
func (comp *ComponentBehavior) Terminated() <-chan struct{} {
	return comp.terminated
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
	return fmt.Sprintf(`{"id":%q, "entity_id":%q, "name":%q, "prototype":%q}`, comp.GetId(), comp.GetEntity().GetId(), comp.GetName(), comp.GetDesc().PT.Prototype())
}

func (comp *ComponentBehavior) init(name string, entity Entity, instance Component) {
	comp.name = name
	comp.entity = entity
	comp.instance = instance
	comp.componentEventTab.Init(false, nil, event.EventRecursion_Allow)
}

func (comp *ComponentBehavior) withContext(ctx context.Context) {
	comp.Context, comp.terminate = context.WithCancel(ctx)
	comp.terminated = make(chan struct{})
}

func (comp *ComponentBehavior) setId(id uid.Id) {
	comp.id = id
}

func (comp *ComponentBehavior) setDesc(desc *ComponentDesc) {
	comp.desc = desc
}

func (comp *ComponentBehavior) setState(state ComponentState) {
	if state <= comp.state {
		return
	}

	comp.state = state

	switch comp.state {
	case ComponentState_Detach:
		comp.terminate()
		comp.componentEventTab.Close()
	case ComponentState_Death:
		close(comp.terminated)
	}
}

func (comp *ComponentBehavior) setReflected(v reflect.Value) {
	comp.reflected = v
}

func (comp *ComponentBehavior) setNonRemovable(b bool) {
	comp.nonRemovable = b
}
