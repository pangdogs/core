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
)

type _TaggedHandleSpan struct {
	head, count int
}

// ManagedHandles 托管事件句柄
type ManagedHandles struct {
	taggedHandleIndex   generic.SliceMap[string, _TaggedHandleSpan]
	untaggedHandleCount int
	handleList          generic.FreeList[Handle]
}

// AddEventHandles 托管事件句柄
func (m *ManagedHandles) AddEventHandles(handles ...Handle) {
	for _, handle := range handles {
		if handle.Bound() {
			m.handleList.PushFront(handle)
			m.untaggedHandleCount++
		}
	}
}

// GetEventHandles 获取已托管的事件句柄
func (m *ManagedHandles) GetEventHandles() []Handle {
	if m.untaggedHandleCount <= 0 {
		return nil
	}

	handles := make([]Handle, 0, m.untaggedHandleCount)
	count := m.untaggedHandleCount

	m.handleList.Traversal(func(slot *generic.FreeSlot[Handle]) bool {
		handles = append(handles, slot.V)
		count--
		return count > 0
	})

	return handles
}

// UnbindEventHandles 解绑定已托管的事件句柄
func (m *ManagedHandles) UnbindEventHandles() {
	if m.untaggedHandleCount <= 0 {
		return
	}

	m.handleList.Traversal(func(slot *generic.FreeSlot[Handle]) bool {
		slot.V.Unbind()
		slot.Free()
		m.untaggedHandleCount--
		return m.untaggedHandleCount > 0
	})
}

// AddTaggedEventHandles 使用标签托管事件句柄
func (m *ManagedHandles) AddTaggedEventHandles(tag string, handles ...Handle) {
	spanIdx, ok := m.taggedHandleIndex.Index(tag)
	if ok {
		span := &m.taggedHandleIndex[spanIdx]
		for _, handle := range handles {
			if handle.Bound() {
				m.handleList.InsertAfter(handle, span.V.head+span.V.count-1)
				span.V.count++
			}
		}
		return
	}

	var span *generic.KV[string, _TaggedHandleSpan]
	for _, handle := range handles {
		if handle.Bound() {
			if span == nil {
				slot := m.handleList.PushBack(handle)
				m.taggedHandleIndex.Add(tag, _TaggedHandleSpan{head: slot.Index(), count: 1})
				spanIdx, ok := m.taggedHandleIndex.Index(tag)
				if !ok {
					exception.Panicf("%w: tagged event handle span not found", ErrEvent)
				}
				span = &m.taggedHandleIndex[spanIdx]
			} else {
				m.handleList.InsertAfter(handle, span.V.head+span.V.count-1)
				span.V.count++
			}
		}
	}
}

// GetTaggedEventHandles 使用标签获取已托管的事件句柄
func (m *ManagedHandles) GetTaggedEventHandles(tag string) []Handle {
	span, ok := m.taggedHandleIndex.Get(tag)
	if !ok {
		return nil
	}

	handles := make([]Handle, 0, span.count)
	count := span.count

	m.handleList.TraversalAt(func(slot *generic.FreeSlot[Handle]) bool {
		handles = append(handles, slot.V)
		count--
		return count > 0
	}, span.head)

	return handles
}

// UnbindTaggedEventHandles 使用标签解绑定已托管的事件句柄
func (m *ManagedHandles) UnbindTaggedEventHandles(tag string) {
	span, ok := m.taggedHandleIndex.Get(tag)
	if !ok {
		return
	}

	count := span.count

	m.handleList.TraversalAt(func(slot *generic.FreeSlot[Handle]) bool {
		slot.V.Unbind()
		slot.Free()
		count--
		return count > 0
	}, span.head)

	m.taggedHandleIndex.Delete(tag)
}

// UnbindAllEventHandles 解绑定所有已托管的事件句柄
func (m *ManagedHandles) UnbindAllEventHandles() {
	m.handleList.TraversalEach(func(slot *generic.FreeSlot[Handle]) {
		slot.V.Unbind()
		slot.Free()
	})
	m.taggedHandleIndex = m.taggedHandleIndex[:0]
	m.untaggedHandleCount = 0
}

// ClearAllUnboundEventHandles 清除所有已失效的事件句柄
func (m *ManagedHandles) ClearAllUnboundEventHandles() {
	if count := m.untaggedHandleCount; count > 0 {
		m.handleList.Traversal(func(slot *generic.FreeSlot[Handle]) bool {
			if !slot.V.Bound() {
				slot.Free()
				m.untaggedHandleCount--
			}
			count--
			return count > 0
		})
	}

	for i := len(m.taggedHandleIndex) - 1; i >= 0; i-- {
		span := &m.taggedHandleIndex[i]
		count := span.V.count

		m.handleList.TraversalAt(func(slot *generic.FreeSlot[Handle]) bool {
			if !slot.V.Bound() {
				slot.Free()
				span.V.count--

				if slot.Index() == span.V.head && span.V.count > 0 {
					headSlot := slot.Next()
					if headSlot == nil {
						exception.Panicf("%w: tagged event handle span head not found", ErrEvent)
					}
					span.V.head = headSlot.Index()
				}
			}
			count--
			return count > 0
		}, span.V.head)

		if span.V.count <= 0 {
			m.taggedHandleIndex.Delete(span.K)
		}
	}
}
