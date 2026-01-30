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

type Meta = generic.SliceMap[string, any]

func New(dict map[string]any) Meta {
	return generic.NewSliceMapFromGoMap(dict)
}

func Build() *MetaCreator {
	return &MetaCreator{}
}

type MetaCreator struct {
	meta Meta
}

func (c *MetaCreator) Add(k string, v any) *MetaCreator {
	c.meta.Add(k, v)
	return c
}

func (c *MetaCreator) TryAdd(k string, v any) *MetaCreator {
	c.meta.TryAdd(k, v)
	return c
}

func (c *MetaCreator) Merge(dict map[string]any) *MetaCreator {
	for k, v := range dict {
		c.meta.Add(k, v)
	}
	return c
}

func (c *MetaCreator) MergeIfAbsent(dict map[string]any) *MetaCreator {
	for k, v := range dict {
		c.meta.TryAdd(k, v)
	}
	return c
}

func (c *MetaCreator) Assign(m Meta) *MetaCreator {
	c.meta = m
	return c
}

func (c *MetaCreator) Delete(k string) *MetaCreator {
	c.meta.Delete(k)
	return c
}

func (c *MetaCreator) Clear() *MetaCreator {
	c.meta = Meta{}
	return c
}

func (c *MetaCreator) New() Meta {
	return c.meta.Clone()
}
