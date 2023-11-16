package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/iface"
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

func getCaller(ctxProvider concurrent.ContextProvider) concurrent.Caller {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetContext())
}

func getRuntimeContext(ctxProvider concurrent.ContextProvider) Context {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetContext())
}

func getServiceContext(ctxProvider concurrent.ContextProvider) service.Context {
	if ctxProvider == nil {
		panic(fmt.Errorf("%w: %w: ctxProvider is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxProvider.GetContext()).getServiceCtx()
}
