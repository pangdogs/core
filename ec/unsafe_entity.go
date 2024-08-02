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
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
	"reflect"
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

// Init 初始化
func (ue _UnsafeEntity) Init(opts EntityOptions) {
	ue.init(opts)
}

// GetOptions 获取实体所有选项
func (ue _UnsafeEntity) GetOptions() *EntityOptions {
	return ue.getOptions()
}

// SetId 设置Id
func (ue _UnsafeEntity) SetId(id uid.Id) {
	ue.setId(id)
}

// SetContext 设置上下文
func (ue _UnsafeEntity) SetContext(ctx iface.Cache) {
	ue.setContext(ctx)
}

// GetVersion 获取组件列表变化版本号
func (ue _UnsafeEntity) GetVersion() int64 {
	return ue.getVersion()
}

// SetState 设置状态
func (ue _UnsafeEntity) SetState(state EntityState) {
	ue.setState(state)
}

// SetReflected 设置反射值
func (ue _UnsafeEntity) SetReflected(v reflect.Value) {
	ue.setReflected(v)
}

// SetTreeNodeState 设置实体树节点状态
func (ue _UnsafeEntity) SetTreeNodeState(state TreeNodeState) {
	ue.setTreeNodeState(state)
}

// SetTreeNodeParent 设置在实体树中的父实体
func (ue _UnsafeEntity) SetTreeNodeParent(parent Entity) {
	ue.setTreeNodeParent(parent)
}

// EnterParentNode 进入父节点
func (ue _UnsafeEntity) EnterParentNode() {
	ue.enterParentNode()
}

// LeaveParentNode 离开父节点
func (ue _UnsafeEntity) LeaveParentNode() {
	ue.leaveParentNode()
}

// EventEntityDestroySelf 事件：实体销毁自身
func (ue _UnsafeEntity) EventEntityDestroySelf() event.IEvent {
	return ue.eventEntityDestroySelf()
}

// CleanManagedHooks 清理所有的托管hook
func (ue _UnsafeEntity) CleanManagedHooks() {
	ue.cleanManagedHooks()
}
