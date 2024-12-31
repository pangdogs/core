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
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityTree 实体树接口
type EntityTree interface {
	ictx.CurrentContextProvider

	// AddNode 新增实体节点，会向实体管理器添加实体
	AddNode(entity ec.Entity, parentId uid.Id) error
	// PruningNode 实体树节点剪枝
	PruningNode(entityId uid.Id)
	// RangeChildren 遍历子实体
	RangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool])
	// ReversedRangeChildren 反向遍历子实体
	ReversedRangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool])
	// FilterChildren 过滤并获取子实体
	FilterChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) []ec.Entity
	// GetChildren 获取所有子实体
	GetChildren(entityId uid.Id) []ec.Entity
	// CountChildren 获取子实体数量
	CountChildren(entityId uid.Id) int
	// IsTop 是否是顶层节点
	IsTop(entityId uid.Id) bool
	// ChangeParent 修改父实体
	ChangeParent(entityId, parentId uid.Id) error
	// GetParent 获取父实体
	GetParent(entityId uid.Id) (ec.Entity, bool)

	IEntityTreeEventTab
}

// AddNode 新增实体节点，会向实体管理器添加实体
func (mgr *_EntityManagerBehavior) AddNode(entity ec.Entity, parentId uid.Id) error {
	if parentId.IsNil() {
		return fmt.Errorf("%w: %w: parentId is nil", ErrEntityManager, exception.ErrArgs)
	}
	return mgr.addEntity(entity, parentId)
}

// PruningNode 实体树节点剪枝
func (mgr *_EntityManagerBehavior) PruningNode(entityId uid.Id) {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return
	}

	if entity.GetState() != ec.EntityState_Alive {
		return
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	mgr.detachParentNode(entity)
	mgr.removeFromParentNode(entity)
}

// RangeChildren 遍历子实体
func (mgr *_EntityManagerBehavior) RangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return
	}

	node.children.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// ReversedRangeChildren 反向遍历子实体
func (mgr *_EntityManagerBehavior) ReversedRangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return
	}

	node.children.ReversedTraversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// FilterChildren 过滤并获取子实体
func (mgr *_EntityManagerBehavior) FilterChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return nil
	}

	var entities []ec.Entity

	node.children.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		entity := iface.Cache2Iface[ec.Entity](n.V.Cache)

		if fun.Exec(entity) {
			entities = append(entities, entity)
		}

		return true
	})

	return entities
}

// GetChildren 获取所有子实体
func (mgr *_EntityManagerBehavior) GetChildren(entityId uid.Id) []ec.Entity {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return nil
	}

	entities := make([]ec.Entity, 0, node.children.Len())

	node.children.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		entities = append(entities, iface.Cache2Iface[ec.Entity](n.V.Cache))
		return true
	})

	return entities
}

// CountChildren 获取子实体数量
func (mgr *_EntityManagerBehavior) CountChildren(entityId uid.Id) int {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return 0
	}
	return node.children.Len()
}

// IsTop 是否是顶层节点
func (mgr *_EntityManagerBehavior) IsTop(entityId uid.Id) bool {
	node, ok := mgr.treeNodes[entityId]
	if !ok {
		return true
	}
	return node.parentAt == nil
}

