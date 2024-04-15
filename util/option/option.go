package option

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
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
