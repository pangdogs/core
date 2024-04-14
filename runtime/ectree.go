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

// ECTree EC树接口
type ECTree interface {
	_ECTree
	concurrent.CurrentContextProvider

	// AddChild 实体加入父实体。切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
	AddChild(parentId, childId uid.Id) error
	// RemoveChild 实体离开父实体，会销毁所有子实体
	RemoveChild(childId uid.Id)
	// RangeChildren 遍历子实体
	RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool])
	// ReverseRangeChildren 反向遍历子实体
	ReverseRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool])
	// CountChildren 获取子实体数量
	CountChildren(parentId uid.Id) int
	// GetParent 获取子实体的父实体
	GetParent(childId uid.Id) (ec.Entity, bool)

	_AutoEventECTreeAddChild    // 事件：EC树中子实体加入父实体
	_AutoEventECTreeRemoveChild // 事件：EC树中子实体离开父实体
}

type _ECTree interface {
	getAndCheckEntity(entityId uid.Id) (ec.Entity, error)
}

type _ECNode struct {
	parent   ec.Entity
	element  *container.Element[iface.FaceAny]
	children *container.List[iface.FaceAny]
}

type _ECTreeBehavior struct {
	ctx                    Context
	ecTree                 map[uid.Id]_ECNode
	eventECTreeAddChild    event.Event
	eventECTreeRemoveChild event.Event
}

func (ecTree *_ECTreeBehavior) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrECTree, exception.ErrArgs))
	}

	ecTree.ctx = ctx
	ecTree.ecTree = map[uid.Id]_ECNode{}

	ctx.ActivateEvent(&ecTree.eventECTreeAddChild, event.EventRecursion_Allow)
	ctx.ActivateEvent(&ecTree.eventECTreeRemoveChild, event.EventRecursion_Allow)

	BindEventEntityMgrRemovingEntity(ecTree.ctx.GetEntityMgr(), ecTree)
}

func (ecTree *_ECTreeBehavior) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Terminated:
		ecTree.eventECTreeAddChild.Close()
		ecTree.eventECTreeRemoveChild.Close()
	}
}

// GetCurrentContext 获取当前上下文
func (ecTree *_ECTreeBehavior) GetCurrentContext() iface.Cache {
	return ecTree.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (ecTree *_ECTreeBehavior) GetConcurrentContext() iface.Cache {
	return ecTree.ctx.GetConcurrentContext()
}

// AddChild 实体加入父实体，在实体加入运行时上下文后才能调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
func (ecTree *_ECTreeBehavior) AddChild(parentId, childId uid.Id) error {
	if parentId == childId {
		return fmt.Errorf("%w: %w: parentId and childId is %q, can't be equal", ErrECTree, exception.ErrArgs, parentId)
	}

	parent, err := ecTree.getAndCheckEntity(parentId)
	if err != nil {
		return err
	}

	child, err := ecTree.getAndCheckEntity(childId)
	if err != nil {
		return err
	}

	if _, ok := ecTree.ecTree[childId]; ok {
		return fmt.Errorf("%w: child entity %q already in ec-tree", ErrECTree, childId)
	}

	parentNode, ok := ecTree.ecTree[parentId]
	if !ok || parentNode.children == nil {
		parentNode.children = container.NewList[iface.FaceAny](ecTree.ctx.getOptions().FaceAnyAllocator, ecTree.ctx)
		ecTree.ecTree[parentId] = parentNode
	}

	ecTree.ecTree[childId] = _ECNode{
		parent:  parent,
		element: parentNode.children.PushBack(iface.MakeFaceAny(child)),
	}

	_child := ec.UnsafeEntity(child)
	_child.SetECParent(parent)
	_child.SetECNodeState(ec.ECNodeState_Attached)

	_EmitEventECTreeAddChild(ecTree, ecTree, parent, child)

	return nil
}

// RemoveChild 实体离开父实体，会销毁所有子实体
func (ecTree *_ECTreeBehavior) RemoveChild(childId uid.Id) {
	node, ok := ecTree.ecTree[childId]
	if !ok {
		return
	}

	child := iface.Cache2Iface[ec.Entity](node.element.Value.Cache)
	_child := ec.UnsafeEntity(child)

	switch child.GetECNodeState() {
	case ec.ECNodeState_Detaching:
		return
	default:
		_child.SetECNodeState(ec.ECNodeState_Detaching)
	}

	if node.children != nil {
		node.children.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
			iface.Cache2Iface[ec.Entity](e.Value.Cache).DestroySelf()
			return true
		})
	}

	_child.SetECNodeState(ec.ECNodeState_Detached)
	_child.SetECParent(nil)

	delete(ecTree.ecTree, childId)
	node.element.Escape()

	_EmitEventECTreeRemoveChild(ecTree, ecTree, node.parent, child)
}

// RangeChildren 遍历子实体
func (ecTree *_ECTreeBehavior) RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeChildren 反向遍历子实体
func (ecTree *_ECTreeBehavior) ReverseRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// CountChildren 获取子实体数量
func (ecTree *_ECTreeBehavior) CountChildren(parentId uid.Id) int {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return 0
	}

	return node.children.Len()
}

// GetParent 获取子实体的父实体
func (ecTree *_ECTreeBehavior) GetParent(childId uid.Id) (ec.Entity, bool) {
	node, ok := ecTree.ecTree[childId]
	if !ok {
		return nil, false
	}

	return node.parent, node.parent != nil
}

// EventECTreeAddChild 事件：EC树中子实体加入父实体
func (ecTree *_ECTreeBehavior) EventECTreeAddChild() event.IEvent {
	return &ecTree.eventECTreeAddChild
}

// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
func (ecTree *_ECTreeBehavior) EventECTreeRemoveChild() event.IEvent {
	return &ecTree.eventECTreeRemoveChild
}

func (ecTree *_ECTreeBehavior) OnEntityMgrRemovingEntity(entityMgr EntityMgr, entity ec.Entity) {
	ecTree.RemoveChild(entity.GetId())
}

func (ecTree *_ECTreeBehavior) getAndCheckEntity(entityId uid.Id) (ec.Entity, error) {
	entity, ok := ecTree.ctx.GetEntityMgr().GetEntity(entityId)
	if !ok {
		return nil, fmt.Errorf("%w: entity %q not exist", ErrECTree, entityId)
	}

	switch entity.GetState() {
	case ec.EntityState_Awake, ec.EntityState_Start, ec.EntityState_Living:
	default:
		return nil, fmt.Errorf("%w: entity %q has invalid state %q", ErrECTree, entityId, entity.GetState())
	}

	return entity, nil
}
