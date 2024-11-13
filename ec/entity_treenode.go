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

import "git.golaxy.org/core/event"

type iTreeNode interface {
	// GetTreeNodeState 获取实体树节点状态
	GetTreeNodeState() TreeNodeState
	// GetTreeNodeParent 获取在实体树中的父实体
	GetTreeNodeParent() (Entity, bool)

	setTreeNodeState(state TreeNodeState)
	setTreeNodeParent(parent Entity)
	enterParentNode()
	leaveParentNode()

	IEntityTreeNodeEventTab
}

// GetTreeNodeState 获取实体树节点状态
func (entity *EntityBehavior) GetTreeNodeState() TreeNodeState {
	return entity.treeNodeState
}

// GetTreeNodeParent 获取在实体树中的父实体
func (entity *EntityBehavior) GetTreeNodeParent() (Entity, bool) {
	return entity.treeNodeParent, entity.treeNodeParent != nil
}

// EventTreeNodeAddChild 事件：实体节点添加子实体
func (entity *EntityBehavior) EventTreeNodeAddChild() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeAddChild()
}

// EventTreeNodeRemoveChild 事件：实体节点删除子实体
func (entity *EntityBehavior) EventTreeNodeRemoveChild() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeRemoveChild()
}

// EventTreeNodeEnterParent 事件：实体加入父实体节点
func (entity *EntityBehavior) EventTreeNodeEnterParent() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeEnterParent()
}

// EventTreeNodeLeaveParent 事件：实体离开父实体节点
func (entity *EntityBehavior) EventTreeNodeLeaveParent() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeLeaveParent()
}

func (entity *EntityBehavior) setTreeNodeState(state TreeNodeState) {
	entity.treeNodeState = state
}

func (entity *EntityBehavior) setTreeNodeParent(parent Entity) {
	entity.treeNodeParent = parent
}

func (entity *EntityBehavior) enterParentNode() {
	if entity.treeNodeParent == nil {
		return
	}

	_EmitEventTreeNodeEnterParentWithInterrupt(entity, func(child, parent Entity) bool {
		return child.GetState() > EntityState_Alive || parent.GetState() > EntityState_Alive
	}, entity.opts.InstanceFace.Iface, entity.treeNodeParent)

	if entity.treeNodeParent == nil {
		return
	}

	_EmitEventTreeNodeAddChildWithInterrupt(entity.treeNodeParent, func(parent, child Entity) bool {
		return parent.GetState() > EntityState_Alive || child.GetState() > EntityState_Alive
	}, entity.treeNodeParent, entity.opts.InstanceFace.Iface)
}

func (entity *EntityBehavior) leaveParentNode() {
	if entity.treeNodeParent == nil {
		return
	}

	_EmitEventTreeNodeRemoveChildWithInterrupt(entity.treeNodeParent, func(parent, child Entity) bool {
		return parent.GetState() >= EntityState_Death || child.GetState() >= EntityState_Death
	}, entity.treeNodeParent, entity.opts.InstanceFace.Iface)

	if entity.treeNodeParent == nil {
		return
	}

	_EmitEventTreeNodeLeaveParentWithInterrupt(entity, func(child, parent Entity) bool {
		return child.GetState() >= EntityState_Death || parent.GetState() >= EntityState_Death
	}, entity.opts.InstanceFace.Iface, entity.treeNodeParent)
}
