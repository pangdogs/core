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
	// PruningNode 实体树节点剪枝，使实体成为根节点
	PruningNode(entityId uid.Id) error
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
	// ChangeParent 修改父实体
	ChangeParent(entityId, parentId uid.Id) error
	// GetParent 获取父实体
	GetParent(entityId uid.Id) (ec.Entity, bool)

	IEntityTreeEventTab
}

// AddNode 新增实体节点，会向实体管理器添加实体
func (mgr *_EntityManagerBehavior) AddNode(entity ec.Entity, parentId uid.Id) error {
	if parentId.IsNil() {
		return fmt.Errorf("%w: %w: parentId is nil", ErrEntityTree, exception.ErrArgs)
	}
	return mgr.addEntity(entity, parentId)
}

// PruningNode 实体树节点剪枝，使实体成为根节点
func (mgr *_EntityManagerBehavior) PruningNode(entityId uid.Id) error {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return fmt.Errorf("%w: entity %q not exist", ErrEntityTree, entityId)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: invalid entity %q tree node state %q", ErrEntityTree, entity.GetId(), entity.GetTreeNodeState())
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	mgr.detachFromParentNode(entity)
	mgr.removeFromParentNode(entity)

	return nil
}

// RangeChildren 遍历子实体
func (mgr *_EntityManagerBehavior) RangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	entityNode, ok := mgr.treeNodes[entityId]
	if !ok || entityNode.children == nil {
		return
	}

	entityNode.children.Traversal(func(childNode *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](childNode.V.Cache))
	})
}

// ReversedRangeChildren 反向遍历子实体
func (mgr *_EntityManagerBehavior) ReversedRangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	entityNode, ok := mgr.treeNodes[entityId]
	if !ok || entityNode.children == nil {
		return
	}

	entityNode.children.ReversedTraversal(func(childNode *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](childNode.V.Cache))
	})
}

// FilterChildren 过滤并获取子实体
func (mgr *_EntityManagerBehavior) FilterChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	entityNode, ok := mgr.treeNodes[entityId]
	if !ok || entityNode.children == nil {
		return nil
	}

	var entities []ec.Entity

	entityNode.children.Traversal(func(childNode *generic.Node[iface.FaceAny]) bool {
		entity := iface.Cache2Iface[ec.Entity](childNode.V.Cache)

		if fun.Exec(entity) {
			entities = append(entities, entity)
		}

		return true
	})

	return entities
}

// GetChildren 获取所有子实体
func (mgr *_EntityManagerBehavior) GetChildren(entityId uid.Id) []ec.Entity {
	entityNode, ok := mgr.treeNodes[entityId]
	if !ok || entityNode.children == nil {
		return nil
	}

	entities := make([]ec.Entity, 0, entityNode.children.Len())

	entityNode.children.Traversal(func(childNode *generic.Node[iface.FaceAny]) bool {
		entities = append(entities, iface.Cache2Iface[ec.Entity](childNode.V.Cache))
		return true
	})

	return entities
}

// CountChildren 获取子实体数量
func (mgr *_EntityManagerBehavior) CountChildren(entityId uid.Id) int {
	entityNode, ok := mgr.treeNodes[entityId]
	if !ok || entityNode.children == nil {
		return 0
	}
	return entityNode.children.Len()
}

// ChangeParent 修改父实体
func (mgr *_EntityManagerBehavior) ChangeParent(entityId, parentId uid.Id) error {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return fmt.Errorf("%w: entity %q not exist", ErrEntityTree, entityId)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if parentId.IsNil() {
		return mgr.PruningNode(entityId)
	}

	parent, ok := mgr.GetEntity(parentId)
	if !ok {
		return fmt.Errorf("%w: parent %q not exist", ErrEntityTree, parent.GetId())
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent %q state %q", ErrEntityTree, parent.GetId(), parent.GetState())
	}

	if parent.GetId() == entity.GetId() {
		return fmt.Errorf("%w: parent and child %q can't be the same", ErrEntityTree, parent.GetId())
	}

	switch entity.GetTreeNodeState() {
	case ec.TreeNodeState_Freedom:
		if err := mgr.appendToParentNode(entity, parent); err != nil {
			return err
		}

		if err := mgr.attachToParentNode(entity, parent); err != nil {
			return err
		}

		return nil

	case ec.TreeNodeState_Attached:
		if currParent, ok := entity.GetTreeNodeParent(); ok {
			if currParent.GetId() == parent.GetId() {
				return nil
			}
		}

		for it, _ := parent.GetTreeNodeParent(); it != nil; it, _ = it.GetTreeNodeParent() {
			if it.GetId() == entity.GetId() {
				return fmt.Errorf("%w: detected a cycle in the tree structure", ErrEntityTree)
			}
		}

		return mgr.changeToParentNode(entity, parent)

	default:
		return fmt.Errorf("%w: invalid entity %q tree node state %q", ErrEntityTree, entity.GetId(), entity.GetTreeNodeState())
	}
}

