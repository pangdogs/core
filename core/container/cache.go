package container

// NewCache 创建链表缓存
func NewCache[T any](size int) *Cache[T] {
	cache := &Cache[T]{}
	cache.init(size)
	return cache
}

// Cache 链表缓存
type Cache[T any] struct {
	heap  []Element[T]
	index int
	size  int
}

// init 初始化
func (cache *Cache[T]) init(size int) {
	if cache == nil {
		return
	}

	cache.size = size
}

// alloc 分配空间
func (cache *Cache[T]) alloc() *Element[T] {
	if cache == nil {
		return &Element[T]{}
	}

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
