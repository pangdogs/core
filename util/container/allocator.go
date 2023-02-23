package container

// Allocator 链表内存分配器
type Allocator[T any] interface {
	// Alloc 分配链表元素
	Alloc() *Element[T]
}

// DefaultAllocator 默认的链表内存分配器
func DefaultAllocator[T any]() Allocator[T] {
	return (*_DefaultAllocator[T])(nil)
}

type _DefaultAllocator[T any] struct{}

// Alloc 分配链表元素
func (*_DefaultAllocator[T]) Alloc() *Element[T] {
	return &Element[T]{}
}
