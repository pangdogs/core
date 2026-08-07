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

// UnsafeEvent 暴露事件控制与派发能力，供生成代码和框架内部使用。
//
// Deprecated: 业务代码应使用生成的派发函数或事件表接口。
func UnsafeEvent(event IEvent) _UnsafeEvent {
	return _UnsafeEvent{
		IEvent: event,
	}
}

type _UnsafeEvent struct {
	IEvent
}

// Ctrl 返回事件控制器。
func (u _UnsafeEvent) Ctrl() IEventCtrl {
	return u.ctrl()
}

// Emit 同步遍历订阅者；fun 返回 false 时停止本次派发。
func (u _UnsafeEvent) Emit(fun func(subscriber iface.Cache) bool) {
	u.emit(fun)
}
