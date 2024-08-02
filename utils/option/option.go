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

package option

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

type Setting[T any] generic.Action1[*T]

func (s Setting[T]) Apply(opts *T) {
	generic.MakeAction1(s).Exec(opts)
}

func Make[T any](defaults Setting[T], settings ...Setting[T]) (opts T) {
	defaults.Apply(&opts)

	for i := range settings {
		settings[i].Apply(&opts)
	}

	return
}

func New[T any](defaults Setting[T], settings ...Setting[T]) *T {
	var opts T

	defaults.Apply(&opts)

	for i := range settings {
		settings[i].Apply(&opts)
	}

	return &opts
}

func Append[T any](opts T, settings ...Setting[T]) T {
	for i := range settings {
		settings[i].Apply(&opts)
	}
	return opts
}

func Change[T any](opts *T, settings ...Setting[T]) *T {
	if opts == nil {
		panic(fmt.Errorf("%w: %w: opts is nil", exception.ErrCore, exception.ErrArgs))
	}
	for i := range settings {
		settings[i].Apply(opts)
	}
	return opts
}
