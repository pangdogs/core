package runtime

import (
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
)

// Current 获取当前运行时上下文
func Current(ctxResolver ContextResolver) Context {
	if ctxResolver == nil {
		panic("nil ctxResolver")
	}

	return util.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getServiceContext(ctxResolver ContextResolver) service.Context {
	return Current(ctxResolver).getServiceCtx()
}
