package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/iface"
)

type (
	CurrentContextResolver    = concurrent.CurrentContextResolver    // 当前上下文获取器
	ConcurrentContextResolver = concurrent.ConcurrentContextResolver // 多线程安全的上下文获取器
)

// Current 获取当前运行时上下文
func Current(ctxResolver CurrentContextResolver) Context {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveCurrentContext())
}

// Concurrent 获取多线程安全的运行时上下文
func Concurrent(ctxResolver ConcurrentContextResolver) ConcurrentContext {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveConcurrentContext())
}

func getCaller(ctxResolver concurrent.ContextResolver) concurrent.Caller {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getRuntimeContext(ctxResolver concurrent.ContextResolver) Context {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveContext())
}

func getServiceContext(ctxResolver concurrent.ContextResolver) service.Context {
	if ctxResolver == nil {
		panic(fmt.Errorf("%w: %w: ctxResolver is nil", ErrContext, exception.ErrArgs))
	}
	return iface.Cache2Iface[Context](ctxResolver.ResolveContext()).getServiceCtx()
}
