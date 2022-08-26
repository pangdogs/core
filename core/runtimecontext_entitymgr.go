package core

import (
	"fmt"
	"github.com/pangdogs/galaxy/core/container"
)

type _EntityQuery interface {
	GetEntity(id uint64) (Entity, bool)
	GetEntityByPersistID(persistID string) (Entity, bool)
	RangeEntities(func(entity Entity) bool)
}

type _EntityReverseQuery interface {
	ReverseRangeEntities(func(entity Entity) bool)
}

type _EntityCountQuery interface {
	GetEntityCount() int
}

type _EntityMgr interface {
	_EntityQuery
	AddEntity(entity Entity)
	RemoveEntity(id uint64)
}

type _EntityMgrEvents interface {
	EventEntityMgrAddEntity() IEvent
	EventEntityMgrRemoveEntity() IEvent
	EventEntityMgrEntityAddComponents() IEvent
	EventEntityMgrEntityRemoveComponent() IEvent
	eventEntityMgrNotifyECTreeRemoveEntity() IEvent
}

// GetEntity ...
func (runtimeCtx *_RuntimeContextBehavior) GetEntity(id uint64) (Entity, bool) {
	e, ok := runtimeCtx.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return Cache2IFace[Entity](e.Element.Value.Cache), true
}

// GetEntityByPersistID ...
func (runtimeCtx *_RuntimeContextBehavior) GetEntityByPersistID(persistID string) (Entity, bool) {
	entity, ok := runtimeCtx.persistentEntityMap[persistID]
	return entity, ok
}

// RangeEntities ...
func (runtimeCtx *_RuntimeContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.Traversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities ...
func (runtimeCtx *_RuntimeContextBehavior) ReverseRangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

// AddEntity ...
func (runtimeCtx *_RuntimeContextBehavior) AddEntity(entity Entity) {
	if entity == nil {
		panic("nil entity")
	}

	if entity.GetRuntimeCtx() != nil {
		panic("entity already added in runtime context")
	}

	entity.setID(runtimeCtx.servCtx.genUID())
	entity.setRuntimeCtx(runtimeCtx.opts.Inheritor.IFace)
	entity.RangeComponents(func(comp Component) bool {
		comp.setID(runtimeCtx.servCtx.genUID())
		return true
	})

	if _, ok := runtimeCtx.entityMap[entity.GetID()]; ok {
		panic(fmt.Errorf("repeated entity '{%d}' in this runtime context", entity.GetID()))
	}

	if entity.GetPersistID() != "" {
		if _, ok := runtimeCtx.persistentEntityMap[entity.GetPersistID()]; ok {
			panic(fmt.Errorf("repeated persistent entity '{%s}' in this runtime context", entity.GetPersistID()))
		}
	}

	entityInfo := _RuntimeCtxEntityInfo{}

	entityInfo.Hooks[0] = BindEvent[EventCompMgrAddComponents[Entity]](entity.EventCompMgrAddComponents(), runtimeCtx)
	entityInfo.Hooks[1] = BindEvent[EventCompMgrRemoveComponent[Entity]](entity.EventCompMgrRemoveComponent(), runtimeCtx)

	entityInfo.Element = runtimeCtx.entityList.PushBack(FaceAny{
		IFace: entity,
		Cache: IFace2Cache(entity),
	})

	runtimeCtx.entityMap[entity.GetID()] = entityInfo

	if entity.GetPersistID() != "" {
		runtimeCtx.persistentEntityMap[entity.GetPersistID()] = entity
	}

	runtimeCtx.CollectGC(entity)

	emitEventEntityMgrAddEntity[RuntimeContext](&runtimeCtx.eventEntityMgrAddEntity, runtimeCtx.opts.Inheritor.IFace, entity)
}

// RemoveEntity ...
func (runtimeCtx *_RuntimeContextBehavior) RemoveEntity(id uint64) {
	e, ok := runtimeCtx.entityMap[id]
	if !ok {
		return
	}

	entity := Cache2IFace[Entity](e.Element.Value.Cache)
	if entity.getInitialing() || entity.getShutting() {
		return
	}

	entity.setShutting(true)
	defer entity.setShutting(false)

	emitEventEntityMgrNotifyECTreeRemoveEntity[RuntimeContext](&runtimeCtx._eventEntityMgrNotifyECTreeRemoveEntity, runtimeCtx.opts.Inheritor.IFace, entity)

	runtimeCtx.ecTree.RemoveChild(id)

	delete(runtimeCtx.entityMap, id)
	e.Element.Escape()

	if entity.GetPersistID() != "" {
		delete(runtimeCtx.persistentEntityMap, entity.GetPersistID())
	}

	for i := range e.Hooks {
		e.Hooks[i].Unbind()
	}

	emitEventEntityMgrRemoveEntity[RuntimeContext](&runtimeCtx.eventEntityMgrRemoveEntity, runtimeCtx.opts.Inheritor.IFace, entity)
}

// GetEntityCount ...
func (runtimeCtx *_RuntimeContextBehavior) GetEntityCount() int {
	return runtimeCtx.entityList.Len()
}

// EventEntityMgrAddEntity ...
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrAddEntity() IEvent {
	return &runtimeCtx.eventEntityMgrAddEntity
}

// EventEntityMgrRemoveEntity ...
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrRemoveEntity() IEvent {
	return &runtimeCtx.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents ...
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrEntityAddComponents() IEvent {
	return &runtimeCtx.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent ...
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrEntityRemoveComponent() IEvent {
	return &runtimeCtx.eventEntityMgrEntityRemoveComponent
}

func (runtimeCtx *_RuntimeContextBehavior) eventEntityMgrNotifyECTreeRemoveEntity() IEvent {
	return &runtimeCtx._eventEntityMgrNotifyECTreeRemoveEntity
}

// OnCompMgrAddComponents ...
func (runtimeCtx *_RuntimeContextBehavior) OnCompMgrAddComponents(entity Entity, components []Component) {
	for i := range components {
		components[i].setID(runtimeCtx.servCtx.genUID())
	}
	emitEventEntityMgrEntityAddComponents(&runtimeCtx.eventEntityMgrEntityAddComponents, runtimeCtx.opts.Inheritor.IFace, entity, components)
}

// OnCompMgrRemoveComponent ...
func (runtimeCtx *_RuntimeContextBehavior) OnCompMgrRemoveComponent(entity Entity, component Component) {
	emitEventEntityMgrEntityRemoveComponent(&runtimeCtx.eventEntityMgrEntityRemoveComponent, runtimeCtx.opts.Inheritor.IFace, entity, component)
}
