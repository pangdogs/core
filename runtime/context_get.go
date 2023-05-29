package runtime

import (
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
)

// Get 获取运行时上下文
func Get(ctxResolver ContextResolver) Context {
	if ctxResolver == nil {
		panic("nil ctxResolver")
	}

	return util.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getServiceContext(ctxResolver ContextResolver) service.Context {
	return Get(ctxResolver).getServiceCtx()
}
