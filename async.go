package golaxy

import (
	"context"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"kit.golaxy.org/golaxy/internal/errors"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/types"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrAsyncAwait            = fmt.Errorf("%w: async/await", errors.ErrGolaxy)
	ErrAllOfAsyncRetFailures = fmt.Errorf("%w: all of async result failures", ErrAsyncAwait)
	ErrAsyncRetClosed        = fmt.Errorf("%w: async result closed", ErrAsyncAwait)
)

// Async 异步执行代码，最多返回一次，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func Async(ctxResolver runtime.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Current(ctxResolver)

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	return ctx.AsyncCall(func() runtime.Ret {
		return segment(ctx)
	})
}

// AsyncVoid 异步执行代码，最多返回一次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncVoid(ctxResolver runtime.ContextResolver, segment func(ctx runtime.Context)) runtime.AsyncRet {
	ctx := runtime.Current(ctxResolver)

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	return ctx.AsyncCall(func() runtime.Ret {
		segment(ctx)
		return runtime.NewRet(nil, nil)
	})
}

// AsyncGo 使用新协程异步执行代码，最多返回一次，有返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncGo(ctxResolver runtime.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Current(ctxResolver)

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer func() {
			if panicErr := types.Panic2Err(recover()); panicErr != nil {
				asyncRet <- runtime.NewRet(nil, fmt.Errorf("%w: %w: %w", errors.ErrPanicked, ErrAsyncAwait, panicErr))
			}
			close(asyncRet)
		}()
		asyncRet <- segment(ctx)
	}()

	return asyncRet
}

