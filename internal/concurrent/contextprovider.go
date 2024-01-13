package concurrent

import (
	"git.golaxy.org/core/util/iface"
)

// ContextProvider 上下文提供者
type ContextProvider interface {
	// GetContext 获取上下文
	GetContext() iface.Cache
}

// CurrentContextProvider 当前上下文提供者
type CurrentContextProvider interface {
	ContextProvider
	// GetCurrentContext 获取当前上下文
	GetCurrentContext() iface.Cache
}

// ConcurrentContextProvider 多线程安全的上下文提供者
type ConcurrentContextProvider interface {
	ContextProvider
	// GetConcurrentContext 获取多线程安全的上下文
	GetConcurrentContext() iface.Cache
}
