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
	"git.golaxy.org/core/utils/iface"
)

// Deprecated: UnsafeEvent 访问事件内部方法
func UnsafeEvent(event IEvent) _UnsafeEvent {
	return _UnsafeEvent{
		IEvent: event,
	}
}

type _UnsafeEvent struct {
	IEvent
}

// Emit 发送事件
func (ue _UnsafeEvent) Emit(fun func(subscriber iface.Cache) bool) {
	ue.emit(fun)
}

// NewHook 创建事件钩子
func (ue _UnsafeEvent) NewHook(subscriberFace iface.FaceAny, priority int32) Hook {
	return ue.newHook(subscriberFace, priority)
}

// RemoveSubscriber 删除订阅者
func (ue _UnsafeEvent) RemoveSubscriber(subscriber any) {
	ue.removeSubscriber(subscriber)
}
