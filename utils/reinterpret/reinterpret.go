package reinterpret

import "git.golaxy.org/core/utils/iface"

// CompositeProvider 支持重新解释类型
type CompositeProvider interface {
	GetCompositeFaceCache() iface.Cache
}

// Cast 重新解释类型
func Cast[T any](cp CompositeProvider) T {
	return iface.Cache2Iface[T](cp.GetCompositeFaceCache())
}
