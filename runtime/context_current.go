package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// Current 获取当前运行时上下文
func Current(provider gctx.CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
}

// Concurrent 获取多线程安全的运行时上下文
func Concurrent(provider gctx.ConcurrentContextProvider) ConcurrentContext {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}

func getCaller(provider gctx.ConcurrentContextProvider) async.Caller {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}

func getRuntimeContext(provider gctx.CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
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
