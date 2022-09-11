package core

import (
	"errors"
	"github.com/pangdogs/galaxy/core/container"
)

// IECTree EC树接口
type IECTree interface {
	// AddChild 子实体（Entity）加入父实体，在实体加入运行时上下文（Runtime Context）后调用，
	//切换父实体时，先调用RemoveChild()离开旧父实体，在调用AddChild()加入新父实体，非线程安全
	AddChild(parentID, childID int64) error

	// RemoveChild 子实体（Entity）离开父实体，在实体从运行时上下文（Runtime Context）中删除前调用，
	//切换父实体时，先调用RemoveChild()离开旧父实体，在调用AddChild()加入新父实体，非线程安全
	RemoveChild(childID int64)

	// RangeChildren 遍历子实体（Entity），非线程安全
	RangeChildren(parentID int64, fun func(child Entity) bool)

	// ReverseRangeChildren 反向遍历子实体（Entity），非线程安全
	ReverseRangeChildren(parentID int64, fun func(child Entity) bool)

	// GetChildCount 获取子实体（Entity）数量，非线程安全
	GetChildCount(parentID int64) int

	// GetParent 获取子实体（Entity）的父实体，非线程安全
	GetParent(childID int64) (Entity, bool)

	// EventECTreeAddChild 事件：EC树中子实体加入父实体
	EventECTreeAddChild() IEvent

	// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
	EventECTreeRemoveChild() IEvent
}

type _ECNode struct {
	Parent          Entity
	ElementInParent *container.Element[FaceAny]
	Children        *container.List[FaceAny]
	Removing        bool
}

// ECTree EC树，除了运行时上下文（Runtime Context）的主EC树以外，自己创建的EC树全部是对实体（Entity）的引用，我们称为EC引用树，
//主要区别是，从主EC树中删除父实体会递归删除并销毁所有子实体，从EC引用树中删除父实体则仅会递归删除所有子实体。
//同个实体可以同时加入多个EC引用树，这个特性可以实现一些特殊的需求。
type ECTree struct {
	runtimeCtx             RuntimeContext
	masterTree             bool
	ecTree                 map[int64]_ECNode
	eventECTreeAddChild    Event
	eventECTreeRemoveChild Event
	inited                 bool
	hook                   Hook
}

// Init 初始化EC树，非线程安全
func (ecTree *ECTree) Init(runtimeCtx RuntimeContext) {
	ecTree.init(runtimeCtx, false)
}

// Shut 销毁EC树，非线程安全
func (ecTree *ECTree) Shut() {
	if !ecTree.masterTree {
		ecTree.hook.Unbind()
	}
	ecTree.eventECTreeAddChild.Close()
	ecTree.eventECTreeRemoveChild.Close()
}

func (ecTree *ECTree) init(runtimeCtx RuntimeContext, masterTree bool) {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	if ecTree.inited {
		panic("repeated init ec-tree invalid")
	}

	ecTree.runtimeCtx = runtimeCtx
	ecTree.masterTree = masterTree
	ecTree.ecTree = map[int64]_ECNode{}
	ecTree.eventECTreeAddChild.Init(false, nil, EventRecursion_Discard, RuntimeContextGetOptions(runtimeCtx).HookCache, runtimeCtx)
	ecTree.eventECTreeRemoveChild.Init(false, nil, EventRecursion_Discard, RuntimeContextGetOptions(runtimeCtx).HookCache, runtimeCtx)
	ecTree.inited = true

	if !ecTree.masterTree {
		ecTree.hook = BindEvent[eventEntityMgrNotifyECTreeRemoveEntity](ecTree.runtimeCtx.eventEntityMgrNotifyECTreeRemoveEntity(), ecTree)
	}
}

func (ecTree *ECTree) onEntityMgrNotifyECTreeRemoveEntity(runtimeCtx RuntimeContext, entity Entity) {
	ecTree.RemoveChild(entity.GetID())
}

