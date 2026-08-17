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
	"reflect"
	"sync/atomic"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
)

// NewEntity 创建处于 Born 状态的实体。
func NewEntity(settings ...option.Setting[EntityOptions]) Entity {
	return UnsafeNewEntity(option.New(With.Default(), settings...))
}

// UnsafeNewEntity 使用已解析的选项创建实体，供框架内部使用。
//
// Deprecated: 请使用 NewEntity。
func UnsafeNewEntity(options EntityOptions) Entity {
	var e Entity

	if !options.InstanceFace.IsNil() {
		e = options.InstanceFace.Iface
	} else {
		e = &EntityBehavior{}
	}
	e.init(options)

	return e
}

// Entity 表示由 Runtime 驱动生命周期的实体。
//
// 除 ConcurrentEntity 暴露的只读能力外，其方法应在所属 Runtime 的运行协程中调用。
type Entity interface {
	iEntity
	iComponentManager
	iTreeNode
	ConcurrentEntity
	corectx.CurrentContextProvider
	reinterpret.InstanceProvider

	// Id 返回实体 ID。
	Id() uid.Id
	// PT 返回实体原型；尚未绑定原型时返回空原型对象。
	PT() EntityPT
	// Scope 返回实体的可查询范围。
	Scope() Scope
	// State 返回当前生命周期状态。
	State() EntityState
	// Reflected 返回实际实体实例的反射值。
	Reflected() reflect.Value
	// Meta 返回实体元数据。
	Meta() meta.Meta
	// Managed 返回随实体销毁自动解绑的事件句柄集合。
	Managed() *event.ManagedHandles
	// Destroy 请求所属 Runtime 销毁实体；重复请求会被忽略。
	Destroy()

	IEntityEventTab
}

type iEntity interface {
	init(options EntityOptions)
	getOptions() *EntityOptions
	setId(id uid.Id)
	setPT(prototype EntityPT)
	setState(state EntityState)
	setReflected(v reflect.Value)
	getProcessedStateBits() *generic.Bits16
	getEnteredHandle() (int, int64)
	setEnteredHandle(idx int, ver int64)
	managedRuntimeUpdateHandle(updateHandle event.Handle)
	managedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle)
	managedUnbindRuntimeHandles()
}

const (
	entityReentrancyGuard_Destroy = iota
)

// EntityBehavior 提供 Entity 的默认实现，扩展实体时应将其匿名嵌入自定义结构体。
type EntityBehavior struct {
	context.Context
	options               EntityOptions
	prototype             EntityPT
	asyncScope            *async.Scope
	terminated            async.Completer
	runtimeCtx            runtimeContext
	componentNameIndex    generic.SliceMap[string, int]
	componentList         generic.FreeList[Component]
	state                 EntityState
	reflected             reflect.Value
	treeNodeState         TreeNodeState
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	enteredIndex          int
	enteredVersion        int64
	managedHandles        event.ManagedHandles
	managedRuntimeHandles [2]event.Handle
	stringerCache         atomic.Pointer[string]

	entityEventTab                 entityEventTab
	entityComponentManagerEventTab entityComponentManagerEventTab
	entityTreeNodeEventTab         entityTreeNodeEventTab
}

// Id 返回实体 ID。
func (entity *EntityBehavior) Id() uid.Id {
	return entity.options.PersistId
}

// PT 返回实体原型；尚未绑定原型时返回空原型对象。
func (entity *EntityBehavior) PT() EntityPT {
	if entity.prototype == nil {
		return noneEntityPT
	}
	return entity.prototype
}

// Scope 返回实体的可查询范围。
func (entity *EntityBehavior) Scope() Scope {
	return entity.options.Scope
}

// State 返回当前生命周期状态。
func (entity *EntityBehavior) State() EntityState {
	return entity.state
}

// Reflected 返回实际实体实例的反射值，并缓存首次解析结果。
func (entity *EntityBehavior) Reflected() reflect.Value {
	if entity.reflected.IsValid() {
		return entity.reflected
	}
	entity.reflected = reflect.ValueOf(entity.getInstance())
	return entity.reflected
}

// Meta 返回实体元数据。
func (entity *EntityBehavior) Meta() meta.Meta {
	return entity.options.Meta
}

// Managed 返回随实体销毁自动解绑的事件句柄集合。
func (entity *EntityBehavior) Managed() *event.ManagedHandles {
	return &entity.managedHandles
}

// Destroy 请求所属 Runtime 销毁实体；实体离开活动阶段后调用无效。
func (entity *EntityBehavior) Destroy() {
	entity.reentrancyGuard.Call(entityReentrancyGuard_Destroy, func() {
		if entity.state > EntityState_Alive {
			return
		}
		_EmitEventEntityDestroy(entity, entity.getInstance())
	})
}

// EventEntityDestroy 返回实体销毁请求事件。
func (entity *EntityBehavior) EventEntityDestroy() event.IEvent {
	return entity.entityEventTab.EventEntityDestroy()
}

// CurrentContextCache 返回实体所属 Runtime 的当前上下文接口缓存。
func (entity *EntityBehavior) CurrentContextCache() iface.Cache {
	return entity.runtimeCtx.CurrentContextCache()
}

// InstanceFaceCache 返回实际实体实例的接口缓存，供类型重解释使用。
func (entity *EntityBehavior) InstanceFaceCache() iface.Cache {
	return entity.options.InstanceFace.Cache
}

func (entity *EntityBehavior) init(options EntityOptions) {
	entity.options = options

	if entity.options.InstanceFace.IsNil() {
		entity.options.InstanceFace = iface.NewFaceT[Entity](entity)
	}
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.options
}

func (entity *EntityBehavior) setId(id uid.Id) {
	entity.options.PersistId = id
}

func (entity *EntityBehavior) setPT(prototype EntityPT) {
	entity.prototype = prototype
}

func (entity *EntityBehavior) setState(state EntityState) {
	if entity.state >= state {
		return
	}

	entity.state = state

	switch entity.state {
	case EntityState_Dead:
		entity.asyncScope.Close()
		entity.entityEventTab.SetEnabled(false)
		entity.entityComponentManagerEventTab.SetEnabled(false)
		entity.entityTreeNodeEventTab.SetEnabled(false)
	case EntityState_Destroyed:
		entity.managedHandles.UnbindAllEventHandles()
		entity.managedUnbindRuntimeHandles()
		entity.terminated.Complete()
	}
}

func (entity *EntityBehavior) setReflected(v reflect.Value) {
	entity.reflected = v
}

func (entity *EntityBehavior) getProcessedStateBits() *generic.Bits16 {
	return &entity.processedStateBits
}

func (entity *EntityBehavior) getEnteredHandle() (int, int64) {
	return entity.enteredIndex, entity.enteredVersion
}

func (entity *EntityBehavior) setEnteredHandle(idx int, ver int64) {
	entity.enteredIndex = idx
	entity.enteredVersion = ver
}

func (entity *EntityBehavior) managedRuntimeUpdateHandle(updateHandle event.Handle) {
	if entity.managedRuntimeHandles[0] != updateHandle {
		entity.managedRuntimeHandles[0].Unbind()
	}
	entity.managedRuntimeHandles[0] = updateHandle
}

func (entity *EntityBehavior) managedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	if entity.managedRuntimeHandles[1] != lateUpdateHandle {
		entity.managedRuntimeHandles[1].Unbind()
	}
	entity.managedRuntimeHandles[1] = lateUpdateHandle
}

func (entity *EntityBehavior) managedUnbindRuntimeHandles() {
	event.UnbindHandles(entity.managedRuntimeHandles[:])
}
