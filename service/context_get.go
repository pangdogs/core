package service

import (
	"github.com/pangdogs/galaxy/ec"
	_ "unsafe"
)

// Get 获取服务上下文
func Get(ctxHolder ec.ContextHolder) Context {
	return getServiceContext(ctxHolder)
}

//go:linkname getServiceContext github.com/pangdogs/galaxy/runtime.getServiceContext
func getServiceContext(ctxHolder ec.ContextHolder) Context
