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

package meta

import (
	"git.golaxy.org/core/utils/generic"
)

// Meta 是以字符串为键、按键升序保存条目的元数据映射。
type Meta = generic.SliceMap[string, any]

// New 从 dict 构造按键升序排列的 Meta；值会复制到新切片中。
func New(dict map[string]any) Meta {
	return generic.NewSliceMapFromGoMap(dict)
}

// Build 创建空的链式元数据构造器。
func Build() *MetaCreator {
	return &MetaCreator{}
}

// MetaCreator 通过链式调用构造 Meta；零值可用。
type MetaCreator struct {
	meta Meta
}

// Add 添加或覆盖键值并返回构造器。
func (c *MetaCreator) Add(k string, v any) *MetaCreator {
	c.meta.Add(k, v)
	return c
}

// TryAdd 仅在键不存在时添加值，并返回构造器。
func (c *MetaCreator) TryAdd(k string, v any) *MetaCreator {
	c.meta.TryAdd(k, v)
	return c
}

// Merge 合并 dict；同名键会被覆盖。
func (c *MetaCreator) Merge(dict map[string]any) *MetaCreator {
	for k, v := range dict {
		c.meta.Add(k, v)
	}
	return c
}

// MergeIfAbsent 合并 dict，但保留已有的同名键。
func (c *MetaCreator) MergeIfAbsent(dict map[string]any) *MetaCreator {
	for k, v := range dict {
		c.meta.TryAdd(k, v)
	}
	return c
}

// Assign 直接绑定 m；m 不会在此时复制。
func (c *MetaCreator) Assign(m Meta) *MetaCreator {
	c.meta = m
	return c
}

// Delete 删除指定键并返回构造器。
func (c *MetaCreator) Delete(k string) *MetaCreator {
	c.meta.Delete(k)
	return c
}

// Clear 清空全部元数据并返回构造器。
func (c *MetaCreator) Clear() *MetaCreator {
	c.meta = Meta{}
	return c
}

// New 返回当前元数据的副本。
func (c *MetaCreator) New() Meta {
	return c.meta.Clone()
}
