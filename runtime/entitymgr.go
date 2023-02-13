package runtime

import (
	"errors"
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetRuntimeCtx 获取运行时上下文
	GetRuntimeCtx() Context
	// GetEntity 查询实体
	GetEntity(id ec.ID) (ec.Entity, bool)
	// RangeEntities 遍历所有实体
	RangeEntities(func(entity ec.Entity) bool)
	// ReverseRangeEntities 反向遍历所有实体
	ReverseRangeEntities(func(entity ec.Entity) bool)
	// GetEntityCount 获取实体数量
	GetEntityCount() int
	// AddEntity 添加实体
	AddEntity(entity ec.Entity, scope ec.Scope) error
	// RemoveEntity 删除实体
	RemoveEntity(id ec.ID)
	// EventEntityMgrAddEntity 事件：实体管理器添加实体
	EventEntityMgrAddEntity() localevent.IEvent
	// EventEntityMgrRemovingEntity 事件：实体管理器开始删除实体
	EventEntityMgrRemovingEntity() localevent.IEvent
	// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
	EventEntityMgrRemoveEntity() localevent.IEvent
	// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
	EventEntityMgrEntityAddComponents() localevent.IEvent
	// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
	EventEntityMgrEntityRemoveComponent() localevent.IEvent
	// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
	EventEntityMgrEntityFirstAccessComponent() localevent.IEvent
}

type _EntityInfo struct {
	Element    *container.Element[util.FaceAny]
	Hooks      [3]localevent.Hook
	GlobalMark bool
}

type _EntityMgr struct {
	runtimeCtx                               Context
	entityMap                                map[ec.ID]_EntityInfo
	entityList                               container.List[util.FaceAny]
	eventEntityMgrAddEntity                  localevent.Event
	eventEntityMgrRemovingEntity             localevent.Event
	eventEntityMgrRemoveEntity               localevent.Event
	eventEntityMgrEntityAddComponents        localevent.Event
	eventEntityMgrEntityRemoveComponent      localevent.Event
	eventEntityMgrEntityFirstAccessComponent localevent.Event
}

func (entityMgr *_EntityMgr) Init(runtimeCtx Context) {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	entityMgr.runtimeCtx = runtimeCtx
	entityMgr.entityList.Init(runtimeCtx.GetFaceCache(), runtimeCtx)
	entityMgr.entityMap = map[ec.ID]_EntityInfo{}

	entityMgr.eventEntityMgrAddEntity.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
	entityMgr.eventEntityMgrRemovingEntity.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
	entityMgr.eventEntityMgrRemoveEntity.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
	entityMgr.eventEntityMgrEntityAddComponents.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
	entityMgr.eventEntityMgrEntityRemoveComponent.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
	entityMgr.eventEntityMgrEntityFirstAccessComponent.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Allow, runtimeCtx.GetHookCache(), runtimeCtx)
}

// GetRuntimeCtx 获取运行时上下文
func (entityMgr *_EntityMgr) GetRuntimeCtx() Context {
	return entityMgr.runtimeCtx
}

// GetEntity 查询实体
func (entityMgr *_EntityMgr) GetEntity(id ec.ID) (ec.Entity, bool) {
	e, ok := entityMgr.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return util.Cache2Iface[ec.Entity](e.Element.Value.Cache), true
}

// RangeEntities 遍历所有实体
func (entityMgr *_EntityMgr) RangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	entityMgr.entityList.Traversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities 反向遍历所有实体
func (entityMgr *_EntityMgr) ReverseRangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	entityMgr.entityList.ReverseTraversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// GetEntityCount 获取实体数量
func (entityMgr *_EntityMgr) GetEntityCount() int {
	return entityMgr.entityList.Len()
}

