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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

type Setting[T any] generic.Action1[*T]

func (s Setting[T]) Apply(options *T) {
	generic.CastAction1(s).UnsafeCall(options)
}

func Make[T any](defaults Setting[T], settings ...Setting[T]) (options T) {
	defaults.Apply(&options)

	for i := range settings {
		settings[i].Apply(&options)
	}

	return
}

func New[T any](defaults Setting[T], settings ...Setting[T]) *T {
	var options T

	defaults.Apply(&options)

	for i := range settings {
		settings[i].Apply(&options)
	}

	return &options
}

func Append[T any](options T, settings ...Setting[T]) T {
	for i := range settings {
		settings[i].Apply(&options)
	}
	return options
}

func Change[T any](options *T, settings ...Setting[T]) *T {
	if options == nil {
		exception.Panicf("%w: %w: options is nil", exception.ErrCore, exception.ErrArgs)
	}
	for i := range settings {
		settings[i].Apply(options)
	}
	return options
}
