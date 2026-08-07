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
	head, tail, count int
}

// ManagedHandles 按无标签或标签分组保存事件句柄，便于批量解绑。
//
// 零值可用；类型本身不提供并发保护。
type ManagedHandles struct {
	taggedHandleIndex   generic.SliceMap[string, _TaggedHandleSpan]
	untaggedHandleCount int
	handleList          generic.FreeList[Handle]
}

// AddEventHandles 记录当前仍有效的无标签句柄；无效句柄会被忽略。
func (m *ManagedHandles) AddEventHandles(handles ...Handle) {
	for _, handle := range handles {
		if handle.Bound() {
			m.handleList.PushFront(handle)
			m.untaggedHandleCount++
		}
	}
}

// GetEventHandles 返回已记录无标签句柄的快照，其中可能包含后来失效的句柄。
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

// UnbindEventHandles 解绑并移除全部无标签句柄。
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

// AddTaggedEventHandles 将当前仍有效的句柄追加到 tag 分组；无效句柄会被忽略。
func (m *ManagedHandles) AddTaggedEventHandles(tag string, handles ...Handle) {
	spanIdx, ok := m.taggedHandleIndex.Index(tag)
	if ok {
		span := &m.taggedHandleIndex[spanIdx]
		for _, handle := range handles {
			if handle.Bound() {
				slot := m.handleList.InsertAfter(handle, span.V.tail)
				if slot == nil {
					exception.Panicf("%w: tagged event handle span tail not found", ErrEvent)
				}
				span.V.tail = slot.Index()
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
				m.taggedHandleIndex.Add(tag, _TaggedHandleSpan{head: slot.Index(), tail: slot.Index(), count: 1})
				spanIdx, ok := m.taggedHandleIndex.Index(tag)
				if !ok {
					exception.Panicf("%w: tagged event handle span not found", ErrEvent)
				}
				span = &m.taggedHandleIndex[spanIdx]
			} else {
				slot := m.handleList.InsertAfter(handle, span.V.tail)
				if slot == nil {
					exception.Panicf("%w: tagged event handle span tail not found", ErrEvent)
				}
				span.V.tail = slot.Index()
				span.V.count++
			}
		}
	}
}

// GetTaggedEventHandles 返回 tag 分组的句柄快照；标签不存在时返回 nil。
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

// UnbindTaggedEventHandles 解绑并移除 tag 分组中的全部句柄。
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

// UnbindAllEventHandles 解绑并清空全部无标签与带标签句柄。
func (m *ManagedHandles) UnbindAllEventHandles() {
	m.handleList.TraversalEach(func(slot *generic.FreeSlot[Handle]) {
		slot.V.Unbind()
		slot.Free()
	})
	m.taggedHandleIndex = m.taggedHandleIndex[:0]
	m.untaggedHandleCount = 0
}

// ClearAllUnboundEventHandles 移除全部已失效句柄，但不影响仍有效的绑定。
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
		remaining := 0
		head, tail := -1, -1

		m.handleList.TraversalAt(func(slot *generic.FreeSlot[Handle]) bool {
			if !slot.V.Bound() {
				slot.Free()
			} else {
				if head < 0 {
					head = slot.Index()
				}
				tail = slot.Index()
				remaining++
			}
			count--
			return count > 0
		}, span.V.head)

		if remaining <= 0 {
			m.taggedHandleIndex.Delete(span.K)
		} else {
			span.V.head = head
			span.V.tail = tail
			span.V.count = remaining
		}
	}
}
