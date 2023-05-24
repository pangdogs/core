package golaxy

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/runtime"
	"sync/atomic"
	"time"
)

// Async 异步执行代码，最多返回一次，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func Async(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() runtime.Ret {
		return segment(ctx)
	})
}

// AsyncVoid 异步执行代码，最多返回一次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
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

// AsyncGo 使用新协程异步执行代码，最多返回一次，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncGo(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				asyncRet <- runtime.NewRet(err, nil)
			}
			close(asyncRet)
		}()
		asyncRet <- segment(ctx)
	}()

	return asyncRet
}

// AsyncGoVoid 使用新协程异步执行代码，最多返回一次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncGoVoid(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context)) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				asyncRet <- runtime.NewRet(err, nil)
			}
			close(asyncRet)
		}()
		segment(ctx)
		asyncRet <- runtime.NewRet(nil, nil)
	}()

	return asyncRet
}

// AsyncTimeAfter 异步等待一段时间后返回，最多返回一次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncTimeAfter(ctx context.Context, dur time.Duration) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := make(chan runtime.Ret, 1)
	timer := time.NewTimer(dur)

	go func() {
		defer func() {
			recover()
			timer.Stop()
			close(asyncRet)
		}()

		select {
		case <-timer.C:
			asyncRet <- runtime.NewRet(nil, nil)
		case <-ctx.Done():
		}
	}()

	return asyncRet
}

// AsyncTimeTick 每隔一段时间后返回，返回多次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncTimeTick(ctx context.Context, dur time.Duration) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := make(chan runtime.Ret, 1)
	tick := time.NewTicker(dur)

	go func() {
		defer func() {
			recover()
			tick.Stop()
			close(asyncRet)
		}()

		for {
			select {
			case <-tick.C:
				select {
				case asyncRet <- runtime.NewRet(nil, nil):
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return asyncRet
}

// AsyncChanRet 异步等待chan返回，支持返回多次，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncChanRet[T any](ctx context.Context, ch <-chan T) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	if ch == nil {
		panic("nil ch")
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer func() {
			recover()
			close(asyncRet)
		}()

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case asyncRet <- runtime.NewRet(nil, v):
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return asyncRet
}

// Await 等待异步结果（async ret）返回，并继续运行后续逻辑
func Await(ctxResolver ec.ContextResolver, asyncRet runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if asyncRet == nil {
		return
	}

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	go func() {
		defer func() {
			recover()
		}()

		for ret := range asyncRet {
			ctx.AsyncCallNoRet(func() { asyncWait(ctx, ret) })
		}
	}()
}

// AwaitAny 等待任意一个异步结果（async ret）成功的一次返回，并继续运行后续逻辑
func AwaitAny(ctxResolver ec.ContextResolver, asyncRets []runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if len(asyncRets) <= 0 {
		return
	}

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	var b atomic.Bool
	waitCtx, cancel := context.WithCancel(ctx)

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func(asyncRet runtime.AsyncRet) {
			defer func() {
				recover()
			}()

			var ret runtime.Ret
			var ok bool

			select {
			case ret, ok = <-asyncRet:
				if !ok || !ret.OK() {
					return
				}
			case <-waitCtx.Done():
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			ctx.AsyncCallNoRet(func() { asyncWait(ctx, ret) })
		}(asyncRet)
	}
}

// AwaitAll 等待所有异步结果（async ret）返回，并继续运行后续逻辑
func AwaitAll(ctxResolver ec.ContextResolver, asyncRets []runtime.AsyncRet, asyncWait func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Get(ctxResolver)

	if len(asyncRets) <= 0 {
		return
	}

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func(asyncRet runtime.AsyncRet) {
			defer func() {
				recover()
			}()

			for ret := range asyncRet {
				ctx.AsyncCallNoRet(func() { asyncWait(ctx, ret) })
			}
		}(asyncRet)
	}
}
