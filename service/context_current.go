package service

import (
	"kit.golaxy.org/golaxy/internal"
	_ "unsafe"
)

// Current 获取当前服务上下文
func Current(ctxResolver internal.ContextResolver) Context {
	return getServiceContext(ctxResolver)
}

//go:linkname getServiceContext kit.golaxy.org/golaxy/runtime.getServiceContext
func getServiceContext(ctxResolver internal.ContextResolver) Context
