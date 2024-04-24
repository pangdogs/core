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

	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.Entity, bool)
	// RangeEntities 遍历所有实体
	RangeEntities(fun generic.Func1[ec.Entity, bool])
	// ReverseRangeEntities 反向遍历所有实体
	ReverseRangeEntities(fun generic.Func1[ec.Entity, bool])
	// CountEntities 获取实体数量
	CountEntities() int
	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)

	_AutoEventEntityMgrAddEntity                  // 事件：实体管理器添加实体
	_AutoEventEntityMgrRemoveEntity               // 事件：实体管理器删除实体
	_AutoEventEntityMgrEntityAddComponents        // 事件：实体管理器中的实体添加组件
	_AutoEventEntityMgrEntityRemoveComponent      // 事件：实体管理器中的实体删除组件
	_AutoEventEntityMgrEntityFirstAccessComponent // 事件：实体管理器中的实体首次访问组件
}

type _EntityInfo struct {
	element *container.Element[iface.FaceAny]
	hooks   [3]event.Hook
}

type _EntityMgrBehavior struct {
	ctx                                      Context
	entityIdx                                map[uid.Id]_EntityInfo
	entityList                               container.List[iface.FaceAny]
	eventEntityMgrAddEntity                  event.Event
	eventEntityMgrRemoveEntity               event.Event
	eventEntityMgrEntityAddComponents        event.Event
	eventEntityMgrEntityRemoveComponent      event.Event
	eventEntityMgrEntityFirstAccessComponent event.Event
}

func (entityMgr *_EntityMgrBehavior) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityMgr, exception.ErrArgs))
	}

	entityMgr.ctx = ctx
	entityMgr.entityIdx = map[uid.Id]_EntityInfo{}

	ctx.ActivateEvent(&entityMgr.eventEntityMgrAddEntity, event.EventRecursion_Allow)
	ctx.ActivateEvent(&entityMgr.eventEntityMgrRemoveEntity, event.EventRecursion_Allow)
	ctx.ActivateEvent(&entityMgr.eventEntityMgrEntityAddComponents, event.EventRecursion_Allow)
	ctx.ActivateEvent(&entityMgr.eventEntityMgrEntityRemoveComponent, event.EventRecursion_Allow)
	ctx.ActivateEvent(&entityMgr.eventEntityMgrEntityFirstAccessComponent, event.EventRecursion_Allow)
}

func (entityMgr *_EntityMgrBehavior) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Starting:
		entityMgr.RangeEntities(func(entity ec.Entity) bool {
			_EmitEventEntityMgrAddEntity(entityMgr, entityMgr, entity)
			return true
		})
	case RunningState_Terminating:
		entityMgr.ReverseRangeEntities(func(entity ec.Entity) bool {
			entity.DestroySelf()
			return true
		})
	case RunningState_Terminated:
		entityMgr.eventEntityMgrAddEntity.Close()
		entityMgr.eventEntityMgrRemoveEntity.Close()
		entityMgr.eventEntityMgrEntityAddComponents.Close()
		entityMgr.eventEntityMgrEntityRemoveComponent.Close()
		entityMgr.eventEntityMgrEntityFirstAccessComponent.Close()
	}
}

