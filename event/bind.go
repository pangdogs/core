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
	"git.golaxy.org/core/utils/iface"
	"github.com/elliotchance/pie/v2"
)

// Bind 将 subscriber 绑定到 event，并返回可独立解绑的句柄。
//
// 订阅者按 priority 升序调用；同优先级保持绑定顺序。未指定 priority 时使用 0。
// event 或 subscriber 无效以及事件已禁用时 panic。
func Bind[T any](event IEvent, subscriber T, priority ...int32) Handle {
	if event == nil {
		exception.Panicf("%w: %w: event is nil", ErrEvent, exception.ErrArgs)
	}
	return event.addSubscriber(iface.NewFaceAny(subscriber), pie.First(priority))
}

// Unbind 解除 subscriber 最近一次对 event 的绑定。
//
// 同一订阅者重复绑定时，重复调用会按绑定顺序的逆序逐个解除。通常应优先保存并使用
// Handle，以避免误解绑；event 为 nil 或 subscriber 的动态类型不可比较时 panic。
func Unbind(event IEvent, subscriber any) {
	if event == nil {
		exception.Panicf("%w: %w: event is nil", ErrEvent, exception.ErrArgs)
	}
	event.removeSubscriber(subscriber)
}
