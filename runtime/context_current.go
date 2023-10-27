package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/iface"
)

// Current 获取当前运行时上下文
func Current(ctxResolver ContextResolver) Context {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, errors.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getServiceContext(ctxResolver ContextResolver) service.Context {
	return Current(ctxResolver).getServiceCtx()
}
