package service

import (
	"kit.golaxy.org/golaxy/ec"
	_ "unsafe"
)

// Get 获取服务上下文
func Get(ctxResolver ec.ContextResolver) Context {
	return getServiceContext(ctxResolver)
}

// TryGet  尝试获取服务上下文
func TryGet(ctxResolver ec.ContextResolver) (Context, bool) {
	serviceCtx, ok := tryGetServiceContext(ctxResolver)
	if !ok {
		return nil, false
	}
	return serviceCtx, true
}

//go:linkname getServiceContext kit.golaxy.org/golaxy/runtime.getServiceContext
func getServiceContext(ctxResolver ec.ContextResolver) Context

//go:linkname tryGetServiceContext kit.golaxy.org/golaxy/runtime.tryGetServiceContext
func tryGetServiceContext(ctxResolver ec.ContextResolver) (Context, bool)
