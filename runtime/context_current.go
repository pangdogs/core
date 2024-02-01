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
func Current(ctxProvider CurrentContextProvider) Context {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetCurrentContext())
}

// Concurrent 获取多线程安全的运行时上下文
func Concurrent(ctxProvider ConcurrentContextProvider) ConcurrentContext {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetConcurrentContext())
}

func getCaller(ctxProvider ConcurrentContextProvider) concurrent.Caller {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetConcurrentContext())
}

func getRuntimeContext(ctxProvider CurrentContextProvider) Context {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetCurrentContext())
}

func getServiceContext(ctxProvider ConcurrentContextProvider) service.Context {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	ctx := iface.Cache2Iface[Context](ctxProvider.GetConcurrentContext())
	if ctx == nil {
		return nil
	}
	return ctx.getServiceCtx()
}
