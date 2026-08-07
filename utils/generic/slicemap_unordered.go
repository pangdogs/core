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

// NewUnorderedSliceMap 按 kvs 顺序构造 UnorderedSliceMap；重复键以后出现的值为准。
func NewUnorderedSliceMap[K comparable, V any](kvs ...UnorderedKV[K, V]) UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		m.Add(kv.K, kv.V)
	}
	return m
}

// NewUnorderedSliceMapFromGoMap 将 Go map 转换为 UnorderedSliceMap；初始顺序不确定。
func NewUnorderedSliceMapFromGoMap[K comparable, V any](dict map[K]V) UnorderedSliceMap[K, V] {
	m := make(UnorderedSliceMap[K, V], 0, len(dict))
	for k, v := range dict {
		m.Add(k, v)
	}
	return m
}

// UnorderedKV 保存一个可比较键值对。
type UnorderedKV[K comparable, V any] struct {
	K K // K 是键。
	V V // V 是值。
}

// UnorderedSliceMap 是按加入顺序保存、使用线性查找的小型切片映射。
//
// “Unordered”表示不按键排序；更新已有键不会改变其位置。零值可用且不提供并发保护。
type UnorderedSliceMap[K comparable, V any] []UnorderedKV[K, V]

// Add 添加或覆盖键值；覆盖不会改变键的位置。
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

// TryAdd 仅在键不存在时追加；成功返回 true。
func (m *UnorderedSliceMap[K, V]) TryAdd(k K, v V) bool {
	idx := slices.IndexFunc(*m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx < 0 {
		*m = append(*m, UnorderedKV[K, V]{K: k, V: v})
	}
	return idx < 0
}

// Delete 删除键并报告该键原先是否存在；其余元素顺序保持不变。
func (m *UnorderedSliceMap[K, V]) Delete(k K) bool {
	idx := slices.IndexFunc(*m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		*m = slices.Delete(*m, idx, idx+1)
	}
	return idx >= 0
}

// Index 返回键的位置；不存在时返回 -1。
func (m UnorderedSliceMap[K, V]) Index(k K) int {
	return slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
}

// Get 返回键对应的值及是否存在。
func (m UnorderedSliceMap[K, V]) Get(k K) (V, bool) {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V, true
	}
	return types.Zero[V](), false
}

// Value 返回键对应的值；不存在时返回 V 的零值。
func (m UnorderedSliceMap[K, V]) Value(k K) V {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	if idx >= 0 {
		return m[idx].V
	}
	return types.Zero[V]()
}

// Exist 报告键是否存在。
func (m UnorderedSliceMap[K, V]) Exist(k K) bool {
	idx := slices.IndexFunc(m, func(kv UnorderedKV[K, V]) bool {
		return kv.K == k
	})
	return idx >= 0
}

// Len 返回键值对数量。
func (m UnorderedSliceMap[K, V]) Len() int {
	return len(m)
}

// Range 按存储顺序遍历；fun 返回 false 时停止。
func (m UnorderedSliceMap[K, V]) Range(fun Func2[K, V, bool]) {
	for _, kv := range m {
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

// Each 按存储顺序遍历全部键值对。
func (m UnorderedSliceMap[K, V]) Each(fun Action2[K, V]) {
	for _, kv := range m {
		fun.UnsafeCall(kv.K, kv.V)
	}
}

// ReversedRange 按存储顺序的逆序遍历；fun 返回 false 时停止。
func (m UnorderedSliceMap[K, V]) ReversedRange(fun Func2[K, V, bool]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		if !fun.UnsafeCall(kv.K, kv.V) {
			return
		}
	}
}

// ReversedEach 按存储顺序的逆序遍历全部键值对。
func (m UnorderedSliceMap[K, V]) ReversedEach(fun Action2[K, V]) {
	for i := len(m) - 1; i >= 0; i-- {
		kv := m[i]
		fun.UnsafeCall(kv.K, kv.V)
	}
}

// Keys 返回按存储顺序排列的键副本。
func (m UnorderedSliceMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for _, kv := range m {
		keys = append(keys, kv.K)
	}
	return keys
}

// Values 返回按存储顺序排列的值副本。
func (m UnorderedSliceMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, kv := range m {
		values = append(values, kv.V)
	}
	return values
}

// Clone 返回映射的浅拷贝。
func (m UnorderedSliceMap[K, V]) Clone() UnorderedSliceMap[K, V] {
	return slices.Clone(m)
}

// ToGoMap 将键值对复制到新的 Go map。
func (m UnorderedSliceMap[K, V]) ToGoMap() map[K]V {
	rv := make(map[K]V, len(m))
	for _, kv := range m {
		rv[kv.K] = kv.V
	}
	return rv
}
