// Package container 实现了一种特殊的链表，可以在遍历时在任意位置添加或删除元素，递归删除元素时仍然能正常工作，但是需要GC支持，主要用于core内部。
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
