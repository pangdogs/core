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

// NewList 创建链表
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// List 链表
type List[T any] struct {
	root Node[T]
	len  int
	ver  int64
}

// Len 链表长度
func (l *List[T]) Len() int {
	return l.len
}

// Version 数据版本
func (l *List[T]) Version() int64 {
	return l.ver
}

// Front 链表头部
func (l *List[T]) Front() *Node[T] {
	if l.len <= 0 {
		return nil
	}
	return l.root.next
}

// Back 链表尾部
func (l *List[T]) Back() *Node[T] {
	if l.len <= 0 {
		return nil
	}
	return l.root.prev
}

// Remove 删除节点
func (l *List[T]) Remove(n *Node[T]) {
	if n != nil {
		n.Escape()
	}
}

// PushFront 在链表头部插入数据
func (l *List[T]) PushFront(value T) *Node[T] {
	l.lazyInit()
	return l.insertValue(value, &l.root)
}

// PushBack 在链表尾部插入数据
func (l *List[T]) PushBack(value T) *Node[T] {
	l.lazyInit()
	return l.insertValue(value, l.root.prev)
}

// InsertBefore 在链表指定位置前插入数据
func (l *List[T]) InsertBefore(value T, at *Node[T]) *Node[T] {
	if !l.check(at) {
		return nil
	}
	return l.insertValue(value, at.prev)
}

// InsertAfter 在链表指定位置后插入数据
func (l *List[T]) InsertAfter(value T, at *Node[T]) *Node[T] {
	if !l.check(at) {
		return nil
	}
	return l.insertValue(value, at)
}

// MoveToFront 移动节点至链表头部
func (l *List[T]) MoveToFront(n *Node[T]) {
	if !l.check(n) || l.root.next == n {
		return
	}
	l.move(n, &l.root)
}

// MoveToBack 移动节点至链表尾部
func (l *List[T]) MoveToBack(n *Node[T]) {
	if !l.check(n) || l.root.prev == n {
		return
	}
	l.move(n, l.root.prev)
}

// MoveBefore 移动节点至链表指定位置前
func (l *List[T]) MoveBefore(n, at *Node[T]) {
	if !l.check(n) || !l.check(at) || n == at {
		return
	}
	l.move(n, at.prev)
}

// MoveAfter 移动节点至链表指定位置后
func (l *List[T]) MoveAfter(n, at *Node[T]) {
	if !l.check(n) || !l.check(at) || n == at {
		return
	}
	l.move(n, at)
}

// PushFrontList 在链表头部插入其他链表，可以传入自身
func (l *List[T]) PushFrontList(other *List[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, n := other.Len(), other.Back(); i > 0; i, n = i-1, n.Prev() {
		l.insertValue(n.V, &l.root)
	}
}

// PushBackList 在链表尾部插入其他链表，可以传入自身
func (l *List[T]) PushBackList(other *List[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, n := other.Len(), other.Front(); i > 0; i, n = i-1, n.Next() {
		l.insertValue(n.V, l.root.prev)
	}
}

// Traversal 遍历节点
func (l *List[T]) Traversal(visitor func(n *Node[T]) bool) {
	if visitor == nil {
		return
	}

	for n := l.Front(); n != nil; n = n.Next() {
		if !visitor(n) {
			break
		}
	}
}

// TraversalAt 从指定位置开始遍历节点
func (l *List[T]) TraversalAt(visitor func(n *Node[T]) bool, at *Node[T]) {
	if visitor == nil || !l.check(at) {
		return
	}

	for n := at; n != nil; n = n.Next() {
		if !visitor(n) {
			break
		}
	}
}

// ReversedTraversal 反向遍历节点
func (l *List[T]) ReversedTraversal(visitor func(n *Node[T]) bool) {
	if visitor == nil {
		return
	}

	for n := l.Back(); n != nil; n = n.Prev() {
		if !visitor(n) {
			break
		}
	}
}

// ReversedTraversalAt 从指定位置开始反向遍历节点
func (l *List[T]) ReversedTraversalAt(visitor func(n *Node[T]) bool, at *Node[T]) {
	if visitor == nil || !l.check(at) {
		return
	}

	for n := at; n != nil; n = n.Prev() {
		if !visitor(n) {
			break
		}
	}
}

// lazyInit 延迟初始化
func (l *List[T]) lazyInit() {
	if l.root.next != nil {
		return
	}
	l.root.next = &l.root
	l.root.prev = &l.root
}

// insertValue 插入数据
func (l *List[T]) insertValue(value T, at *Node[T]) *Node[T] {
	l.lazyInit()
	return l.insert(newNode(value), at)
}

// insert 插入节点
func (l *List[T]) insert(n, at *Node[T]) *Node[T] {
	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n
	n.list = l
	l.len++
	l.ver++
	n.ver = l.ver
	return n
}

// remove 删除节点
func (l *List[T]) remove(n *Node[T]) {
	l.lazyInit()

	n.prev.next = n.next
	n.next.prev = n.prev
	n.escaped = true
	l.len--
	l.ver++
}

// move 移动节点
func (l *List[T]) move(n, at *Node[T]) *Node[T] {
	l.lazyInit()

	if n == at {
		return n
	}
	n.prev.next = n.next
	n.next.prev = n.prev

	n.prev = at
	n.next = at.next
	n.prev.next = n
	n.next.prev = n

	l.ver++

	return n
}

// check 检查节点
func (l *List[T]) check(n *Node[T]) bool {
	return n != nil && !n.escaped && n.list == l && n.prev.next == n && n.next.prev == n
}
