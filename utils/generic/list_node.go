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

func newNode[T any](value T) *Node[T] {
	return &Node[T]{
		V: value,
	}
}

// Node 链表节点
type Node[T any] struct {
	V          T
	next, prev *Node[T]
	list       *List[T]
	ver        int64
	escaped    bool
}

// Version 创建时的数据版本
func (n *Node[T]) Version() int64 {
	return n.ver
}

// Next 下一个节点
func (n *Node[T]) Next() *Node[T] {
	for next := n.getNext(); next != nil; next = next.getNext() {
		if !next.escaped {
			return next
		}
	}
	return nil
}

// Prev 前一个节点
func (n *Node[T]) Prev() *Node[T] {
	for prev := n.getPrev(); prev != nil; prev = prev.getPrev() {
		if !prev.escaped {
			return prev
		}
	}
	return nil
}

// Escape 从链表中删除
func (n *Node[T]) Escape() {
	if n.list == nil || !n.list.check(n) {
		return
	}
	n.list.remove(n)
}

// Escaped 是否已从链表中删除
func (n *Node[T]) Escaped() bool {
	return n.escaped
}

// getNext 下一个节点，包含正在删除的节点
func (n *Node[T]) getNext() *Node[T] {
	if n.list == nil {
		return nil
	}
	if next := n.next; next != &n.list.root {
		return next
	}
	return nil
}

// getPrev 前一个节点，包含正在删除的节点
func (n *Node[T]) getPrev() *Node[T] {
	if n.list == nil {
		return nil
	}
	if prev := n.prev; prev != &n.list.root {
		return prev
	}
	return nil
}
