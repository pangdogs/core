package gctx

import (
	"git.golaxy.org/core/utils/iface"
)

// CurrentContextProvider 当前上下文提供者
type CurrentContextProvider interface {
	ConcurrentContextProvider
	// GetCurrentContext 获取当前上下文
	GetCurrentContext() iface.Cache
}

// ConcurrentContextProvider 多线程安全的上下文提供者
type ConcurrentContextProvider interface {
	// GetConcurrentContext 获取多线程安全的上下文
	GetConcurrentContext() iface.Cache
}
