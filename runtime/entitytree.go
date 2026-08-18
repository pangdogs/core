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
	// ForestNodeID 是所有根实体共用的虚拟父节点 ID；它是保留哨兵，不应修改。
	ForestNodeID = uid.From("d5rh7sbr1n96c63fs3vg")
	// forestNodeIdx 是虚拟森林节点在内部索引中的保留值。
	forestNodeIdx = -1
)

// EntityTree 管理当前运行时实体之间的父子关系。
// 树操作不提供并发保护，应在所属运行时 goroutine 中执行。
type EntityTree interface {
	corectx.CurrentContextProvider

	// MakeRoot 将自由实体作为根节点加入实体树。
	MakeRoot(entityID uid.ID) error
	// AddChild 将自由实体 childID 挂到 parentID 下。
	AddChild(parentID, childID uid.ID) error
	// RemoveNode 按后序递归移除整个子树的树关系；实体本身不会被销毁。
	RemoveNode(childID uid.ID) error
	// DetachNode 将节点从当前父实体移到虚拟森林节点下，使其成为根节点。
	DetachNode(childID uid.ID) error
	// MoveNode 将节点移动到新的父节点下。
	MoveNode(childID, parentID uid.ID) error
	// IsFree 报告实体是否尚未加入实体树。
	IsFree(entityID uid.ID) (bool, error)
	// IsRoot 报告实体是否直接挂在虚拟森林节点下。
	IsRoot(entityID uid.ID) (bool, error)
	// IsLeaf 报告实体是否没有子节点。
	IsLeaf(entityID uid.ID) (bool, error)
	// GetParent 返回父实体；根节点没有实体父节点，因此返回错误。
	GetParent(childID uid.ID) (ec.Entity, error)
	// RangeChildren 按加入顺序遍历直接子节点，回调返回 false 时停止。
	RangeChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) error
	// EachChildren 按加入顺序遍历全部直接子节点。
	EachChildren(parentID uid.ID, fun generic.Action1[ec.Entity]) error
	// ReversedRangeChildren 按加入顺序逆向遍历直接子节点，回调返回 false 时停止。
	ReversedRangeChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) error
	// ReversedEachChildren 按加入顺序逆向遍历全部直接子节点。
	ReversedEachChildren(parentID uid.ID, fun generic.Action1[ec.Entity]) error
	// FilterChildren 按加入顺序返回符合条件的直接子节点。
	FilterChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error)
	// ListChildren 按加入顺序返回直接子节点切片。
	ListChildren(parentID uid.ID) ([]ec.Entity, error)
	// CountChildren 返回直接子节点数。
	CountChildren(parentID uid.ID) (int, error)

	IEntityTreeEventTab
}

// MakeRoot 将自由实体作为根节点加入实体树。
func (mgr *_EntityManager) MakeRoot(entityID uid.ID) error {
	return mgr.AddChild(ForestNodeID, entityID)
}

// AddChild 将自由实体 childID 挂到 parentID 下。
func (mgr *_EntityManager) AddChild(parentID, childID uid.ID) error {
	parentSlotIdx, parentTreeNode := mgr.getTreeNode(parentID)
	if parentSlotIdx < 0 {
		if parentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not exists", ErrEntityTree, parentID)
		}
	} else {
		if parentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
		}

		parentEntity := mgr.entityList.Get(parentSlotIdx).V

		if parentEntity.State() < ec.EntityState_Awaking || parentEntity.State() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentID, parentEntity.State())
		}
	}

	childSlotIdx, childTreeNode := mgr.getTreeNode(childID)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childID)
	}
	if childTreeNode != nil {
		return fmt.Errorf("%w: child entity %q already in the entity-tree", ErrEntityTree, childID)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childID, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Free {
		return fmt.Errorf("%w: child entity %q is in an unexpected tree node state %q", ErrEntityTree, childID, childEntity.TreeNodeState())
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
		caller := newTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			_EmitEventEntityTreeAddNode(mgr, mgr, parentID, childID)
		}) {
			return nil
		}

		if parentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(parentEntity).EmitEventTreeNodeAddChild(childID)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeAttachParent(parentID)
		}) {
			return nil
		}
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Attached)

	return nil
}

