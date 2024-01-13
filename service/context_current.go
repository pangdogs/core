package service

import (
	"git.golaxy.org/core/internal/concurrent"
	_ "unsafe"
)

type (
	ContextProvider = concurrent.ContextProvider // 上下文提供者
)

//go:linkname getServiceContext git.golaxy.org/core/runtime.getServiceContext
func getServiceContext(ctxProvider concurrent.ContextProvider) Context

// Current 获取服务上下文
func Current(ctxProvider ContextProvider) Context {
	return getServiceContext(ctxProvider)
}
