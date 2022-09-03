package core

import (
	"fmt"
	"github.com/pangdogs/galaxy/core/container"
)

type _RuntimeContextEntityMgr interface {
	GetEntity(id int64) (Entity, bool)
	RangeEntities(func(entity Entity) bool)
	ReverseRangeEntities(func(entity Entity) bool)
	GetEntityCount() int
	AddEntity(entity Entity)
	RemoveEntity(id int64)
	EventEntityMgrAddEntity() IEvent
	EventEntityMgrRemoveEntity() IEvent
	EventEntityMgrEntityAddComponents() IEvent
	EventEntityMgrEntityRemoveComponent() IEvent
	eventEntityMgrNotifyECTreeRemoveEntity() IEvent
}

// GetEntity ...
func (runtimeCtx *_RuntimeContextBehavior) GetEntity(id int64) (Entity, bool) {
	e, ok := runtimeCtx.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return Cache2IFace[Entity](e.Element.Value.Cache), true
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

	if entity.GetID() <= 0 {
		entity.setID(runtimeCtx.servCtx.GenUID())
	}

	entity.setRuntimeCtx(runtimeCtx.opts.Inheritor.IFace)
	entity.RangeComponents(func(comp Component) bool {
		if comp.GetID() <= 0 {
			comp.setID(runtimeCtx.servCtx.GenUID())
		}
		return true
	})

	if _, ok := runtimeCtx.entityMap[entity.GetID()]; ok {
		panic(fmt.Errorf("repeated entity '{%d}' in this runtime context", entity.GetID()))
	}

	entityInfo := _RuntimeCtxEntityInfo{}

	entityInfo.Hooks[0] = BindEvent[EventCompMgrAddComponents[Entity]](entity.EventCompMgrAddComponents(), runtimeCtx)
	entityInfo.Hooks[1] = BindEvent[EventCompMgrRemoveComponent[Entity]](entity.EventCompMgrRemoveComponent(), runtimeCtx)

	entityInfo.Element = runtimeCtx.entityList.PushBack(FaceAny{
		IFace: entity,
		Cache: IFace2Cache(entity),
	})

	runtimeCtx.entityMap[entity.GetID()] = entityInfo

	runtimeCtx.CollectGC(entity.getGC())

	emitEventEntityMgrAddEntity[RuntimeContext](&runtimeCtx.eventEntityMgrAddEntity, runtimeCtx.opts.Inheritor.IFace, entity)
}

// RemoveEntity ...
func (runtimeCtx *_RuntimeContextBehavior) RemoveEntity(id int64) {
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
		if components[i].GetID() <= 0 {
			components[i].setID(runtimeCtx.servCtx.GenUID())
		}
	}
	emitEventEntityMgrEntityAddComponents(&runtimeCtx.eventEntityMgrEntityAddComponents, runtimeCtx.opts.Inheritor.IFace, entity, components)
}

// OnCompMgrRemoveComponent ...
func (runtimeCtx *_RuntimeContextBehavior) OnCompMgrRemoveComponent(entity Entity, component Component) {
	emitEventEntityMgrEntityRemoveComponent(&runtimeCtx.eventEntityMgrEntityRemoveComponent, runtimeCtx.opts.Inheritor.IFace, entity, component)
}
