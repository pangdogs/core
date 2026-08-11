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
	"fmt"

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

// ErrFreeList 是 FreeList 错误的共同根错误。
var ErrFreeList = fmt.Errorf("%w: free-list", exception.ErrCore)

// NewFreeList 创建空槽链表；FreeList 的零值同样可用。
func NewFreeList[T any]() *FreeList[T] {
	return &FreeList[T]{}
}

type freeSlotState uint8

const (
	freeSlotState_Freed freeSlotState = iota
	freeSlotState_Active
	freeSlotState_Orphaned
)

// FreeSlot 是 FreeList 中可复用的槽位。
//
// Index 会被后续元素复用，应与 Version 一起作为句柄校验。扩容可能使已保存的槽位指针
// 指向旧切片，因此槽位指针只适合短期使用。
type FreeSlot[T any] struct {
	V               T
	idx, prev, next int
	pendingFreeNext int
	list            *FreeList[T]
	ver             int64
	state           freeSlotState
}

// Version 返回该槽位当前占用或移动版本。
func (s *FreeSlot[T]) Version() int64 {
	return s.ver
}

// Prev 返回链表顺序中的前一槽位；不存在或当前槽位已释放时返回 nil。
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

// Next 返回链表顺序中的后一槽位；不存在或当前槽位已释放时返回 nil。
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

// Index 返回可复用的槽位索引。
func (s *FreeSlot[T]) Index() int {
	return s.idx
}

// Free 释放当前版本的槽位；重复调用无效。
// 遍历期间释放会先将槽位置为悬空，待最外层遍历结束后回收。
func (s *FreeSlot[T]) Free() {
	if s.list == nil {
		return
	}
	s.list.ReleaseIfVersion(s.idx, s.ver)
}

// Orphaned 报告槽位是否已标记释放、正等待遍历结束后回收。
func (s *FreeSlot[T]) Orphaned() bool {
	return s.state == freeSlotState_Orphaned
}

// Freed 报告槽位是否已回收到空闲链。
func (s *FreeSlot[T]) Freed() bool {
	return s.state == freeSlotState_Freed
}

// FreeList 是以可复用切片槽位实现的有序链表。
//
// 它提供稳定索引与版本组合句柄，并允许遍历回调释放元素。零值可用、首次使用后不可
// 复制，且不提供并发保护。
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

// Cap 返回当前已分配的槽位数量。
func (l *FreeList[T]) Cap() int {
	return len(l.slots)
}

// Reserve 确保底层至少能容纳 capacity 个槽位；不会改变现有元素的链表顺序。
// capacity 为负数时 panic。
func (l *FreeList[T]) Reserve(capacity int) {
	if capacity < 0 {
		exception.Panicf("%w: %w: capacity %d is negative", ErrFreeList, exception.ErrArgs, capacity)
	}
	if capacity <= len(l.slots) {
		return
	}

	if l.ver == 0 {
		l.init(capacity)
		return
	}

	slots := make([]FreeSlot[T], capacity)
	copy(slots, l.slots)
	l.slots = slots
}

// Len 返回链表中的槽位数；遍历期间包含尚未回收的悬空槽位。
func (l *FreeList[T]) Len() int {
	return l.len
}

// Version 返回链表最近一次结构变更的版本。
func (l *FreeList[T]) Version() int64 {
	return l.ver
}

// OrphanCount 返回等待遍历结束后回收的槽位数。
func (l *FreeList[T]) OrphanCount() int {
	return l.orphanCount
}

// Depth 返回当前嵌套遍历深度。
func (l *FreeList[T]) Depth() int {
	return l.depth
}

// Front 返回链表头槽位；空表返回 nil。遍历期间返回值可能已悬空。
func (l *FreeList[T]) Front() *FreeSlot[T] {
	if l.ver <= 0 || l.head < 0 {
		return nil
	}
	return &l.slots[l.head]
}

// Back 返回链表尾槽位；空表返回 nil。遍历期间返回值可能已悬空。
func (l *FreeList[T]) Back() *FreeSlot[T] {
	if l.ver <= 0 || l.tail < 0 {
		return nil
	}
	return &l.slots[l.tail]
}