// ChangeParent 修改父实体
func (mgr *_EntityManagerBehavior) ChangeParent(entityId, parentId uid.Id) error {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return fmt.Errorf("%w: entity not exist", ErrEntityManager)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityManager, entity.GetState())
	}

	if parentId.IsNil() {
		mgr.PruningNode(entityId)
		return nil
	}

	parent, ok := mgr.GetEntity(parentId)
	if !ok {
		return fmt.Errorf("%w: parent not exist", ErrEntityManager)
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent state %q", ErrEntityManager, parent.GetState())
	}

	if parent.GetId() == entity.GetId() {
		return fmt.Errorf("%w: parent and child cannot be the same", ErrEntityManager)
	}

	switch entity.GetTreeNodeState() {
	case ec.TreeNodeState_Freedom:
		mgr.addToParentNode(entity, parent)
		mgr.attachParentNode(entity, parent)
	case ec.TreeNodeState_Attached:
		if p, ok := entity.GetTreeNodeParent(); ok {
			if p.GetId() == parent.GetId() {
				return nil
			}
		}

		for p, _ := parent.GetTreeNodeParent(); p != nil; p, _ = p.GetTreeNodeParent() {
			if p.GetId() == entity.GetId() {
				return fmt.Errorf("%w: detected a cycle in the tree structure", ErrEntityManager)
			}
		}

		mgr.changeParentNode(entity, parent)
	default:
		return fmt.Errorf("%w: invalid entity tree node state %q", ErrEntityManager, entity.GetTreeNodeState())
	}

	return nil
}

// GetParent 获取父实体
func (mgr *_EntityManagerBehavior) GetParent(entityId uid.Id) (ec.Entity, bool) {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return nil, false
	}
	return entity.GetTreeNodeParent()
}

func (mgr *_EntityManagerBehavior) addToParentNode(entity, parent ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Freedom {
		return
	}

	mgr.enterParent(entity, parent)
}

func (mgr *_EntityManagerBehavior) attachParentNode(entity, parent ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attaching {
		return
	}

	ec.UnsafeEntity(entity).EnterParentNode()

	_EmitEventEntityTreeAddNodeWithInterrupt(mgr, func(entityTree EntityTree, parent, child ec.Entity) bool {
		return parent.GetState() > ec.EntityState_Alive || child.GetState() > ec.EntityState_Alive
	}, mgr, parent, entity)

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Attached)
}

func (mgr *_EntityManagerBehavior) detachParentNode(entity ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Detaching {
		return
	}

	parent, ok := entity.GetTreeNodeParent()
	if !ok {
		return
	}

	_EmitEventEntityTreeRemoveNodeWithInterrupt(mgr, func(entityTree EntityTree, parent, child ec.Entity) bool {
		return parent.GetState() >= ec.EntityState_Death || child.GetState() >= ec.EntityState_Death
	}, mgr, parent, entity)

	ec.UnsafeEntity(entity).LeaveParentNode()
}

func (mgr *_EntityManagerBehavior) removeFromParentNode(entity ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	mgr.leaveParent(entity)
}

func (mgr *_EntityManagerBehavior) changeParentNode(entity, parent ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	mgr.detachParentNode(entity)

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	mgr.enterParent(entity, parent)
	mgr.attachParentNode(entity, parent)
}

func (mgr *_EntityManagerBehavior) enterParent(entity, parent ec.Entity) {
	parentNode, ok := mgr.treeNodes[parent.GetId()]
	if !ok {
		parentNode = &_TreeNode{}
		mgr.treeNodes[parent.GetId()] = parentNode
	}
	if parentNode.children == nil {
		parentNode.children = generic.NewList[iface.FaceAny]()
	}

	node, ok := mgr.treeNodes[entity.GetId()]
	if !ok {
		node = &_TreeNode{}
		mgr.treeNodes[entity.GetId()] = node
	}

	if node.parentAt != nil {
		node.parentAt.Escape()
		node.parentAt = nil
	}

	node.parentAt = parentNode.children.PushBack(iface.MakeFaceAny(entity))

	ec.UnsafeEntity(entity).SetTreeNodeParent(parent)
	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Attaching)
}

func (mgr *_EntityManagerBehavior) leaveParent(entity ec.Entity) {
	ec.UnsafeEntity(entity).SetTreeNodeParent(nil)
	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Freedom)

	node, ok := mgr.treeNodes[entity.GetId()]
	if ok {
		if node.parentAt != nil {
			node.parentAt.Escape()
			node.parentAt = nil
		}

		if node.children == nil || node.children.Len() <= 0 {
			delete(mgr.treeNodes, entity.GetId())
		}
	}
}
