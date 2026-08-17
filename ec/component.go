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
	"reflect"
	"sync/atomic"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// Component 表示依附于实体、由实体所属 Runtime 驱动生命周期的组件。
//
// 除 ConcurrentComponent 暴露的能力外，Component 的方法应在所属 Runtime 的
// 运行 goroutine 中调用。
type Component interface {
	iComponent
	ConcurrentComponent
	corectx.CurrentContextProvider

	// Id 返回组件 ID；未启用组件唯一 ID 时通常与 Entity ID 相同。
	Id() uid.Id
	// Builtin 返回组件的原型描述；实体原型内建组件包含其位置与配置。
	// 由 ComponentPT 独立构造的组件通常返回 Offset=-1 的描述，未绑定原型时返回空描述。
	Builtin() BuiltinComponent
	// Name 返回组件在 Entity 中的名称。
	Name() string
	// Entity 返回组件所依附的实体。
	Entity() Entity
	// State 返回当前生命周期状态。
	State() ComponentState
	// Reflected 返回实际组件实例的反射值。
	Reflected() reflect.Value
	// Removable 报告组件是否允许动态删除。
	Removable() bool
	// Enabled 报告组件是否已被标记为启用。
	Enabled() bool
	// SetEnabled 请求切换启用状态；重复设置或组件已进入 Detaching 及后续状态时无效。
	SetEnabled(b bool)
	// Managed 返回随组件销毁自动解绑的事件句柄集合。
	Managed() *event.ManagedHandles
	// Destroy 请求从实体删除组件；不可删除的组件会忽略该请求。
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

// ComponentBehavior 提供 Component 的默认实现，扩展组件时应将其匿名嵌入自定义结构体。
type ComponentBehavior struct {
	id                    uid.Id
	builtin               *BuiltinComponent
	name                  string
	entity                Entity
	asyncScope            atomic.Pointer[componentAsyncScopeState]
	instance              Component
	state                 ComponentState
	reflected             reflect.Value
	removable             bool
	enabled               bool
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	attachedIndex         int
	attachedVersion       int64
	managedHandles        event.ManagedHandles
	managedRuntimeHandles [2]event.Handle
	stringerCache         atomic.Pointer[string]

	componentEventTab componentEventTab
}

// Id 返回组件 ID。
func (comp *ComponentBehavior) Id() uid.Id {
	return comp.id
}

// Builtin 返回组件的原型描述；未绑定组件原型时返回空描述。
func (comp *ComponentBehavior) Builtin() BuiltinComponent {
	if comp.builtin == nil {
		return *noneBuiltinComponent
	}
	return *comp.builtin
}

// Name 返回组件在实体中的名称。
func (comp *ComponentBehavior) Name() string {
	return comp.name
}

// Entity 返回组件所依附的实体。
func (comp *ComponentBehavior) Entity() Entity {
	return comp.entity
}

// State 返回当前生命周期状态。
func (comp *ComponentBehavior) State() ComponentState {
	return comp.state
}

// Reflected 返回实际组件实例的反射值，并缓存首次解析结果。
func (comp *ComponentBehavior) Reflected() reflect.Value {
	if comp.reflected.IsValid() {
		return comp.reflected
	}
	comp.reflected = reflect.ValueOf(comp.instance)
	return comp.reflected
}

// Removable 报告组件是否允许动态删除。
func (comp *ComponentBehavior) Removable() bool {
	return comp.removable
}

// Enabled 报告组件是否已被标记为启用。
func (comp *ComponentBehavior) Enabled() bool {
	return comp.enabled
}

// SetEnabled 请求切换启用状态；标记改变后先派发组件事件，再通知实体组件管理器推进生命周期。
// Runtime 仅在所属 Entity 处于 Awaking 至 Alive 时处理该生命周期请求。
func (comp *ComponentBehavior) SetEnabled(b bool) {
	comp.reentrancyGuard.Call(componentReentrancyGuard_SetEnable, func() {
		if comp.state > ComponentState_Alive {
			return
		}

		if comp.enabled == b {
			return
		}
		comp.enabled = b

		_EmitEventComponentEnableChanged(comp, comp.instance, b)

		if comp.entity != nil {
			comp.entity.onComponentEnableChangedIfVersion(comp.attachedIndex, comp.attachedVersion)
		}
	})
}

// Managed 返回随组件销毁自动解绑的事件句柄集合。
func (comp *ComponentBehavior) Managed() *event.ManagedHandles {
	return &comp.managedHandles
}

// Destroy 请求从实体删除组件；不可删除或已进入 Detaching 及后续状态时无效。
// 有效请求会先派发组件销毁请求事件，再同步进入实体组件管理器的移除流程。
func (comp *ComponentBehavior) Destroy() {
	comp.reentrancyGuard.Call(componentReentrancyGuard_Destroy, func() {
		if comp.state > ComponentState_Alive {
			return
		}

		if !comp.Removable() {
			return
		}

		_EmitEventComponentDestroy(comp, comp.instance)

		if comp.entity != nil {
			comp.entity.onComponentDestroyIfVersion(comp.attachedIndex, comp.attachedVersion)
		}
	})
}

// EventComponentEnableChanged 返回启用标记变更事件；派发时 Runtime 尚未处理对应生命周期。
func (comp *ComponentBehavior) EventComponentEnableChanged() event.IEvent {
	return comp.componentEventTab.EventComponentEnableChanged()
}

// EventComponentDestroy 返回显式销毁请求事件；派发时组件尚未进入 Detaching。
// 该事件不表示移除完成，随所属 Entity 销毁也不会触发。
func (comp *ComponentBehavior) EventComponentDestroy() event.IEvent {
	return comp.componentEventTab.EventComponentDestroy()
}

// CurrentContextCache 返回所属实体的当前上下文接口缓存。
func (comp *ComponentBehavior) CurrentContextCache() iface.Cache {
	return comp.entity.CurrentContextCache()
}

func (comp *ComponentBehavior) init(name string, entity Entity, instance Component) {
	comp.name = name
	comp.entity = entity
	comp.instance = instance
	comp.removable = comp.Builtin().Removable
	comp.enabled = true
}

func (comp *ComponentBehavior) setId(id uid.Id) {
	comp.id = id
}

func (comp *ComponentBehavior) setBuiltin(builtin *BuiltinComponent) {
	comp.builtin = builtin
}

func (comp *ComponentBehavior) setState(state ComponentState) {
	switch state {
	case ComponentState_Idle:
		switch comp.state {
		case ComponentState_Enabling, ComponentState_Starting, ComponentState_Alive:
			break
		default:
			return
		}
	case ComponentState_Starting:
		switch comp.state {
		case ComponentState_Enabling, ComponentState_Idle:
			break
		default:
			return
		}
	default:
		if comp.state >= state {
			return
		}
	}

	comp.state = state

	switch comp.state {
	case ComponentState_Dead:
		comp.closeAsyncScope()
		comp.componentEventTab.SetEnabled(false)
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
