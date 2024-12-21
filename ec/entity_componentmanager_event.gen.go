/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

// Code generated by eventc event; DO NOT EDIT.

package ec

import (
	event "git.golaxy.org/core/event"
)

type iAutoEventComponentManagerAddComponents interface {
	EventComponentManagerAddComponents() event.IEvent
}

func BindEventComponentManagerAddComponents(auto iAutoEventComponentManagerAddComponents, subscriber EventComponentManagerAddComponents, priority ...int32) event.Hook {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	return event.Bind[EventComponentManagerAddComponents](auto.EventComponentManagerAddComponents(), subscriber, priority...)
}

func _EmitEventComponentManagerAddComponents(auto iAutoEventComponentManagerAddComponents, entity Entity, components []Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerAddComponents()).Emit(func(subscriber event.Cache) bool {
		event.Cache2Iface[EventComponentManagerAddComponents](subscriber).OnComponentManagerAddComponents(entity, components)
		return true
	})
}

func _EmitEventComponentManagerAddComponentsWithInterrupt(auto iAutoEventComponentManagerAddComponents, interrupt func(entity Entity, components []Component) bool, entity Entity, components []Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerAddComponents()).Emit(func(subscriber event.Cache) bool {
		if interrupt != nil {
			if interrupt(entity, components) {
				return false
			}
		}
		event.Cache2Iface[EventComponentManagerAddComponents](subscriber).OnComponentManagerAddComponents(entity, components)
		return true
	})
}

func HandleEventComponentManagerAddComponents(fun func(entity Entity, components []Component)) EventComponentManagerAddComponentsHandler {
	return EventComponentManagerAddComponentsHandler(fun)
}

type EventComponentManagerAddComponentsHandler func(entity Entity, components []Component)

func (h EventComponentManagerAddComponentsHandler) OnComponentManagerAddComponents(entity Entity, components []Component) {
	h(entity, components)
}

type iAutoEventComponentManagerRemoveComponent interface {
	EventComponentManagerRemoveComponent() event.IEvent
}

func BindEventComponentManagerRemoveComponent(auto iAutoEventComponentManagerRemoveComponent, subscriber EventComponentManagerRemoveComponent, priority ...int32) event.Hook {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	return event.Bind[EventComponentManagerRemoveComponent](auto.EventComponentManagerRemoveComponent(), subscriber, priority...)
}

func _EmitEventComponentManagerRemoveComponent(auto iAutoEventComponentManagerRemoveComponent, entity Entity, component Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerRemoveComponent()).Emit(func(subscriber event.Cache) bool {
		event.Cache2Iface[EventComponentManagerRemoveComponent](subscriber).OnComponentManagerRemoveComponent(entity, component)
		return true
	})
}

func _EmitEventComponentManagerRemoveComponentWithInterrupt(auto iAutoEventComponentManagerRemoveComponent, interrupt func(entity Entity, component Component) bool, entity Entity, component Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerRemoveComponent()).Emit(func(subscriber event.Cache) bool {
		if interrupt != nil {
			if interrupt(entity, component) {
				return false
			}
		}
		event.Cache2Iface[EventComponentManagerRemoveComponent](subscriber).OnComponentManagerRemoveComponent(entity, component)
		return true
	})
}

func HandleEventComponentManagerRemoveComponent(fun func(entity Entity, component Component)) EventComponentManagerRemoveComponentHandler {
	return EventComponentManagerRemoveComponentHandler(fun)
}

type EventComponentManagerRemoveComponentHandler func(entity Entity, component Component)

func (h EventComponentManagerRemoveComponentHandler) OnComponentManagerRemoveComponent(entity Entity, component Component) {
	h(entity, component)
}

type iAutoEventComponentManagerFirstTouchComponent interface {
	EventComponentManagerFirstTouchComponent() event.IEvent
}

func BindEventComponentManagerFirstTouchComponent(auto iAutoEventComponentManagerFirstTouchComponent, subscriber EventComponentManagerFirstTouchComponent, priority ...int32) event.Hook {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	return event.Bind[EventComponentManagerFirstTouchComponent](auto.EventComponentManagerFirstTouchComponent(), subscriber, priority...)
}

func _EmitEventComponentManagerFirstTouchComponent(auto iAutoEventComponentManagerFirstTouchComponent, entity Entity, component Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerFirstTouchComponent()).Emit(func(subscriber event.Cache) bool {
		event.Cache2Iface[EventComponentManagerFirstTouchComponent](subscriber).OnComponentManagerFirstTouchComponent(entity, component)
		return true
	})
}

func _EmitEventComponentManagerFirstTouchComponentWithInterrupt(auto iAutoEventComponentManagerFirstTouchComponent, interrupt func(entity Entity, component Component) bool, entity Entity, component Component) {
	if auto == nil {
		event.Panicf("%w: %w: auto is nil", event.ErrEvent, event.ErrArgs)
	}
	event.UnsafeEvent(auto.EventComponentManagerFirstTouchComponent()).Emit(func(subscriber event.Cache) bool {
		if interrupt != nil {
			if interrupt(entity, component) {
				return false
			}
		}
		event.Cache2Iface[EventComponentManagerFirstTouchComponent](subscriber).OnComponentManagerFirstTouchComponent(entity, component)
		return true
	})
}

func HandleEventComponentManagerFirstTouchComponent(fun func(entity Entity, component Component)) EventComponentManagerFirstTouchComponentHandler {
	return EventComponentManagerFirstTouchComponentHandler(fun)
}

type EventComponentManagerFirstTouchComponentHandler func(entity Entity, component Component)

func (h EventComponentManagerFirstTouchComponentHandler) OnComponentManagerFirstTouchComponent(entity Entity, component Component) {
	h(entity, component)
}