// RemoveNode 按后序递归移除整个子树的树关系；实体本身不会被销毁。
func (mgr *_EntityManager) RemoveNode(childID uid.ID) error {
	childSlotIdx, childTreeNode := mgr.getTreeNode(childID)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childID)
	}
	if childTreeNode == nil {
		return fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childID)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childID, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childID, childEntity.TreeNodeState())
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	parentID := ForestNodeID
	parentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var parentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		parentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		parentEntity = mgr.entityList.Get(childTreeNode.parent).V
		parentID = parentEntity.ID()
	}

	{
		caller := newTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
				entity := mgr.entityList.Get(slot.V).V
				mgr.RemoveNode(entity.ID())
			})
		}) {
			return nil
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(parentID)
		}) {
			return nil
		}

		if parentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(parentEntity).EmitEventTreeNodeRemoveChild(childID)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			_EmitEventEntityTreeRemoveNode(mgr, mgr, parentID, childID)
		}) {
			return nil
		}
	}

	delete(mgr.entityTreeNodes, childSlotIdx)
	parentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Free)

	return nil
}

// DetachNode 将节点从当前父实体移到虚拟森林节点下，使其成为根节点。
func (mgr *_EntityManager) DetachNode(childID uid.ID) error {
	return mgr.MoveNode(childID, ForestNodeID)
}

// MoveNode 将节点移动到新的父节点下。
func (mgr *_EntityManager) MoveNode(childID, parentID uid.ID) error {
	toParentSlotIdx, toParentTreeNode := mgr.getTreeNode(parentID)
	if toParentSlotIdx < 0 {
		if toParentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not exists", ErrEntityTree, parentID)
		}
	} else {
		if toParentTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
		}

		toParentEntity := mgr.entityList.Get(toParentSlotIdx).V

		if toParentEntity.State() < ec.EntityState_Awaking || toParentEntity.State() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentID, toParentEntity.State())
		}
	}

	childSlotIdx, childTreeNode := mgr.getTreeNode(childID)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childID)
	}
	if childTreeNode == nil {
		return fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childID)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childID, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childID, childEntity.TreeNodeState())
	}

	for ancestorSlotIdx := toParentSlotIdx; ancestorSlotIdx >= 0; {
		if ancestorSlotIdx == childSlotIdx {
			return fmt.Errorf("%w: moving child entity %q under parent entity %q would create a cycle", ErrEntityTree, childID, parentID)
		}

		ancestorTreeNode := mgr.entityTreeNodes[ancestorSlotIdx]
		if ancestorTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q has an invalid ancestor chain", ErrEntityTree, parentID)
		}
		ancestorSlotIdx = ancestorTreeNode.parent
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Moving)

	fromParentID := ForestNodeID
	fromParentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var fromParentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		fromParentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		fromParentEntity = mgr.entityList.Get(childTreeNode.parent).V
		fromParentID = fromParentEntity.ID()
	}

	toParentID := parentID
	var toParentEntity ec.Entity
	if toParentSlotIdx >= 0 {
		toParentEntity = mgr.entityList.Get(toParentSlotIdx).V
	}

	{
		caller := newTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(fromParentID)
		}) {
			return nil
		}

		if fromParentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(fromParentEntity).EmitEventTreeNodeRemoveChild(childID)
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
			_EmitEventEntityTreeMoveNode(mgr, mgr, childID, fromParentID, toParentID)
		}) {
			return nil
		}

		if toParentEntity != nil {
			if !caller.Call(func() {
				ec.UnsafeEntity(toParentEntity).EmitEventTreeNodeAddChild(childID)
			}) {
				return nil
			}
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeAttachParent(parentID)
		}) {
			return nil
		}

		if !caller.Call(func() {
			ec.UnsafeEntity(childEntity).EmitEventTreeNodeMoveTo(fromParentID, toParentID)
		}) {
			return nil
		}
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Attached)

	return nil
}

// IsFree 报告实体是否尚未加入实体树。
func (mgr *_EntityManager) IsFree(entityID uid.ID) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityID)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityID)
	}
	return treeNode == nil, nil
}

// IsRoot 报告实体是否直接挂在虚拟森林节点下。
func (mgr *_EntityManager) IsRoot(entityID uid.ID) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityID)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityID)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityID)
	}
	return treeNode.parent == forestNodeIdx, nil
}

// IsLeaf 报告实体是否没有子节点。
func (mgr *_EntityManager) IsLeaf(entityID uid.ID) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityID)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityID)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityID)
	}
	return treeNode.children.Len()-treeNode.children.OrphanCount() <= 0, nil
}

