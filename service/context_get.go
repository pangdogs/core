package service

import (
	"github.com/golaxy-kit/golaxy/ec"
	_ "unsafe"
)

// Get 获取服务上下文
func Get(ctxHolder ec.ContextHolder) Context {
	return getServiceContext(ctxHolder)
}

// TryGet  尝试获取服务上下文
func TryGet(ctxHolder ec.ContextHolder) (Context, bool) {
	serviceCtx, ok := tryGetServiceContext(ctxHolder)
	if !ok {
		return nil, false
	}
	return serviceCtx, true
}

//go:linkname getServiceContext github.com/golaxy-kit/golaxy/runtime.getServiceContext
func getServiceContext(ctxHolder ec.ContextHolder) Context

//go:linkname tryGetServiceContext github.com/golaxy-kit/golaxy/runtime.tryGetServiceContext
func tryGetServiceContext(ctxHolder ec.ContextHolder) (Context, bool)
