package generic

func _NewElement[T any](value T) *Element[T] {
	return &Element[T]{
		Value: value,
	}
}

// Element 元素
type Element[T any] struct {
	_next, _prev *Element[T]
	list         *List[T]
	ver          int64
	escaped      bool
	Value        T
}

// Version 创建时的数据变化版本
func (e *Element[T]) Version() int64 {
	return e.ver
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
	if e.list == nil || e.escaped {
		return
	}
	e.list.remove(e)
	e.escaped = true
}

// Escaped 是否已从链表中删除
func (e *Element[T]) Escaped() bool {
	return e.escaped
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