// GetParent 返回父实体；根节点没有实体父节点，因此返回错误。
func (mgr *_EntityManager) GetParent(childID uid.ID) (ec.Entity, error) {
	slotIdx, treeNode := mgr.getTreeNode(childID)
	if slotIdx < 0 {
		return nil, fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childID)
	}
	if treeNode == nil {
		return nil, fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childID)
	}
	if treeNode.parent == forestNodeIdx {
		return nil, fmt.Errorf("%w: child entity %q is root node", ErrEntityTree, childID)
	}
	return mgr.entityList.Get(treeNode.parent).V, nil
}

// RangeChildren 按加入顺序遍历直接子节点，回调返回 false 时停止。
func (mgr *_EntityManager) RangeChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}
	treeNode.children.Traversal(func(slot *generic.FreeSlot[int]) bool {
		return fun(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// EachChildren 按加入顺序遍历全部直接子节点。
func (mgr *_EntityManager) EachChildren(parentID uid.ID, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}
	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedRangeChildren 按加入顺序逆向遍历直接子节点，回调返回 false 时停止。
func (mgr *_EntityManager) ReversedRangeChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}
	treeNode.children.ReversedTraversal(func(slot *generic.FreeSlot[int]) bool {
		return fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedEachChildren 按加入顺序逆向遍历全部直接子节点。
func (mgr *_EntityManager) ReversedEachChildren(parentID uid.ID, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}
	treeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// FilterChildren 按加入顺序返回符合条件的直接子节点。
func (mgr *_EntityManager) FilterChildren(parentID uid.ID, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error) {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return nil, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
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

// ListChildren 按加入顺序返回直接子节点切片。
func (mgr *_EntityManager) ListChildren(parentID uid.ID) ([]ec.Entity, error) {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return nil, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}

	entities := make([]ec.Entity, 0, treeNode.children.Len()-treeNode.children.OrphanCount())

	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		entities = append(entities, mgr.entityList.Get(slot.V).V)
	})

	return entities, nil
}

// CountChildren 返回直接子节点数。
func (mgr *_EntityManager) CountChildren(parentID uid.ID) (int, error) {
	_, treeNode := mgr.getTreeNode(parentID)
	if treeNode == nil {
		return 0, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentID)
	}
	return treeNode.children.Len() - treeNode.children.OrphanCount(), nil
}

func (mgr *_EntityManager) onEntityDestroyRemoveNode(childID uid.ID) {
	childSlotIdx, childTreeNode := mgr.getTreeNode(childID)
	if childSlotIdx < 0 {
		return
	}
	if childTreeNode == nil {
		return
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	parentID := ForestNodeID
	parentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var parentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		parentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		parentEntity = mgr.entityList.Get(childTreeNode.parent).V
		parentID = parentEntity.ID()
	}

	childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		entity := mgr.entityList.Get(slot.V).V
		mgr.onEntityDestroyRemoveNode(entity.ID())
	})

	ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(parentID)

	if parentEntity != nil {
		ec.UnsafeEntity(parentEntity).EmitEventTreeNodeRemoveChild(childID)
	}

	_EmitEventEntityTreeRemoveNode(mgr, mgr, parentID, childID)

	delete(mgr.entityTreeNodes, childSlotIdx)
	parentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Free)
}

func (mgr *_EntityManager) getTreeNode(entityID uid.ID) (int, *_TreeNode) {
	if entityID == ForestNodeID {
		return forestNodeIdx, mgr.entityTreeNodes[forestNodeIdx]
	}

	slotIdx, ok := mgr.entityIDIndex[entityID]
	if !ok {
		return -2, nil
	}

	treeNode, ok := mgr.entityTreeNodes[slotIdx]
	if !ok {
		return slotIdx, nil
	}

	return slotIdx, treeNode
}

func newTreeNodeCaller(entity ec.Entity) _TreeNodeCaller {
	return _TreeNodeCaller{entity: entity, state: entity.TreeNodeState()}
}

type _TreeNodeCaller struct {
	entity ec.Entity
	state  ec.TreeNodeState
}

func (c _TreeNodeCaller) Call(fun func()) bool {
	if c.entity.TreeNodeState() != c.state {
		return false
	}

	fun()

	return c.entity.TreeNodeState() == c.state
}
