package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/iface"
)

type (
	CurrentContextProvider    = concurrent.CurrentContextProvider    // 当前上下文提供者
	ConcurrentContextProvider = concurrent.ConcurrentContextProvider // 多线程安全的上下文提供者
)

// Current 获取当前运行时上下文
func Current(provider CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
}

// Concurrent 获取多线程安全的运行时上下文
func Concurrent(provider ConcurrentContextProvider) ConcurrentContext {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}

func getCaller(provider ConcurrentContextProvider) concurrent.Caller {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetConcurrentContext())
}

func getRuntimeContext(provider CurrentContextProvider) Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](provider.GetCurrentContext())
}

func getServiceContext(provider ConcurrentContextProvider) service.Context {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrContext, exception.ErrArgs))
	}
	ctx := iface.Cache2Iface[Context](provider.GetConcurrentContext())
	if ctx == nil {
		return nil
	}
	return ctx.getServiceCtx()
}
