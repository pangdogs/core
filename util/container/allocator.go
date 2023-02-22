package container

// NewAllocator 创建链表内存分配器
func NewAllocator[T any](size int) Allocator[T] {
	if size <= 0 {
		panic("size less equal 0 is invalid")
	}
	return &_Allocator[T]{
		size: size,
	}
}

// Allocator 链表内存分配器
type Allocator[T any] interface {
	// Alloc 分配链表元素
	Alloc() *Element[T]
}

type _Allocator[T any] struct {
	heap  []Element[T]
	index int
	size  int
}

// Alloc 分配链表元素
func (allocator *_Allocator[T]) Alloc() *Element[T] {
	if allocator.index >= len(allocator.heap) {
		allocator.index = 0
		allocator.heap = make([]Element[T], allocator.size)
	}

	e := &allocator.heap[allocator.index]
	allocator.index++

	return e
}
