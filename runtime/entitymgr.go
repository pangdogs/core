package runtime

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/uid"
)

// EntityMgr 实体管理器接口
type EntityMgr interface {
	concurrent.CurrentContextProvider

	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)
	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.Entity, bool)
	// ContainsEntity 实体是否存在
	ContainsEntity(id uid.Id) bool
	// RangeEntities 遍历所有实体
	RangeEntities(fun generic.Func1[ec.Entity, bool])
	// ReversedRangeEntities 反向遍历所有实体
	ReversedRangeEntities(fun generic.Func1[ec.Entity, bool])
	// FilterEntities 过滤并获取实体
	FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity
	// GetEntities 获取所有实体
	GetEntities() []ec.Entity
	// CountEntities 获取实体数量
	CountEntities() int

	iAutoEventEntityMgrAddEntity                  // 事件：实体管理器添加实体
	iAutoEventEntityMgrRemoveEntity               // 事件：实体管理器删除实体
	iAutoEventEntityMgrEntityAddComponents        // 事件：实体管理器中的实体添加组件
	iAutoEventEntityMgrEntityRemoveComponent      // 事件：实体管理器中的实体删除组件
	iAutoEventEntityMgrEntityFirstAccessComponent // 事件：实体管理器中的实体首次访问组件
}

type _EntityInfo struct {
	element *container.Element[iface.FaceAny]
	hooks   [3]event.Hook
}

type _EntityMgrBehavior struct {
	ctx                                      Context
	entityIdx                                map[uid.Id]_EntityInfo
	entityList                               container.List[iface.FaceAny]
	treeNodes                                map[uid.Id]_TreeNode
	eventEntityMgrAddEntity                  event.Event
	eventEntityMgrRemoveEntity               event.Event
	eventEntityMgrEntityAddComponents        event.Event
	eventEntityMgrEntityRemoveComponent      event.Event
	eventEntityMgrEntityFirstAccessComponent event.Event
	eventEntityTreeAddChild                  event.Event
	eventEntityTreeRemoveChild               event.Event
}

func (mgr *_EntityMgrBehavior) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityMgr, exception.ErrArgs))
	}

	mgr.ctx = ctx
	mgr.entityIdx = map[uid.Id]_EntityInfo{}
	mgr.treeNodes = map[uid.Id]_TreeNode{}

	ctx.ActivateEvent(&mgr.eventEntityMgrAddEntity, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityMgrRemoveEntity, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityMgrEntityAddComponents, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityMgrEntityRemoveComponent, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityMgrEntityFirstAccessComponent, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityTreeAddChild, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.eventEntityTreeRemoveChild, event.EventRecursion_Allow)
}

func (mgr *_EntityMgrBehavior) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Starting:
		mgr.RangeEntities(func(entity ec.Entity) bool {
			_EmitEventEntityMgrAddEntity(mgr, mgr, entity)
			return true
		})
	case RunningState_Terminating:
		mgr.ReversedRangeEntities(func(entity ec.Entity) bool {
			entity.DestroySelf()
			return true
		})
	case RunningState_Terminated:
		mgr.eventEntityMgrAddEntity.Close()
		mgr.eventEntityMgrRemoveEntity.Close()
		mgr.eventEntityMgrEntityAddComponents.Close()
		mgr.eventEntityMgrEntityRemoveComponent.Close()
		mgr.eventEntityMgrEntityFirstAccessComponent.Close()
		mgr.eventEntityTreeAddChild.Close()
		mgr.eventEntityTreeRemoveChild.Close()
	}
}

