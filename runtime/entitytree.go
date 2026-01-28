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

package runtime

import (
	"fmt"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
)

var (
	// ForestNodeId 实体树森林节点ID
	ForestNodeId = uid.From("d5rh7sbr1n96c63fs3vg")
	// forestNodeIdx 实体树森林节点索引
	forestNodeIdx = -1
)

// EntityTree 实体树接口
type EntityTree interface {
	corectx.CurrentContextProvider

	// MakeRoot 创建根节点
	MakeRoot(entityId uid.Id) error
	// AddChild 新增子节点
	AddChild(parentId, childId uid.Id) error
	// RemoveNode 删除子节点，会后序遍历递归删除所有子节点
	RemoveNode(childId uid.Id) error
	// DetachNode 脱离父节点，成为根节点
	DetachNode(childId uid.Id) error
	// MoveNode 修改父节点
	MoveNode(childId, parentId uid.Id) error
	// IsFreedom 是否是自由节点
	IsFreedom(entityId uid.Id) (bool, error)
	// IsRoot 是否是根节点
	IsRoot(entityId uid.Id) (bool, error)
	// IsLeaf 是否是叶子节点
	IsLeaf(entityId uid.Id) (bool, error)
	// GetParent 获取父实体
	GetParent(childId uid.Id) (ec.Entity, error)
	// RangeChildren 遍历所有子节点
	RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error
	// EachChildren 遍历每个子节点
	EachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error
	// ReversedRangeChildren 反向遍历所有子节点
	ReversedRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error
	// ReversedEachChildren 反向遍历每个子节点
	ReversedEachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error
	// FilterChildren 过滤并获取子节点
	FilterChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error)
	// ListChildren 获取所有子节点
	ListChildren(parentId uid.Id) ([]ec.Entity, error)
	// CountChildren 获取子节点数量
	CountChildren(parentId uid.Id) (int, error)

	IEntityTreeEventTab
}

// MakeRoot 创建根节点
func (mgr *_EntityManagerBehavior) MakeRoot(entityId uid.Id) error {
	return mgr.AddChild(ForestNodeId, entityId)
}

// AddChild 新增子节点
func (mgr *_EntityManagerBehavior) AddChild(parentId, childId uid.Id) error {
	parentSlotIdx, parentTreeNode := mgr.getTreeNode(parentId)
	if parentSlotIdx < 0 {
		if parentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not exists", ErrEntityTree, parentId)
		}
	} else {
		if parentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
		}

		parentEntity := mgr.entityList.Get(parentSlotIdx).V

		if parentEntity.GetState() < ec.EntityState_Awake || parentEntity.GetState() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentId, parentEntity.GetState())
		}
	}

	childSlotIdx, childTreeNode := mgr.getTreeNode(childId)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childId)
	}
	if childTreeNode != nil {
		return fmt.Errorf("%w: child entity %q already in the entity-tree", ErrEntityTree, childId)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.GetState() < ec.EntityState_Awake || childEntity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.GetState())
	}

	if childEntity.GetTreeNodeState() != ec.TreeNodeState_Freedom {
		return fmt.Errorf("%w: child entity %q is in an unexpected tree node state %q", ErrEntityTree, childId, childEntity.GetTreeNodeState())
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Attaching)

	treeNode := &_TreeNode{parent: parentSlotIdx}
	mgr.entityTreeNodes[childSlotIdx] = treeNode
	attachedSlot := parentTreeNode.children.PushBack(childSlotIdx)
	treeNode.attachedIndex = attachedSlot.Index()
	treeNode.attachedVersion = attachedSlot.Version()

	var parentEntity ec.Entity
	if parentSlotIdx >= 0 {
		parentEntity = mgr.entityList.Get(parentSlotIdx).V
	}

	{
		caller := makeTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			_EmitEventEntityTreeAddNode(mgr, mgr, parentId, childId)
		}) {
			return nil
		}

		if parentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(parentEntity).EmitEventTreeNodeAddChild(childId)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeAttachParent(parentId)
		}) {
			return nil
		}
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Attached)

	return nil
}

