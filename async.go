package core

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/util/generic"
	"time"
	_ "unsafe"
)

//go:linkname getRuntimeContext git.golaxy.org/core/runtime.getRuntimeContext
func getRuntimeContext(ctxProvider concurrent.CurrentContextProvider) runtime.Context

// Async 异步执行代码，有返回值
func Async(ctxProvider runtime.CurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, runtime.Ret], va ...any) runtime.AsyncRet {
	ctx := getRuntimeContext(ctxProvider)
	return ctx.Call(func(va ...any) runtime.Ret {
		ctx := va[0].(runtime.Context)
		fun := va[1].(generic.FuncVar1[runtime.Context, any, runtime.Ret])
		funVa := va[2].([]any)
		return fun.Exec(ctx, funVa...)
	}, ctx, fun, va)
}

// AsyncVoid 异步执行代码，无返回值
func AsyncVoid(ctxProvider runtime.CurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], va ...any) runtime.AsyncRet {
	ctx := getRuntimeContext(ctxProvider)
	return ctx.CallVoid(func(va ...any) {
		ctx := va[0].(runtime.Context)
		fun := va[1].(generic.ActionVar1[runtime.Context, any])
		funVa := va[2].([]any)
		fun.Exec(ctx, funVa...)
	}, ctx, fun, va)
}

// Go 使用新线程执行代码，有返回值
func Go(ctx context.Context, fun generic.FuncVar1[context.Context, any, runtime.Ret], va ...any) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(fun generic.FuncVar1[context.Context, any, runtime.Ret], ctx context.Context, va []any, asyncRet chan runtime.Ret) {
		ret, panicErr := fun.Invoke(ctx, va...)
		if panicErr != nil {
			ret.Error = panicErr
		}
		asyncRet <- ret
		close(asyncRet)
	}(fun, ctx, va, asyncRet)

	return asyncRet
}

// GoVoid 使用新线程执行代码，无返回值
func GoVoid(ctx context.Context, fun generic.ActionVar1[context.Context, any], va ...any) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(fun generic.ActionVar1[context.Context, any], ctx context.Context, va []any, asyncRet chan runtime.Ret) {
		asyncRet <- runtime.MakeRet(nil, fun.Invoke(ctx, va...))
		close(asyncRet)
	}(fun, ctx, va, asyncRet)

	return asyncRet
}

// TimeAfter 定时器，指定时长
func TimeAfter(ctx context.Context, dur time.Duration) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(ctx context.Context, dur time.Duration, asyncRet chan runtime.Ret) {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- runtime.MakeRet(nil, nil)
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}(ctx, dur, asyncRet)

	return asyncRet
}

// TimeAt 定时器，指定时间点
func TimeAt(ctx context.Context, at time.Time) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(ctx context.Context, at time.Time, asyncRet chan runtime.Ret) {
		timer := time.NewTimer(time.Until(at))
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- runtime.MakeRet(nil, nil)
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}(ctx, at, asyncRet)

	return asyncRet
}

// TimeTick 心跳器
func TimeTick(ctx context.Context, dur time.Duration) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(ctx context.Context, dur time.Duration, asyncRet chan runtime.Ret) {
		tick := time.NewTicker(dur)
		defer tick.Stop()

	loop:
		for {
			select {
			case <-tick.C:
				select {
				case asyncRet <- runtime.MakeRet(nil, nil):
				case <-ctx.Done():
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}

		close(asyncRet)
	}(ctx, dur, asyncRet)

	return asyncRet
}

// ReadChan 读取channel
func ReadChan[T any](ctx context.Context, ch <-chan T) runtime.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	if ch == nil {
		panic(fmt.Errorf("%w: %w: ch is nil", ErrGolaxy, ErrArgs))
	}

	asyncRet := concurrent.MakeAsyncRet()

	go func(ctx context.Context, ch <-chan T, asyncRet chan runtime.Ret) {
	loop:
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					break loop
				}
				select {
				case asyncRet <- runtime.MakeRet(v, nil):
				case <-ctx.Done():
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}
		close(asyncRet)
	}(ctx, ch, asyncRet)

	return asyncRet
}
