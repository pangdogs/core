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
	"cmp"
	"slices"

	"git.golaxy.org/core/utils/types"
)

// NewSliceMap 从 kvs 构造按键升序排列的 SliceMap；重复键以后出现的值为准。
func NewSliceMap[K cmp.Ordered, V any](kvs ...KV[K, V]) SliceMap[K, V] {
	m := make(SliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return m
}

// NewSliceMapFromGoMap 将 Go map 转换为按键升序排列的 SliceMap。
func NewSliceMapFromGoMap[K cmp.Ordered, V any](dict map[K]V) SliceMap[K, V] {
	m := make(SliceMap[K, V], 0, len(dict))
	for k, v := range dict {
		m.Add(k, v)
	}
	return m
}

// KV 保存一个有序键值对。
type KV[K cmp.Ordered, V any] struct {
	K K // K 是键。
	V V // V 是值。
}

// SliceMap 是按键升序存储的小型切片映射。
//
// 查询使用二分搜索，增删需要移动切片元素。零值可用；类型不提供并发保护。
type SliceMap[K cmp.Ordered, V any] []KV[K, V]

// Add 添加或覆盖键值，并保持键升序。
func (m *SliceMap[K, V]) Add(k K, v V) {
	idx, ok := slices.BinarySearchFunc(*m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	if ok {
		(*m)[idx] = KV[K, V]{K: k, V: v}
	} else {
		*m = slices.Insert(*m, idx, KV[K, V]{K: k, V: v})
	}
}

// TryAdd 仅在键不存在时添加；成功返回 true。
func (m *SliceMap[K, V]) TryAdd(k K, v V) bool {
	idx, ok := slices.BinarySearchFunc(*m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	if !ok {
		*m = slices.Insert(*m, idx, KV[K, V]{K: k, V: v})
	}
	return !ok
}

// Delete 删除键并报告该键原先是否存在。
func (m *SliceMap[K, V]) Delete(k K) bool {
	idx, ok := slices.BinarySearchFunc(*m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	if ok {
		*m = slices.Delete(*m, idx, idx+1)
	}
	return ok
}

// Index 返回键在有序切片中的位置及是否存在。
func (m SliceMap[K, V]) Index(k K) (int, bool) {
	return slices.BinarySearchFunc(m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
}

// Get 返回键对应的值及是否存在。
func (m SliceMap[K, V]) Get(k K) (V, bool) {
	idx, ok := slices.BinarySearchFunc(m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	if ok {
		return m[idx].V, true
	}
	return types.Zero[V](), false
}

// Value 返回键对应的值；不存在时返回 V 的零值。
func (m SliceMap[K, V]) Value(k K) V {
	idx, ok := slices.BinarySearchFunc(m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	if ok {
		return m[idx].V
	}
	return types.Zero[V]()
}

// Exist 报告键是否存在。
func (m SliceMap[K, V]) Exist(k K) bool {
	_, ok := slices.BinarySearchFunc(m, KV[K, V]{K: k}, func(a, b KV[K, V]) int {
		return cmp.Compare(a.K, b.K)
	})
	return ok
}

// Len 返回键值对数量。
func (m SliceMap[K, V]) Len() int {
	return len(m)
}

// Range 按键升序遍历；fun 返回 false 时停止。
func (m SliceMap[K, V]) Range(fun Func2[K, V, bool]) {
	for _, kv := range m {
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

// Each 按键升序遍历全部键值对。
func (m SliceMap[K, V]) Each(fun Action2[K, V]) {
	for _, kv := range m {
		fun.UnsafeCall(kv.K, kv.V)
	}
}

// ReversedRange 按键降序遍历；fun 返回 false 时停止。
func (m SliceMap[K, V]) ReversedRange(fun Func2[K, V, bool]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

// ReversedEach 按键降序遍历全部键值对。
func (m SliceMap[K, V]) ReversedEach(fun Action2[K, V]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		fun.UnsafeCall(kv.K, kv.V)
	}
}

// Keys 返回按升序排列的键副本。
func (m SliceMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for _, kv := range m {
		keys = append(keys, kv.K)
	}
	return keys
}

// Values 返回按键升序对应的值副本。
func (m SliceMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, kv := range m {
		values = append(values, kv.V)
	}
	return values
}

// Clone 返回映射的浅拷贝。
func (m SliceMap[K, V]) Clone() SliceMap[K, V] {
	return slices.Clone(m)
}

// ToUnorderedSliceMap 转换为保持当前键顺序的 UnorderedSliceMap。
func (m SliceMap[K, V]) ToUnorderedSliceMap() UnorderedSliceMap[K, V] {
	rv := make(UnorderedSliceMap[K, V], 0, len(m))
	for _, kv := range m {
		rv.Add(kv.K, kv.V)
	}
	return rv
}

// ToGoMap 将键值对复制到新的 Go map。
func (m SliceMap[K, V]) ToGoMap() map[K]V {
	rv := make(map[K]V, len(m))
	for _, kv := range m {
		rv[kv.K] = kv.V
	}
	return rv
}