// GetCurrentContext 获取当前上下文
func (mgr *_EntityMgrBehavior) GetCurrentContext() iface.Cache {
	return mgr.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (mgr *_EntityMgrBehavior) GetConcurrentContext() iface.Cache {
	return mgr.ctx.GetConcurrentContext()
}

// AddEntity 添加实体
func (mgr *_EntityMgrBehavior) AddEntity(entity ec.Entity) error {
	return mgr.addEntity(entity, uid.Nil)
}

// RemoveEntity 删除实体
func (mgr *_EntityMgrBehavior) RemoveEntity(id uid.Id) {
	mgr.removeEntity(id)
}

// GetEntity 查询实体
func (mgr *_EntityMgrBehavior) GetEntity(id uid.Id) (ec.Entity, bool) {
	entityInfo, ok := mgr.entityIdx[id]
	if !ok {
		return nil, false
	}

	if entityInfo.element.Escaped() {
		return nil, false
	}

	return iface.Cache2Iface[ec.Entity](entityInfo.element.Value.Cache), true
}

// ContainsEntity 实体是否存在
func (mgr *_EntityMgrBehavior) ContainsEntity(id uid.Id) bool {
	_, ok := mgr.entityIdx[id]
	return ok
}

// RangeEntities 遍历所有实体
func (mgr *_EntityMgrBehavior) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReversedRangeEntities 反向遍历所有实体
func (mgr *_EntityMgrBehavior) ReversedRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.ReversedTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// FilterEntities 过滤并获取实体
func (mgr *_EntityMgrBehavior) FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	var entities []ec.Entity

	mgr.entityList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		entity := iface.Cache2Iface[ec.Entity](e.Value.Cache)

		if fun.Exec(entity) {
			entities = append(entities, entity)
		}

		return true
	})

	return entities
}

// GetEntities 获取所有实体
func (mgr *_EntityMgrBehavior) GetEntities() []ec.Entity {
	entities := make([]ec.Entity, 0, mgr.entityList.Len())

	mgr.entityList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		entities = append(entities, iface.Cache2Iface[ec.Entity](e.Value.Cache))
		return true
	})

	return entities
}

// CountEntities 获取实体数量
func (mgr *_EntityMgrBehavior) CountEntities() int {
	return mgr.entityList.Len()
}

// EventEntityMgrAddEntity 事件：实体管理器添加实体
func (mgr *_EntityMgrBehavior) EventEntityMgrAddEntity() event.IEvent {
	return &mgr.eventEntityMgrAddEntity
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
func (mgr *_EntityMgrBehavior) EventEntityMgrRemoveEntity() event.IEvent {
	return &mgr.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
func (mgr *_EntityMgrBehavior) EventEntityMgrEntityAddComponents() event.IEvent {
	return &mgr.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
func (mgr *_EntityMgrBehavior) EventEntityMgrEntityRemoveComponent() event.IEvent {
	return &mgr.eventEntityMgrEntityRemoveComponent
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
func (mgr *_EntityMgrBehavior) EventEntityMgrEntityFirstAccessComponent() event.IEvent {
	return &mgr.eventEntityMgrEntityFirstAccessComponent
}

func (mgr *_EntityMgrBehavior) OnComponentMgrAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		comp := components[i]

		if comp.GetId().IsNil() {
			ec.UnsafeComponent(comp).SetId(uid.New())
		}
	}

	_EmitEventEntityMgrEntityAddComponentsWithInterrupt(mgr, func(entityMgr EntityMgr, entity ec.Entity, components []ec.Component) bool {
		return entity.GetState() > ec.EntityState_Living
	}, mgr, entity, components)
}

func (mgr *_EntityMgrBehavior) OnComponentMgrRemoveComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityMgrEntityRemoveComponentWithInterrupt(mgr, func(entityMgr EntityMgr, entity ec.Entity, component ec.Component) bool {
		return entity.GetState() > ec.EntityState_Living
	}, mgr, entity, component)
}

func (mgr *_EntityMgrBehavior) OnComponentMgrFirstAccessComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityMgrEntityFirstAccessComponentWithInterrupt(mgr, func(entityMgr EntityMgr, entity ec.Entity, component ec.Component) bool {
		return entity.GetState() > ec.EntityState_Living
	}, mgr, entity, component)
}

