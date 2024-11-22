package event

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// Cache 引用iface.Cache
type Cache = iface.Cache

// Cache2Iface 引用iface.Cache2Iface
func Cache2Iface[T any](c Cache) T {
	return iface.Cache2Iface[T](c)
}

// Panicf 引用exception.Panicf
func Panicf(format string, args ...any) {
	exception.Panicf(format, args...)
}
