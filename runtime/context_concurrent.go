package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentContext 多线程安全的运行时上下文接口
type ConcurrentContext interface {
	gctx.ConcurrentContextProvider
	gctx.Context
	async.Caller
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
}

// Concurrent 获取多线程安全的运行时上下文
func Concurrent(provider gctx.ConcurrentContextProvider) ConcurrentContext {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}

func getServiceContext(provider gctx.ConcurrentContextProvider) service.Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	ctx := iface.Cache2Iface[Context](provider.GetConcurrentContext())
	if ctx == nil {
		return nil
	}
	return ctx.getServiceCtx()
}

func getCaller(provider gctx.ConcurrentContextProvider) async.Caller {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}
