/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"time"
	_ "unsafe"
)

//go:linkname getCaller git.golaxy.org/core/runtime.getCaller
func getCaller(provider ictx.ConcurrentContextProvider) async.Caller

// Async 异步执行代码，有返回值
func Async(provider ictx.ConcurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, async.Ret], va ...any) async.AsyncRet {
	ctx := getCaller(provider)
	return ctx.Call(func(va ...any) async.Ret {
		ctx := va[0].(runtime.Context)
		fun := va[1].(generic.FuncVar1[runtime.Context, any, async.Ret])
		funVa := va[2].([]any)
		return fun.Exec(ctx, funVa...)
	}, ctx, fun, va)
}

// AsyncVoid 异步执行代码，无返回值
func AsyncVoid(provider ictx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], va ...any) async.AsyncRet {
	ctx := getCaller(provider)
	return ctx.CallVoid(func(va ...any) {
		ctx := va[0].(runtime.Context)
		fun := va[1].(generic.ActionVar1[runtime.Context, any])
		funVa := va[2].([]any)
		fun.Exec(ctx, funVa...)
	}, ctx, fun, va)
}

// Go 使用新线程执行代码，有返回值
func Go(ctx context.Context, fun generic.FuncVar1[context.Context, any, async.Ret], va ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func(fun generic.FuncVar1[context.Context, any, async.Ret], ctx context.Context, va []any, asyncRet chan async.Ret) {
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
func GoVoid(ctx context.Context, fun generic.ActionVar1[context.Context, any], va ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func(fun generic.ActionVar1[context.Context, any], ctx context.Context, va []any, asyncRet chan async.Ret) {
		asyncRet <- async.MakeRet(nil, fun.Invoke(ctx, va...))
		close(asyncRet)
	}(fun, ctx, va, asyncRet)

	return asyncRet
}

// TimeAfter 定时器，指定时长
func TimeAfter(ctx context.Context, dur time.Duration) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func(ctx context.Context, dur time.Duration, asyncRet chan async.Ret) {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- async.VoidRet
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}(ctx, dur, asyncRet)

	return asyncRet
}

// TimeAt 定时器，指定时间点
func TimeAt(ctx context.Context, at time.Time) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func(ctx context.Context, at time.Time, asyncRet chan async.Ret) {
		timer := time.NewTimer(time.Until(at))
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- async.VoidRet
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}(ctx, at, asyncRet)

	return asyncRet
}

// TimeTick 心跳器
func TimeTick(ctx context.Context, dur time.Duration) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func(ctx context.Context, dur time.Duration, asyncRet chan async.Ret) {
		tick := time.NewTicker(dur)
		defer tick.Stop()

	loop:
		for {
			select {
			case <-tick.C:
				select {
				case asyncRet <- async.VoidRet:
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
func ReadChan[T any](ctx context.Context, ch <-chan T) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	if ch == nil {
		panic(fmt.Errorf("%w: %w: ch is nil", ErrCore, ErrArgs))
	}

	asyncRet := async.MakeAsyncRet()

	go func(ctx context.Context, ch <-chan T, asyncRet chan async.Ret) {
	loop:
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					break loop
				}
				select {
				case asyncRet <- async.MakeRet(v, nil):
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
