package generic

// NewList 创建链表
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// List 链表，可以在遍历时在任意位置添加或删除元素，递归添加或删除元素时仍然能正常工作，非线程安全。
type List[T any] struct {
	root Element[T]
	len  int
	ver  int64
}

// Len 链表长度
func (l *List[T]) Len() int {
	return l.len
}

// Version 数据变化版本
func (l *List[T]) Version() int64 {
	return l.ver
}

// Front 链表头部
func (l *List[T]) Front() *Element[T] {
	if l.len <= 0 {
		return nil
	}
	return l.root._next
}

// Back 链表尾部
func (l *List[T]) Back() *Element[T] {
	if l.len <= 0 {
		return nil
	}
	return l.root._prev
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
	l.lazyInit()
	return l.insertValue(value, &l.root)
}

// PushBack 在链表尾部插入数据
func (l *List[T]) PushBack(value T) *Element[T] {
	l.lazyInit()
	return l.insertValue(value, l.root._prev)
}

// InsertBefore 在链表指定位置前插入数据
func (l *List[T]) InsertBefore(value T, at *Element[T]) *Element[T] {
	if !l.check(at) {
		return nil
	}
	return l.insertValue(value, at._prev)
}

// InsertAfter 在链表指定位置后插入数据
func (l *List[T]) InsertAfter(value T, at *Element[T]) *Element[T] {
	if !l.check(at) {
		return nil
	}
	return l.insertValue(value, at)
}

// MoveToFront 移动元素至链表头部
func (l *List[T]) MoveToFront(e *Element[T]) {
	if !l.check(e) || l.root._next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack 移动元素至链表尾部
func (l *List[T]) MoveToBack(e *Element[T]) {
	if !l.check(e) || l.root._prev == e {
		return
	}
	l.move(e, l.root._prev)
}

// MoveBefore 移动元素至链表指定位置前
func (l *List[T]) MoveBefore(e, at *Element[T]) {
	if !l.check(e) || !l.check(at) || e == at {
		return
	}
	l.move(e, at._prev)
}

// MoveAfter 移动元素至链表指定位置后
func (l *List[T]) MoveAfter(e, at *Element[T]) {
	if !l.check(e) || !l.check(at) || e == at {
		return
	}
	l.move(e, at)
}

// PushFrontList 在链表头部插入其他链表，可以传入自身
func (l *List[T]) PushFrontList(other *List[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// PushBackList 在链表尾部插入其他链表，可以传入自身
func (l *List[T]) PushBackList(other *List[T]) {
	if other == nil {
		return
	}
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root._prev)
	}
}

// Traversal 遍历元素
func (l *List[T]) Traversal(visitor func(e *Element[T]) bool) {
	if visitor == nil {
		return
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if !visitor(e) {
			break
		}
	}
}

// TraversalAt 从指定位置开始遍历元素
func (l *List[T]) TraversalAt(visitor func(e *Element[T]) bool, at *Element[T]) {
	if visitor == nil || !l.check(at) {
		return
	}

	for e := at; e != nil; e = e.Next() {
		if !visitor(e) {
			break
		}
	}
}

// ReversedTraversal 反向遍历元素
func (l *List[T]) ReversedTraversal(visitor func(e *Element[T]) bool) {
	if visitor == nil {
		return
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		if !visitor(e) {
			break
		}
	}
}

// ReversedTraversalAt 从指定位置开始反向遍历元素
func (l *List[T]) ReversedTraversalAt(visitor func(e *Element[T]) bool, at *Element[T]) {
	if visitor == nil || !l.check(at) {
		return
	}

	for e := at; e != nil; e = e.Prev() {
		if !visitor(e) {
			break
		}
	}
}

// lazyInit 延迟初始化
func (l *List[T]) lazyInit() {
	if l.root._next != nil {
		return
	}
	l.root._next = &l.root
	l.root._prev = &l.root
}

// insertValue 插入数据
func (l *List[T]) insertValue(value T, at *Element[T]) *Element[T] {
	return l.insert(_NewElement(value), at)
}

// insert 插入元素
func (l *List[T]) insert(e, at *Element[T]) *Element[T] {
	l.lazyInit()
	e._prev = at
	e._next = at._next
	e._prev._next = e
	e._next._prev = e
	e.list = l
	l.len++
	l.ver++
	e.ver = l.ver
	return e
}

// remove 删除元素
func (l *List[T]) remove(e *Element[T]) {
	l.lazyInit()
	e._prev._next = e._next
	e._next._prev = e._prev
	l.len--
	l.ver++
}

// move 移动元素
func (l *List[T]) move(e, at *Element[T]) *Element[T] {
	l.lazyInit()

	if e == at {
		return e
	}
	e._prev._next = e._next
	e._next._prev = e._prev

	e._prev = at
	e._next = at._next
	e._prev._next = e
	e._next._prev = e

	l.ver++

	return e
}

func (l *List[T]) check(e *Element[T]) bool {
	return e != nil && e.list == l && !e.escaped
}
