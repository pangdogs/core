package service

import (
	"kit.golaxy.org/golaxy/internal/concurrent"
	_ "unsafe"
)

type (
	ContextResolver = concurrent.ContextResolver // 上下文获取器
)

//go:linkname getServiceContext kit.golaxy.org/golaxy/runtime.getServiceContext
func getServiceContext(ctxResolver concurrent.ContextResolver) Context

// Current 获取服务上下文
func Current(ctxResolver ContextResolver) Context {
	return getServiceContext(ctxResolver)
}
