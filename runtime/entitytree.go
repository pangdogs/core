package runtime

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityTree 实体树接口
type EntityTree interface {
	gctx.CurrentContextProvider

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

	iAutoEventEntityTreeAddNode    // 事件：新增实体树节点
	iAutoEventEntityTreeRemoveNode // 事件：删除实体树节点
}

// AddNode 新增实体节点，会向实体管理器添加实体
func (mgr *_EntityMgrBehavior) AddNode(entity ec.Entity, parentId uid.Id) error {
	if parentId.IsNil() {
		return fmt.Errorf("%w: %w: parentId is nil", ErrEntityMgr, exception.ErrArgs)
	}
	return mgr.addEntity(entity, parentId)
}

// PruningNode 实体树节点剪枝
func (mgr *_EntityMgrBehavior) PruningNode(entityId uid.Id) {
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
func (mgr *_EntityMgrBehavior) RangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return
	}

	node.children.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// ReversedRangeChildren 反向遍历子实体
func (mgr *_EntityMgrBehavior) ReversedRangeChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return
	}

	node.children.ReversedTraversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// FilterChildren 过滤并获取子实体
func (mgr *_EntityMgrBehavior) FilterChildren(entityId uid.Id, fun generic.Func1[ec.Entity, bool]) []ec.Entity {
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
func (mgr *_EntityMgrBehavior) GetChildren(entityId uid.Id) []ec.Entity {
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
func (mgr *_EntityMgrBehavior) CountChildren(entityId uid.Id) int {
	node, ok := mgr.treeNodes[entityId]
	if !ok || node.children == nil {
		return 0
	}
	return node.children.Len()
}

// IsTop 是否是顶层节点
func (mgr *_EntityMgrBehavior) IsTop(entityId uid.Id) bool {
	node, ok := mgr.treeNodes[entityId]
	if !ok {
		return false
	}
	return node.parentAt == nil
}

// ChangeParent 修改父实体
func (mgr *_EntityMgrBehavior) ChangeParent(entityId, parentId uid.Id) error {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return fmt.Errorf("%w: entity not exist", ErrEntityMgr)
	}

	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityMgr, entity.GetState())
	}

	if parentId.IsNil() {
		mgr.PruningNode(entityId)
		return nil
	}

	parent, ok := mgr.GetEntity(parentId)
	if !ok {
		return fmt.Errorf("%w: parent not exist", ErrEntityMgr)
	}

	if parent.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: invalid parent state %q", ErrEntityMgr, parent.GetState())
	}

	if parent.GetId() == entity.GetId() {
		return fmt.Errorf("%w: parent and child cannot be the same", ErrEntityMgr)
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
				return fmt.Errorf("%w: detected a cycle in the tree structure", ErrEntityMgr)
			}
		}

		mgr.changeParentNode(entity, parent)
	default:
		return fmt.Errorf("%w: invalid entity tree node state %q", ErrEntityMgr, entity.GetTreeNodeState())
	}

	return nil
}

// GetParent 获取父实体
func (mgr *_EntityMgrBehavior) GetParent(entityId uid.Id) (ec.Entity, bool) {
	entity, ok := mgr.GetEntity(entityId)
	if !ok {
		return nil, false
	}
	return entity.GetTreeNodeParent()
}

// EventEntityTreeAddNode 事件：新增实体树节点
func (mgr *_EntityMgrBehavior) EventEntityTreeAddNode() event.IEvent {
	return &mgr.eventEntityTreeAddChild
}

// EventEntityTreeRemoveNode 事件：删除实体树节点
func (mgr *_EntityMgrBehavior) EventEntityTreeRemoveNode() event.IEvent {
	return &mgr.eventEntityTreeRemoveChild
}

func (mgr *_EntityMgrBehavior) addToParentNode(entity, parent ec.Entity) {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	if parent == nil {
		panic(fmt.Errorf("%w: %w: parent is nil", ErrEntityMgr, exception.ErrArgs))
	}

	if entity.GetState() > ec.EntityState_Alive || parent.GetState() > ec.EntityState_Alive {
		return
	}

	if entity.GetTreeNodeState() != ec.TreeNodeState_Freedom {
		return
	}

	mgr.enterParent(entity, parent)
}

func (mgr *_EntityMgrBehavior) attachParentNode(entity, parent ec.Entity) {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	if parent == nil {
		panic(fmt.Errorf("%w: %w: parent is nil", ErrEntityMgr, exception.ErrArgs))
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

func (mgr *_EntityMgrBehavior) detachParentNode(entity ec.Entity) {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
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

func (mgr *_EntityMgrBehavior) removeFromParentNode(entity ec.Entity) {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	mgr.leaveParent(entity)
}

func (mgr *_EntityMgrBehavior) changeParentNode(entity, parent ec.Entity) {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	if parent == nil {
		panic(fmt.Errorf("%w: %w: parent is nil", ErrEntityMgr, exception.ErrArgs))
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

func (mgr *_EntityMgrBehavior) enterParent(entity, parent ec.Entity) {
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

func (mgr *_EntityMgrBehavior) leaveParent(entity ec.Entity) {
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
