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
	"reflect"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// Component 组件接口
type Component interface {
	iComponent
	corectx.CurrentContextProvider
	fmt.Stringer

	// GetId 获取组件Id
	GetId() uid.Id
	// GetBuiltin 获取实体原型中的组件信息
	GetBuiltin() BuiltinComponent
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetRemovable 是否可以删除
	GetRemovable() bool
	// GetEnable 获取组件是否启用
	GetEnable() bool
	// SetEnable 设置组件是否启用
	SetEnable(b bool)
	// Managed 托管事件句柄
	Managed() *event.ManagedHandles
	// Destroy 销毁
	Destroy()

	IComponentEventTab
}

type iComponent interface {
	init(name string, entity Entity, instance Component)
	setId(id uid.Id)
	setBuiltin(builtin *BuiltinComponent)
	setState(state ComponentState)
	setReflected(v reflect.Value)
	setRemovable(b bool)
	getCallingStateBits() *generic.Bits16
	getProcessedStateBits() *generic.Bits16
	getAttachedHandle() (int, int64)
	setAttachedHandle(idx int, ver int64)
	managedRuntimeUpdateHandle(updateHandle event.Handle)
	managedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle)
	managedUnbindRuntimeHandles()
}

const (
	componentReentrancyGuard_SetEnable = iota
	componentReentrancyGuard_Destroy
)

// ComponentBehavior 组件行为，在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                    uid.Id
	builtin               *BuiltinComponent
	name                  string
	entity                Entity
	instance              Component
	state                 ComponentState
	reflected             reflect.Value
	removable             bool
	enable                bool
	callingStateBits      generic.Bits16
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	attachedIndex         int
	attachedVersion       int64
	managedHandles        event.ManagedHandles
	managedRuntimeHandles [2]event.Handle
	stringerCache         string

	componentEventTab componentEventTab
}

// GetId 获取组件Id
func (comp *ComponentBehavior) GetId() uid.Id {
	return comp.id
}

// GetBuiltin 获取实体原型中的组件信息
func (comp *ComponentBehavior) GetBuiltin() BuiltinComponent {
	if comp.builtin == nil {
		return *noneBuiltinComponent
	}
	return *comp.builtin
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

// GetRemovable 是否可以删除
func (comp *ComponentBehavior) GetRemovable() bool {
	return comp.removable
}

// GetEnable 获取组件是否启用
func (comp *ComponentBehavior) GetEnable() bool {
	return comp.enable
}

// SetEnable 设置组件是否启用
func (comp *ComponentBehavior) SetEnable(b bool) {
	comp.reentrancyGuard.Call(componentReentrancyGuard_SetEnable, func() {
		if comp.state > ComponentState_Alive {
			return
		}

		if comp.enable == b {
			return
		}
		comp.enable = b

		_EmitEventComponentEnableChanged(comp, comp.instance, b)

		if comp.entity != nil {
			comp.entity.onComponentEnableChangedIfVersion(comp.attachedIndex, comp.attachedVersion)
		}
	})
}

// Managed 获取托管事件句柄
func (comp *ComponentBehavior) Managed() *event.ManagedHandles {
	return &comp.managedHandles
}

// Destroy 销毁
func (comp *ComponentBehavior) Destroy() {
	comp.reentrancyGuard.Call(componentReentrancyGuard_Destroy, func() {
		if comp.state > ComponentState_Alive {
			return
		}

		if !comp.GetRemovable() {
			return
		}

		_EmitEventComponentDestroy(comp, comp.instance)

		if comp.entity != nil {
			comp.entity.removeComponentIfVersion(comp.attachedIndex, comp.attachedVersion)
		}
	})
}

// EventComponentEnableChanged 事件：组件启用状态改变
func (comp *ComponentBehavior) EventComponentEnableChanged() event.IEvent {
	return comp.componentEventTab.EventComponentEnableChanged()
}

// EventComponentDestroy 事件：组件销毁
func (comp *ComponentBehavior) EventComponentDestroy() event.IEvent {
	return comp.componentEventTab.EventComponentDestroy()
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
	if comp.stringerCache == "" {
		comp.stringerCache = fmt.Sprintf(`{"id":%q, "entity_id":%q, "name":%q, "prototype":%q}`, comp.GetId(), comp.GetEntity().GetId(), comp.GetName(), comp.GetBuiltin().PT.Prototype())
	}
	return comp.stringerCache
}

func (comp *ComponentBehavior) init(name string, entity Entity, instance Component) {
	comp.name = name
	comp.entity = entity
	comp.instance = instance
	comp.removable = true
	comp.enable = true
	comp.setState(ComponentState_Birth)
}

func (comp *ComponentBehavior) setId(id uid.Id) {
	comp.id = id
}

func (comp *ComponentBehavior) setBuiltin(builtin *BuiltinComponent) {
	comp.builtin = builtin
}

func (comp *ComponentBehavior) setState(state ComponentState) {
	switch state {
	case ComponentState_Idle, ComponentState_Alive:
		break
	default:
		if comp.processedStateBits.Is(int(state)) {
			return
		}
	}

	comp.state = state
	comp.processedStateBits.Set(int(state), true)

	switch comp.state {
	case ComponentState_Death:
		comp.componentEventTab.SetEnable(false)
	case ComponentState_Destroyed:
		comp.managedHandles.UnbindAllEventHandles()
		comp.managedUnbindRuntimeHandles()
	}
}

func (comp *ComponentBehavior) setReflected(v reflect.Value) {
	comp.reflected = v
}

func (comp *ComponentBehavior) setRemovable(b bool) {
	comp.removable = b
}

func (comp *ComponentBehavior) getCallingStateBits() *generic.Bits16 {
	return &comp.callingStateBits
}

func (comp *ComponentBehavior) getProcessedStateBits() *generic.Bits16 {
	return &comp.processedStateBits
}

func (comp *ComponentBehavior) getAttachedHandle() (int, int64) {
	return comp.attachedIndex, comp.attachedVersion
}

func (comp *ComponentBehavior) setAttachedHandle(idx int, ver int64) {
	comp.attachedIndex = idx
	comp.attachedVersion = ver
}

func (comp *ComponentBehavior) managedRuntimeUpdateHandle(updateHandle event.Handle) {
	if comp.managedRuntimeHandles[0] != updateHandle {
		comp.managedRuntimeHandles[0].Unbind()
	}
	comp.managedRuntimeHandles[0] = updateHandle
}

func (comp *ComponentBehavior) managedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	if comp.managedRuntimeHandles[1] != lateUpdateHandle {
		comp.managedRuntimeHandles[1].Unbind()
	}
	comp.managedRuntimeHandles[1] = lateUpdateHandle
}

func (comp *ComponentBehavior) managedUnbindRuntimeHandles() {
	event.UnbindHandles(comp.managedRuntimeHandles[:])
}
