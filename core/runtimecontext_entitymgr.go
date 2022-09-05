package core

import (
	"errors"
	"fmt"
	"github.com/pangdogs/galaxy/core/container"
)

// _RuntimeContextEntityMgr 运行时上下文（Runtime Context）的实体（Entity）管理器
type _RuntimeContextEntityMgr interface {
	// GetEntity 查询实体
	GetEntity(id int64) (Entity, bool)

	// RangeEntities 遍历所有实体
	RangeEntities(func(entity Entity) bool)

	// ReverseRangeEntities 反向遍历所有实体
	ReverseRangeEntities(func(entity Entity) bool)

	// GetEntityCount 获取实体数量
	GetEntityCount() int

	// AddEntity 添加实体
	AddEntity(entity Entity) error

	// RemoveEntity 删除实体
	RemoveEntity(id int64)

	// EventEntityMgrAddEntity 事件：运行时上下文（Runtime Context）添加实体（Entity）
	EventEntityMgrAddEntity() IEvent

	// EventEntityMgrRemoveEntity 事件：运行时上下文（Runtime Context）删除实体（Entity）
	EventEntityMgrRemoveEntity() IEvent

	// EventEntityMgrEntityAddComponents 事件：运行时上下文（Runtime Context）中的实体（Entity）添加组件（Component）
	EventEntityMgrEntityAddComponents() IEvent

	// EventEntityMgrEntityRemoveComponent 事件：运行时上下文（Runtime Context）中的实体（Entity）删除组件（Component）
	EventEntityMgrEntityRemoveComponent() IEvent

	eventEntityMgrNotifyECTreeRemoveEntity() IEvent
}

// GetEntity 查询实体
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

// RangeEntities 遍历所有实体
func (runtimeCtx *_RuntimeContextBehavior) RangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.Traversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities 反向遍历所有实体
func (runtimeCtx *_RuntimeContextBehavior) ReverseRangeEntities(fun func(entity Entity) bool) {
	if fun == nil {
		return
	}

	runtimeCtx.entityList.ReverseTraversal(func(e *container.Element[FaceAny]) bool {
		return fun(Cache2IFace[Entity](e.Value.Cache))
	})
}

// GetEntityCount 获取实体数量
func (runtimeCtx *_RuntimeContextBehavior) GetEntityCount() int {
	return runtimeCtx.entityList.Len()
}

// AddEntity 添加实体
func (runtimeCtx *_RuntimeContextBehavior) AddEntity(entity Entity) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetRuntimeCtx() != nil {
		return errors.New("entity already added in runtime context")
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
		return fmt.Errorf("repeated entity '%d' in this runtime context", entity.GetID())
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

	return nil
}

// RemoveEntity 删除实体
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

// EventEntityMgrAddEntity 事件：运行时上下文（Runtime Context）添加实体（Entity）
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrAddEntity() IEvent {
	return &runtimeCtx.eventEntityMgrAddEntity
}

// EventEntityMgrRemoveEntity 事件：运行时上下文（Runtime Context）删除实体（Entity）
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrRemoveEntity() IEvent {
	return &runtimeCtx.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：运行时上下文（Runtime Context）中的实体（Entity）添加组件（Component）
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrEntityAddComponents() IEvent {
	return &runtimeCtx.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：运行时上下文（Runtime Context）中的实体（Entity）删除组件（Component）
func (runtimeCtx *_RuntimeContextBehavior) EventEntityMgrEntityRemoveComponent() IEvent {
	return &runtimeCtx.eventEntityMgrEntityRemoveComponent
}

func (runtimeCtx *_RuntimeContextBehavior) eventEntityMgrNotifyECTreeRemoveEntity() IEvent {
	return &runtimeCtx._eventEntityMgrNotifyECTreeRemoveEntity
}

// OnCompMgrAddComponents 事件回调：实体的组件管理器加入一些组件
func (runtimeCtx *_RuntimeContextBehavior) OnCompMgrAddComponents(entity Entity, components []Component) {
	for i := range components {
		if components[i].GetID() <= 0 {
			components[i].setID(runtimeCtx.servCtx.GenUID())
		}
	}
	emitEventEntityMgrEntityAddComponents(&runtimeCtx.eventEntityMgrEntityAddComponents, runtimeCtx.opts.Inheritor.IFace, entity, components)
}

// OnCompMgrRemoveComponent 事件回调：实体的组件管理器删除组件
func (runtimeCtx *_RuntimeContextBehavior) OnCompMgrRemoveComponent(entity Entity, component Component) {
	emitEventEntityMgrEntityRemoveComponent(&runtimeCtx.eventEntityMgrEntityRemoveComponent, runtimeCtx.opts.Inheritor.IFace, entity, component)
}