// RemoveNode 删除子节点，会后序遍历递归删除所有子节点
func (mgr *_EntityManagerBehavior) RemoveNode(childId uid.Id) error {
	childSlotIdx, childTreeNode := mgr.getTreeNode(childId)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childId)
	}
	if childTreeNode == nil {
		return fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childId)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.GetState() < ec.EntityState_Awake || childEntity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.GetState())
	}

	if childEntity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childId, childEntity.GetTreeNodeState())
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	parentId := ForestNodeId
	parentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var parentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		parentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		parentEntity = mgr.entityList.Get(childTreeNode.parent).V
		parentId = parentEntity.GetId()
	}

	{
		caller := makeTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
				entity := mgr.entityList.Get(slot.V).V
				mgr.RemoveNode(entity.GetId())
			})
		}) {
			return nil
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(parentId)
		}) {
			return nil
		}

		if parentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(parentEntity).EmitEventTreeNodeRemoveChild(childId)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			_EmitEventEntityTreeRemoveNode(mgr, mgr, parentId, childId)
		}) {
			return nil
		}
	}

	delete(mgr.entityTreeNodes, childSlotIdx)
	parentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Freedom)

	return nil
}

// DetachNode 脱离父节点，成为根节点
func (mgr *_EntityManagerBehavior) DetachNode(childId uid.Id) error {
	return mgr.MoveNode(childId, ForestNodeId)
}

// MoveNode 修改父节点
func (mgr *_EntityManagerBehavior) MoveNode(childId, parentId uid.Id) error {
	toParentSlotIdx, toParentTreeNode := mgr.getTreeNode(parentId)
	if toParentSlotIdx < 0 {
		if toParentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not exists", ErrEntityTree, parentId)
		}
	} else {
		if toParentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
		}

		toParentEntity := mgr.entityList.Get(toParentSlotIdx).V

		if toParentEntity.GetState() < ec.EntityState_Awake || toParentEntity.GetState() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentId, toParentEntity.GetState())
		}
	}

	childSlotIdx, childTreeNode := mgr.getTreeNode(childId)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childId)
	}
	if childTreeNode == nil {
		return fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childId)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.GetState() < ec.EntityState_Awake || childEntity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.GetState())
	}

	if childEntity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childId, childEntity.GetTreeNodeState())
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Moving)

	fromParentId := ForestNodeId
	fromParentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var fromParentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		fromParentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		fromParentEntity = mgr.entityList.Get(childTreeNode.parent).V
		fromParentId = fromParentEntity.GetId()
	}

	toParentId := parentId
	var toParentEntity ec.Entity
	if toParentSlotIdx >= 0 {
		toParentEntity = mgr.entityList.Get(toParentSlotIdx).V
	}

	{
		caller := makeTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(fromParentId)
		}) {
			return nil
		}

		if fromParentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(fromParentEntity).EmitEventTreeNodeRemoveChild(childId)
			}) {
				return nil
			}
		}

		fromParentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)
		attachedSlot := toParentTreeNode.children.PushBack(childSlotIdx)
		childTreeNode.parent = toParentSlotIdx
		childTreeNode.attachedIndex = attachedSlot.Index()
		childTreeNode.attachedVersion = attachedSlot.Version()

		if !caller.Call(func() {
			_EmitEventEntityTreeMoveNode(mgr, mgr, childId, fromParentId, toParentId)
		}) {
			return nil
		}

		if toParentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(toParentEntity).EmitEventTreeNodeAddChild(childId)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeAttachParent(parentId)
		}) {
			return nil
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeMoveTo(fromParentId, toParentId)
		}) {
			return nil
		}
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Attached)

	return nil
}

// IsFreedom 是否是自由节点
func (mgr *_EntityManagerBehavior) IsFreedom(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	return treeNode == nil, nil
}

// IsRoot 是否是根节点
func (mgr *_EntityManagerBehavior) IsRoot(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityId)
	}
	return treeNode.parent == forestNodeIdx, nil
}

// IsLeaf 是否是叶子节点
func (mgr *_EntityManagerBehavior) IsLeaf(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityId)
	}
	return treeNode.children.Len()-treeNode.children.OrphanCount() <= 0, nil
}