// AddEntity 添加实体
func (entityMgr *_EntityMgr) AddEntity(entity ec.Entity, scope ec.Scope) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	switch scope {
	case ec.Scope_Local, ec.Scope_Global:
	default:
		return errors.New("scope invalid")
	}

	if entity.GetState() != ec.EntityState_Birth {
		return errors.New("entity state not birth invalid")
	}

	if entity.GetID() != util.Zero[ec.ID]() {
		if _, ok := entityMgr.entityMap[entity.GetID()]; ok {
			return fmt.Errorf("entity id already existed")
		}
	}

	runtimeCtx := entityMgr.runtimeCtx
	serviceCtx := runtimeCtx.GetServiceCtx()
	_entity := ec.UnsafeEntity(entity)

	if _entity.GetID() == util.Zero[ec.ID]() {
		_entity.SetID(serviceCtx.GenPersistID())
	}
	_entity.SetSerialNo(serviceCtx.GenSerialNo())
	_entity.SetContext(util.Iface2Cache[Context](runtimeCtx))

	_entity.RangeComponents(func(comp ec.Component) bool {
		_comp := ec.UnsafeComponent(comp)

		if _comp.GetID() == util.Zero[ec.ID]() {
			_comp.SetID(serviceCtx.GenPersistID())
		}
		_comp.SetSerialNo(serviceCtx.GenSerialNo())

		return true
	})

	if scope == ec.Scope_Global {
		_, loaded, err := entityMgr.GetRuntimeCtx().GetServiceCtx().GetEntityMgr().GetOrAddEntity(entity)
		if err != nil {
			return err
		}
		if loaded {
			return errors.New("add entity to service context failed, entity id already existed")
		}
	}

	entityInfo := _EntityInfo{}
	entityInfo.Element = entityMgr.entityList.PushBack(util.NewFacePair[any](entity, entity))
	entityInfo.Hooks[0] = localevent.BindEvent[ec.EventCompMgrAddComponents](entity.EventCompMgrAddComponents(), entityMgr)
	entityInfo.Hooks[1] = localevent.BindEvent[ec.EventCompMgrRemoveComponent](entity.EventCompMgrRemoveComponent(), entityMgr)
	if _entity.GetOptions().EnableComponentAwakeByAccess {
		entityInfo.Hooks[2] = localevent.BindEvent[ec.EventCompMgrFirstAccessComponent](entity.EventCompMgrFirstAccessComponent(), entityMgr)
	}
	entityInfo.GlobalMark = scope == ec.Scope_Global

	entityMgr.entityMap[entity.GetID()] = entityInfo

	_entity.SetState(ec.EntityState_Entry)
	_entity.SetGCCollector(runtimeCtx)
	runtimeCtx.CollectGC(_entity.GetInnerGC())

	emitEventEntityMgrAddEntity(&entityMgr.eventEntityMgrAddEntity, entityMgr, entity)

	return nil
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgr) RemoveEntity(id ec.ID) {
	entityInfo, ok := entityMgr.entityMap[id]
	if !ok {
		return
	}

	entity := ec.UnsafeEntity(util.Cache2Iface[ec.Entity](entityInfo.Element.Value.Cache))
	if entity.GetState() >= ec.EntityState_Leave {
		return
	}

	entity.SetState(ec.EntityState_Leave)

	runtimeCtx := entityMgr.runtimeCtx
	serviceCtx := runtimeCtx.GetServiceCtx()

	if entityInfo.GlobalMark {
		serviceCtx.GetEntityMgr().RemoveEntityWithSerialNo(entity.GetID(), entity.GetSerialNo())
	}

	emitEventEntityMgrRemovingEntity(&entityMgr.eventEntityMgrRemovingEntity, entityMgr, entity.Entity)

	delete(entityMgr.entityMap, id)
	entityInfo.Element.Escape()

	for i := range entityInfo.Hooks {
		entityInfo.Hooks[i].Unbind()
	}

	emitEventEntityMgrRemoveEntity(&entityMgr.eventEntityMgrRemoveEntity, entityMgr, entity.Entity)
}

// EventEntityMgrAddEntity 事件：实体管理器添加实体
func (entityMgr *_EntityMgr) EventEntityMgrAddEntity() localevent.IEvent {
	return &entityMgr.eventEntityMgrAddEntity
}

// EventEntityMgrRemovingEntity 事件：实体管理器开始删除实体
func (entityMgr *_EntityMgr) EventEntityMgrRemovingEntity() localevent.IEvent {
	return &entityMgr.eventEntityMgrRemovingEntity
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
func (entityMgr *_EntityMgr) EventEntityMgrRemoveEntity() localevent.IEvent {
	return &entityMgr.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityAddComponents() localevent.IEvent {
	return &entityMgr.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityRemoveComponent() localevent.IEvent {
	return &entityMgr.eventEntityMgrEntityRemoveComponent
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityFirstAccessComponent() localevent.IEvent {
	return &entityMgr.eventEntityMgrEntityFirstAccessComponent
}

func (entityMgr *_EntityMgr) OnCompMgrAddComponents(entity ec.Entity, components []ec.Component) {
	serviceCtx := entityMgr.runtimeCtx.GetServiceCtx()

	for i := range components {
		_comp := ec.UnsafeComponent(components[i])

		if _comp.GetID() == util.Zero[ec.ID]() {
			_comp.SetID(serviceCtx.GenPersistID())
		}
		_comp.SetSerialNo(serviceCtx.GenSerialNo())
	}

	emitEventEntityMgrEntityAddComponents(&entityMgr.eventEntityMgrEntityAddComponents, entityMgr, entity, components)
}

func (entityMgr *_EntityMgr) OnCompMgrRemoveComponent(entity ec.Entity, component ec.Component) {
	emitEventEntityMgrEntityRemoveComponent(&entityMgr.eventEntityMgrEntityRemoveComponent, entityMgr, entity, component)
}

func (entityMgr *_EntityMgr) OnCompMgrFirstAccessComponent(entity ec.Entity, component ec.Component) {
	emitEventEntityMgrEntityFirstAccessComponent(&entityMgr.eventEntityMgrEntityFirstAccessComponent, entityMgr, entity, component)
}
