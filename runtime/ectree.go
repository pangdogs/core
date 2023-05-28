package runtime

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// IECTree EC树接口
type IECTree interface {
	internal.ContextResolver
	// AddChild 子实体加入父实体，在实体加入运行时上下文后调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
	AddChild(parentId, childId uid.Id) error
	// RemoveChild 子实体离开父实体，在实体从运行时上下文中删除前调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
	RemoveChild(childId uid.Id)
	// RangeChildren 遍历子实体
	RangeChildren(parentId uid.Id, fun func(child ec.Entity) bool)
	// ReverseRangeChildren 反向遍历子实体
	ReverseRangeChildren(parentId uid.Id, fun func(child ec.Entity) bool)
	// GetChildCount 获取子实体数量
	GetChildCount(parentId uid.Id) int
	// GetParent 获取子实体的父实体
	GetParent(childId uid.Id) (ec.Entity, bool)
	// EventECTreeAddChild 事件：EC树中子实体加入父实体
	EventECTreeAddChild() localevent.IEvent
	// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
	EventECTreeRemoveChild() localevent.IEvent
}

type _ECNode struct {
	Parent          ec.Entity
	ElementInParent *container.Element[util.FaceAny]
	Children        *container.List[util.FaceAny]
	Removing        bool
}

// ECTree EC树，除了运行时上下文的主EC树以外，自己创建的EC树全部是对实体的引用，称为EC引用树，
// 主要区别是，从主EC树中删除父实体会递归删除并销毁所有子实体，从EC引用树中删除父实体则仅会递归删除所有子实体。
// 同个实体可以同时加入多个EC引用树，这个特性可以实现一些特殊的需求。
type ECTree struct {
	ctx                    Context
	masterTree             bool
	ecTree                 map[uid.Id]_ECNode
	eventECTreeAddChild    localevent.Event
	eventECTreeRemoveChild localevent.Event
	hook                   localevent.Hook
	inited                 bool
}

// Init 初始化EC树
func (ecTree *ECTree) Init(ctx Context) {
	ecTree.init(ctx, false)
}

// Shut 销毁EC树
func (ecTree *ECTree) Shut() {
	ecTree.hook.Unbind()
	ecTree.eventECTreeAddChild.Close()
	ecTree.eventECTreeRemoveChild.Close()
}

func (ecTree *ECTree) init(ctx Context, masterTree bool) {
	if ctx == nil {
		panic("nil ctx")
	}

	if ecTree.inited {
		panic("ec-tree initialized")
	}

	ecTree.inited = true

	ecTree.ctx = ctx
	ecTree.masterTree = masterTree
	ecTree.ecTree = map[uid.Id]_ECNode{}
	ecTree.eventECTreeAddChild.Init(ctx.GetAutoRecover(), ctx.GetReportError(), localevent.EventRecursion_Allow, ctx.getOptions().HookAllocator, ctx)
	ecTree.eventECTreeRemoveChild.Init(ctx.GetAutoRecover(), ctx.GetReportError(), localevent.EventRecursion_Allow, ctx.getOptions().HookAllocator, ctx)

	var priority int32
	if !ecTree.masterTree {
		priority = -1
	}
	ecTree.hook = localevent.BindEventWithPriority[EventEntityMgrRemovingEntity](ecTree.ctx.GetEntityMgr().EventEntityMgrRemovingEntity(), ecTree, priority)
}

func (ecTree *ECTree) OnEntityMgrRemovingEntity(entityMgr IEntityMgr, entity ec.Entity) {
	ecTree.RemoveChild(entity.GetId())
}

// ResolveContext 解析上下文
func (ecTree *ECTree) ResolveContext() util.IfaceCache {
	return ecTree.ctx.ResolveContext()
}

