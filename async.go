package golaxy

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/runtime"
	"sync/atomic"
	"time"
)

// Async 异步执行代码，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func Async(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() runtime.Ret {
		return segment(ctx)
	})
}

// AsyncVoid 异步执行代码，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncVoid(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context)) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() runtime.Ret {
		segment(ctx)
		return runtime.NewRet(nil, nil)
	})
}

// Await 等待异步结果（async ret）返回，并继续运行后续逻辑
func Await(ctxResolver ec.ContextResolver, asyncRet runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if asyncRet == nil {
		return
	}

	go func() {
		defer func() {
			recover()
		}()

		ret, ok := <-asyncRet
		if !ok {
			ret = runtime.NewRet(errors.New("asyncRet closed"), nil)
		}

		ctx.AsyncCallNoRet(func() {
			asyncWait(ctx, ret)
		})
	}()
}

// AwaitAny 等待任意一个异步结果（async ret）返回，并继续运行后续逻辑
func AwaitAny(ctxResolver ec.ContextResolver, asyncRets []runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if len(asyncRets) <= 0 {
		return
	}

	var b atomic.Bool

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func() {
			defer func() {
				recover()
			}()

			ret, ok := <-asyncRet
			if !ok {
				return
			}

			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			ctx.AsyncCallNoRet(func() {
				asyncWait(ctx, ret)
			})
		}()
	}
}

// AwaitAll 等待所有异步结果（async ret）返回，并继续运行后续逻辑
func AwaitAll(ctxResolver ec.ContextResolver, asyncRets []runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if len(asyncRets) <= 0 {
		return
	}

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func() {
			defer func() {
				recover()
			}()

			ret, ok := <-asyncRet
			if !ok {
				ret = runtime.NewRet(errors.New("asyncRet closed"), nil)
			}

			ctx.AsyncCallNoRet(func() {
				asyncWait(ctx, ret)
			})
		}()
	}
}

// AwaitTimeAfterFunc 等待一段时间，并继续运行后续逻辑
func AwaitTimeAfterFunc(ctxResolver ec.ContextResolver, dur time.Duration, segment func(ctx runtime.Context)) {
	ctx := runtime.Get(ctxResolver)
	time.AfterFunc(dur, func() { AsyncVoid(ctx, segment) })
}
