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
	"git.golaxy.org/core/ec/ictx"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
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
	// ManagedAddHooks 托管事件钩子（event.Hook），在组件销毁时自动解绑定
	ManagedAddHooks(hooks ...event.Hook)
	// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在组件销毁时自动解绑定
	ManagedAddTagHooks(tag string, hooks ...event.Hook)
	// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
	ManagedGetTagHooks(tag string) []event.Hook
	// ManagedCleanTagHooks 清理根据标签托管的事件钩子（event.Hook）
	ManagedCleanTagHooks(tag string)
	// DestroySelf 销毁自身
	DestroySelf()

	IComponentEventTab
}

type iComponent interface {
	init(name string, entity Entity, instance Component)
	withContext(ctx context.Context)
	setId(id uid.Id)
	setBuiltin(builtin *BuiltinComponent)
	setState(state ComponentState)
	setReflected(v reflect.Value)
	setRemovable(b bool)
	getCallingStateBits() *types.Bits16
	getProcessedStateBits() *types.Bits16
	managedCleanAllHooks()
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	context.Context
	terminate          context.CancelFunc
	terminated         chan async.Ret
	id                 uid.Id
	builtin            *BuiltinComponent
	name               string
	entity             Entity
	instance           Component
	state              ComponentState
	reflected          reflect.Value
	removable          bool
	enable             bool
	callingStateBits   types.Bits16
	processedStateBits types.Bits16
	managedHooks       []event.Hook
	managedTagHooks    generic.SliceMap[string, []event.Hook]

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
	if comp.enable == b {
		return
	}

	comp.enable = b

	_EmitEventComponentEnableChanged(comp, comp.instance, b)
}

// DestroySelf 销毁自身
func (comp *ComponentBehavior) DestroySelf() {
	_EmitEventComponentDestroySelf(comp, comp.instance)
}

// EventComponentEnableChanged 事件：组件启用状态改变
func (comp *ComponentBehavior) EventComponentEnableChanged() event.IEvent {
	return comp.componentEventTab.EventComponentEnableChanged()
}

// EventComponentDestroySelf 事件：组件销毁自身
func (comp *ComponentBehavior) EventComponentDestroySelf() event.IEvent {
	return comp.componentEventTab.EventComponentDestroySelf()
}

// Terminated 已停止
func (comp *ComponentBehavior) Terminated() async.AsyncRet {
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
	return fmt.Sprintf(`{"id":%q, "entity_id":%q, "name":%q, "prototype":%q}`, comp.GetId(), comp.GetEntity().GetId(), comp.GetName(), comp.GetBuiltin().PT.Prototype())
}

func (comp *ComponentBehavior) init(name string, entity Entity, instance Component) {
	comp.name = name
	comp.entity = entity
	comp.instance = instance
	comp.removable = true
	comp.enable = true
	comp.componentEventTab.Init(false, nil, event.EventRecursion_Allow)
	comp.setState(ComponentState_Birth)
}

func (comp *ComponentBehavior) withContext(ctx context.Context) {
	comp.Context, comp.terminate = context.WithCancel(ctx)
	comp.terminated = async.MakeAsyncRet()
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
		if comp.processedStateBits.Is(int8(state)) {
			return
		}
	}

	comp.state = state
	comp.processedStateBits.Set(int8(state), true)

	switch comp.state {
	case ComponentState_Death:
		comp.terminate()
		comp.componentEventTab.Close()
	case ComponentState_Destroyed:
		comp.managedCleanAllHooks()
		async.Return(comp.terminated, async.VoidRet)
	}
}

func (comp *ComponentBehavior) setReflected(v reflect.Value) {
	comp.reflected = v
}

func (comp *ComponentBehavior) setRemovable(b bool) {
	comp.removable = b
}

func (comp *ComponentBehavior) getCallingStateBits() *types.Bits16 {
	return &comp.callingStateBits
}

func (comp *ComponentBehavior) getProcessedStateBits() *types.Bits16 {
	return &comp.processedStateBits
}
