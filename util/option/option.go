package option

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/generic"
)

type Setting[T any] generic.Action1[*T]

func (s Setting[T]) Cast() generic.Action1[*T] {
	return generic.Action1[*T](s)
}

func Make[T any](defaults Setting[T], settings ...Setting[T]) (opts T) {
	defaults.Cast().Exec(&opts)

	for i := range settings {
		settings[i].Cast().Exec(&opts)
	}

	return
}

func New[T any](defaults Setting[T], settings ...Setting[T]) *T {
	var opts T

	defaults.Cast().Exec(&opts)

	for i := range settings {
		settings[i].Cast().Exec(&opts)
	}

	return &opts
}

func Append[T any](opts T, settings ...Setting[T]) T {
	for i := range settings {
		settings[i].Cast().Exec(&opts)
	}
	return opts
}

func Change[T any](opts *T, settings ...Setting[T]) *T {
	if opts == nil {
		panic(fmt.Errorf("%w: opts is nil", internal.ErrArgs))
	}
	for i := range settings {
		settings[i].Cast().Exec(opts)
	}
	return opts
}
