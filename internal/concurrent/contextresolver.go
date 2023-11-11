package concurrent

import (
	"kit.golaxy.org/golaxy/util/iface"
)

// ContextResolver 上下文获取器
type ContextResolver interface {
	// ResolveContext 解析上下文
	ResolveContext() iface.Cache
}

// CurrentContextResolver 当前上下文获取器
type CurrentContextResolver interface {
	ContextResolver
	// ResolveCurrentContext 解析当前上下文
	ResolveCurrentContext() iface.Cache
}

// ConcurrentContextResolver 多线程安全的上下文获取器
type ConcurrentContextResolver interface {
	ContextResolver
	// ResolveConcurrentContext 解析多线程安全的上下文
	ResolveConcurrentContext() iface.Cache
}
