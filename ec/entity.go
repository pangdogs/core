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
	"reflect"

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

// NewEntity 创建实体
func NewEntity(settings ...option.Setting[EntityOptions]) Entity {
	return UnsafeNewEntity(option.Make(With.Default(), settings...))
}

// Deprecated: UnsafeNewEntity 内部创建实体
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

// Entity 实体接口
type Entity interface {
	iEntity
	iConcurrentEntity
	iContext
	iComponentManager
	iTreeNode
	corectx.CurrentContextProvider
	reinterpret.InstanceProvider
	fmt.Stringer

	// GetId 获取实体Id
	GetId() uid.Id
	// GetPT 获取实体原型信息
	GetPT() EntityPT
	// GetScope 获取可访问作用域
	GetScope() Scope
	// GetState 获取实体状态
	GetState() EntityState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetMeta 获取Meta信息
	GetMeta() meta.Meta
	// Managed 托管事件句柄
	Managed() *event.ManagedHandles
	// Destroy 销毁
	Destroy()

	IEntityEventTab
}

type iEntity interface {
	init(options EntityOptions)
	withContext(ctx context.Context)
	getOptions() *EntityOptions
	setId(id uid.Id)
	setPT(prototype EntityPT)
	setContext(ctx iface.Cache)
	setState(state EntityState)
	setReflected(v reflect.Value)
	getCallingStateBits() *generic.Bits16
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

// EntityBehavior 实体行为，在扩展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	context.Context
	terminate             context.CancelFunc
	terminated            chan async.Ret
	options               EntityOptions
	prototype             EntityPT
	context               iface.Cache
	componentNameIndex    generic.SliceMap[string, int]
	componentList         generic.FreeList[Component]
	state                 EntityState
	reflected             reflect.Value
	treeNodeState         TreeNodeState
	callingStateBits      generic.Bits16
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	enteredIndex          int
	enteredVersion        int64
	managedHandles        event.ManagedHandles
	managedRuntimeHandles [2]event.Handle

	entityEventTab                 entityEventTab
	entityComponentManagerEventTab entityComponentManagerEventTab
	entityTreeNodeEventTab         entityTreeNodeEventTab
}

// GetId 获取实体Id
func (entity *EntityBehavior) GetId() uid.Id {
	return entity.options.PersistId
}

// GetPT 获取实体原型
func (entity *EntityBehavior) GetPT() EntityPT {
	if entity.prototype == nil {
		return noneEntityPT
	}
	return entity.prototype
}

// GetScope 获取可访问作用域
func (entity *EntityBehavior) GetScope() Scope {
	return entity.options.Scope
}

// GetState 获取实体状态
func (entity *EntityBehavior) GetState() EntityState {
	return entity.state
}

// GetReflected 获取反射值
func (entity *EntityBehavior) GetReflected() reflect.Value {
	if entity.reflected.IsValid() {
		return entity.reflected
	}
	entity.reflected = reflect.ValueOf(entity.getInstance())
	return entity.reflected
}

// GetMeta 获取Meta信息
func (entity *EntityBehavior) GetMeta() meta.Meta {
	return entity.options.Meta
}

// Managed 托管事件句柄
func (entity *EntityBehavior) Managed() *event.ManagedHandles {
	return &entity.managedHandles
}

// Destroy 销毁
func (entity *EntityBehavior) Destroy() {
	entity.reentrancyGuard.Call(entityReentrancyGuard_Destroy, func() {
		if entity.state > EntityState_Alive {
			return
		}
		_EmitEventEntityDestroy(entity, entity.getInstance())
	})
}

// EventEntityDestroy 事件：实体销毁
func (entity *EntityBehavior) EventEntityDestroy() event.IEvent {
	return entity.entityEventTab.EventEntityDestroy()
}

// GetCurrentContext 获取当前上下文
func (entity *EntityBehavior) GetCurrentContext() iface.Cache {
	return entity.context
}

// GetConcurrentContext 解析线程安全的上下文
func (entity *EntityBehavior) GetConcurrentContext() iface.Cache {
	return entity.context
}

// GetInstanceFaceCache 支持重新解释类型
func (entity *EntityBehavior) GetInstanceFaceCache() iface.Cache {
	return entity.options.InstanceFace.Cache
}

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "prototype":%q}`, entity.GetId(), entity.GetPT().Prototype())
}

func (entity *EntityBehavior) init(options EntityOptions) {
	entity.options = options

	if entity.options.InstanceFace.IsNil() {
		entity.options.InstanceFace = iface.MakeFaceT[Entity](entity)
	}

	entity.setState(EntityState_Birth)
}

func (entity *EntityBehavior) withContext(ctx context.Context) {
	entity.Context, entity.terminate = context.WithCancel(ctx)
	entity.terminated = async.MakeAsyncRet()
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

func (entity *EntityBehavior) setContext(ctx iface.Cache) {
	entity.context = ctx
}

func (entity *EntityBehavior) setState(state EntityState) {
	if entity.processedStateBits.Is(int(state)) {
		return
	}

	entity.state = state
	entity.processedStateBits.Set(int(state), true)

	switch entity.state {
	case EntityState_Death:
		entity.terminate()
		entity.entityEventTab.SetEnable(false)
		entity.entityComponentManagerEventTab.SetEnable(false)
		entity.entityTreeNodeEventTab.SetEnable(false)
	case EntityState_Destroyed:
		entity.managedHandles.UnbindAllEventHandles()
		entity.managedUnbindRuntimeHandles()
		async.Return(entity.terminated, async.VoidRet)
	}
}

func (entity *EntityBehavior) setReflected(v reflect.Value) {
	entity.reflected = v
}

func (entity *EntityBehavior) getCallingStateBits() *generic.Bits16 {
	return &entity.callingStateBits
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

func (entity *EntityBehavior) getInstance() Entity {
	return entity.options.InstanceFace.Iface
}