// Get 返回底层指定索引的槽位；索引无效时返回 nil。
// 返回的槽位可能尚未使用、已悬空或已释放，调用方必须检查其状态和版本。
func (l *FreeList[T]) Get(idx int) *FreeSlot[T] {
	if l.ver <= 0 || idx < 0 || idx >= len(l.slots) {
		return nil
	}
	return &l.slots[idx]
}

// Release 释放指定索引当前对应的槽位；索引无效或已释放时无效。
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

// ReleaseIfVersion 仅在索引与版本同时匹配时释放槽位。
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

// ReleaseOrphans 回收全部悬空槽位；仍处于遍历中时无效。
func (l *FreeList[T]) ReleaseOrphans() {
	if l.ver <= 0 {
		return
	}
	l.releaseOrphans()
}

// PushFront 在链表头部插入值并返回新槽位。
func (l *FreeList[T]) PushFront(value T) *FreeSlot[T] {
	l.lazyInit()
	return l.appendValue(value, -1)
}

// PushBack 在链表尾部插入值并返回新槽位。
func (l *FreeList[T]) PushBack(value T) *FreeSlot[T] {
	l.lazyInit()
	return l.appendValue(value, l.tail)
}

// PopFront 移除并返回头部值；空表返回 T 的零值与 false。
func (l *FreeList[T]) PopFront() (T, bool) {
	slot := l.Front()
	if slot == nil {
		return types.Zero[T](), false
	}
	v := slot.V
	slot.Free()
	return v, true
}

// PopBack 移除并返回尾部值；空表返回 T 的零值与 false。
func (l *FreeList[T]) PopBack() (T, bool) {
	slot := l.Back()
	if slot == nil {
		return types.Zero[T](), false
	}
	v := slot.V
	slot.Free()
	return v, true
}

// InsertBefore 在索引 at 对应槽位前插入值；位置无效时返回 nil。
func (l *FreeList[T]) InsertBefore(value T, at int) *FreeSlot[T] {
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return nil
	}
	return l.appendValue(value, slotAt.prev)
}

// InsertAfter 在索引 at 对应槽位后插入值；位置无效时返回 nil。
func (l *FreeList[T]) InsertAfter(value T, at int) *FreeSlot[T] {
	slotAt := l.Get(at)
	if slotAt == nil || slotAt.Freed() {
		return nil
	}
	return l.appendValue(value, at)
}

// MoveToFront 将指定槽位移动到链表头；位置无效时无效。
func (l *FreeList[T]) MoveToFront(idx int) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	l.moveAfter(slot, -1)
}

// MoveToBack 将指定槽位移动到链表尾；位置无效时无效。
func (l *FreeList[T]) MoveToBack(idx int) {
	slot := l.Get(idx)
	if slot == nil || slot.Freed() {
		return
	}
	l.moveAfter(slot, l.tail)
}

// MoveBefore 将 idx 对应槽位移动到 at 对应槽位之前；任一位置无效时无效。
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

// MoveAfter 将 idx 对应槽位移动到 at 对应槽位之后；任一位置无效时无效。
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

// PushFrontList 将 other 的活动值副本插入头部并保持原顺序；other 可以是自身。
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

// PushBackList 将 other 的活动值副本追加到尾部并保持原顺序；other 可以是自身。
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

// Traversal 从头到尾遍历活动槽位；visitor 返回 false 时停止。
// visitor 为 nil 时无效，回调中释放的槽位会延迟到最外层遍历结束后回收。
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

// TraversalEach 从头到尾遍历全部活动槽位。
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

// TraversalAt 从索引 at 对应槽位开始向后遍历；visitor 返回 false 时停止。
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

// TraversalEachAt 从索引 at 对应槽位开始向后遍历全部活动槽位。
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

// ReversedTraversal 从尾到头遍历活动槽位；visitor 返回 false 时停止。
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

// ReversedTraversalAt 从索引 at 对应槽位开始向前遍历；visitor 返回 false 时停止。
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

// ReversedTraversalEach 从尾到头遍历全部活动槽位。
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

// ReversedTraversalEachAt 从索引 at 对应槽位开始向前遍历全部活动槽位。
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

// Clone 返回仅包含活动值的浅拷贝；nil 接收者返回 nil。
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

// ToSlice 按链表顺序返回全部活动值的切片副本。
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
	l.init(8)
}

func (l *FreeList[T]) init(capacity int) {
	l.slots = make([]FreeSlot[T], capacity)
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
			exception.Panicf("%w: no free slot", ErrFreeList)
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
