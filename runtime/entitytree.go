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
	// ForestNodeId 是所有根实体共用的虚拟父节点 ID；它是保留哨兵，不应修改。
	ForestNodeId = uid.From("d5rh7sbr1n96c63fs3vg")
	// forestNodeIdx 是虚拟森林节点在内部索引中的保留值。
	forestNodeIdx = -1
)

// EntityTree 管理当前运行时实体之间的父子关系。
// 树操作不提供并发保护，应在所属运行时 goroutine 中执行。
type EntityTree interface {
	corectx.CurrentContextProvider

	// MakeRoot 将自由实体作为根节点加入实体树。
	MakeRoot(entityId uid.Id) error
	// AddChild 将自由实体 childId 挂到 parentId 下。
	AddChild(parentId, childId uid.Id) error
	// RemoveNode 按后序递归移除整个子树的树关系；实体本身不会被销毁。
	RemoveNode(childId uid.Id) error
	// DetachNode 将节点从当前父实体移到虚拟森林节点下，使其成为根节点。
	DetachNode(childId uid.Id) error
	// MoveNode 将节点移动到新的父节点下。
	MoveNode(childId, parentId uid.Id) error
	// IsFree 报告实体是否尚未加入实体树。
	IsFree(entityId uid.Id) (bool, error)
	// IsRoot 报告实体是否直接挂在虚拟森林节点下。
	IsRoot(entityId uid.Id) (bool, error)
	// IsLeaf 报告实体是否没有子节点。
	IsLeaf(entityId uid.Id) (bool, error)
	// GetParent 返回父实体；根节点没有实体父节点，因此返回错误。
	GetParent(childId uid.Id) (ec.Entity, error)
	// RangeChildren 按加入顺序遍历直接子节点，回调返回 false 时停止。
	RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error
	// EachChildren 按加入顺序遍历全部直接子节点。
	EachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error
	// ReversedRangeChildren 按加入顺序逆向遍历直接子节点，回调返回 false 时停止。
	ReversedRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error
	// ReversedEachChildren 按加入顺序逆向遍历全部直接子节点。
	ReversedEachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error
	// FilterChildren 按加入顺序返回符合条件的直接子节点。
	FilterChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error)
	// ListChildren 按加入顺序返回直接子节点切片。
	ListChildren(parentId uid.Id) ([]ec.Entity, error)
	// CountChildren 返回直接子节点数。
	CountChildren(parentId uid.Id) (int, error)

	IEntityTreeEventTab
}

// MakeRoot 将自由实体作为根节点加入实体树。
func (mgr *_EntityManager) MakeRoot(entityId uid.Id) error {
	return mgr.AddChild(ForestNodeId, entityId)
}

// AddChild 将自由实体 childId 挂到 parentId 下。
func (mgr *_EntityManager) AddChild(parentId, childId uid.Id) error {
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

		if parentEntity.State() < ec.EntityState_Awaking || parentEntity.State() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentId, parentEntity.State())
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

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Free {
		return fmt.Errorf("%w: child entity %q is in an unexpected tree node state %q", ErrEntityTree, childId, childEntity.TreeNodeState())
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

// RemoveNode 按后序递归移除整个子树的树关系；实体本身不会被销毁。
func (mgr *_EntityManager) RemoveNode(childId uid.Id) error {
	childSlotIdx, childTreeNode := mgr.getTreeNode(childId)
	if childSlotIdx < 0 {
		return fmt.Errorf("%w: child entity %q not exists", ErrEntityTree, childId)
	}
	if childTreeNode == nil {
		return fmt.Errorf("%w: child entity %q not in the entity-tree", ErrEntityTree, childId)
	}

	childEntity := mgr.entityList.Get(childSlotIdx).V

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childId, childEntity.TreeNodeState())
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	parentId := ForestNodeId
	parentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var parentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		parentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		parentEntity = mgr.entityList.Get(childTreeNode.parent).V
		parentId = parentEntity.Id()
	}

	{
		caller := newTreeNodeCaller(childEntity)

		if !caller.Call(func() {
			childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
				entity := mgr.entityList.Get(slot.V).V
				mgr.RemoveNode(entity.Id())
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

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Free)

	return nil
}

// DetachNode 将节点从当前父实体移到虚拟森林节点下，使其成为根节点。
func (mgr *_EntityManager) DetachNode(childId uid.Id) error {
	return mgr.MoveNode(childId, ForestNodeId)
}

// MoveNode 将节点移动到新的父节点下。
func (mgr *_EntityManager) MoveNode(childId, parentId uid.Id) error {
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

		if toParentEntity.State() < ec.EntityState_Awaking || toParentEntity.State() > ec.EntityState_Alive {
			return fmt.Errorf("%w: parent entity %q is in an unexpected state %q", ErrEntityTree, parentId, toParentEntity.State())
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

	if childEntity.State() < ec.EntityState_Awaking || childEntity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: child entity %q is in an unexpected state %q", ErrEntityTree, childId, childEntity.State())
	}

	if childEntity.TreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: child entity %q has an unexpected tree node state %q", ErrEntityTree, childId, childEntity.TreeNodeState())
	}

	for ancestorSlotIdx := toParentSlotIdx; ancestorSlotIdx >= 0; {
		if ancestorSlotIdx == childSlotIdx {
			return fmt.Errorf("%w: moving child entity %q under parent entity %q would create a cycle", ErrEntityTree, childId, parentId)
		}

		ancestorTreeNode := mgr.entityTreeNodes[ancestorSlotIdx]
		if ancestorTreeNode == nil {
			return fmt.Errorf("%w: parent entity %q has an invalid ancestor chain", ErrEntityTree, parentId)
		}
		ancestorSlotIdx = ancestorTreeNode.parent
	}

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Moving)

	fromParentId := ForestNodeId
	fromParentTreeNode := mgr.entityTreeNodes[forestNodeIdx]
	var fromParentEntity ec.Entity
	if childTreeNode.parent >= 0 {
		fromParentTreeNode = mgr.entityTreeNodes[childTreeNode.parent]
		fromParentEntity = mgr.entityList.Get(childTreeNode.parent).V
		fromParentId = fromParentEntity.Id()
	}

	toParentId := parentId
	var toParentEntity ec.Entity
	if toParentSlotIdx >= 0 {
		toParentEntity = mgr.entityList.Get(toParentSlotIdx).V
	}

	{
		caller := newTreeNodeCaller(childEntity)

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

// IsFree 报告实体是否尚未加入实体树。
func (mgr *_EntityManager) IsFree(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	return treeNode == nil, nil
}

// IsRoot 报告实体是否直接挂在虚拟森林节点下。
func (mgr *_EntityManager) IsRoot(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityId)
	}
	return treeNode.parent == forestNodeIdx, nil
}

