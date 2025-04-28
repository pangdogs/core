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

package event

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
)

// EventRecursion 发生事件递归时的处理方式（事件递归：事件发送过程中，在订阅者的逻辑中，再次发送这个事件）
type EventRecursion int32

const (
	EventRecursion_Allow    EventRecursion = iota // 允许事件递归，可能会无限递归
	EventRecursion_Disallow                       // 不允许事件递归，递归时会panic
	EventRecursion_Discard                        // 丢弃递归的事件，不会再发送给任何订阅者
	EventRecursion_Truncate                       // 截断递归的事件，不会再发送给当前订阅者，但是会发送给其他订阅者
	EventRecursion_Deepest                        // 深度优先处理递归的事件，会中断当前事件发送过程，并在新的事件发送过程中，不会再次发送给这个订阅者
)

var (
	// EventRecursionLimit 事件递归次数上限，超过此上限会panic
	EventRecursionLimit = int32(128)
)

// IEvent 事件接口
type IEvent interface {
	ctrl() IEventCtrl
	emit(fun generic.Func1[iface.Cache, bool])
	newHook(subscriberFace iface.FaceAny, priority int32) Hook
	removeSubscriber(subscriber any)
}

// Event 事件
type Event struct {
	subscribers    generic.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	emitted        int32
	emitDepth      int32
	inited         bool
	enabled        bool
}

// Init 初始化事件
func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion) {
	if event.inited {
		exception.Panicf("%w: event is already initialized", ErrEvent)
	}

	event.autoRecover = autoRecover
	event.reportError = reportError
	event.eventRecursion = eventRecursion
	event.inited = true

	event.Enable()
}

// Enable 启用事件
func (event *Event) Enable() {
	if !event.inited {
		exception.Panicf("%w: event not initialized", ErrEvent)
	}
	event.enabled = true
}

// Disable 关闭事件
func (event *Event) Disable() {
	event.UnbindAll()
	event.enabled = false
}

// UnbindAll 解绑定所有订阅者
func (event *Event) UnbindAll() {
	event.subscribers.Traversal(func(node *generic.Node[Hook]) bool {
		node.V.Unbind()
		return true
	})
}

func (event *Event) ctrl() IEventCtrl {
	return event
}

func (event *Event) emit(fun generic.Func1[iface.Cache, bool]) {
	if !event.enabled {
		return
	}

	if event.emitted >= EventRecursionLimit {
		exception.Panicf("%w: recursive event calls(%d) cause stack overflow", ErrEvent, event.emitted)
	}

	switch event.eventRecursion {
	case EventRecursion_Discard:
		if event.emitted > 0 {
			return
		}
	}

	event.emitted++
	defer func() { event.emitted-- }()
	event.emitDepth = event.emitted
	ver := event.subscribers.Version()

	event.subscribers.Traversal(func(node *generic.Node[Hook]) bool {
		if !event.enabled {
			return false
		}

		if node.V.subscriberFace.IsNil() || node.Version() > ver {
			return true
		}

		switch event.eventRecursion {
		case EventRecursion_Allow:
			break
		case EventRecursion_Disallow:
			if node.V.received > 0 {
				exception.Panicf("%w: recursive event disallowed", ErrEvent)
			}
		case EventRecursion_Truncate:
			if node.V.received > 0 {
				return true
			}
		case EventRecursion_Deepest:
			if event.emitDepth != event.emitted {
				return false
			}
			if node.V.received > 0 {
				return true
			}
		}

		node.V.received++
		defer func() { node.V.received-- }()

		ret, panicErr := fun.Call(event.autoRecover, event.reportError, node.V.subscriberFace.Cache)
		if panicErr != nil {
			return true
		}

		return ret
	})
}

func (event *Event) newHook(subscriberFace iface.FaceAny, priority int32) Hook {
	if !event.enabled {
		exception.Panicf("%w: event disabled", ErrEvent)
	}

	if subscriberFace.IsNil() {
		exception.Panicf("%w: %w: subscriberFace is nil", ErrEvent, exception.ErrArgs)
	}

	hook := Hook{
		subscriberFace: subscriberFace,
		priority:       priority,
	}

	var at *generic.Node[Hook]

	event.subscribers.ReversedTraversal(func(other *generic.Node[Hook]) bool {
		if hook.priority >= other.V.priority {
			at = other
			return false
		}
		return true
	})

	if at != nil {
		hook.at = event.subscribers.InsertAfter(Hook{}, at)
	} else {
		hook.at = event.subscribers.PushFront(Hook{})
	}

	hook.at.V = hook

	return hook
}

func (event *Event) removeSubscriber(subscriber any) {
	event.subscribers.ReversedTraversal(func(other *generic.Node[Hook]) bool {
		if other.V.subscriberFace.Iface == subscriber {
			other.Escape()
			return false
		}
		return true
	})
}
