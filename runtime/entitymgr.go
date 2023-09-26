package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	internal.ContextResolver
	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.Entity, bool)
	// RangeEntities 遍历所有实体
	RangeEntities(func(entity ec.Entity) bool)
	// ReverseRangeEntities 反向遍历所有实体
	ReverseRangeEntities(func(entity ec.Entity) bool)
	// CountEntities 获取实体数量
	CountEntities() int
	// AddEntity 添加实体
	AddEntity(entity ec.Entity, scope ec.Scope) error
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)
	// EventEntityMgrAddEntity 事件：实体管理器添加实体
	EventEntityMgrAddEntity() event.IEvent
	// EventEntityMgrRemovingEntity 事件：实体管理器开始删除实体
	EventEntityMgrRemovingEntity() event.IEvent
	// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
	EventEntityMgrRemoveEntity() event.IEvent
	// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
	EventEntityMgrEntityAddComponents() event.IEvent
	// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
	EventEntityMgrEntityRemoveComponent() event.IEvent
	// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
	EventEntityMgrEntityFirstAccessComponent() event.IEvent
}

type _EntityInfo struct {
	Element    *container.Element[iface.FaceAny]
	Hooks      [3]event.Hook
	GlobalMark bool
}

type _EntityMgr struct {
	ctx                                      Context
	entityMap                                map[uid.Id]_EntityInfo
	entityList                               container.List[iface.FaceAny]
	eventEntityMgrAddEntity                  event.Event
	eventEntityMgrRemovingEntity             event.Event
	eventEntityMgrRemoveEntity               event.Event
	eventEntityMgrEntityAddComponents        event.Event
	eventEntityMgrEntityRemoveComponent      event.Event
	eventEntityMgrEntityFirstAccessComponent event.Event
}

func (entityMgr *_EntityMgr) Init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityMgr, internal.ErrArgs))
	}

	entityMgr.ctx = ctx
	entityMgr.entityList.Init(ctx.GetFaceAnyAllocator(), ctx)
	entityMgr.entityMap = map[uid.Id]_EntityInfo{}

	entityMgr.eventEntityMgrAddEntity.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
	entityMgr.eventEntityMgrRemovingEntity.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
	entityMgr.eventEntityMgrRemoveEntity.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
	entityMgr.eventEntityMgrEntityAddComponents.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
	entityMgr.eventEntityMgrEntityRemoveComponent.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
	entityMgr.eventEntityMgrEntityFirstAccessComponent.Init(ctx.GetAutoRecover(), ctx.GetReportError(), event.EventRecursion_Allow, ctx.GetHookAllocator(), ctx)
}

// ResolveContext 解析上下文
func (entityMgr *_EntityMgr) ResolveContext() iface.Cache {
	return entityMgr.ctx.ResolveContext()
}

// GetEntity 查询实体
func (entityMgr *_EntityMgr) GetEntity(id uid.Id) (ec.Entity, bool) {
	e, ok := entityMgr.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return iface.Cache2Iface[ec.Entity](e.Element.Value.Cache), true
}