// GetCurrentContext 获取当前上下文
func (entityMgr *_EntityMgrBehavior) GetCurrentContext() iface.Cache {
	return entityMgr.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (entityMgr *_EntityMgrBehavior) GetConcurrentContext() iface.Cache {
	return entityMgr.ctx.GetConcurrentContext()
}

// GetEntity 查询实体
func (entityMgr *_EntityMgrBehavior) GetEntity(id uid.Id) (ec.Entity, bool) {
	entityInfo, ok := entityMgr.entityIdx[id]
	if !ok {
		return nil, false
	}

	if entityInfo.element.Escaped() {
		return nil, false
	}

	return iface.Cache2Iface[ec.Entity](entityInfo.element.Value.Cache), true
}

// RangeEntities 遍历所有实体
func (entityMgr *_EntityMgrBehavior) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	entityMgr.entityList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities 反向遍历所有实体
func (entityMgr *_EntityMgrBehavior) ReverseRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	entityMgr.entityList.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// CountEntities 获取实体数量
func (entityMgr *_EntityMgrBehavior) CountEntities() int {
	return entityMgr.entityList.Len()
}

// AddEntity 添加实体
func (entityMgr *_EntityMgrBehavior) AddEntity(entity ec.Entity) error {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs))
	}

	switch entity.GetScope() {
	case ec.Scope_Local, ec.Scope_Global:
	default:
		return fmt.Errorf("%w: %w: invalid scope", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetState() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityMgr, entity.GetState())
	}

	if !entity.GetId().IsNil() {
		if _, ok := entityMgr.entityIdx[entity.GetId()]; ok {
			return fmt.Errorf("%w: entity already existed", ErrEntityMgr)
		}
	}

	_entity := ec.UnsafeEntity(entity)

	if _entity.GetId().IsNil() {
		_entity.SetId(uid.New())
	}
	_entity.SetContext(iface.Iface2Cache[Context](entityMgr.ctx))

	_entity.RangeComponents(func(comp ec.Component) bool {
		_comp := ec.UnsafeComponent(comp)

		if _comp.GetId().IsNil() {
			_comp.SetId(uid.New())
		}

		return true
	})

	if entity.GetScope() == ec.Scope_Global {
		_, loaded, err := service.Current(entityMgr).GetEntityMgr().GetOrAddEntity(entity)
		if err != nil {
			return err
		}
		if loaded {
			return fmt.Errorf("%w: %w", ErrEntityMgr, err)
		}
	}

	entityInfo := _EntityInfo{
		element: entityMgr.entityList.PushBack(iface.MakeFaceAny(entity)),
		hooks:   [3]event.Hook{ec.BindEventCompMgrAddComponents(entity, entityMgr), ec.BindEventCompMgrRemoveComponent(entity, entityMgr)},
	}
	if _entity.GetOptions().AwakeOnFirstAccess {
		entityInfo.hooks[2] = ec.BindEventCompMgrFirstAccessComponent(entity, entityMgr)
	}

	entityMgr.entityIdx[entity.GetId()] = entityInfo

	_entity.SetState(ec.EntityState_Enter)

	_EmitEventEntityMgrAddEntity(entityMgr, entityMgr, entity)
	return nil
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgrBehavior) RemoveEntity(id uid.Id) {
	entityInfo, ok := entityMgr.entityIdx[id]
	if !ok {
		return
	}

	entity := ec.UnsafeEntity(iface.Cache2Iface[ec.Entity](entityInfo.element.Value.Cache))
	if entity.GetState() >= ec.EntityState_Leave {
		return
	}

	entity.SetState(ec.EntityState_Leave)

	delete(entityMgr.entityIdx, id)
	entityInfo.element.Escape()

	for i := range entityInfo.hooks {
		entityInfo.hooks[i].Unbind()
	}

	if entity.GetScope() == ec.Scope_Global {
		service.Current(entityMgr).GetEntityMgr().RemoveEntity(entity.GetId())
	}

	_EmitEventEntityMgrRemoveEntity(entityMgr, entityMgr, entity.Entity)
}

// EventEntityMgrAddEntity 事件：实体管理器添加实体
func (entityMgr *_EntityMgrBehavior) EventEntityMgrAddEntity() event.IEvent {
	return &entityMgr.eventEntityMgrAddEntity
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
func (entityMgr *_EntityMgrBehavior) EventEntityMgrRemoveEntity() event.IEvent {
	return &entityMgr.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
func (entityMgr *_EntityMgrBehavior) EventEntityMgrEntityAddComponents() event.IEvent {
	return &entityMgr.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
func (entityMgr *_EntityMgrBehavior) EventEntityMgrEntityRemoveComponent() event.IEvent {
	return &entityMgr.eventEntityMgrEntityRemoveComponent
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
func (entityMgr *_EntityMgrBehavior) EventEntityMgrEntityFirstAccessComponent() event.IEvent {
	return &entityMgr.eventEntityMgrEntityFirstAccessComponent
}

func (entityMgr *_EntityMgrBehavior) OnCompMgrAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		_comp := ec.UnsafeComponent(components[i])

		if _comp.GetId().IsNil() {
			_comp.SetId(uid.New())
		}
	}

	_EmitEventEntityMgrEntityAddComponents(entityMgr, entityMgr, entity, components)
}

func (entityMgr *_EntityMgrBehavior) OnCompMgrRemoveComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityMgrEntityRemoveComponent(entityMgr, entityMgr, entity, component)
}

func (entityMgr *_EntityMgrBehavior) OnCompMgrFirstAccessComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityMgrEntityFirstAccessComponent(entityMgr, entityMgr, entity, component)
}
