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

// Deprecated: UnsafeEntity 访问实体内部函数
func UnsafeEntity(entity Entity) _UnsafeEntity {
	return _UnsafeEntity{
		Entity: entity,
	}
}

type _UnsafeEntity struct {
	Entity
}

// WithContext 传递上下文
func (u _UnsafeEntity) WithContext(ctx context.Context) {
	u.withContext(ctx)
}

// Options 获取实体所有选项
func (u _UnsafeEntity) Options() *EntityOptions {
	return u.getOptions()
}

// SetId 设置Id
func (u _UnsafeEntity) SetId(id uid.Id) {
	u.setId(id)
}

// SetPT 设置实体原型信息
func (u _UnsafeEntity) SetPT(prototype EntityPT) {
	u.setPT(prototype)
}

// SetContext 设置上下文
func (u _UnsafeEntity) SetContext(ctx iface.Cache) {
	u.setContext(ctx)
}

// SetState 设置状态
func (u _UnsafeEntity) SetState(state EntityState) {
	u.setState(state)
}

// SetReflected 设置反射值
func (u _UnsafeEntity) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// ProcessedStateBits 获取已处理状态标志位
func (u _UnsafeEntity) ProcessedStateBits() *generic.Bits16 {
	return u.getProcessedStateBits()
}

// EnteredHandle 获取加入运行时时的句柄
func (u _UnsafeEntity) EnteredHandle() (int, int64) {
	return u.getEnteredHandle()
}

// SetEnteredHandle 设置加入运行时时的句柄
func (u _UnsafeEntity) SetEnteredHandle(idx int, ver int64) {
	u.setEnteredHandle(idx, ver)
}

// ManagedRuntimeUpdateHandle 托管运行时更新句柄
func (u _UnsafeEntity) ManagedRuntimeUpdateHandle(updateHandle event.Handle) {
	u.managedRuntimeUpdateHandle(updateHandle)
}

// ManagedRuntimeLateUpdateHandle 托管运行时延迟更新句柄
func (u _UnsafeEntity) ManagedRuntimeLateUpdateHandle(lateUpdateHandle event.Handle) {
	u.managedRuntimeLateUpdateHandle(lateUpdateHandle)
}

// ManagedUnbindRuntimeHandles 解绑定托管的运行时句柄
func (u _UnsafeEntity) ManagedUnbindRuntimeHandles() {
	u.managedUnbindRuntimeHandles()
}

// Version 获取实体组件变化版本号
func (u _UnsafeEntity) Version() int64 {
	return u.getVersion()
}

// ComponentNameIndex 获取实体组件名称索引
func (u _UnsafeEntity) ComponentNameIndex() *generic.SliceMap[string, int] {
	return u.getComponentNameIndex()
}

// ComponentList 获取实体组件链表
func (u _UnsafeEntity) ComponentList() *generic.FreeList[Component] {
	return u.getComponentList()
}

// SetTreeNodeState 设置实体树节点状态
func (u _UnsafeEntity) SetTreeNodeState(state TreeNodeState) {
	u.setTreeNodeState(state)
}

// EmitEventTreeNodeAddChild 发送实体树节点添加子实体事件
func (u _UnsafeEntity) EmitEventTreeNodeAddChild(childId uid.Id) {
	u.emitEventTreeNodeAddChild(childId)
}

// EmitEventTreeNodeRemoveChild 发送实体树节点删除子实体事件
func (u _UnsafeEntity) EmitEventTreeNodeRemoveChild(childId uid.Id) {
	u.emitEventTreeNodeRemoveChild(childId)
}

// EmitEventTreeNodeAttachParent 发送实体树节点加入父节点事件
func (u _UnsafeEntity) EmitEventTreeNodeAttachParent(parentId uid.Id) {
	u.emitEventTreeNodeAttachParent(parentId)
}

// EmitEventTreeNodeDetachParent 发送实体树节点离开父节点事件
func (u _UnsafeEntity) EmitEventTreeNodeDetachParent(parentId uid.Id) {
	u.emitEventTreeNodeDetachParent(parentId)
}

// EmitEventTreeNodeMoveTo 发送实体树节点切换父节点事件
func (u _UnsafeEntity) EmitEventTreeNodeMoveTo(fromParentId, toParentId uid.Id) {
	u.emitEventTreeNodeMoveTo(fromParentId, toParentId)
}
