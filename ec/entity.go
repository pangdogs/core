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
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/types"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// NewEntity 创建实体
func NewEntity(settings ...option.Setting[EntityOptions]) Entity {
	return UnsafeNewEntity(option.Make(With.Default(), settings...))
}

// Deprecated: UnsafeNewEntity 内部创建实体
func UnsafeNewEntity(options EntityOptions) Entity {
	if !options.InstanceFace.IsNil() {
		options.InstanceFace.Iface.init(options)
		return options.InstanceFace.Iface
	}

	e := &EntityBehavior{}
	e.init(options)

	return e.opts.InstanceFace.Iface
}

// Entity 实体接口
type Entity interface {
	iEntity
	iConcurrentEntity
	iContext
	iComponentManager
	iTreeNode
	ictx.CurrentContextProvider
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
	// ManagedAddHooks 托管事件钩子（event.Hook），在实体销毁时自动解绑定
	ManagedAddHooks(hooks ...event.Hook)
	// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在实体销毁时自动解绑定
	ManagedAddTagHooks(tag string, hooks ...event.Hook)
	// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
	ManagedGetTagHooks(tag string) []event.Hook
	// ManagedUnbindTagHooks 根据标签解绑定托管的事件钩子（event.Hook）
	ManagedUnbindTagHooks(tag string)
	// DestroySelf 销毁自身
	DestroySelf()

	IEntityEventTab
}

type iEntity interface {
	init(opts EntityOptions)
	withContext(ctx context.Context)
	getOptions() *EntityOptions
	setId(id uid.Id)
	setPT(prototype EntityPT)
	setContext(ctx iface.Cache)
	getVersion() int64
	setState(state EntityState)
	setReflected(v reflect.Value)
	getProcessedStateBits() *types.Bits16
	managedUnbindAllHooks()
}

// EntityBehavior 实体行为，在扩展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	context.Context
	terminate          context.CancelFunc
	terminated         chan async.Ret
	opts               EntityOptions
	prototype          EntityPT
	context            iface.Cache
	componentNameIndex generic.SliceMap[string, *generic.Node[Component]]
	components         generic.List[Component]
	state              EntityState
	reflected          reflect.Value
	treeNodeState      TreeNodeState
	treeNodeParent     Entity
	callingStateBits   types.Bits16
	processedStateBits types.Bits16
	managedHooks       []event.Hook
	managedTagHooks    generic.SliceMap[string, []event.Hook]

	entityEventTab                 entityEventTab
	entityComponentManagerEventTab entityComponentManagerEventTab
	entityTreeNodeEventTab         entityTreeNodeEventTab
}

// GetId 获取实体Id
func (entity *EntityBehavior) GetId() uid.Id {
	return entity.opts.PersistId
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
	return entity.opts.Scope
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
	entity.reflected = reflect.ValueOf(entity.opts.InstanceFace.Iface)
	return entity.reflected
}

// GetMeta 获取Meta信息
func (entity *EntityBehavior) GetMeta() meta.Meta {
	return entity.opts.Meta
}

// DestroySelf 销毁自身
func (entity *EntityBehavior) DestroySelf() {
	_EmitEventEntityDestroySelf(entity, entity.opts.InstanceFace.Iface)
}

// EventEntityDestroySelf 事件：实体销毁自身
func (entity *EntityBehavior) EventEntityDestroySelf() event.IEvent {
	return entity.entityEventTab.EventEntityDestroySelf()
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
	return entity.opts.InstanceFace.Cache
}

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "prototype":%q}`, entity.GetId(), entity.GetPT().Prototype())
}

func (entity *EntityBehavior) init(opts EntityOptions) {
	entity.opts = opts

	if entity.opts.InstanceFace.IsNil() {
		entity.opts.InstanceFace = iface.MakeFaceT[Entity](entity)
	}

	entity.entityEventTab.Init(false, nil, event.EventRecursion_Allow)
	entity.entityComponentManagerEventTab.Init(false, nil, event.EventRecursion_Allow)
	entity.entityTreeNodeEventTab.Init(false, nil, event.EventRecursion_Allow)

	entity.setState(EntityState_Birth)
}

func (entity *EntityBehavior) withContext(ctx context.Context) {
	entity.Context, entity.terminate = context.WithCancel(ctx)
	entity.terminated = async.MakeAsyncRet()
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setId(id uid.Id) {
	entity.opts.PersistId = id
}

func (entity *EntityBehavior) setPT(prototype EntityPT) {
	entity.prototype = prototype
}

func (entity *EntityBehavior) setContext(ctx iface.Cache) {
	entity.context = ctx
}

func (entity *EntityBehavior) getVersion() int64 {
	return entity.components.Version()
}

func (entity *EntityBehavior) setState(state EntityState) {
	if entity.processedStateBits.Is(int8(state)) {
		return
	}

	entity.state = state
	entity.processedStateBits.Set(int8(state), true)

	switch entity.state {
	case EntityState_Death:
		entity.terminate()
		entity.entityEventTab.Disable()
		entity.entityComponentManagerEventTab.Disable()
		entity.entityTreeNodeEventTab.Disable()
	case EntityState_Destroyed:
		entity.managedUnbindAllHooks()
		async.Return(entity.terminated, async.VoidRet)
	}
}

func (entity *EntityBehavior) setReflected(v reflect.Value) {
	entity.reflected = v
}

func (entity *EntityBehavior) getProcessedStateBits() *types.Bits16 {
	return &entity.processedStateBits
}
