package runtime

import (
	"errors"
	"fmt"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// _EntityMgr 运行时上下文的实体管理器
type _EntityMgr interface {
	// GetEntity 查询实体
	GetEntity(id int64) (ec.Entity, bool)

	// RangeEntities 遍历所有实体
	RangeEntities(func(entity ec.Entity) bool)

	// ReverseRangeEntities 反向遍历所有实体
	ReverseRangeEntities(func(entity ec.Entity) bool)

	// GetEntityCount 获取实体数量
	GetEntityCount() int

	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error

	// RemoveEntity 删除实体
	RemoveEntity(id int64)

	// EventEntityMgrAddEntity 事件：运行时上下文添加实体
	EventEntityMgrAddEntity() localevent.IEvent

	// EventEntityMgrRemoveEntity 事件：运行时上下文删除实体
	EventEntityMgrRemoveEntity() localevent.IEvent

	// EventEntityMgrEntityAddComponents 事件：运行时上下文中的实体添加组件
	EventEntityMgrEntityAddComponents() localevent.IEvent

	// EventEntityMgrEntityRemoveComponent 事件：运行时上下文中的实体删除组件
	EventEntityMgrEntityRemoveComponent() localevent.IEvent

	eventEntityMgrNotifyECTreeRemoveEntity() localevent.IEvent
}

// GetEntity 查询实体
func (ctx *ContextBehavior) GetEntity(id int64) (ec.Entity, bool) {
	e, ok := ctx.entityMap[id]
	if !ok {
		return nil, false
	}

	if e.Element.Escaped() {
		return nil, false
	}

	return util.Cache2Iface[ec.Entity](e.Element.Value.Cache), true
}

// RangeEntities 遍历所有实体
func (ctx *ContextBehavior) RangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	ctx.entityList.Traversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// ReverseRangeEntities 反向遍历所有实体
func (ctx *ContextBehavior) ReverseRangeEntities(fun func(entity ec.Entity) bool) {
	if fun == nil {
		return
	}

	ctx.entityList.ReverseTraversal(func(e *container.Element[util.FaceAny]) bool {
		return fun(util.Cache2Iface[ec.Entity](e.Value.Cache))
	})
}

// GetEntityCount 获取实体数量
func (ctx *ContextBehavior) GetEntityCount() int {
	return ctx.entityList.Len()
}

// AddEntity 添加实体
func (ctx *ContextBehavior) AddEntity(entity ec.Entity) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	_entity := ec.UnsafeEntity(entity)

	if _entity.GetContext() != util.NilIfaceCache {
		return errors.New("entity already added in runtime context")
	}

	if entity.GetID() <= 0 {
		_entity.SetID(ctx.serviceCtx.GenUID())
	}

	_entity.SetContext(ctx.opts.Inheritor.Cache)
	_entity.SetGCCollector(ctx.opts.Inheritor.Iface)

	entity.RangeComponents(func(comp ec.Component) bool {
		_comp := ec.UnsafeComponent(comp)

		if _comp.GetID() <= 0 {
			_comp.SetID(ctx.serviceCtx.GenUID())
		}

		_comp.SetPrimary(true)

		return true
	})

	if _, ok := ctx.entityMap[entity.GetID()]; ok {
		return fmt.Errorf("repeated entity '%d' in this runtime context", entity.GetID())
	}

	entityInfo := _EntityInfo{}

	entityInfo.Hooks[0] = localevent.BindEvent[ec.EventCompMgrAddComponents](entity.EventCompMgrAddComponents(), ctx)
	entityInfo.Hooks[1] = localevent.BindEvent[ec.EventCompMgrRemoveComponent](entity.EventCompMgrRemoveComponent(), ctx)

	entityInfo.Element = ctx.entityList.PushBack(util.FaceAny{
		Iface: entity,
		Cache: util.Iface2Cache(entity),
	})

	ctx.entityMap[entity.GetID()] = entityInfo

	ctx.CollectGC(_entity.GetInnerGC())

	emitEventEntityMgrAddEntity(&ctx.eventEntityMgrAddEntity, ctx.opts.Inheritor.Iface, entity)

	return nil
}

// RemoveEntity 删除实体
func (ctx *ContextBehavior) RemoveEntity(id int64) {
	e, ok := ctx.entityMap[id]
	if !ok {
		return
	}

	entity := util.Cache2Iface[ec.Entity](e.Element.Value.Cache)
	_entity := ec.UnsafeEntity(entity)

	if _entity.GetInitialing() || _entity.GetShutting() {
		return
	}

	_entity.SetShutting(true)
	defer _entity.SetShutting(false)

	emitEventEntityMgrNotifyECTreeRemoveEntity(&ctx._eventEntityMgrNotifyECTreeRemoveEntity, ctx.opts.Inheritor.Iface, entity)

	ctx.ecTree.RemoveChild(id)

	delete(ctx.entityMap, id)
	e.Element.Escape()

	for i := range e.Hooks {
		e.Hooks[i].Unbind()
	}

	emitEventEntityMgrRemoveEntity(&ctx.eventEntityMgrRemoveEntity, ctx.opts.Inheritor.Iface, entity)
}

// EventEntityMgrAddEntity 事件：运行时上下文添加实体
func (ctx *ContextBehavior) EventEntityMgrAddEntity() localevent.IEvent {
	return &ctx.eventEntityMgrAddEntity
}

// EventEntityMgrRemoveEntity 事件：运行时上下文删除实体
func (ctx *ContextBehavior) EventEntityMgrRemoveEntity() localevent.IEvent {
	return &ctx.eventEntityMgrRemoveEntity
}

// EventEntityMgrEntityAddComponents 事件：运行时上下文中的实体添加组件
func (ctx *ContextBehavior) EventEntityMgrEntityAddComponents() localevent.IEvent {
	return &ctx.eventEntityMgrEntityAddComponents
}

// EventEntityMgrEntityRemoveComponent 事件：运行时上下文中的实体删除组件
func (ctx *ContextBehavior) EventEntityMgrEntityRemoveComponent() localevent.IEvent {
	return &ctx.eventEntityMgrEntityRemoveComponent
}

func (ctx *ContextBehavior) eventEntityMgrNotifyECTreeRemoveEntity() localevent.IEvent {
	return &ctx._eventEntityMgrNotifyECTreeRemoveEntity
}

// OnCompMgrAddComponents 事件回调：实体的组件管理器加入一些组件
func (ctx *ContextBehavior) OnCompMgrAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		if components[i].GetID() <= 0 {
			ec.UnsafeComponent(components[i]).SetID(ctx.serviceCtx.GenUID())
		}
	}
	emitEventEntityMgrEntityAddComponents(&ctx.eventEntityMgrEntityAddComponents, ctx.opts.Inheritor.Iface, entity, components)
}

// OnCompMgrRemoveComponent 事件回调：实体的组件管理器删除组件
func (ctx *ContextBehavior) OnCompMgrRemoveComponent(entity ec.Entity, component ec.Component) {
	emitEventEntityMgrEntityRemoveComponent(&ctx.eventEntityMgrEntityRemoveComponent, ctx.opts.Inheritor.Iface, entity, component)
}
