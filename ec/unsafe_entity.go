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
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
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
func (u _UnsafeEntity) Init(opts EntityOptions) {
	u.init(opts)
}

// WithContext 传递上下文
func (u _UnsafeEntity) WithContext(ctx context.Context) {
	u.withContext(ctx)
}

// GetOptions 获取实体所有选项
func (u _UnsafeEntity) GetOptions() *EntityOptions {
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

// GetVersion 获取组件列表变化版本号
func (u _UnsafeEntity) GetVersion() int64 {
	return u.getVersion()
}

// SetState 设置状态
func (u _UnsafeEntity) SetState(state EntityState) {
	u.setState(state)
}

// SetReflected 设置反射值
func (u _UnsafeEntity) SetReflected(v reflect.Value) {
	u.setReflected(v)
}

// GetProcessedStateBits 获取已处理状态标志位
func (u _UnsafeEntity) GetProcessedStateBits() *types.Bits16 {
	return u.getProcessedStateBits()
}

// GetComponentNameIndex 获取组件名称索引
func (u _UnsafeEntity) GetComponentNameIndex() *generic.SliceMap[string, *generic.Node[Component]] {
	return u.getComponentNameIndex()
}

// GetComponentList 获取组件列表
func (u _UnsafeEntity) GetComponentList() *generic.List[Component] {
	return u.getComponentList()
}

// RemoveComponentByRef 使用组件引用删除组件
func (u _UnsafeEntity) RemoveComponentByRef(comp Component) {
	u.removeComponentByRef(comp)
}

// SetTreeNodeState 设置实体树节点状态
func (u _UnsafeEntity) SetTreeNodeState(state TreeNodeState) {
	u.setTreeNodeState(state)
}

// SetTreeNodeParent 设置在实体树中的父实体
func (u _UnsafeEntity) SetTreeNodeParent(parent Entity) {
	u.setTreeNodeParent(parent)
}

// EnterParentNode 进入父节点
func (u _UnsafeEntity) EnterParentNode() {
	u.enterParentNode()
}

// LeaveParentNode 离开父节点
func (u _UnsafeEntity) LeaveParentNode() {
	u.leaveParentNode()
}
