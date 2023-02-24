// Package container 实现了一种特殊的链表，可以在遍历时在任意位置添加或删除元素，递归添加或删除元素时仍然能正常工作，但是需要GC支持，主要用于ec包内部。
package container

// Element 元素
type Element[T any] struct {
	_next, _prev *Element[T]
	list         *List[T]
	escaped      bool
	Value        T
}

// next 下一个元素，包含正在删除的元素
func (e *Element[T]) next() *Element[T] {
	if n := e._next; e.list != nil && n != &e.list.root {
		return n
	}
	return nil
}

// prev 前一个元素，包含正在删除的元素
func (e *Element[T]) prev() *Element[T] {
	if p := e._prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Next 下一个元素
func (e *Element[T]) Next() *Element[T] {
	for n := e.next(); n != nil; n = n.next() {
		if !n.escaped {
			return n
		}
	}
	return nil
}

// Prev 前一个元素
func (e *Element[T]) Prev() *Element[T] {
	for p := e.prev(); p != nil; p = p.prev() {
		if !p.escaped {
			return p
		}
	}
	return nil
}

// Escape 从链表中删除
func (e *Element[T]) Escape() {
	if e.list != nil {
		if !e.escaped {
			e.escaped = true
			e.list.markGC()
		}
	}
}

// Escaped 是否已从链表中删除
func (e *Element[T]) Escaped() bool {
	return e.escaped
}

// release 从链表中释放
func (e *Element[T]) release() {
	*e = Element[T]{}
}

// Released 是否已从链表中释放
func (e *Element[T]) Released() bool {
	return e.list == nil
}

// NewList 创建链表
func NewList[T any](allocator Allocator[T], gcCollector GCCollector) *List[T] {
	return new(List[T]).Init(allocator, gcCollector)
}

// List 链表，非线程安全，支持在遍历过程中删除元素
type List[T any] struct {
	allocator   Allocator[T]
	gcCollector GCCollector
	root        Element[T]
	cap, gcLen  int
}

// Init 初始化
func (l *List[T]) Init(allocator Allocator[T], gcCollector GCCollector) *List[T] {
	if allocator == nil {
		allocator = DefaultAllocator[T]()
	}
	l.allocator = allocator
	l.gcCollector = gcCollector
	l.root._next = &l.root
	l.root._prev = &l.root
	l.cap = 0
	l.gcLen = 0
	return l
}

// SetAllocator 设置链表内存分配器
func (l *List[T]) SetAllocator(allocator Allocator[T]) {
	if allocator == nil {
		allocator = DefaultAllocator[T]()
	}

	if l.allocator == allocator {
		return
	}

	l.allocator = allocator
}

// GetAllocator 获取链表内存分配器
func (l *List[T]) GetAllocator() Allocator[T] {
	return l.allocator
}

// SetGCCollector 设置GC收集器
func (l *List[T]) SetGCCollector(gcCollector GCCollector) {
	if l.gcCollector == gcCollector {
		return
	}

	l.gcCollector = gcCollector

	if l.gcCollector != nil {
		l.gcCollector.CollectGC(l)
	}
}

// GetGCCollector 获取GC收集器
func (l *List[T]) GetGCCollector() GCCollector {
	return l.gcCollector
}

// GC 执行GC
func (l *List[T]) GC() {
	if l.gcLen <= 0 {
		return
	}

	for e := l.Front(); e != nil; {
		if e.escaped {
			t := e.next()
			l.release(e)
			e = t
		} else {
			e = e.next()
		}
	}
}

// NeedGC 是否需要GC
func (l *List[T]) NeedGC() bool {
	return l.gcLen > 0
}

// markGC 标记需要GC
func (l *List[T]) markGC() {
	l.gcLen++
	if l.gcLen == 1 && l.gcCollector != nil {
		l.gcCollector.CollectGC(l)
	}
}

// Cap 链表容量，包含已标记需要GC的元素
func (l *List[T]) Cap() int {
	return l.cap
}

// Len 链表长度
func (l *List[T]) Len() int {
	return l.cap - l.gcLen
}

// Front 链表头部
func (l *List[T]) Front() *Element[T] {
	if l.cap == 0 {
		return nil
	}
	return l.root._next
}

// Back 链表尾部
func (l *List[T]) Back() *Element[T] {
	if l.cap == 0 {
		return nil
	}
	return l.root._prev
}

// insert 插入元素
func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	e._prev = at
	e._next = at._next
	e._prev._next = e
	e._next._prev = e
	e.list = l
	l.cap++
	return e
}