// AsyncGoVoid 使用新协程异步执行代码，最多返回一次，无返回值，返回的异步结果（async ret）可以给Await()等待并继续后续逻辑运行
func AsyncGoVoid(ctxResolver runtime.ContextResolver, segment func(ctx runtime.Context)) runtime.AsyncRet {
	ctx := runtime.Current(ctxResolver)

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer func() {
			if panicErr := types.Panic2Err(recover()); panicErr != nil {
				asyncRet <- runtime.NewRet(nil, fmt.Errorf("%w: %w: %w", errors.ErrPanicked, ErrAsyncAwait, panicErr))
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

	go func() {
		timer := time.NewTimer(dur)

		defer func() {
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

	go func() {
		tick := time.NewTicker(dur)

		defer func() {
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
		panic(fmt.Errorf("%w: %w: ch is nil", ErrAsyncAwait, ErrArgs))
	}

	asyncRet := make(chan runtime.Ret, 1)

	go func() {
		defer close(asyncRet)

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case asyncRet <- runtime.NewRet(v, nil):
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

// Await 异步等待异步结果（async ret）返回，并继续运行后续逻辑
func Await(ctxResolver runtime.ContextResolver, asyncRet runtime.AsyncRet, segment func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Current(ctxResolver)

	if asyncRet == nil {
		panic(fmt.Errorf("%w: %w: asyncRet is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	go func() {
		for ret := range asyncRet {
			ctx.AsyncCallVoid(func() { segment(ctx, ret) })
		}
		ctx.AsyncCallVoid(func() { segment(ctx, runtime.NewRet(nil, ErrAsyncRetClosed)) })
	}()
}

// AwaitAny 异步等待任意一个异步结果（async ret）成功的一次返回，并继续运行后续逻辑
func AwaitAny(ctxResolver runtime.ContextResolver, asyncRets []runtime.AsyncRet, segment func(ctx runtime.Context, ret runtime.Ret)) {
	ctx := runtime.Current(ctxResolver)

	if len(asyncRets) <= 0 {
		panic(fmt.Errorf("%w: %w: asyncRets is empty", ErrAsyncAwait, errors.ErrArgs))
	}

	if pie.Contains(asyncRets, nil) {
		panic(fmt.Errorf("%w: %w: asyncRets contain nil elements", ErrAsyncAwait, errors.ErrArgs))
	}

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	var wg sync.WaitGroup
	var b atomic.Bool
	waitCtx, cancel := context.WithCancel(ctx)

	for _, asyncRet := range asyncRets {
		wg.Add(1)
		go func(asyncRet runtime.AsyncRet) {
			defer wg.Done()

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

			ctx.AsyncCallVoid(func() { segment(ctx, ret) })
		}(asyncRet)
	}

	go func() {
		wg.Wait()
		if !b.Load() {
			ctx.AsyncCallVoid(func() { segment(ctx, runtime.NewRet(nil, ErrAllOfAsyncRetFailures)) })
		}
	}()
}

// AwaitAll 异步等待所有异步结果（async ret）返回，并继续运行后续逻辑
func AwaitAll(ctxResolver runtime.ContextResolver, asyncRets []runtime.AsyncRet, segment func(ctx runtime.Context, rets []runtime.Ret)) {
	ctx := runtime.Current(ctxResolver)

	if len(asyncRets) <= 0 {
		panic(fmt.Errorf("%w: %w: asyncRets is empty", ErrAsyncAwait, errors.ErrArgs))
	}

	if pie.Contains(asyncRets, nil) {
		panic(fmt.Errorf("%w: %w: asyncRets contain nil elements", ErrAsyncAwait, errors.ErrArgs))
	}

	if segment == nil {
		panic(fmt.Errorf("%w: %w: segment is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	var wg sync.WaitGroup
	rets := make([]runtime.Ret, len(asyncRets))

	for i, asyncRet := range asyncRets {
		wg.Add(1)
		go func(asyncRet runtime.AsyncRet, ret *runtime.Ret) {
			defer wg.Done()

			r, ok := <-asyncRet
			if !ok {
				r.Error = ErrAsyncRetClosed
			}

			*ret = r
		}(asyncRet, &rets[i])
	}

	go func() {
		wg.Wait()
		ctx.AsyncCallVoid(func() { segment(ctx, rets) })
	}()
}

// Wait 同步等待异步结果（async ret）返回，并继续运行后续逻辑
func Wait(asyncRet runtime.AsyncRet) runtime.Ret {
	if asyncRet == nil {
		panic(fmt.Errorf("%w: %w: asyncRet is nil", ErrAsyncAwait, errors.ErrArgs))
	}

	ret, ok := <-asyncRet
	if !ok {
		ret.Error = ErrAsyncRetClosed
	}

	return ret
}

// WaitAny 同步等待任意一个异步结果（async ret）成功的一次返回，并继续运行后续逻辑
func WaitAny(asyncRets []runtime.AsyncRet) runtime.Ret {
	if len(asyncRets) <= 0 {
		panic(fmt.Errorf("%w: %w: asyncRets is empty", ErrAsyncAwait, errors.ErrArgs))
	}

	if pie.Contains(asyncRets, nil) {
		panic(fmt.Errorf("%w: %w: asyncRets contain nil elements", ErrAsyncAwait, errors.ErrArgs))
	}

	var wg sync.WaitGroup
	var b atomic.Bool
	waitCtx, cancel := context.WithCancel(context.Background())
	var ret runtime.Ret

	for _, asyncRet := range asyncRets {
		wg.Add(1)
		go func(asyncRet runtime.AsyncRet) {
			defer wg.Done()

			var r runtime.Ret
			var ok bool

			select {
			case r, ok = <-asyncRet:
				if !ok || !r.OK() {
					return
				}
			case <-waitCtx.Done():
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			ret = r
		}(asyncRet)
	}

	wg.Wait()

	if !b.Load() {
		return runtime.NewRet(nil, ErrAllOfAsyncRetFailures)
	}

	return ret
}

// WaitAll 同步等待所有异步结果（async ret）返回，并继续运行后续逻辑
func WaitAll(asyncRets []runtime.AsyncRet) []runtime.Ret {
	if len(asyncRets) <= 0 {
		panic(fmt.Errorf("%w: %w: asyncRets is empty", ErrAsyncAwait, errors.ErrArgs))
	}

	if pie.Contains(asyncRets, nil) {
		panic(fmt.Errorf("%w: %w: asyncRets contain nil elements", ErrAsyncAwait, errors.ErrArgs))
	}

	var wg sync.WaitGroup
	rets := make([]runtime.Ret, len(asyncRets))

	for i, asyncRet := range asyncRets {
		wg.Add(1)
		go func(asyncRet runtime.AsyncRet, ret *runtime.Ret) {
			defer wg.Done()

			r, ok := <-asyncRet
			if !ok {
				r.Error = ErrAsyncRetClosed
			}

			*ret = r
		}(asyncRet, &rets[i])
	}

	wg.Wait()

	return rets
}
