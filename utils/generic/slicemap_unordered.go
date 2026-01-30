/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package generic

import (
	"slices"

	"git.golaxy.org/core/utils/types"
)

func NewUnorderedSliceMap[K comparable, V any](kvs ...UnorderedKV[K, V]) *UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return &m
}

func NewUnorderedSliceMapVal[K comparable, V any](kvs ...UnorderedKV[K, V]) UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return m
}

func NewUnorderedSliceMapFromGoMap[K comparable, V any](dict map[K]V) *UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(dict))
	for k, v := range dict {
		m.Add(k, v)
	}
	return &m
}

func NewUnorderedSliceMapValFromGoMap[K comparable, V any](dict map[K]V) UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(dict))
	for k, v := range dict {
		m.Add(k, v)
	}
	return m
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

func (m UnorderedSliceMap[K, V]) Index(k K) int {
	return slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
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
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

func (m UnorderedSliceMap[K, V]) Each(fun Action2[K, V]) {
	for _, kv := range m {
		fun.UnsafeCall(kv.K, kv.V)
	}
}

func (m UnorderedSliceMap[K, V]) ReversedRange(fun Func2[K, V, bool]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

func (m UnorderedSliceMap[K, V]) ReversedEach(fun Action2[K, V]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		fun.UnsafeCall(kv.K, kv.V)
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
	rv := make(map[K]V, len(m))
	for _, kv := range m {
		rv[kv.K] = kv.V
	}
	return rv
}
