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
	"sync"

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
	return UnsafeNewEntity(option.New(With.Default(), settings...))
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

	// Id 获取实体Id
	Id() uid.Id
	// PT 获取实体原型信息
	PT() EntityPT
	// Scope 获取可访问作用域
	Scope() Scope
	// State 获取实体状态
	State() EntityState
	// Reflected 获取反射值
	Reflected() reflect.Value
	// Meta 获取Meta信息
	Meta() meta.Meta
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
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	enteredIndex          int
	enteredVersion        int64
	managedHandles        event.ManagedHandles
	managedRuntimeHandles [2]event.Handle
	stringerOnce          sync.Once
	stringerCache         string

	entityEventTab                 entityEventTab
	entityComponentManagerEventTab entityComponentManagerEventTab
	entityTreeNodeEventTab         entityTreeNodeEventTab
}

// Id 获取实体Id
func (entity *EntityBehavior) Id() uid.Id {
	return entity.options.PersistId
}

// PT 获取实体原型
func (entity *EntityBehavior) PT() EntityPT {
	if entity.prototype == nil {
		return noneEntityPT
	}
	return entity.prototype
}

// Scope 获取可访问作用域
func (entity *EntityBehavior) Scope() Scope {
	return entity.options.Scope
}

// State 获取实体状态
func (entity *EntityBehavior) State() EntityState {
	return entity.state
}

// Reflected 获取反射值
func (entity *EntityBehavior) Reflected() reflect.Value {
	if entity.reflected.IsValid() {
		return entity.reflected
	}
	entity.reflected = reflect.ValueOf(entity.getInstance())
	return entity.reflected
}

// Meta 获取Meta信息
func (entity *EntityBehavior) Meta() meta.Meta {
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

// CurrentContext 获取当前上下文
func (entity *EntityBehavior) CurrentContext() iface.Cache {
	return entity.context
}

// ConcurrentContext 解析线程安全的上下文
func (entity *EntityBehavior) ConcurrentContext() iface.Cache {
	return entity.context
}

// InstanceFaceCache 支持重新解释类型
func (entity *EntityBehavior) InstanceFaceCache() iface.Cache {
	return entity.options.InstanceFace.Cache
}

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	entity.stringerOnce.Do(func() {
		entity.stringerCache = fmt.Sprintf(`{"id":%q,"prototype":%q}`, entity.Id(), entity.PT().Prototype())
	})
	return entity.stringerCache
}

func (entity *EntityBehavior) init(options EntityOptions) {
	entity.options = options

	if entity.options.InstanceFace.IsNil() {
		entity.options.InstanceFace = iface.NewFaceT[Entity](entity)
	}
}

func (entity *EntityBehavior) withContext(ctx context.Context) {
	entity.Context, entity.terminate = context.WithCancel(ctx)
	entity.terminated = async.NewAsyncRet()
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
	if entity.state >= state {
		return
	}

	entity.state = state

	switch entity.state {
	case EntityState_Death:
		entity.terminate()
		entity.entityEventTab.SetEnabled(false)
		entity.entityComponentManagerEventTab.SetEnabled(false)
		entity.entityTreeNodeEventTab.SetEnabled(false)
	case EntityState_Destroyed:
		entity.managedHandles.UnbindAllEventHandles()
		entity.managedUnbindRuntimeHandles()
		async.Return(entity.terminated, async.VoidRet)
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

func (entity *EntityBehavior) getInstance() Entity {
	return entity.options.InstanceFace.Iface
}