// RangeEntities 遍历所有实体
func (entityMgr *_EntityMgr) RangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	entityMgr.entityList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities 反向遍历所有实体
func (entityMgr *_EntityMgr) ReverseRangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	entityMgr.entityList.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		return fun(iface.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// CountEntities 获取实体数量
func (entityMgr *_EntityMgr) CountEntities() int {
	return entityMgr.entityList.Len()
}

// AddEntity 添加实体
func (entityMgr *_EntityMgr) AddEntity(entity ec.Entity, scope ec.Scope) error {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, internal.ErrArgs))
	}

	switch scope {
	case ec.Scope_Local, ec.Scope_Global:
	default:
		return fmt.Errorf("%w: %w: invalid scope", ErrEntityMgr, internal.ErrArgs)
	}

	if entity.GetState() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityMgr, entity.GetState())
	}

	if !entity.GetId().IsNil() {
		if _, ok := entityMgr.entityMap[entity.GetId()]; ok {
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

	if scope == ec.Scope_Global {
		_, loaded, err := service.Current(entityMgr).GetEntityMgr().GetOrAddEntity(entity)
		if err != nil {
			return err
		}
		if loaded {
			return fmt.Errorf("%w: %w", ErrEntityMgr, err)
		}
	}

	entityInfo := _EntityInfo{}
	entityInfo.Element = entityMgr.entityList.PushBack(iface.NewFacePair[any](entity, entity))
	entityInfo.Hooks[0] = ec.BindEventCompMgrAddComponents(entity, entityMgr)
	entityInfo.Hooks[1] = ec.BindEventCompMgrRemoveComponent(entity, entityMgr)
	if _entity.GetOptions().ComponentAwakeByAccess {
		entityInfo.Hooks[2] = ec.BindEventCompMgrFirstAccessComponent(entity, entityMgr)
	}
	entityInfo.GlobalMark = scope == ec.Scope_Global

	entityMgr.entityMap[entity.GetId()] = entityInfo

	_entity.SetState(ec.EntityState_Entry)

	if _entity.GetGCCollector() == nil {
		_entity.SetGCCollector(entityMgr.ctx)
	}

	emitEventEntityMgrAddEntity(entityMgr, entityMgr, entity)

	return nil
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgr) RemoveEntity(id uid.Id) {
	entityInfo, ok := entityMgr.entityMap[id]
	if !ok {
		return
	}

	entity := ec.UnsafeEntity(iface.Cache2Iface[ec.Entity](entityInfo.Element.Value.Cache))
	if entity.GetState() >= ec.EntityState_Leave {
		return
	}

	entity.SetState(ec.EntityState_Leave)

	if entityInfo.GlobalMark {
		service.Current(entityMgr).GetEntityMgr().RemoveEntity(entity.GetId())
	}

	emitEventEntityMgrRemovingEntity(entityMgr, entityMgr, entity.Entity)

	delete(entityMgr.entityMap, id)
	entityInfo.Element.Escape()

	for i := range entityInfo.Hooks {
		entityInfo.Hooks[i].Unbind()
	}

	emitEventEntityMgrRemoveEntity(entityMgr, entityMgr, entity.Entity)
}

// EventEntityMgrAddEntity 事件：实体管理器添加实体
func (entityMgr *_EntityMgr) EventEntityMgrAddEntity() event.IEvent {
	return &entityMgr.eventEntityMgrAddEntity
}

// EventEntityMgrRemovingEntity 事件：实体管理器开始删除实体
func (entityMgr *_EntityMgr) EventEntityMgrRemovingEntity() event.IEvent {
	return &entityMgr.eventEntityMgrRemovingEntity
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
func (entityMgr *_EntityMgr) EventEntityMgrRemoveEntity() event.IEvent {
	return &entityMgr.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityAddComponents() event.IEvent {
	return &entityMgr.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityRemoveComponent() event.IEvent {
	return &entityMgr.eventEntityMgrEntityRemoveComponent
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
func (entityMgr *_EntityMgr) EventEntityMgrEntityFirstAccessComponent() event.IEvent {
	return &entityMgr.eventEntityMgrEntityFirstAccessComponent
}

func (entityMgr *_EntityMgr) OnCompMgrAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		_comp := ec.UnsafeComponent(components[i])

		if _comp.GetId().IsNil() {
			_comp.SetId(uid.New())
		}
	}

	emitEventEntityMgrEntityAddComponents(entityMgr, entityMgr, entity, components)
}

func (entityMgr *_EntityMgr) OnCompMgrRemoveComponent(entity ec.Entity, component ec.Component) {
	emitEventEntityMgrEntityRemoveComponent(entityMgr, entityMgr, entity, component)
}

func (entityMgr *_EntityMgr) OnCompMgrFirstAccessComponent(entity ec.Entity, component ec.Component) {
	emitEventEntityMgrEntityFirstAccessComponent(entityMgr, entityMgr, entity, component)
}