func (mgr *_EntityMgrBehavior) addEntity(entity ec.Entity, parentId uid.Id) error {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	parent, err := func() (ec.Entity, error) {
		if parentId.IsNil() {
			return nil, nil
		}
		parent, ok := mgr.GetEntity(parentId)
		if !ok {
			return nil, fmt.Errorf("%w: parent not exist", ErrEntityMgr)
		}
		if parent.GetState() > ec.EntityState_Living {
			return nil, fmt.Errorf("%w: invalid parent state %q", ErrEntityMgr, parent.GetState())
		}
		if parent.GetId() == entity.GetId() {
			return nil, fmt.Errorf("%w: parent and child cannot be the same", ErrEntityMgr)
		}
		return parent, nil
	}()
	if err != nil {
		return err
	}

	switch entity.GetScope() {
	case ec.Scope_Local, ec.Scope_Global:
	default:
		return fmt.Errorf("%w: %w: invalid scope", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetState() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityMgr, entity.GetState())
	}

	if entity.GetId().IsNil() {
		ec.UnsafeEntity(entity).SetId(uid.New())
	}
	ec.UnsafeEntity(entity).SetContext(iface.Iface2Cache[Context](mgr.ctx))

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetId().IsNil() {
			ec.UnsafeComponent(comp).SetId(uid.New())
		}
		return true
	})

	if mgr.ContainsEntity(entity.GetId()) {
		return fmt.Errorf("%w: entity already exists in entity-mgr", ErrEntityMgr)
	}

	if parent != nil {
		if _, ok := mgr.treeNodes[entity.GetId()]; ok {
			return fmt.Errorf("%w: entity already exists in entity-tree", ErrEntityTree)
		}
	}

	if entity.GetScope() == ec.Scope_Global {
		_, loaded, err := service.Current(mgr).GetEntityMgr().GetOrAddEntity(entity)
		if err != nil {
			return err
		}
		if loaded {
			return fmt.Errorf("%w: %w", ErrEntityMgr, err)
		}
	}

	entityInfo := _EntityInfo{
		element: mgr.entityList.PushBack(iface.MakeFaceAny(entity)),
		hooks: [3]event.Hook{
			ec.BindEventComponentMgrAddComponents(entity, mgr),
			ec.BindEventComponentMgrRemoveComponent(entity, mgr),
		},
	}
	if ec.UnsafeEntity(entity).GetOptions().AwakeOnFirstAccess {
		entityInfo.hooks[2] = ec.BindEventComponentMgrFirstAccessComponent(entity, mgr)
	}
	mgr.entityIdx[entity.GetId()] = entityInfo

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Enter)

	if parent != nil {
		mgr.addToParentNode(entity, parent)
	}

	_EmitEventEntityMgrAddEntityWithInterrupt(mgr, func(entityMgr EntityMgr, entity ec.Entity) bool {
		return entity.GetState() > ec.EntityState_Living
	}, mgr, entity)

	if parent != nil {
		mgr.attachParentNode(entity, parent)
	}

	return nil
}

func (mgr *_EntityMgrBehavior) removeEntity(id uid.Id) {
	entityInfo, ok := mgr.entityIdx[id]
	if !ok {
		return
	}

	entity := iface.Cache2Iface[ec.Entity](entityInfo.element.Value.Cache)

	if entity.GetState() > ec.EntityState_Living {
		return
	}
	ec.UnsafeEntity(entity).SetState(ec.EntityState_Leave)

	if entity.GetTreeNodeState() == ec.TreeNodeState_Attached {
		ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)
	}

	mgr.ReversedRangeChildren(entity.GetId(), func(child ec.Entity) bool {
		child.DestroySelf()
		return true
	})

	mgr.detachParentNode(entity)

	_EmitEventEntityMgrRemoveEntity(mgr, mgr, entity)

	mgr.removeFromParentNode(entity)

	delete(mgr.entityIdx, id)
	entityInfo.element.Escape()
	event.Clean(entityInfo.hooks[:])

	if entity.GetScope() == ec.Scope_Global {
		service.Current(mgr).GetEntityMgr().RemoveEntity(entity.GetId())
	}
}
