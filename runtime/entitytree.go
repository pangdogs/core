package runtime

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/uid"
)

// EntityTree 实体树接口
type EntityTree interface {
	iEntityTree
	concurrent.CurrentContextProvider

	// AddNode 新增实体树节点。切换父实体时，RemoveNode()离开旧父实体，AddNode()加入新父实体
	AddNode(parentId, childId uid.Id) error
	// RemoveNode 删除实体树节点，会销毁实体节点所有的子实体
	RemoveNode(parentId uid.Id)
	// RangeChildren 遍历子实体
	RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool])
	// ReverseRangeChildren 反向遍历子实体
	ReverseRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool])
	// CountChildren 获取子实体数量
	CountChildren(parentId uid.Id) int
	// GetParent 获取子实体的父实体
	GetParent(childId uid.Id) (ec.Entity, bool)

	_AutoEventEntityTreeAddNode    // 事件：新增实体树节点
	_AutoEventEntityTreeRemoveNode // 事件：删除实体树节点
}

type iEntityTree interface {
	getAndCheckEntity(entityId uid.Id) (ec.Entity, error)
}

type _TreeNode struct {
	parent   ec.Entity
	element  *container.Element[iface.FaceAny]
	children *container.List[iface.FaceAny]
}

type _EntityTreeBehavior struct {
	ctx                        Context
	treeNodes                  map[uid.Id]_TreeNode
	hookEntityMgrRemoveEntity  event.Hook
	eventEntityTreeAddChild    event.Event
	eventEntityTreeRemoveChild event.Event
}

func (entityTree *_EntityTreeBehavior) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityTree, exception.ErrArgs))
	}

	entityTree.ctx = ctx
	entityTree.treeNodes = map[uid.Id]_TreeNode{}
	entityTree.hookEntityMgrRemoveEntity = BindEventEntityMgrRemoveEntity(entityTree.ctx.GetEntityMgr(), entityTree)

	ctx.ActivateEvent(&entityTree.eventEntityTreeAddChild, event.EventRecursion_Allow)
	ctx.ActivateEvent(&entityTree.eventEntityTreeRemoveChild, event.EventRecursion_Allow)
}

func (entityTree *_EntityTreeBehavior) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Terminated:
		entityTree.hookEntityMgrRemoveEntity.Unbind()
		entityTree.eventEntityTreeAddChild.Close()
		entityTree.eventEntityTreeRemoveChild.Close()
	}
}

// GetCurrentContext 获取当前上下文
func (entityTree *_EntityTreeBehavior) GetCurrentContext() iface.Cache {
	return entityTree.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (entityTree *_EntityTreeBehavior) GetConcurrentContext() iface.Cache {
	return entityTree.ctx.GetConcurrentContext()
}

// AddNode 新增实体树节点。切换父实体时，RemoveNode()离开旧父实体，AddNode()加入新父实体
func (entityTree *_EntityTreeBehavior) AddNode(parentId, childId uid.Id) error {
	if parentId == childId {
		return fmt.Errorf("%w: %w: parentId and childId is %q, can't be equal", ErrEntityTree, exception.ErrArgs, parentId)
	}

	parent, err := entityTree.getAndCheckEntity(parentId)
	if err != nil {
		return err
	}

	child, err := entityTree.getAndCheckEntity(childId)
	if err != nil {
		return err
	}

	if _, ok := entityTree.treeNodes[childId]; ok {
		return fmt.Errorf("%w: child entity %q already in entity-tree", ErrEntityTree, childId)
	}

	parentNode, ok := entityTree.treeNodes[parentId]
	if !ok || parentNode.children == nil {
		parentNode.children = container.NewList[iface.FaceAny]()
		entityTree.treeNodes[parentId] = parentNode
	}

	entityTree.treeNodes[childId] = _TreeNode{
		parent:  parent,
		element: parentNode.children.PushBack(iface.MakeFaceAny(child)),
	}

	_child := ec.UnsafeEntity(child)
	_child.SetTreeNodeParent(parent)
	_child.SetTreeNodeState(ec.TreeNodeState_Attached)

	_EmitEventEntityTreeAddNode(entityTree, entityTree, parent, child)
	return nil
}

// RemoveNode 删除实体树节点，会销毁实体节点所有的子实体
func (entityTree *_EntityTreeBehavior) RemoveNode(parentId uid.Id) {
	node, ok := entityTree.treeNodes[parentId]
	if !ok {
		return
	}

	child := ec.UnsafeEntity(iface.Cache2Iface[ec.Entity](node.element.Value.Cache))
	if child.GetTreeNodeState() != ec.TreeNodeState_Attached {
		return
	}

	child.SetTreeNodeState(ec.TreeNodeState_Detaching)

	if node.children != nil {
		node.children.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
			iface.Cache2Iface[ec.Entity](e.Value.Cache).DestroySelf()
			return true
		})
	}

	child.SetTreeNodeState(ec.TreeNodeState_Detached)
	child.SetTreeNodeParent(nil)

	delete(entityTree.treeNodes, parentId)
	node.element.Escape()

	_EmitEventEntityTreeRemoveNode(entityTree, entityTree, node.parent, child.Entity)
}

// RangeChildren 遍历子实体
func (entityTree *_EntityTreeBehavior) RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := entityTree.treeNodes[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeChildren 反向遍历子实体
func (entityTree *_EntityTreeBehavior) ReverseRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := entityTree.treeNodes[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// CountChildren 获取子实体数量
func (entityTree *_EntityTreeBehavior) CountChildren(parentId uid.Id) int {
	node, ok := entityTree.treeNodes[parentId]
	if !ok || node.children == nil {
		return 0
	}

	return node.children.Len()
}

// GetParent 获取子实体的父实体
func (entityTree *_EntityTreeBehavior) GetParent(childId uid.Id) (ec.Entity, bool) {
	node, ok := entityTree.treeNodes[childId]
	if !ok {
		return nil, false
	}

	return node.parent, node.parent != nil
}

// EventEntityTreeAddNode 事件：新增实体树节点
func (entityTree *_EntityTreeBehavior) EventEntityTreeAddNode() event.IEvent {
	return &entityTree.eventEntityTreeAddChild
}

// EventEntityTreeRemoveNode 事件：删除实体树节点
func (entityTree *_EntityTreeBehavior) EventEntityTreeRemoveNode() event.IEvent {
	return &entityTree.eventEntityTreeRemoveChild
}

func (entityTree *_EntityTreeBehavior) OnEntityMgrRemoveEntity(entityMgr EntityMgr, entity ec.Entity) {
	entityTree.RemoveNode(entity.GetId())
}

func (entityTree *_EntityTreeBehavior) getAndCheckEntity(entityId uid.Id) (ec.Entity, error) {
	entity, ok := entityTree.ctx.GetEntityMgr().GetEntity(entityId)
	if !ok {
		return nil, fmt.Errorf("%w: entity %q not exist", ErrEntityTree, entityId)
	}

	switch entity.GetState() {
	case ec.EntityState_Living:
	default:
		return nil, fmt.Errorf("%w: entity %q has invalid state %q", ErrEntityTree, entityId, entity.GetState())
	}

	return entity, nil
}
