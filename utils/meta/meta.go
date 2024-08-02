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

import "git.golaxy.org/core/utils/generic"

type Meta = generic.SliceMap[string, any]

type _MetaCtor struct {
	meta Meta
}

func (c _MetaCtor) Add(k string, v any) _MetaCtor {
	c.meta.Add(k, v)
	return c
}

func (c _MetaCtor) Delete(k string) _MetaCtor {
	c.meta.Delete(k)
	return c
}

func (c _MetaCtor) Combine(m map[string]any) _MetaCtor {
	for k, v := range m {
		c.meta.TryAdd(k, v)
	}
	return c
}

func (c _MetaCtor) Override(m map[string]any) _MetaCtor {
	for k, v := range m {
		c.meta.Add(k, v)
	}
	return c
}

func (c _MetaCtor) Clean() _MetaCtor {
	c.meta = c.meta[:0]
	return c
}

func (c _MetaCtor) Get() Meta {
	return c.meta
}

func Make() _MetaCtor {
	return _MetaCtor{}
}
