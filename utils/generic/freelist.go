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

package generic

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

// NewFreeList 创建自由链表
func NewFreeList[T any]() *FreeList[T] {
	return &FreeList[T]{}
}

type freeSlotState uint8

const (
	freeSlotState_Freed freeSlotState = iota
	freeSlotState_Active
	freeSlotState_Orphaned
)

// FreeSlot 自由链表槽位
type FreeSlot[T any] struct {
	V               T
	idx, prev, next int
	pendingFreeNext int
	list            *FreeList[T]
	ver             int64
	state           freeSlotState
}

// Version 被占用时的数据版本号
func (s *FreeSlot[T]) Version() int64 {
	return s.ver
}

// Prev 前一个槽位
func (s *FreeSlot[T]) Prev() *FreeSlot[T] {
	if s.list == nil || s.Freed() {
		return nil
	}
	slotPrev := s.list.Get(s.prev)
	if slotPrev == nil || slotPrev.Freed() {
		return nil
	}
	return slotPrev
}

// Next 下一个槽位
func (s *FreeSlot[T]) Next() *FreeSlot[T] {
	if s.list == nil || s.Freed() {
		return nil
	}
	slotNext := s.list.Get(s.next)
	if slotNext == nil || slotNext.Freed() {
		return nil
	}
	return slotNext
}

// Index 槽位索引
func (s *FreeSlot[T]) Index() int {
	return s.idx
}

// Free 释放
func (s *FreeSlot[T]) Free() {
	if s.list == nil {
		return
	}
	s.list.ReleaseIfVersion(s.idx, s.ver)
}

// Orphaned 是否悬空准备释放
func (s *FreeSlot[T]) Orphaned() bool {
	return s.state == freeSlotState_Orphaned
}

// Freed 是否已被释放
func (s *FreeSlot[T]) Freed() bool {
	return s.state == freeSlotState_Freed
}

// FreeList 自由链表
type FreeList[T any] struct {
	_               noCopy
	slots           []FreeSlot[T]
	head            int
	tail            int
	unused          int
	pendingFreeHead int
	freeHead        int
	len             int
	ver             int64
	orphanCount     int
	depth           int
}

// Cap 总容量
func (l *FreeList[T]) Cap() int {
	return len(l.slots)
}

// Len 链表长度
func (l *FreeList[T]) Len() int {
	return l.len
}

// Version 数据版本号
func (l *FreeList[T]) Version() int64 {
	return l.ver
}

// OrphanCount 悬空槽位数量
func (l *FreeList[T]) OrphanCount() int {
	return l.orphanCount
}

// Depth 遍历递归深度
func (l *FreeList[T]) Depth() int {
	return l.depth
}

// Front 链表头部
func (l *FreeList[T]) Front() *FreeSlot[T] {
	if l.ver <= 0 || l.head < 0 {
		return nil
	}
	return &l.slots[l.head]
}

// Back 链表尾部
func (l *FreeList[T]) Back() *FreeSlot[T] {
	if l.ver <= 0 || l.tail < 0 {
		return nil
	}
	return &l.slots[l.tail]
}

// Get 获取槽位
func (l *FreeList[T]) Get(idx int) *FreeSlot[T] {
	if l.ver <= 0 || idx < 0 || idx >= len(l.slots) {
		return nil
	}
	return &l.slots[idx]
}

// Release 释放槽位
func (l *FreeList[T]) Release(idx int) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	if l.depth > 0 {
		l.orphan(slot)
		return
	}
	l.release(slot)
}

// ReleaseIfVersion 数据版本号匹配时释放槽位
func (l *FreeList[T]) ReleaseIfVersion(idx int, ver int64) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() || slot.Version() != ver {
		return
	}
	if l.depth > 0 {
		l.orphan(slot)
		return
	}
	l.release(slot)
}

// ReleaseOrphans 释放悬空槽位
func (l *FreeList[T]) ReleaseOrphans() {
	if l.ver <= 0 {
		return
	}
	l.releaseOrphans()
}

// PushFront 在链表头部插入数据
func (l *FreeList[T]) PushFront(value T) *FreeSlot[T] {
	l.lazyInit()
	return l.appendValue(value, -1)
}

// PushBack 在链表尾部插入数据
func (l *FreeList[T]) PushBack(value T) *FreeSlot[T] {
	l.lazyInit()
	return l.appendValue(value, l.tail)
}

// PopFront 弹出链表头部数据
func (l *FreeList[T]) PopFront() (T, bool) {
	slot := l.Front()
	if slot == nil {
		return types.Zero[T](), false
	}
	v := slot.V
	slot.Free()
	return v, true
}

// PopBack 弹出链表尾部数据
func (l *FreeList[T]) PopBack() (T, bool) {
	slot := l.Back()
	if slot == nil {
		return types.Zero[T](), false
	}
	v := slot.V
	slot.Free()
	return v, true
}

