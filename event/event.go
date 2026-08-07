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

// EventRecursion 定义订阅者在派发过程中再次派发同一事件时的处理策略。
type EventRecursion int8

const (
	EventRecursion_Allow        EventRecursion = iota // EventRecursion_Allow 允许递归派发，仍受递归深度上限约束。
	EventRecursion_Disallow                           // EventRecursion_Disallow 在递归派发时 panic。
	EventRecursion_Discard                            // EventRecursion_Discard 丢弃整次递归派发。
	EventRecursion_SkipReceived                       // EventRecursion_SkipReceived 跳过仍位于当前递归调用栈中的订阅者。
	EventRecursion_ReceiveOnce                        // EventRecursion_ReceiveOnce 使每个订阅者在一次顶层派发中至多接收一次。
)

var (
	// EventRecursionLimit 是同一事件的最大递归深度；达到上限时 panic。
	EventRecursionLimit = 128
)

// IEvent 是生成代码和 Bind 使用的信号式事件接口。
//
// 它不提供并发保护，绑定、解绑与派发必须由调用方串行化。
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

// Event 保存一个进程内同步信号式事件的订阅者列表与派发策略。
//
// Event 的零值可用，默认允许递归且不会恢复订阅者 panic。Event 不支持并发访问。
type Event struct {
	autoRecover bool
	reportError chan error
	recursion   EventRecursion
	disabled    bool
	subscribers generic.FreeList[_Subscriber]
	emitted     int64
}

// PanicHandling 返回订阅者 panic 的恢复与上报设置。
func (event *Event) PanicHandling() (autoRecover bool, reportError chan error) {
	return event.autoRecover, event.reportError
}

// SetPanicHandling 设置订阅者 panic 的处理方式。
//
// autoRecover 为 true 时，panic 会被转换为带堆栈的错误并尝试非阻塞写入 reportError；
// 写入失败或 reportError 为 nil 时错误只作为本次调用结果被内部丢弃，派发继续。
func (event *Event) SetPanicHandling(autoRecover bool, reportError chan error) {
	event.autoRecover = autoRecover
	event.reportError = reportError
}

// Recursion 返回当前递归派发策略。
func (event *Event) Recursion() EventRecursion {
	return event.recursion
}

// SetRecursion 设置递归派发策略。
func (event *Event) SetRecursion(recursion EventRecursion) {
	event.recursion = recursion
}

// Enabled 报告事件是否启用。
func (event *Event) Enabled() bool {
	return !event.disabled
}

// SetEnabled 设置事件是否启用；禁用时会立即解绑全部订阅者。
func (event *Event) SetEnabled(b bool) {
	if !event.disabled == b {
		return
	}

	event.disabled = !b

	if event.disabled {
		event.UnbindAll()
	}
}

// UnbindAll 解绑全部订阅者，并使已有 Handle 失效。
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