// insertValue 插入数据
func (l *List[T]) insertValue(value T, at *Element[T]) *Element[T] {
	e := l.allocator.Alloc()
	e.Value = value
	return l.insert(e, at)
}

// release 释放元素
func (l *List[T]) release(e *Element[T]) {
	e._prev._next = e._next
	e._next._prev = e._prev
	e.release()
	l.cap--
	l.gcLen--
}

// move 移动元素
func (l *List[T]) move(e, at *Element[T]) *Element[T] {
	if e == at {
		return e
	}
	e._prev._next = e._next
	e._next._prev = e._prev

	e._prev = at
	e._next = at._next
	e._prev._next = e
	e._next._prev = e

	return e
}

// Remove 删除元素
func (l *List[T]) Remove(e *Element[T]) T {
	if e.list == l {
		e.Escape()
	}
	return e.Value
}

// PushFront 在链表头部插入数据
func (l *List[T]) PushFront(value T) *Element[T] {
	return l.insertValue(value, &l.root)
}

// PushBack 在链表尾部插入数据
func (l *List[T]) PushBack(value T) *Element[T] {
	return l.insertValue(value, l.root._prev)
}

// InsertBefore 在链表指定位置前插入数据
func (l *List[T]) InsertBefore(value T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(value, mark._prev)
}

// InsertAfter 在链表指定位置后插入数据
func (l *List[T]) InsertAfter(value T, mark *Element[T]) *Element[T] {
	if mark.list != l {
		return nil
	}
	return l.insertValue(value, mark)
}

// MoveToFront 移动元素至链表头部
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e.list != l || l.root._next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack 移动元素至链表尾部
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e.list != l || l.root._prev == e {
		return
	}
	l.move(e, l.root._prev)
}

// MoveBefore 移动元素至链表指定位置前
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark._prev)
}

// MoveAfter 移动元素至链表指定位置后
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

// PushBackList 在链表尾部插入其他链表
func (l *List[T]) PushBackList(other *List[T]) {
	for i, e := other.Cap(), other.Front(); i > 0; i, e = i-1, e.next() {
		l.insertValue(e.Value, l.root._prev)
	}
}

// PushFrontList 在链表头部插入其他链表
func (l *List[T]) PushFrontList(other *List[T]) {
	for i, e := other.Cap(), other.Back(); i > 0; i, e = i-1, e.prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// Traversal 遍历元素
func (l *List[T]) Traversal(visitor func(e *Element[T]) bool) {
	if visitor == nil {
		return
	}

	for e := l.Front(); e != nil; e = e.next() {
		if !e.escaped && !visitor(e) {
			break
		}
	}
}

// TraversalAt 从指定位置开始遍历元素
func (l *List[T]) TraversalAt(visitor func(e *Element[T]) bool, mark *Element[T]) {
	if visitor == nil || mark.list != l {
		return
	}

	for e := mark; e != nil; e = e.next() {
		if !e.escaped && !visitor(e) {
			break
		}
	}
}

// ReverseTraversal 反向遍历元素
func (l *List[T]) ReverseTraversal(visitor func(e *Element[T]) bool) {
	if visitor == nil {
		return
	}

	for e := l.Back(); e != nil; e = e.prev() {
		if !e.escaped && !visitor(e) {
			break
		}
	}
}

// ReverseTraversalAt 从指定位置开始反向遍历元素
func (l *List[T]) ReverseTraversalAt(visitor func(e *Element[T]) bool, mark *Element[T]) {
	if visitor == nil || mark.list != l {
		return
	}

	for e := mark; e != nil; e = e.prev() {
		if !e.escaped && !visitor(e) {
			break
		}
	}
}
