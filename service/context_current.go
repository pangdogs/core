package service

import (
	"git.golaxy.org/core/internal/concurrent"
	_ "unsafe"
)

type (
	ConcurrentContextProvider = concurrent.ConcurrentContextProvider // 多线程安全的上下文提供者
)

//go:linkname getServiceContext git.golaxy.org/core/runtime.getServiceContext
func getServiceContext(ctxProvider concurrent.ConcurrentContextProvider) Context

// Current 获取服务上下文
func Current(ctxProvider ConcurrentContextProvider) Context {
	return getServiceContext(ctxProvider)
}
