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
type EventRecursion int8

const (
	EventRecursion_Allow        EventRecursion = iota // 允许事件递归，可能会无限递归
	EventRecursion_Disallow                           // 不允许事件递归，递归时会panic
	EventRecursion_Discard                            // 丢弃递归的事件，不会再发送给任何订阅者
	EventRecursion_SkipReceived                       // 发送递归事件时跳过已接收事件的订阅者
	EventRecursion_ReceiveOnce                        // 订阅者在整个事件发送过程中只接收一次
)

var (
	// EventRecursionLimit 事件递归次数上限，超过此上限会panic
	EventRecursionLimit = 128
)

// IEvent 事件接口
type IEvent interface {
	ctrl() IEventCtrl
	emit(fun generic.Func1[iface.Cache, bool])
	addSubscriber(subscriberFace iface.FaceAny, priority int32) Handle
	removeSubscriber(subscriber any)
}

type _Subscriber struct {
	face            iface.FaceAny
	priority        int32
	receivedDepth   int32
	receivedEmitted int64
}

// Event 事件
type Event struct {
	autoRecover bool
	reportError chan error
	recursion   EventRecursion
	disabled    bool
	subscribers generic.FreeList[_Subscriber]
	emitted     int64
}

// GetPanicHandling 获取panic时的处理方式
func (event *Event) GetPanicHandling() (autoRecover bool, reportError chan error) {
	return event.autoRecover, event.reportError
}

// SetPanicHandling 设置panic时的处理方式
func (event *Event) SetPanicHandling(autoRecover bool, reportError chan error) {
	event.autoRecover = autoRecover
	event.reportError = reportError
}

// GetRecursion 获取发生事件递归时的处理方式
func (event *Event) GetRecursion() EventRecursion {
	return event.recursion
}

// SetRecursion 设置发生事件递归时的处理方式
func (event *Event) SetRecursion(recursion EventRecursion) {
	event.recursion = recursion
}

// GetEnable 获取事件是否启用
func (event *Event) GetEnable() bool {
	return !event.disabled
}

// SetEnable 设置事件是否启用
func (event *Event) SetEnable(b bool) {
	if !event.disabled == b {
		return
	}

	event.disabled = !b

	if event.disabled {
		event.UnbindAll()
	}
}

// UnbindAll 解绑定所有订阅者
func (event *Event) UnbindAll() {
	event.subscribers.TraversalEach(func(slot *generic.FreeSlot[_Subscriber]) { slot.Free() })
}

func (event *Event) ctrl() IEventCtrl {
	return event
}

func (event *Event) emit(fun generic.Func1[iface.Cache, bool]) {
	if event.disabled {
		return
	}

	emitDepth := event.subscribers.Depth()
	if emitDepth > 0 {
		if emitDepth >= EventRecursionLimit {
			exception.Panicf("%w: recursive event calls(%d) cause stack overflow", ErrEvent, event.subscribers.Depth())
		}
	} else {
		event.emitted++
	}

	switch event.recursion {
	case EventRecursion_Discard:
		if emitDepth > 0 {
			return
		}
	case EventRecursion_Disallow:
		if emitDepth > 0 {
			exception.Panicf("%w: recursive event disallowed", ErrEvent)
		}
	}

	ver := event.subscribers.Version()

	event.subscribers.Traversal(func(slot *generic.FreeSlot[_Subscriber]) bool {
		if event.disabled {
			return false
		}

		if slot.V.face.IsNil() || slot.Version() > ver {
			return true
		}

		switch event.recursion {
		case EventRecursion_SkipReceived:
			if slot.V.receivedDepth > 0 {
				return true
			}
		case EventRecursion_ReceiveOnce:
			if slot.V.receivedEmitted >= event.emitted {
				return true
			}
		}

		slot.V.receivedDepth++
		defer func() { slot.V.receivedDepth-- }()

		slot.V.receivedEmitted = event.emitted

		ret, panicErr := fun.Call(event.autoRecover, event.reportError, slot.V.face.Cache)
		if panicErr != nil {
			return true
		}

		return ret
	})
}

func (event *Event) addSubscriber(subscriberFace iface.FaceAny, priority int32) Handle {
	if event.disabled {
		exception.Panicf("%w: event disabled", ErrEvent)
	}

	if subscriberFace.IsNil() {
		exception.Panicf("%w: %w: subscriberFace is nil", ErrEvent, exception.ErrArgs)
	}

	var at *generic.FreeSlot[_Subscriber]
	event.subscribers.ReversedTraversal(func(slot *generic.FreeSlot[_Subscriber]) bool {
		if priority >= slot.V.priority {
			at = slot
			return false
		}
		return true
	})

	var slot *generic.FreeSlot[_Subscriber]
	if at != nil {
		slot = event.subscribers.InsertAfter(_Subscriber{face: subscriberFace, priority: priority}, at.Index())
	} else {
		slot = event.subscribers.PushFront(_Subscriber{face: subscriberFace, priority: priority})
	}

	return Handle{
		event: event,
		idx:   slot.Index(),
		ver:   slot.Version(),
	}
}

func (event *Event) removeSubscriber(subscriber any) {
	event.subscribers.ReversedTraversal(func(slot *generic.FreeSlot[_Subscriber]) bool {
		if slot.V.face.Iface == subscriber {
			slot.Free()
			return false
		}
		return true
	})
}
