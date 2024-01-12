// Code generated by eventcode --decl_file=entitymgr_event.go gen_event --package=runtime; DO NOT EDIT.

package runtime

import (
	"fmt"
	event "kit.golaxy.org/golaxy/event"
	iface "kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/ec"
)

type iAutoEventEntityMgrAddEntity interface {
	EventEntityMgrAddEntity() event.IEvent
}

func BindEventEntityMgrAddEntity(auto iAutoEventEntityMgrAddEntity, subscriber EventEntityMgrAddEntity, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrAddEntity](auto.EventEntityMgrAddEntity(), subscriber, priority...)
}

func emitEventEntityMgrAddEntity(auto iAutoEventEntityMgrAddEntity, entityMgr EntityMgr, entity ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrAddEntity()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrAddEntity](subscriber).OnEntityMgrAddEntity(entityMgr, entity)
		return true
	})
}

type iAutoEventEntityMgrRemovingEntity interface {
	EventEntityMgrRemovingEntity() event.IEvent
}

func BindEventEntityMgrRemovingEntity(auto iAutoEventEntityMgrRemovingEntity, subscriber EventEntityMgrRemovingEntity, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrRemovingEntity](auto.EventEntityMgrRemovingEntity(), subscriber, priority...)
}

func emitEventEntityMgrRemovingEntity(auto iAutoEventEntityMgrRemovingEntity, entityMgr EntityMgr, entity ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrRemovingEntity()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrRemovingEntity](subscriber).OnEntityMgrRemovingEntity(entityMgr, entity)
		return true
	})
}

type iAutoEventEntityMgrRemoveEntity interface {
	EventEntityMgrRemoveEntity() event.IEvent
}

func BindEventEntityMgrRemoveEntity(auto iAutoEventEntityMgrRemoveEntity, subscriber EventEntityMgrRemoveEntity, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrRemoveEntity](auto.EventEntityMgrRemoveEntity(), subscriber, priority...)
}

func emitEventEntityMgrRemoveEntity(auto iAutoEventEntityMgrRemoveEntity, entityMgr EntityMgr, entity ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrRemoveEntity()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrRemoveEntity](subscriber).OnEntityMgrRemoveEntity(entityMgr, entity)
		return true
	})
}

type iAutoEventEntityMgrEntityAddComponents interface {
	EventEntityMgrEntityAddComponents() event.IEvent
}

func BindEventEntityMgrEntityAddComponents(auto iAutoEventEntityMgrEntityAddComponents, subscriber EventEntityMgrEntityAddComponents, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrEntityAddComponents](auto.EventEntityMgrEntityAddComponents(), subscriber, priority...)
}

func emitEventEntityMgrEntityAddComponents(auto iAutoEventEntityMgrEntityAddComponents, entityMgr EntityMgr, entity ec.Entity, components []ec.Component) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrEntityAddComponents()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrEntityAddComponents](subscriber).OnEntityMgrEntityAddComponents(entityMgr, entity, components)
		return true
	})
}

type iAutoEventEntityMgrEntityRemoveComponent interface {
	EventEntityMgrEntityRemoveComponent() event.IEvent
}

func BindEventEntityMgrEntityRemoveComponent(auto iAutoEventEntityMgrEntityRemoveComponent, subscriber EventEntityMgrEntityRemoveComponent, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrEntityRemoveComponent](auto.EventEntityMgrEntityRemoveComponent(), subscriber, priority...)
}

func emitEventEntityMgrEntityRemoveComponent(auto iAutoEventEntityMgrEntityRemoveComponent, entityMgr EntityMgr, entity ec.Entity, component ec.Component) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrEntityRemoveComponent()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrEntityRemoveComponent](subscriber).OnEntityMgrEntityRemoveComponent(entityMgr, entity, component)
		return true
	})
}

type iAutoEventEntityMgrEntityFirstAccessComponent interface {
	EventEntityMgrEntityFirstAccessComponent() event.IEvent
}

func BindEventEntityMgrEntityFirstAccessComponent(auto iAutoEventEntityMgrEntityFirstAccessComponent, subscriber EventEntityMgrEntityFirstAccessComponent, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.BindEvent[EventEntityMgrEntityFirstAccessComponent](auto.EventEntityMgrEntityFirstAccessComponent(), subscriber, priority...)
}

func emitEventEntityMgrEntityFirstAccessComponent(auto iAutoEventEntityMgrEntityFirstAccessComponent, entityMgr EntityMgr, entity ec.Entity, component ec.Component) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityMgrEntityFirstAccessComponent()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityMgrEntityFirstAccessComponent](subscriber).OnEntityMgrEntityFirstAccessComponent(entityMgr, entity, component)
		return true
	})
}