// GetParent 获取父实体
func (mgr *_EntityManagerBehavior) GetParent(childId uid.Id) (ec.Entity, error) {
	slotIdx, treeNode := mgr.getTreeNode(childId)
	if slotIdx < 0 {
		return nil, fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childId)
	}
	if treeNode == nil {
		return nil, fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childId)
	}
	if treeNode.parent == forestNodeIdx {
		return nil, fmt.Errorf("%w: child entity %q is root node", ErrEntityTree, childId)
	}
	return mgr.entityList.Get(treeNode.parent).V, nil
}

// RangeChildren 遍历所有子节点
func (mgr *_EntityManagerBehavior) RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.Traversal(func(slot *generic.FreeSlot[int]) bool {
		return fun(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// EachChildren 遍历每个子节点
func (mgr *_EntityManagerBehavior) EachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedRangeChildren 反向遍历所有子节点
func (mgr *_EntityManagerBehavior) ReversedRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.ReversedTraversal(func(slot *generic.FreeSlot[int]) bool {
		return fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedEachChildren 反向遍历每个子节点
func (mgr *_EntityManagerBehavior) ReversedEachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// FilterChildren 过滤并获取子节点
func (mgr *_EntityManagerBehavior) FilterChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error) {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return nil, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}

	var entities []ec.Entity

	ver := treeNode.children.Version()
	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		if slot.Version() > ver {
			return
		}
		entity := mgr.entityList.Get(slot.V).V
		if fun.UnsafeCall(entity) {
			entities = append(entities, entity)
		}
	})

	return entities, nil
}

// ListChildren 获取所有子节点
func (mgr *_EntityManagerBehavior) ListChildren(parentId uid.Id) ([]ec.Entity, error) {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return nil, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}

	entities := make([]ec.Entity, 0, treeNode.children.Len()-treeNode.children.OrphanCount())

	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		entities = append(entities, mgr.entityList.Get(slot.V).V)
	})

	return entities, nil
}

// CountChildren 获取子节点数量
func (mgr *_EntityManagerBehavior) CountChildren(parentId uid.Id) (int, error) {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return 0, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	return treeNode.children.Len() - treeNode.children.OrphanCount(), nil
}

func (mgr *_EntityManagerBehavior) onEntityDestroyRemoveNode(childId uid.Id) {
	childSlotIdx, childTreeNode := mgr.getTreeNode(childId)
	if childSlotIdx < 0 {
		return
	}
	if childTreeNode == nil {
		return
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	parentId := ForestNodeId
	parentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var parentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		parentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		parentEntity = mgr.entityList.Get(childTreeNode.parent).V
		parentId = parentEntity.GetId()
	}

	childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		entity := mgr.entityList.Get(slot.V).V
		mgr.onEntityDestroyRemoveNode(entity.GetId())
	})

	ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(parentId)

	if parentEntity != nil {
		ec.UnsafeEntity(parentEntity).EmitEventTreeNodeRemoveChild(childId)
	}

	_EmitEventEntityTreeRemoveNode(mgr, mgr, parentId, childId)

	delete(mgr.entityTreeNodes, childSlotIdx)
	parentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Freedom)
}

func (mgr *_EntityManagerBehavior) getTreeNode(entityId uid.Id) (int, *_TreeNode) {
	if entityId == ForestNodeId {
		return forestNodeIdx, mgr.entityTreeNodes[forestNodeIdx]
	}

	slotIdx, ok := mgr.entityIdIndex[entityId]
	if !ok {
		return -2, nil
	}

	treeNode, ok := mgr.entityTreeNodes[slotIdx]
	if !ok {
		return slotIdx, nil
	}

	return slotIdx, treeNode
}

func makeTreeNodeCaller(entity ec.Entity) _TreeNodeCaller {
	return _TreeNodeCaller{entity: entity, state: entity.GetTreeNodeState()}
}

type _TreeNodeCaller struct {
	entity ec.Entity
	state  ec.TreeNodeState
}

func (c _TreeNodeCaller) Call(fun func()) bool {
	if c.entity.GetTreeNodeState() != c.state {
		return false
	}

	fun()

	return c.entity.GetTreeNodeState() == c.state
}
