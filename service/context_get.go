package service

import (
	"kit.golaxy.org/golaxy/internal"
	_ "unsafe"
)

// Get 获取服务上下文
func Get(ctxResolver internal.ContextResolver) Context {
	return getServiceContext(ctxResolver)
}

//go:linkname getServiceContext kit.golaxy.org/golaxy/runtime.getServiceContext
func getServiceContext(ctxResolver internal.ContextResolver) Context
