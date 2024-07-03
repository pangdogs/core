package generic

import (
	"git.golaxy.org/core/utils/types"
	"slices"
)

func MakeUnorderedSliceMap[K comparable, V any](kvs ...UnorderedKV[K, V]) UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return m
}

func NewUnorderedSliceMap[K comparable, V any](kvs ...UnorderedKV[K, V]) *UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return &m
}

func MakeUnorderedSliceMapFromGoMap[K comparable, V any](m map[K]V) UnorderedSliceMap[K, V] {
	sm := make(UnorderedSliceMap[K, V], 0, len(m))
	for k, v := range m {
		sm.Add(k, v)
	}
	return sm
}

func NewUnorderedSliceMapFromGoMap[K comparable, V any](m map[K]V) *UnorderedSliceMap[K, V] {
	sm := make(UnorderedSliceMap[K, V], 0, len(m))
	for k, v := range m {
		sm.Add(k, v)
	}
	return &sm
}

type UnorderedKV[K comparable, V any] struct {
	K K
	V V
}

type UnorderedSliceMap[K comparable, V any] []UnorderedKV[K, V]

func (m *UnorderedSliceMap[K, V]) Add(k K, v V) {
	idx := slices.IndexFunc(*m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		(*m)[idx] = UnorderedKV[K, V]{K: k, V: v}
	} else {
		*m = append(*m, UnorderedKV[K, V]{K: k, V: v})
	}
}

func (m *UnorderedSliceMap[K, V]) TryAdd(k K, v V) bool {
	idx := slices.IndexFunc(*m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx < 0 {
		*m = append(*m, UnorderedKV[K, V]{K: k, V: v})
	}
	return idx < 0
}

func (m *UnorderedSliceMap[K, V]) Delete(k K) bool {
	idx := slices.IndexFunc(*m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		*m = slices.Delete(*m, idx, idx+1)
	}
	return idx >= 0
}

func (m UnorderedSliceMap[K, V]) Get(k K) (V, bool) {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V, true
	}
	return types.ZeroT[V](), false
}

func (m UnorderedSliceMap[K, V]) Value(k K) V {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V
	}
	return types.ZeroT[V]()
}

func (m UnorderedSliceMap[K, V]) Exist(k K) bool {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	return idx >= 0
}

func (m UnorderedSliceMap[K, V]) Len() int {
	return len(m)
}

func (m UnorderedSliceMap[K, V]) Range(fun Func2[K, V, bool]) {
	for _, kv := range m {
		if !fun.Exec(kv.K, kv.V) {
			return
		}
	}
}

func (m UnorderedSliceMap[K, V]) Each(fun Action2[K, V]) {
	for _, kv := range m {
		fun.Exec(kv.K, kv.V)
	}
}

func (m UnorderedSliceMap[K, V]) ReversedRange(fun Func2[K, V, bool]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		if !fun.Exec(kv.K, kv.V) {
			return
		}
	}
}

func (m UnorderedSliceMap[K, V]) ReversedEach(fun Action2[K, V]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		fun.Exec(kv.K, kv.V)
	}
}

func (m UnorderedSliceMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for _, kv := range m {
		keys = append(keys, kv.K)
	}
	return keys
}

func (m UnorderedSliceMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, kv := range m {
		values = append(values, kv.V)
	}
	return values
}

func (m UnorderedSliceMap[K, V]) Clone() UnorderedSliceMap[K, V] {
	return slices.Clone(m)
}

func (m UnorderedSliceMap[K, V]) ToGoMap() map[K]V {
	gm := make(map[K]V, len(m))
	for _, kv := range m {
		gm[kv.K] = kv.V
	}
	return gm
}
