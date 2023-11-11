package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// IECTree EC树接口
type IECTree interface {
	concurrent.CurrentContextResolver

	// TryAddChild 尝试加入父实体
	TryAddChild(parentId, childId uid.Id) error
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

	iAutoEventECTreeAddChild    // 事件：EC树中子实体加入父实体
	iAutoEventECTreeRemoveChild // 事件：EC树中子实体离开父实体
}

// _ECNode EC树节点
type _ECNode struct {
	parent   ec.Entity
	element  *container.Element[iface.FaceAny]
	children *container.List[iface.FaceAny]
}

// _ECTree EC树
type _ECTree struct {
	ctx                    Context
	ecTree                 map[uid.Id]_ECNode
	eventECTreeAddChild    event.Event
	eventECTreeRemoveChild event.Event
}

func (ecTree *_ECTree) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrECTree, exception.ErrArgs))
	}

	ecTree.ctx = ctx
	ecTree.ecTree = map[uid.Id]_ECNode{}
	ecTree.eventECTreeAddChild.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.getOptions().HookAllocator, ctx)
	ecTree.eventECTreeRemoveChild.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.getOptions().HookAllocator, ctx)

	BindEventEntityMgrRemovingEntity(ecTree.ctx.GetEntityMgr(), ecTree)
}

func (ecTree *_ECTree) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Terminated:
		ecTree.eventECTreeAddChild.Close()
		ecTree.eventECTreeRemoveChild.Close()
	}
}

// ResolveContext 解析上下文
func (ecTree *_ECTree) ResolveContext() iface.Cache {
	return ecTree.ctx.ResolveContext()
}

// ResolveCurrentContext 解析当前上下文
func (ecTree *_ECTree) ResolveCurrentContext() iface.Cache {
	return ecTree.ctx.ResolveCurrentContext()
}

// TryAddChild 尝试加入父实体
func (ecTree *_ECTree) TryAddChild(parentId, childId uid.Id) error {
	_, _, err := ecTree.tryAddChild(parentId, childId)
	return err
}

// AddChild 实体加入父实体，在实体加入运行时上下文后才能调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
func (ecTree *_ECTree) AddChild(parentId, childId uid.Id) error {
	parent, child, err := ecTree.tryAddChild(parentId, childId)
	if err != nil {
		return err
	}

	parentNode, ok := ecTree.ecTree[parentId]
	if !ok || parentNode.children == nil {
		parentNode.children = container.NewList[iface.FaceAny](ecTree.ctx.getOptions().FaceAnyAllocator, ecTree.ctx)
		ecTree.ecTree[parentId] = parentNode
	}

	ecTree.ecTree[childId] = _ECNode{
		parent:  parent,
		element: parentNode.children.PushBack(iface.MakeFacePair[any](child, child)),
	}

	_child := ec.UnsafeEntity(child)
	_child.SetECParent(parent)
	_child.SetECNodeState(ec.ECNodeState_Attached)

	emitEventECTreeAddChild(ecTree, ecTree, parent, child)

	return nil
}

// RemoveChild 实体离开父实体，会销毁所有子实体
func (ecTree *_ECTree) RemoveChild(childId uid.Id) {
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

	emitEventECTreeRemoveChild(ecTree, ecTree, node.parent, child)
}

// RangeChildren 遍历子实体
func (ecTree *_ECTree) RangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeChildren 反向遍历子实体
func (ecTree *_ECTree) ReverseRangeChildren(parentId uid.Id, fun generic.Func1[ec.Entity, bool]) {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return
	}

	node.children.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// CountChildren 获取子实体数量
func (ecTree *_ECTree) CountChildren(parentId uid.Id) int {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.children == nil {
		return 0
	}

	return node.children.Len()
}

// GetParent 获取子实体的父实体
func (ecTree *_ECTree) GetParent(childId uid.Id) (ec.Entity, bool) {
	node, ok := ecTree.ecTree[childId]
	if !ok {
		return nil, false
	}

	return node.parent, node.parent != nil
}

// EventECTreeAddChild 事件：EC树中子实体加入父实体
func (ecTree *_ECTree) EventECTreeAddChild() event.IEvent {
	return &ecTree.eventECTreeAddChild
}

// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
func (ecTree *_ECTree) EventECTreeRemoveChild() event.IEvent {
	return &ecTree.eventECTreeRemoveChild
}

func (ecTree *_ECTree) OnEntityMgrRemovingEntity(entityMgr IEntityMgr, entity ec.Entity) {
	ecTree.RemoveChild(entity.GetId())
}

func (ecTree *_ECTree) tryAddChild(parentId, childId uid.Id) (parent, child ec.Entity, err error) {
	if parentId == childId {
		return nil, nil, fmt.Errorf("%w: %w: parentId and childId can't be equal", ErrECTree, exception.ErrArgs)
	}

	parent, ok := ecTree.ctx.GetEntityMgr().GetEntity(parentId)
	if !ok {
		return nil, nil, fmt.Errorf("%w: parent entity not exist", ErrECTree)
	}

	switch parent.GetState() {
	case ec.EntityState_Awake, ec.EntityState_Start, ec.EntityState_Living:
	default:
		return nil, nil, fmt.Errorf("%w: invalid parent entity state %q", ErrECTree, parent.GetState())
	}

	child, ok = ecTree.ctx.GetEntityMgr().GetEntity(childId)
	if !ok {
		return nil, nil, fmt.Errorf("%w: child entity not exist", ErrECTree)
	}

	switch child.GetState() {
	case ec.EntityState_Awake, ec.EntityState_Start, ec.EntityState_Living:
	default:
		return nil, nil, fmt.Errorf("%w: invalid child entity state %q", ErrECTree, parent.GetState())
	}

	if _, ok = ecTree.ecTree[childId]; ok {
		return nil, nil, fmt.Errorf("%w: child entity already in ec-tree", ErrECTree)
	}

	return
}
