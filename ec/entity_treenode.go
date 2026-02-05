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
	"git.golaxy.org/core/utils/uid"
)

type iTreeNode interface {
	iiTreeNode

	// TreeNodeState 获取实体树节点状态
	TreeNodeState() TreeNodeState

	IEntityTreeNodeEventTab
}

type iiTreeNode interface {
	setTreeNodeState(state TreeNodeState)
	emitEventTreeNodeAddChild(childId uid.Id)
	emitEventTreeNodeRemoveChild(childId uid.Id)
	emitEventTreeNodeAttachParent(parentId uid.Id)
	emitEventTreeNodeDetachParent(parentId uid.Id)
	emitEventTreeNodeMoveTo(fromParentId, toParentId uid.Id)
}

// TreeNodeState 获取实体树节点状态
func (entity *EntityBehavior) TreeNodeState() TreeNodeState {
	return entity.treeNodeState
}

// EventTreeNodeAddChild 事件：实体节点添加子实体
func (entity *EntityBehavior) EventTreeNodeAddChild() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeAddChild()
}

// EventTreeNodeRemoveChild 事件：实体节点删除子实体
func (entity *EntityBehavior) EventTreeNodeRemoveChild() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeRemoveChild()
}

// EventTreeNodeAttachParent 事件：实体加入父实体节点
func (entity *EntityBehavior) EventTreeNodeAttachParent() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeAttachParent()
}

// EventTreeNodeDetachParent 事件：实体离开父实体节点
func (entity *EntityBehavior) EventTreeNodeDetachParent() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeDetachParent()
}

// EventTreeNodeMoveTo 事件：实体切换父节点
func (entity *EntityBehavior) EventTreeNodeMoveTo() event.IEvent {
	return entity.entityTreeNodeEventTab.EventTreeNodeMoveTo()
}

func (entity *EntityBehavior) setTreeNodeState(state TreeNodeState) {
	entity.treeNodeState = state
}

func (entity *EntityBehavior) emitEventTreeNodeAddChild(childId uid.Id) {
	_EmitEventTreeNodeAddChild(entity, entity.getInstance(), childId)
}

func (entity *EntityBehavior) emitEventTreeNodeRemoveChild(childId uid.Id) {
	_EmitEventTreeNodeRemoveChild(entity, entity.getInstance(), childId)
}

func (entity *EntityBehavior) emitEventTreeNodeAttachParent(parentId uid.Id) {
	_EmitEventTreeNodeAttachParent(entity, entity.getInstance(), parentId)
}

func (entity *EntityBehavior) emitEventTreeNodeDetachParent(parentId uid.Id) {
	_EmitEventTreeNodeDetachParent(entity, entity.getInstance(), parentId)
}

func (entity *EntityBehavior) emitEventTreeNodeMoveTo(fromParentId, toParentId uid.Id) {
	_EmitEventTreeNodeMoveTo(entity, entity.getInstance(), fromParentId, toParentId)
}
