// Code generated by eventc event; DO NOT EDIT.

package ec

import (
	"fmt"
	event "git.golaxy.org/core/event"
	iface "git.golaxy.org/core/utils/iface"
)

type iAutoEventComponentDestroySelf interface {
	EventComponentDestroySelf() event.IEvent
}

func BindEventComponentDestroySelf(auto iAutoEventComponentDestroySelf, subscriber EventComponentDestroySelf, priority ...int32) event.Hook {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	return event.Bind[EventComponentDestroySelf](auto.EventComponentDestroySelf(), subscriber, priority...)
}

func _EmitEventComponentDestroySelf(auto iAutoEventComponentDestroySelf, comp Component) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventComponentDestroySelf()).Emit(func(subscriber iface.Cache) bool {
		iface.Cache2Iface[EventComponentDestroySelf](subscriber).OnComponentDestroySelf(comp)
		return true
	})
}

func _EmitEventComponentDestroySelfWithInterrupt(auto iAutoEventComponentDestroySelf, interrupt func(comp Component) bool, comp Component) {
	if auto == nil {
		panic(fmt.Errorf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs))
	}
	event.UnsafeEvent(auto.EventComponentDestroySelf()).Emit(func(subscriber iface.Cache) bool {
		if interrupt != nil {
			if interrupt(comp) {
				return false
			}
		}
		iface.Cache2Iface[EventComponentDestroySelf](subscriber).OnComponentDestroySelf(comp)
		return true
	})
}

func HandleEventComponentDestroySelf(fun func(comp Component)) EventComponentDestroySelfHandler {
	return EventComponentDestroySelfHandler(fun)
}

type EventComponentDestroySelfHandler func(comp Component)

func (h EventComponentDestroySelfHandler) OnComponentDestroySelf(comp Component) {
	h(comp)
}
