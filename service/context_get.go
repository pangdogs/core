package service

import (
	"kit.golaxy.org/golaxy/ec"
	_ "unsafe"
)

// Get 获取服务上下文
func Get(ctxResolver ec.ContextResolver) Context {
	return getServiceContext(ctxResolver)
}

//go:linkname getServiceContext kit.golaxy.org/golaxy/runtime.getServiceContext
func getServiceContext(ctxResolver ec.ContextResolver) Context
