package generic

import (
	"git.golaxy.org/core/utils/types"
	"slices"
)

func MakeSliceMap[K comparable, V any](kvs ...KV[K, V]) SliceMap[K, V] {
	kvs = slices.Clone(kvs)
	for i := len(kvs) - 1; i >= 0; i-- {
		it := kvs[i]

		if !slices.ContainsFunc(kvs[:i], func(kv KV[K, V]) bool {
			return kv.K == it.K
		}) {
			continue
		}

		kvs = append(kvs[:i], kvs[i+1:]...)
	}
	return kvs
}

func NewSliceMap[K comparable, V any](kvs ...KV[K, V]) *SliceMap[K, V] {
	kvs = slices.Clone(kvs)
	for i := len(kvs) - 1; i >= 0; i-- {
		it := kvs[i]

		if !slices.ContainsFunc(kvs[:i], func(kv KV[K, V]) bool {
			return kv.K == it.K
		}) {
			continue
		}

		kvs = append(kvs[:i], kvs[i+1:]...)
	}
	return (*SliceMap[K, V])(&kvs)
}

func MakeSliceMapFromGoMap[K comparable, V any](m map[K]V) SliceMap[K, V] {
	sm := make(SliceMap[K, V], 0, len(m))
	for k, v := range m {
		sm = append(sm, KV[K, V]{K: k, V: v})
	}
	return sm
}

func NewSliceMapFromGoMap[K comparable, V any](m map[K]V) *SliceMap[K, V] {
	sm := make(SliceMap[K, V], 0, len(m))
	for k, v := range m {
		sm = append(sm, KV[K, V]{K: k, V: v})
	}
	return &sm
}

type KV[K comparable, V any] struct {
	K K
	V V
}

type SliceMap[K comparable, V any] []KV[K, V]

func (m *SliceMap[K, V]) Add(k K, v V) {
	idx := slices.IndexFunc(*m, func(kv KV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		(*m)[idx] = KV[K, V]{K: k, V: v}
		return
	}
	*m = append(*m, KV[K, V]{K: k, V: v})
}

func (m *SliceMap[K, V]) TryAdd(k K, v V) bool {
	idx := slices.IndexFunc(*m, func(kv KV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return false
	}
	*m = append(*m, KV[K, V]{K: k, V: v})
	return true
}

func (m *SliceMap[K, V]) Delete(k K) bool {
	var ok bool
	*m = slices.DeleteFunc(*m, func(kv KV[K, V]) bool {
		ok = kv.K == k
		return ok
	})
	return ok
}

func (m SliceMap[K, V]) Get(k K) (V, bool) {
	idx := slices.IndexFunc(m, func(kv KV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V, true
	}
	return types.ZeroT[V](), false
}

func (m SliceMap[K, V]) Value(k K) V {
	idx := slices.IndexFunc(m, func(kv KV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V
	}
	return types.ZeroT[V]()
}

func (m SliceMap[K, V]) Exist(k K) bool {
	return slices.ContainsFunc(m, func(kv KV[K, V]) bool {
		return kv.K == k
	})
}

func (m SliceMap[K, V]) ToGoMap() map[K]V {
	rv := make(map[K]V, len(m))
	for _, kv := range m {
		rv[kv.K] = kv.V
	}
	return rv
}
