package runtime

import (
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
)

// Get 获取运行时上下文
func Get(ctxResolver ec.ContextResolver) Context {
	if ctxResolver == nil {
		panic("nil ctxResolver")
	}

	ctx := ec.UnsafeContextResolver(ctxResolver).GetContext()
	if ctx == util.NilIfaceCache {
		panic("nil context")
	}

	return util.Cache2Iface[Context](ctx)
}

// TryGet 尝试获取运行时上下文
func TryGet(ctxResolver ec.ContextResolver) (Context, bool) {
	if ctxResolver == nil {
		return nil, false
	}

	ctx := ec.UnsafeContextResolver(ctxResolver).GetContext()
	if ctx == util.NilIfaceCache {
		return nil, false
	}

	return util.Cache2Iface[Context](ctx), true
}

func getServiceContext(ctxResolver ec.ContextResolver) service.Context {
	return Get(ctxResolver).GetServiceCtx()
}

func tryGetServiceContext(ctxResolver ec.ContextResolver) (service.Context, bool) {
	runtimeCtx, ok := TryGet(ctxResolver)
	if !ok {
		return nil, false
	}
	return runtimeCtx.GetServiceCtx(), true
}