// AddChild 子实体（Entity）加入父实体，在实体加入运行时上下文（Runtime Context）后调用，
//切换父实体时，先调用RemoveChild()离开旧父实体，在调用AddChild()加入新父实体，非线程安全
func (ecTree *ECTree) AddChild(parentID, childID int64) error {
	if parentID == childID {
		return errors.New("parentID equal childID invalid")
	}

	parent, ok := ecTree.runtimeCtx.GetEntity(parentID)
	if !ok {
		return errors.New("parent not exist")
	}

	child, ok := ecTree.runtimeCtx.GetEntity(childID)
	if !ok {
		return errors.New("child not exist")
	}

	if _, ok = ecTree.ecTree[childID]; ok {
		return errors.New("child already in ec-tree")
	}

	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		node.Children = container.NewList[FaceAny](RuntimeContextGetOptions(ecTree.runtimeCtx).FaceCache, ecTree.runtimeCtx)
		ecTree.ecTree[parentID] = node
	}

	element := node.Children.PushBack(FaceAny{
		Iface: child,
		Cache: Iface2Cache[Entity](child),
	})

	ecTree.ecTree[childID] = _ECNode{
		Parent:          parent,
		ElementInParent: element,
	}

	if ecTree.masterTree {
		child.setParent(parent)
	}

	emitEventECTreeAddChild(ecTree.EventECTreeAddChild(), ecTree, parent, child)

	return nil
}

// RemoveChild 子实体（Entity）离开父实体，在实体从运行时上下文（Runtime Context）中删除前调用，
//切换父实体时，先调用RemoveChild()离开旧父实体，在调用AddChild()加入新父实体，非线程安全
func (ecTree *ECTree) RemoveChild(childID int64) {
	node, ok := ecTree.ecTree[childID]
	if !ok {
		return
	}

	if node.Removing {
		return
	}

	node.Removing = true
	ecTree.ecTree[childID] = node

	if node.Children != nil {
		node.Children.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
			if ecTree.masterTree {
				Cache2Iface[Entity](e.Value.Cache).DestroySelf()
			} else {
				ecTree.RemoveChild(Cache2Iface[Entity](e.Value.Cache).GetID())
			}
			return true
		})
	}

	delete(ecTree.ecTree, childID)
	node.ElementInParent.Escape()

	child := Cache2Iface[Entity](node.ElementInParent.Value.Cache)

	if ecTree.masterTree {
		child.setParent(nil)
	}

	emitEventECTreeRemoveChild(ecTree.EventECTreeRemoveChild(), ecTree, node.Parent, child)
}

// RangeChildren 遍历子实体（Entity），非线程安全
func (ecTree *ECTree) RangeChildren(parentID int64, fun func(child Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return
	}

	node.Children.Traversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2Iface[Entity](e.Value.Cache))
	})
}

// ReverseRangeChildren 反向遍历子实体（Entity），非线程安全
func (ecTree *ECTree) ReverseRangeChildren(parentID int64, fun func(child Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return
	}

	node.Children.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2Iface[Entity](e.Value.Cache))
	})
}

// GetChildCount 获取子实体（Entity）数量，非线程安全
func (ecTree *ECTree) GetChildCount(parentID int64) int {
	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return 0
	}

	return node.Children.Len()
}

// GetParent 获取子实体（Entity）的父实体，非线程安全
func (ecTree *ECTree) GetParent(childID int64) (Entity, bool) {
	node, ok := ecTree.ecTree[childID]
	if !ok {
		return nil, false
	}

	return node.Parent, node.Parent != nil
}

// EventECTreeAddChild 事件：EC树中子实体加入父实体
func (ecTree *ECTree) EventECTreeAddChild() IEvent {
	return &ecTree.eventECTreeAddChild
}

// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
func (ecTree *ECTree) EventECTreeRemoveChild() IEvent {
	return &ecTree.eventECTreeRemoveChild
}
