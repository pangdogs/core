package core

import (
	"fmt"
	"github.com/pangdogs/galaxy/core/container"
)

func (runtimeCtx *RuntimeContextBehavior) GetEntity(id uint64) (Entity, bool) {
	e, ok := runtimeCtx.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return Cache2IFace[Entity](e.Element.Value.Cache), true
}

func (runtimeCtx *RuntimeContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.Traversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

func (runtimeCtx *RuntimeContextBehavior) ReverseRangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

func (runtimeCtx *RuntimeContextBehavior) AddEntity(entity Entity) {
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

	entityInfo := _RuntimeCtxEntityInfo{}

	entityInfo.Hooks[0] = BindEvent[EventCompMgrAddComponents[Entity]](entity.EventCompMgrAddComponents(), runtimeCtx)
	entityInfo.Hooks[1] = BindEvent[EventCompMgrRemoveComponent[Entity]](entity.EventCompMgrRemoveComponent(), runtimeCtx)

	entityInfo.Element = runtimeCtx.entityList.PushBack(FaceAny{
		IFace: entity,
		Cache: IFace2Cache(entity),
	})

	runtimeCtx.entityMap[entity.GetID()] = entityInfo

	runtimeCtx.CollectGC(entity)

	emitEventEntityMgrAddEntity[RuntimeContext](&runtimeCtx.eventEntityMgrAddEntity, runtimeCtx.opts.Inheritor.IFace, entity)
}

func (runtimeCtx *RuntimeContextBehavior) RemoveEntity(id uint64) {
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

	runtimeCtx.ecTree.RemoveChild(id)

	delete(runtimeCtx.entityMap, id)
	e.Element.Escape()

	for i := range e.Hooks {
		e.Hooks[i].Unbind()
	}

	emitEventEntityMgrRemoveEntity[RuntimeContext](&runtimeCtx.eventEntityMgrRemoveEntity, runtimeCtx.opts.Inheritor.IFace, entity)
}

func (runtimeCtx *RuntimeContextBehavior) GetEntityCount() int {
	return runtimeCtx.entityList.Len()
}

func (runtimeCtx *RuntimeContextBehavior) EventEntityMgrAddEntity() IEvent {
	return &runtimeCtx.eventEntityMgrAddEntity
}

func (runtimeCtx *RuntimeContextBehavior) EventEntityMgrRemoveEntity() IEvent {
	return &runtimeCtx.eventEntityMgrRemoveEntity
}

func (runtimeCtx *RuntimeContextBehavior) EventEntityMgrEntityAddComponents() IEvent {
	return &runtimeCtx.eventEntityMgrEntityAddComponents
}

func (runtimeCtx *RuntimeContextBehavior) EventEntityMgrEntityRemoveComponent() IEvent {
	return &runtimeCtx.eventEntityMgrEntityRemoveComponent
}

func (runtimeCtx *RuntimeContextBehavior) OnCompMgrAddComponents(entity Entity, components []Component) {
	for i := range components {
		components[i].setID(runtimeCtx.servCtx.genUID())
	}
	emitEventEntityMgrEntityAddComponents(&runtimeCtx.eventEntityMgrEntityAddComponents, runtimeCtx.opts.Inheritor.IFace, entity, components)
}

func (runtimeCtx *RuntimeContextBehavior) OnCompMgrRemoveComponent(entity Entity, component Component) {
	emitEventEntityMgrEntityRemoveComponent(&runtimeCtx.eventEntityMgrEntityRemoveComponent, runtimeCtx.opts.Inheritor.IFace, entity, component)
}
