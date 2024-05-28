package service

import (
	"git.golaxy.org/core/internal/gctx"
	_ "unsafe"
)

//go:linkname getServiceContext git.golaxy.org/core/runtime.getServiceContext
func getServiceContext(provider gctx.ConcurrentContextProvider) Context

// Current 获取服务上下文
func Current(provider gctx.ConcurrentContextProvider) Context {
	return getServiceContext(provider)
}
