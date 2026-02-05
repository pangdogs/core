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

// Handle 事件句柄，由BindEvent()创建并返回的绑定句柄，请勿自己创建
type Handle struct {
	event *Event
	idx   int
	ver   int64
}

// Unbind 解绑定事件与订阅者
func (handle Handle) Unbind() {
	if handle.event != nil {
		handle.event.subscribers.ReleaseIfVersion(handle.idx, handle.ver)
	}
}

// Bound 是否已绑定事件
func (handle Handle) Bound() bool {
	if handle.event == nil {
		return false
	}
	slot := handle.event.subscribers.Get(handle.idx)
	return slot != nil && !slot.Orphaned() && !slot.Freed() && slot.Version() == handle.ver
}

// UnbindHandles 解绑定事件句柄（Handle）
func UnbindHandles(handles []Handle) {
	for i := range handles {
		handles[i].Unbind()
	}
}