// GetParent 获取父实体
func (mgr *_EntityManagerBehavior) GetParent(entityId uid.Id) (ec.Entity, bool) {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return nil, false
	}
	return entity.GetTreeNodeParent()
}

func (mgr *_EntityManagerBehavior) changeToParentNode(entity, parent ec.Entity) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityTree, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityTree, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent %q state %q", ErrEntityTree, parent.GetId(), parent.GetState())
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return fmt.Errorf("%w: invalid entity %q tree node state %q", ErrEntityTree, entity.GetId(), entity.GetTreeNodeState())
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)

	mgr.detachFromParentNode(entity)

	if err := mgr.appendToParentNode(entity, parent); err != nil {
		return err
	}

	if err := mgr.attachToParentNode(entity, parent); err != nil {
		return err
	}

	return nil
}

func (mgr *_EntityManagerBehavior) appendToParentNode(entity, parent ec.Entity) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityTree, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityTree, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent %q state %q", ErrEntityTree, parent.GetId(), parent.GetState())
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Freedom {
		return fmt.Errorf("%w: invalid entity %q tree node state %q", ErrEntityTree, entity.GetId(), entity.GetTreeNodeState())
	}

	parentNode, ok := mgr.treeNodes[parent.GetId()]
	if !ok {
		parentNode = &_TreeNode{}
		mgr.treeNodes[parent.GetId()] = parentNode
	}
	if parentNode.children == nil {
		parentNode.children = generic.NewList[iface.FaceAny]()
	}

	entityNode, ok := mgr.treeNodes[entity.GetId()]
	if !ok {
		entityNode = &_TreeNode{}
		mgr.treeNodes[entity.GetId()] = entityNode
	}
	if entityNode.parentAt != nil {
		entityNode.parentAt.Escape()
		entityNode.parentAt = nil
	}
	entityNode.parentAt = parentNode.children.PushBack(iface.MakeFaceAny(entity))

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Attaching)
	ec.UnsafeEntity(entity).SetTreeNodeParent(parent)

	return nil
}

func (mgr *_EntityManagerBehavior) removeFromParentNode(entity ec.Entity) {
	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Freedom)
	ec.UnsafeEntity(entity).SetTreeNodeParent(nil)

	entityNode, ok := mgr.treeNodes[entity.GetId()]
	if ok {
		if entityNode.parentAt != nil {
			entityNode.parentAt.Escape()
			entityNode.parentAt = nil
		}

		if entityNode.children == nil || entityNode.children.Len() <= 0 {
			delete(mgr.treeNodes, entity.GetId())
		}
	}
}

func (mgr *_EntityManagerBehavior) attachToParentNode(entity, parent ec.Entity) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityTree, exception.ErrArgs)
	}

	if parent == nil {
		exception.Panicf("%w: %w: parent is nil", ErrEntityTree, exception.ErrArgs)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent %q state %q", ErrEntityTree, parent.GetId(), parent.GetState())
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Attaching {
		return fmt.Errorf("%w: invalid entity %q tree node state %q", ErrEntityTree, entity.GetId(), entity.GetTreeNodeState())
	}

	ec.UnsafeEntity(entity).EnterParentNode()

	_EmitEventEntityTreeAddNodeWithInterrupt(mgr, func(entityTree EntityTree, parent, child ec.Entity) bool {
		return parent.GetState() > ec.EntityState_Alive || child.GetState() > ec.EntityState_Alive
	}, mgr, parent, entity)

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityTree, entity.GetId(), entity.GetState())
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent %q state %q", ErrEntityTree, parent.GetId(), parent.GetState())
	}

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Attached)

	return nil
}

func (mgr *_EntityManagerBehavior) detachFromParentNode(entity ec.Entity) {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityTree, exception.ErrArgs)
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Detaching {
		return
	}

	parent, ok := entity.GetTreeNodeParent()
	if !ok {
		return
	}

	_EmitEventEntityTreeRemoveNodeWithInterrupt(mgr, func(entityTree EntityTree, parent, child ec.Entity) bool {
		return parent.GetState() >= ec.EntityState_Destroyed || child.GetState() >= ec.EntityState_Destroyed
	}, mgr, parent, entity)

	ec.UnsafeEntity(entity).LeaveParentNode()
}