// AddChild 子实体加入父实体，在实体加入运行时上下文后调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
func (ecTree *ECTree) AddChild(parentId, childId uid.Id) error {
	if parentId == childId {
		return errors.New("parentId equal childId is invalid")
	}

	parent, ok := ecTree.ctx.GetEntityMgr().GetEntity(parentId)
	if !ok {
		return errors.New("parent not exist")
	}

	switch parent.GetState() {
	case ec.EntityState_Init, ec.EntityState_Inited, ec.EntityState_Living:
	default:
		return errors.New("parent state not init or start or living is invalid")
	}

	child, ok := ecTree.ctx.GetEntityMgr().GetEntity(childId)
	if !ok {
		return errors.New("child not exist")
	}

	switch child.GetState() {
	case ec.EntityState_Init, ec.EntityState_Inited, ec.EntityState_Living:
	default:
		return errors.New("parent state not init or start or living is invalid")
	}

	if _, ok = ecTree.ecTree[childId]; ok {
		return errors.New("child id already existed")
	}

	node, ok := ecTree.ecTree[parentId]
	if !ok || node.Children == nil {
		node.Children = container.NewList[util.FaceAny](ecTree.ctx.getOptions().FaceAnyAllocator, ecTree.ctx)
		ecTree.ecTree[parentId] = node
	}

	element := node.Children.PushBack(util.NewFacePair[any](child, child))

	ecTree.ecTree[childId] = _ECNode{
		Parent:          parent,
		ElementInParent: element,
	}

	if ecTree.masterTree {
		ec.UnsafeEntity(child).SetParent(parent)
	}

	emitEventECTreeAddChild(ecTree.EventECTreeAddChild(), ecTree, parent, child)

	return nil
}

// RemoveChild 子实体离开父实体，在实体从运行时上下文中删除前调用，切换父实体时，先调用RemoveChild()离开旧父实体，再调用AddChild()加入新父实体
func (ecTree *ECTree) RemoveChild(childId uid.Id) {
	node, ok := ecTree.ecTree[childId]
	if !ok {
		return
	}

	if node.Removing {
		return
	}

	node.Removing = true
	ecTree.ecTree[childId] = node

	if node.Children != nil {
		node.Children.ReverseTraversal(func(e *container.Element[util.FaceAny]) bool {
			if ecTree.masterTree {
				util.Cache2Iface[ec.Entity](e.Value.Cache).DestroySelf()
			} else {
				ecTree.RemoveChild(util.Cache2Iface[ec.Entity](e.Value.Cache).GetId())
			}
			return true
		})
	}

	delete(ecTree.ecTree, childId)
	node.ElementInParent.Escape()

	child := util.Cache2Iface[ec.Entity](node.ElementInParent.Value.Cache)

	if ecTree.masterTree {
		ec.UnsafeEntity(child).SetParent(nil)
	}

	emitEventECTreeRemoveChild(ecTree.EventECTreeRemoveChild(), ecTree, node.Parent, child)
}

// RangeChildren 遍历子实体
func (ecTree *ECTree) RangeChildren(parentId uid.Id, fun func(child ec.Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentId]
	if !ok || node.Children == nil {
		return
	}

	node.Children.Traversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeChildren 反向遍历子实体
func (ecTree *ECTree) ReverseRangeChildren(parentId uid.Id, fun func(child ec.Entity) bool) {
	if fun == nil {
		return
	}

	node, ok := ecTree.ecTree[parentId]
	if !ok || node.Children == nil {
		return
	}

	node.Children.ReverseTraversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// GetChildCount 获取子实体数量
func (ecTree *ECTree) GetChildCount(parentId uid.Id) int {
	node, ok := ecTree.ecTree[parentId]
	if !ok || node.Children == nil {
		return 0
	}

	return node.Children.Len()
}

// GetParent 获取子实体的父实体
func (ecTree *ECTree) GetParent(childId uid.Id) (ec.Entity, bool) {
	node, ok := ecTree.ecTree[childId]
	if !ok {
		return nil, false
	}

	return node.Parent, node.Parent != nil
}

// EventECTreeAddChild 事件：EC树中子实体加入父实体
func (ecTree *ECTree) EventECTreeAddChild() localevent.IEvent {
	return &ecTree.eventECTreeAddChild
}

// EventECTreeRemoveChild 事件：EC树中子实体离开父实体
func (ecTree *ECTree) EventECTreeRemoveChild() localevent.IEvent {
	return &ecTree.eventECTreeRemoveChild
}
