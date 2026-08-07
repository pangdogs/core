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

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// UnsafeEntity 暴露实体状态机、组件表与运行时句柄等框架内部能力。
//
// Deprecated: 仅供框架内部使用。
func UnsafeEntity(entity Entity) _UnsafeEntity {
	return _UnsafeEntity{
		Entity: entity,
	}
}

type _UnsafeEntity struct {
	Entity
}

// WithContext 创建由 ctx 派生的实体生命周期上下文。
func (u _UnsafeEntity) WithContext(ctx context.Context) {
	u.withContext(ctx)
}

// Options 返回实体当前持有的构造选项。
func (u _UnsafeEntity) Options() *EntityOptions {
	return u.getOptions()
}

// Instance 返回实际实体实例。
func (u _UnsafeEntity) Instance() Entity {
	return u.getInstance()
}

// SetId 设置实体 ID。
func (u _UnsafeEntity) SetId(id uid.Id) {
	u.setId(id)
}

// SetPT 绑定实体原型。
func (u _UnsafeEntity) SetPT(prototype EntityPT) {
	u.setPT(prototype)
}

// SetContext 设置实体所属 Runtime 的上下文缓存。
func (u _UnsafeEntity) SetContext(ctx iface.Cache) {
	u.setContext(ctx)
}

// SetState 推进实体生命周期状态。
func (u _UnsafeEntity) SetState(state EntityState) {
	u.setState(state)
}

// SetReflected 缓存实际实体实例的反射值。
func (u _UnsafeEntity) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// ProcessedStateBits 返回生命周期阶段的已处理标志位。
func (u _UnsafeEntity) ProcessedStateBits() *generic.Bits16 {
	return u.getProcessedStateBits()
}

// EnteredHandle 返回实体在 Runtime 实体表中的位置与版本。
func (u _UnsafeEntity) EnteredHandle() (int, int64) {
	return u.getEnteredHandle()
}

// SetEnteredHandle 设置实体在 Runtime 实体表中的位置与版本。
func (u _UnsafeEntity) SetEnteredHandle(idx int, ver int64) {
	u.setEnteredHandle(idx, ver)
}

// ManagedRuntimeUpdateHandle 替换并托管 Runtime 更新事件句柄。
func (u _UnsafeEntity) ManagedRuntimeUpdateHandle(updateHandle event.Handle) {
	u.managedRuntimeUpdateHandle(updateHandle)
}

// ManagedRuntimeLateUpdateHandle 替换并托管 Runtime 后置更新事件句柄。
func (u _UnsafeEntity) ManagedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	u.managedRuntimeLateUpdateHandle(lateUpdateHandle)
}

// ManagedUnbindRuntimeHandles 解绑全部托管的 Runtime 更新事件句柄。
func (u _UnsafeEntity) ManagedUnbindRuntimeHandles() {
	u.managedUnbindRuntimeHandles()
}

// Version 返回实体组件表的当前版本。
func (u _UnsafeEntity) Version() int64 {
	return u.getVersion()
}

// ComponentNameIndex 返回实体组件名称索引。
func (u _UnsafeEntity) ComponentNameIndex() *generic.SliceMap[string, int] {
	return u.getComponentNameIndex()
}

// ComponentList 返回实体组件槽表。
func (u _UnsafeEntity) ComponentList() *generic.FreeList[Component] {
	return u.getComponentList()
}

// SetTreeNodeState 设置实体树节点状态。
func (u _UnsafeEntity) SetTreeNodeState(state TreeNodeState) {
	u.setTreeNodeState(state)
}

// EmitEventTreeNodeAddChild 派发直接子实体添加事件。
func (u _UnsafeEntity) EmitEventTreeNodeAddChild(childId uid.Id) {
	u.emitEventTreeNodeAddChild(childId)
}

// EmitEventTreeNodeRemoveChild 派发直接子实体移除事件。
func (u _UnsafeEntity) EmitEventTreeNodeRemoveChild(childId uid.Id) {
	u.emitEventTreeNodeRemoveChild(childId)
}

// EmitEventTreeNodeAttachParent 派发接入父节点事件。
func (u _UnsafeEntity) EmitEventTreeNodeAttachParent(parentId uid.Id) {
	u.emitEventTreeNodeAttachParent(parentId)
}

// EmitEventTreeNodeDetachParent 派发脱离父节点事件。
func (u _UnsafeEntity) EmitEventTreeNodeDetachParent(parentId uid.Id) {
	u.emitEventTreeNodeDetachParent(parentId)
}

// EmitEventTreeNodeMoveTo 派发父节点变更事件。
func (u _UnsafeEntity) EmitEventTreeNodeMoveTo(fromParentId, toParentId uid.Id) {
	u.emitEventTreeNodeMoveTo(fromParentId, toParentId)
}
