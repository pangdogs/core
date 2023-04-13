package runtime

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
)

// Get 获取运行时上下文
func Get(ctxResolver ec.ContextResolver) Context {
	if ctxResolver == nil {
		panic("nil ctxResolver")
	}

	return util.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getServiceContext(ctxResolver ec.ContextResolver) service.Context {
	return Get(ctxResolver).GetServiceCtx()
}
