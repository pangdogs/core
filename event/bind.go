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

// Bind 绑定事件与订阅者，可以设置优先级调整回调先后顺序（升序）
func Bind[T any](event IEvent, subscriber T, priority ...int32) Handle {
	if event == nil {
		exception.Panicf("%w: %w: event is nil", ErrEvent, exception.ErrArgs)
	}
	return event.addSubscriber(iface.MakeFaceAny(subscriber), pie.First(priority))
}

// Unbind 解绑定事件与订阅者，在同个订阅者多次绑定事件的情况下，会以逆序依次解除，正常情况下应该使用事件句柄（Handle）解绑定，不应该使用该函数
func Unbind(event IEvent, subscriber any) {
	if event == nil {
		exception.Panicf("%w: %w: event is nil", ErrEvent, exception.ErrArgs)
	}
	event.removeSubscriber(subscriber)
}
