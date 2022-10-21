package runtime

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
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

func getServiceContext(ctxHolder ec.ContextHolder) service.Context {
	return Get(ctxHolder).GetServiceCtx()
}
