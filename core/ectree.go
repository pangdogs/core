package core

import (
	"errors"
	"github.com/pangdogs/galaxy/core/container"
)

type IECTree interface {
	AddChild(parentID uint64, childID uint64) error
	RemoveChild(childID uint64)
	RangeChildren(parentID uint64, fun func(child Entity) bool)
	ReverseRangeChildren(parentID uint64, fun func(child Entity) bool)
	GetChildCount(parentID uint64) int
	GetParent(childID uint64) (Entity, bool)
	EventECTreeAddChild() IEvent
	EventECTreeRemoveChild() IEvent
}

type _ECNode struct {
	Parent          Entity
	ElementInParent *container.Element[FaceAny]
	Children        *container.List[FaceAny]
	Removing        bool
}

type ECTree struct {
	runtimeCtx             RuntimeContext
	masterTree             bool
	ecTree                 map[uint64]_ECNode
	eventECTreeAddChild    Event
	eventECTreeRemoveChild Event
	inited                 bool
	hook                   Hook
}

func (ecTree *ECTree) Init(runtimeCtx RuntimeContext) {
	ecTree.init(runtimeCtx, false)
}

func (ecTree *ECTree) Shut() {
	if !ecTree.masterTree {
		ecTree.hook.Unbind()
	}
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
	ecTree.ecTree = map[uint64]_ECNode{}
	ecTree.eventECTreeAddChild.Init(false, nil, EventRecursion_Discard, RuntimeContextGetOptions(runtimeCtx).HookCache, runtimeCtx)
	ecTree.eventECTreeRemoveChild.Init(false, nil, EventRecursion_Discard, RuntimeContextGetOptions(runtimeCtx).HookCache, runtimeCtx)
	ecTree.inited = true

	if !ecTree.masterTree {
		ecTree.hook = BindEvent[eventEntityMgrNotifyECTreeRemoveEntity[RuntimeContext]](ecTree.runtimeCtx.eventEntityMgrNotifyECTreeRemoveEntity(), ecTree)
	}
}

func (ecTree *ECTree) onEntityMgrNotifyECTreeRemoveEntity(entityMgr RuntimeContext, entity Entity) {
	ecTree.RemoveChild(entity.GetID())
}

func (ecTree *ECTree) AddChild(parentID uint64, childID uint64) error {
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
		IFace: child,
		Cache: IFace2Cache[Entity](child),
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

func (ecTree *ECTree) RemoveChild(childID uint64) {
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
				Cache2IFace[Entity](e.Value.Cache).DestroySelf()
			} else {
				ecTree.RemoveChild(Cache2IFace[Entity](e.Value.Cache).GetID())
			}
			return true
		})
	}

	delete(ecTree.ecTree, childID)
	node.ElementInParent.Escape()

	child := Cache2IFace[Entity](node.ElementInParent.Value.Cache)

	if ecTree.masterTree {
		child.setParent(nil)
	}

	emitEventECTreeRemoveChild(ecTree.EventECTreeRemoveChild(), ecTree, node.Parent, child)
}

func (ecTree *ECTree) RangeChildren(parentID uint64, fun func(child Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return
	}

	node.Children.Traversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

func (ecTree *ECTree) ReverseRangeChildren(parentID uint64, fun func(child Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return
	}

	node.Children.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

func (ecTree *ECTree) GetChildCount(parentID uint64) int {
	node, ok := ecTree.ecTree[parentID]
	if !ok || node.Children == nil {
		return 0
	}

	return node.Children.Len()
}

func (ecTree *ECTree) GetParent(childID uint64) (Entity, bool) {
	node, ok := ecTree.ecTree[childID]
	if !ok {
		return nil, false
	}

	return node.Parent, node.Parent != nil
}

func (ecTree *ECTree) EventECTreeAddChild() IEvent {
	return &ecTree.eventECTreeAddChild
}

func (ecTree *ECTree) EventECTreeRemoveChild() IEvent {
	return &ecTree.eventECTreeRemoveChild
}
