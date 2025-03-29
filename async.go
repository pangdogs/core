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
	"git.golaxy.org/core/ec/ictx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"time"
)

// CallAsync 异步执行代码，有返回值
func CallAsync(provider ictx.ConcurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, async.Ret], args ...any) async.AsyncRet {
	ctx := runtime.UnsafeConcurrentContext(runtime.Concurrent(provider)).GetContext()
	return ctx.CallAsync(func(...any) async.Ret { return fun.UnsafeCall(ctx, args...) })
}

// CallVoidAsync 异步执行代码，无返回值
func CallVoidAsync(provider ictx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], args ...any) async.AsyncRet {
	ctx := runtime.UnsafeConcurrentContext(runtime.Concurrent(provider)).GetContext()
	return ctx.CallVoidAsync(func(...any) { fun.UnsafeCall(ctx, args...) })
}

// GoAsync 使用新线程执行代码，有返回值
func GoAsync(ctx context.Context, fun generic.FuncVar1[context.Context, any, async.Ret], args ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		ret, panicErr := fun.SafeCall(ctx, args...)
		if panicErr != nil {
			ret.Error = panicErr
		}
		async.Return(asyncRet, ret)
	}()

	return asyncRet
}

// GoVoidAsync 使用新线程执行代码，无返回值
func GoVoidAsync(ctx context.Context, fun generic.ActionVar1[context.Context, any], args ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		async.Return(asyncRet, async.MakeRet(nil, fun.SafeCall(ctx, args...)))
	}()

	return asyncRet
}

// TimeAfterAsync 定时器，指定时长
func TimeAfterAsync(ctx context.Context, dur time.Duration) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case <-timer.C:
			async.YieldReturn(ctx, asyncRet, async.VoidRet)
		case <-ctx.Done():
			break
		}

		async.YieldBreak(asyncRet)
	}()

	return asyncRet
}

// TimeAtAsync 定时器，指定时间点
func TimeAtAsync(ctx context.Context, at time.Time) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		timer := time.NewTimer(time.Until(at))
		defer timer.Stop()

		select {
		case <-timer.C:
			async.YieldReturn(ctx, asyncRet, async.VoidRet)
		case <-ctx.Done():
			break
		}

		async.YieldBreak(asyncRet)
	}()

	return asyncRet
}

// TimeTickAsync 心跳器
func TimeTickAsync(ctx context.Context, dur time.Duration) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		tick := time.NewTicker(dur)
		defer tick.Stop()

	loop:
		for {
			select {
			case <-tick.C:
				if !async.YieldReturn(ctx, asyncRet, async.VoidRet) {
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}

		async.YieldBreak(asyncRet)
	}()

	return asyncRet
}

// ReadChanAsync 读取channel转换为AsyncRet
func ReadChanAsync(ctx context.Context, ch <-chan any) async.AsyncRet {
	return ReadChanAsyncT[any](ctx, ch)
}

// ReadChanAsyncT 读取channel转换为AsyncRet
func ReadChanAsyncT[T any](ctx context.Context, ch <-chan T) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	if ch == nil {
		exception.Panicf("%w: %w: ch is nil", ErrCore, ErrArgs)
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
	loop:
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					break loop
				}
				if !async.YieldReturn(ctx, asyncRet, async.MakeRet(v, nil)) {
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}
		async.YieldBreak(asyncRet)
	}()

	return asyncRet
}