// IsLeaf 报告实体是否没有子节点。
func (mgr *_EntityManager) IsLeaf(entityId uid.Id) (bool, error) {
	slotIdx, treeNode := mgr.getTreeNode(entityId)
	if slotIdx < 0 {
		return false, fmt.Errorf("%w: entity %q not exists", ErrEntityTree, entityId)
	}
	if treeNode == nil {
		return false, fmt.Errorf("%w: entity %q not in the entity-tree", ErrEntityTree, entityId)
	}
	return treeNode.children.Len()-treeNode.children.OrphanCount() <= 0, nil
}

// GetParent 返回父实体；根节点没有实体父节点，因此返回错误。
func (mgr *_EntityManager) GetParent(childId uid.Id) (ec.Entity, error) {
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

// RangeChildren 按加入顺序遍历直接子节点，回调返回 false 时停止。
func (mgr *_EntityManager) RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.Traversal(func(slot *generic.FreeSlot[int]) bool {
		return fun(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// EachChildren 按加入顺序遍历全部直接子节点。
func (mgr *_EntityManager) EachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.TraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedRangeChildren 按加入顺序逆向遍历直接子节点，回调返回 false 时停止。
func (mgr *_EntityManager) ReversedRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.ReversedTraversal(func(slot *generic.FreeSlot[int]) bool {
		return fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// ReversedEachChildren 按加入顺序逆向遍历全部直接子节点。
func (mgr *_EntityManager) ReversedEachChildren(parentId uid.Id, fun generic.Action1[ec.Entity]) error {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	treeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		fun.UnsafeCall(mgr.entityList.Get(slot.V).V)
	})
	return nil
}

// FilterChildren 按加入顺序返回符合条件的直接子节点。
func (mgr *_EntityManager) FilterChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) ([]ec.Entity, error) {
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

// ListChildren 按加入顺序返回直接子节点切片。
func (mgr *_EntityManager) ListChildren(parentId uid.Id) ([]ec.Entity, error) {
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

// CountChildren 返回直接子节点数。
func (mgr *_EntityManager) CountChildren(parentId uid.Id) (int, error) {
	_, treeNode := mgr.getTreeNode(parentId)
	if treeNode == nil {
		return 0, fmt.Errorf("%w: parent entity %q not in the entity-tree", ErrEntityTree, parentId)
	}
	return treeNode.children.Len() - treeNode.children.OrphanCount(), nil
}

func (mgr *_EntityManager) onEntityDestroyRemoveNode(childId uid.Id) {
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
		parentId = parentEntity.Id()
	}

	childTreeNode.children.ReversedTraversalEach(func(slot *generic.FreeSlot[int]) {
		entity := mgr.entityList.Get(slot.V).V
		mgr.onEntityDestroyRemoveNode(entity.Id())
	})

	ec.UnsafeEntity(childEntity).EmitEventTreeNodeDetachParent(parentId)

	if parentEntity != nil {
		ec.UnsafeEntity(parentEntity).EmitEventTreeNodeRemoveChild(childId)
	}

	_EmitEventEntityTreeRemoveNode(mgr, mgr, parentId, childId)

	delete(mgr.entityTreeNodes, childSlotIdx)
	parentTreeNode.children.ReleaseIfVersion(childTreeNode.attachedIndex, childTreeNode.attachedVersion)

	ec.UnsafeEntity(childEntity).SetTreeNodeState(ec.TreeNodeState_Free)
}

func (mgr *_EntityManager) getTreeNode(entityId uid.Id) (int, *_TreeNode) {
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
