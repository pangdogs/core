package event

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

type Cache = iface.Cache

func Cache2Iface[T any](c Cache) T {
	return iface.Cache2Iface[T](c)
}

func Panicf(format string, args ...any) {
	exception.Panicf(format, args...)
}
