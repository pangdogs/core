package runtime

import (
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/service"
	"github.com/galaxy-kit/galaxy-go/util"
)

// Get 获取运行时上下文
func Get(ctxHolder ec.ContextHolder) Context {
	if ctxHolder == nil {
		panic("nil ctxHolder")
	}

	ctx := ec.UnsafeContextHolder(ctxHolder).GetContext()
	if ctx == util.NilIfaceCache {
		panic("nil context")
	}

	return util.Cache2Iface[Context](ctx)
}

// TryGet 尝试获取运行时上下文
func TryGet(ctxHolder ec.ContextHolder) (Context, bool) {
	if ctxHolder == nil {
		return nil, false
	}

	ctx := ec.UnsafeContextHolder(ctxHolder).GetContext()
	if ctx == util.NilIfaceCache {
		return nil, false
	}

	return util.Cache2Iface[Context](ctx), true
}

func getServiceContext(ctxHolder ec.ContextHolder) service.Context {
	return Get(ctxHolder).GetServiceCtx()
}

func tryGetServiceContext(ctxHolder ec.ContextHolder) (service.Context, bool) {
	runtimeCtx, ok := TryGet(ctxHolder)
	if !ok {
		return nil, false
	}
	return runtimeCtx.GetServiceCtx(), true
}
