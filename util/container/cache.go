package container

// NewCache 创建链表缓存
func NewCache[T any](size int) Cache[T] {
	return &_Cache[T]{
		size: size,
	}
}

// Cache 链表缓存
type Cache[T any] interface {
	Alloc() *Element[T]
}

type _Cache[T any] struct {
	heap  []Element[T]
	index int
	size  int
}

func (cache *_Cache[T]) Alloc() *Element[T] {
	if cache.index >= len(cache.heap) {
		if cache.size <= 0 {
			return &Element[T]{}
		}

		cache.index = 0
		cache.heap = make([]Element[T], cache.size)
	}

	e := &cache.heap[cache.index]
	cache.index++

	return e
}