// InsertBefore 在链表指定位置前插入数据
func (l *FreeList[T]) InsertBefore(value T, at int) *FreeSlot[T] {
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return nil
	}
	return l.appendValue(value, slotAt.prev)
}

// InsertAfter 在链表指定位置后插入数据
func (l *FreeList[T]) InsertAfter(value T, at int) *FreeSlot[T] {
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return nil
	}
	return l.appendValue(value, at)
}

// MoveToFront 移动槽位至链表头部
func (l *FreeList[T]) MoveToFront(idx int) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	l.moveAfter(slot, -1)
}

// MoveToBack 移动槽位至链表尾部
func (l *FreeList[T]) MoveToBack(idx int) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	l.moveAfter(slot, l.tail)
}

// MoveBefore 移动槽位至链表指定位置前
func (l *FreeList[T]) MoveBefore(idx, at int) {
	if idx == at {
		return
	}
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.moveAfter(slot, slotAt.prev)
}

// MoveAfter 移动槽位至链表指定位置后
func (l *FreeList[T]) MoveAfter(idx, at int) {
	if idx == at {
		return
	}
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.moveAfter(slot, at)
}

// PushFrontList 在链表头部插入其他链表，跳过悬空节点，可以传入自身
func (l *FreeList[T]) PushFrontList(other *FreeList[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, n := other.Len(), other.Back(); i > 0; i, n = i-1, n.Prev() {
		if !n.Orphaned() {
			l.appendValue(n.V, -1)
		}
	}
}

// PushBackList 在链表尾部插入其他链表，跳过悬空节点，可以传入自身
func (l *FreeList[T]) PushBackList(other *FreeList[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, n := other.Len(), other.Front(); i > 0; i, n = i-1, n.Next() {
		if !n.Orphaned() {
			l.appendValue(n.V, l.tail)
		}
	}
}

// Traversal 遍历槽位，跳过悬空节点
func (l *FreeList[T]) Traversal(visitor func(slot *FreeSlot[T]) bool) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := l.Front(); s != nil; s = s.Next() {
		if s.Orphaned() {
			continue
		}
		if !visitor(s) {
			break
		}
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// TraversalEach 遍历每个槽位，跳过悬空节点
func (l *FreeList[T]) TraversalEach(visitor func(slot *FreeSlot[T])) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := l.Front(); s != nil; s = s.Next() {
		if s.Orphaned() {
			continue
		}
		visitor(s)
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// TraversalAt 从指定位置开始遍历槽位，跳过悬空节点
func (l *FreeList[T]) TraversalAt(visitor func(slot *FreeSlot[T]) bool, at int) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := slotAt; s != nil; s = s.Next() {
		if s.Orphaned() {
			continue
		}
		if !visitor(s) {
			break
		}
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// TraversalEachAt 从指定位置开始遍历每个槽位，跳过悬空节点
func (l *FreeList[T]) TraversalEachAt(visitor func(slot *FreeSlot[T]), at int) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := slotAt; s != nil; s = s.Next() {
		if s.Orphaned() {
			continue
		}
		visitor(s)
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// ReversedTraversal 反向遍历槽位，跳过悬空节点
func (l *FreeList[T]) ReversedTraversal(visitor func(slot *FreeSlot[T]) bool) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := l.Back(); s != nil; s = s.Prev() {
		if s.Orphaned() {
			continue
		}
		if !visitor(s) {
			break
		}
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// ReversedTraversalAt 从指定位置开始反向遍历槽位，跳过悬空节点
func (l *FreeList[T]) ReversedTraversalAt(visitor func(slot *FreeSlot[T]) bool, at int) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := slotAt; s != nil; s = s.Prev() {
		if s.Orphaned() {
			continue
		}
		if !visitor(s) {
			break
		}
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// ReversedTraversalEach 反向遍历槽位，跳过悬空节点
func (l *FreeList[T]) ReversedTraversalEach(visitor func(slot *FreeSlot[T])) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := l.Back(); s != nil; s = s.Prev() {
		if s.Orphaned() {
			continue
		}
		visitor(s)
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// ReversedTraversalEachAt 从指定位置开始反向遍历每个槽位，跳过悬空节点
func (l *FreeList[T]) ReversedTraversalEachAt(visitor func(slot *FreeSlot[T]), at int) {
	if l.ver <= 0 || visitor == nil {
		return
	}
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return
	}
	l.depth++
	defer l.traversalReleaseOrphans()
	ver := l.ver
	for s := slotAt; s != nil; s = s.Prev() {
		if s.Orphaned() {
			continue
		}
		visitor(s)
		if ver != l.ver {
			s = l.Get(s.idx)
		}
	}
}

// Clone 拷贝链表，跳过悬空节点
func (l *FreeList[T]) Clone() *FreeList[T] {
	if l == nil {
		return nil
	}

	copied := NewFreeList[T]()
	if l.ver <= 0 {
		return copied
	}

	copied.lazyInit()
	l.TraversalEach(func(slot *FreeSlot[T]) {
		copied.appendValue(slot.V, copied.tail)
	})

	return copied
}

// ToSlice 链表所有数据转换为切片，跳过悬空节点
func (l *FreeList[T]) ToSlice() []T {
	slice := make([]T, 0, l.Len()-l.OrphanCount())
	l.TraversalEach(func(slot *FreeSlot[T]) {
		slice = append(slice, slot.V)
	})
	return slice
}

func (l *FreeList[T]) lazyInit() {
	if l.ver != 0 {
		return
	}
	l.slots = make([]FreeSlot[T], 8)
	l.head = -1
	l.tail = -1
	l.unused = 0
	l.pendingFreeHead = -1
	l.freeHead = -1
	l.len = 0
	l.ver++
	l.orphanCount = 0
	l.depth = 0
}

func (l *FreeList[T]) appendValue(value T, at int) *FreeSlot[T] {
	slotsCap := len(l.slots)
	if l.freeHead < 0 && l.unused >= slotsCap {
		var slots []FreeSlot[T]
		if slotsCap < 1024 {
			slots = make([]FreeSlot[T], slotsCap*2)
		} else {
			slots = make([]FreeSlot[T], slotsCap+slotsCap/4)
		}
		copy(slots, l.slots)
		l.slots = slots
		slotsCap = len(slots)
	}

	var slot *FreeSlot[T]
	if l.freeHead >= 0 {
		slot = &l.slots[l.freeHead]
		l.freeHead = slot.next
	} else {
		if l.unused >= slotsCap {
			exception.Panic("FreeList: no free slot")
		}
		slot = &l.slots[l.unused]
		slot.list = l
		slot.idx = l.unused
		l.unused++
	}

	slot.V = value
	slot.pendingFreeNext = -1
	slot.state = freeSlotState_Active

	if at < 0 {
		if l.head < 0 {
			slot.prev = -1
			slot.next = -1
			l.head = slot.idx
			l.tail = slot.idx
		} else {
			slot.prev = -1
			slot.next = l.head
			l.slots[l.head].prev = slot.idx
			l.head = slot.idx
		}
	} else {
		slotAt := &l.slots[at]

		slot.prev = at
		slot.next = slotAt.next
		slotAt.next = slot.idx
		if slot.next >= 0 {
			l.slots[slot.next].prev = slot.idx
		} else {
			l.tail = slot.idx
		}
	}

	l.len++
	l.ver++
	slot.ver = l.ver

	return slot
}

func (l *FreeList[T]) moveAfter(slot *FreeSlot[T], at int) {
	if slot.idx == at || l.len < 2 {
		return
	}

	if slot.prev < 0 {
		l.head = slot.next
		if l.head >= 0 {
			l.slots[l.head].prev = -1
		}
	} else {
		l.slots[slot.prev].next = slot.next
	}
	if slot.next < 0 {
		l.tail = slot.prev
		if l.tail >= 0 {
			l.slots[l.tail].next = -1
		}
	} else {
		l.slots[slot.next].prev = slot.prev
	}

	if at < 0 {
		slot.prev = -1
		slot.next = l.head
		l.slots[l.head].prev = slot.idx
		l.head = slot.idx
	} else {
		slotAt := &l.slots[at]

		slot.prev = at
		slot.next = slotAt.next
		slotAt.next = slot.idx
		if slot.next >= 0 {
			l.slots[slot.next].prev = slot.idx
		} else {
			l.tail = slot.idx
		}
	}

	l.ver++
	slot.ver = l.ver
}

func (l *FreeList[T]) release(slot *FreeSlot[T]) {
	if slot.prev < 0 {
		l.head = slot.next
		if l.head >= 0 {
			l.slots[l.head].prev = -1
		}
	} else {
		l.slots[slot.prev].next = slot.next
	}
	if slot.next < 0 {
		l.tail = slot.prev
		if l.tail >= 0 {
			l.slots[l.tail].next = -1
		}
	} else {
		l.slots[slot.next].prev = slot.prev
	}
	slot.V = types.Zero[T]()
	slot.next = l.freeHead
	slot.state = freeSlotState_Freed
	l.freeHead = slot.idx
	l.ver++
	l.len--
}

func (l *FreeList[T]) orphan(slot *FreeSlot[T]) {
	if slot.Orphaned() {
		return
	}
	slot.V = types.Zero[T]()
	slot.state = freeSlotState_Orphaned
	slot.pendingFreeNext = l.pendingFreeHead
	l.pendingFreeHead = slot.idx
	l.orphanCount++
}

func (l *FreeList[T]) traversalReleaseOrphans() {
	l.depth--
	l.releaseOrphans()
}

func (l *FreeList[T]) releaseOrphans() {
	if l.depth > 0 || l.pendingFreeHead < 0 {
		return
	}

	for idx := l.pendingFreeHead; idx >= 0; {
		slot := &l.slots[idx]
		next := slot.pendingFreeNext
		l.release(slot)
		l.orphanCount--
		idx = next
	}
	l.pendingFreeHead = -1
}
