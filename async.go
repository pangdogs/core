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
)

// Async 异步执行代码，有返回值
func Async(provider ictx.ConcurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, async.Ret], args ...any) async.AsyncRet {
	ctx := runtime.UnsafeConcurrentContext(runtime.Concurrent(provider)).GetContext()
	return ctx.Call(func(...any) async.Ret { return fun.Exec(ctx, args...) })
}

// AsyncVoid 异步执行代码，无返回值
func AsyncVoid(provider ictx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], args ...any) async.AsyncRet {
	ctx := runtime.UnsafeConcurrentContext(runtime.Concurrent(provider)).GetContext()
	return ctx.CallVoid(func(...any) { fun.Exec(ctx, args...) })
}

// Go 使用新线程执行代码，有返回值
func Go(ctx context.Context, fun generic.FuncVar1[context.Context, any, async.Ret], args ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		ret, panicErr := fun.Invoke(ctx, args...)
		if panicErr != nil {
			ret.Error = panicErr
		}
		asyncRet <- ret
		close(asyncRet)
	}()

	return asyncRet
}

// GoVoid 使用新线程执行代码，无返回值
func GoVoid(ctx context.Context, fun generic.ActionVar1[context.Context, any], args ...any) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		asyncRet <- async.MakeRet(nil, fun.Invoke(ctx, args...))
		close(asyncRet)
	}()

	return asyncRet
}

// TimeAfter 定时器，指定时长
func TimeAfter(ctx context.Context, dur time.Duration) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- async.VoidRet
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}()

	return asyncRet
}

// TimeAt 定时器，指定时间点
func TimeAt(ctx context.Context, at time.Time) async.AsyncRet {
	if ctx == nil {
		ctx = context.Background()
	}

	asyncRet := async.MakeAsyncRet()

	go func() {
		timer := time.NewTimer(time.Until(at))
		defer timer.Stop()

		select {
		case <-timer.C:
			asyncRet <- async.VoidRet
		case <-ctx.Done():
			break
		}

		close(asyncRet)
	}()

	return asyncRet
}

// TimeTick 心跳器
func TimeTick(ctx context.Context, dur time.Duration) async.AsyncRet {
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
	}()

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

	go func() {
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
	}()

	return asyncRet
}
