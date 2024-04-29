// Code generated by eventcode --decl_file=entitytree_event.go gen_event --package=runtime; DO NOT EDIT.

package runtime

import (
	"fmt"
	event "git.golaxy.org/core/event"
	iface "git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/ec"
)

type iAutoEventEntityTreeAddNode interface {
	EventEntityTreeAddNode() event.IEvent
}

func BindEventEntityTreeAddNode(auto iAutoEventEntityTreeAddNode, subscriber EventEntityTreeAddNode, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.Bind[EventEntityTreeAddNode](auto.EventEntityTreeAddNode(), subscriber, priority...)
}

func _EmitEventEntityTreeAddNode(auto iAutoEventEntityTreeAddNode, entityTree EntityTree, parent, child ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityTreeAddNode()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityTreeAddNode](subscriber).OnEntityTreeAddNode(entityTree, parent, child)
		return true
	})
}

func _EmitEventEntityTreeAddNodeWithInterrupt(auto iAutoEventEntityTreeAddNode, interrupt func(entityTree EntityTree, parent, child ec.Entity) bool, entityTree EntityTree, parent, child ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityTreeAddNode()).Emit(func(subscriber iface.Cache) bool {
		if interrupt != nil {
			if interrupt(entityTree, parent, child) {
				return false
			}
		}
		iface.Cache2Iface[EventEntityTreeAddNode](subscriber).OnEntityTreeAddNode(entityTree, parent, child)
		return true
	})
}

func HandleEventEntityTreeAddNode(fun func(entityTree EntityTree, parent, child ec.Entity)) EventEntityTreeAddNodeHandler {
	return EventEntityTreeAddNodeHandler(fun)
}

type EventEntityTreeAddNodeHandler func(entityTree EntityTree, parent, child ec.Entity)

func (h EventEntityTreeAddNodeHandler) OnEntityTreeAddNode(entityTree EntityTree, parent, child ec.Entity) {
	h(entityTree, parent, child)
}

type iAutoEventEntityTreeRemoveNode interface {
	EventEntityTreeRemoveNode() event.IEvent
}

func BindEventEntityTreeRemoveNode(auto iAutoEventEntityTreeRemoveNode, subscriber EventEntityTreeRemoveNode, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.Bind[EventEntityTreeRemoveNode](auto.EventEntityTreeRemoveNode(), subscriber, priority...)
}

func _EmitEventEntityTreeRemoveNode(auto iAutoEventEntityTreeRemoveNode, entityTree EntityTree, parent, child ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityTreeRemoveNode()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventEntityTreeRemoveNode](subscriber).OnEntityTreeRemoveNode(entityTree, parent, child)
		return true
	})
}

func _EmitEventEntityTreeRemoveNodeWithInterrupt(auto iAutoEventEntityTreeRemoveNode, interrupt func(entityTree EntityTree, parent, child ec.Entity) bool, entityTree EntityTree, parent, child ec.Entity) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventEntityTreeRemoveNode()).Emit(func(subscriber iface.Cache) bool {
		if interrupt != nil {
			if interrupt(entityTree, parent, child) {
				return false
			}
		}
		iface.Cache2Iface[EventEntityTreeRemoveNode](subscriber).OnEntityTreeRemoveNode(entityTree, parent, child)
		return true
	})
}

func HandleEventEntityTreeRemoveNode(fun func(entityTree EntityTree, parent, child ec.Entity)) EventEntityTreeRemoveNodeHandler {
	return EventEntityTreeRemoveNodeHandler(fun)
}

type EventEntityTreeRemoveNodeHandler func(entityTree EntityTree, parent, child ec.Entity)

func (h EventEntityTreeRemoveNodeHandler) OnEntityTreeRemoveNode(entityTree EntityTree, parent, child ec.Entity) {
	h(entityTree, parent, child)
}