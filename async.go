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
	"time"

	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// Submit 将有返回值函数投递到 provider 所属 Runtime，并返回执行结果 Future。
func Submit(provider corectx.ConcurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, async.Result], args ...any) async.Future {
	return runtime.Concurrent(provider).Submit(fun, args...)
}

// SubmitDelegate 将有返回值委托投递到 provider 所属 Runtime。
func SubmitDelegate(provider corectx.ConcurrentContextProvider, fun generic.DelegateVar1[runtime.Context, any, async.Result], args ...any) async.Future {
	return runtime.Concurrent(provider).SubmitDelegate(fun, args...)
}

// SubmitVoid 将无业务返回值函数投递到 provider 所属 Runtime，并返回可等待错误的 Future。
func SubmitVoid(provider corectx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], args ...any) async.Future {
	return runtime.Concurrent(provider).SubmitVoid(fun, args...)
}

// SubmitDelegateVoid 将无业务返回值委托投递到 provider 所属 Runtime。
func SubmitDelegateVoid(provider corectx.ConcurrentContextProvider, fun generic.DelegateVoidVar1[runtime.Context, any], args ...any) async.Future {
	return runtime.Concurrent(provider).SubmitDelegateVoid(fun, args...)
}

// Post 将无返回值函数投递到 provider 所属 Runtime，不创建 Future。
// 仅返回队列关闭、容量不足等同步入队错误。
func Post(provider corectx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], args ...any) error {
	return runtime.Concurrent(provider).Post(fun, args...)
}

// PostDelegate 将无返回值委托投递到 provider 所属 Runtime，不创建 Future。
func PostDelegate(provider corectx.ConcurrentContextProvider, fun generic.DelegateVoidVar1[runtime.Context, any], args ...any) error {
	return runtime.Concurrent(provider).PostDelegate(fun, args...)
}

// Spawn 在 scope 中启动后台 goroutine。fun 不得直接访问 Runtime 局部状态。
func Spawn(scope *async.Scope, fun generic.FuncVar1[context.Context, any, async.Result], args ...any) async.Future {
	return async.Spawn(scope, func(ctx context.Context) async.Result {
		return fun.UnsafeCall(ctx, args...)
	})
}

// SpawnVoid 在 scope 中启动无业务返回值的后台 goroutine。
func SpawnVoid(scope *async.Scope, fun generic.ActionVar1[context.Context, any], args ...any) async.Future {
	return async.SpawnVoid(scope, func(ctx context.Context) {
		fun.UnsafeCall(ctx, args...)
	})
}

// After 在 dur 后以当前时间完成 Future；ctx 取消时以 ctx.Err 完成。
func After(ctx context.Context, dur time.Duration) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}
	if dur < 0 {
		dur = 0
	}
	promise, future := async.NewPromise()
	timer := time.AfterFunc(dur, func() {
		promise.Resolve(async.NewResult(time.Now(), nil))
	})
	stopContext := context.AfterFunc(ctx, func() {
		promise.Resolve(async.NewResult(nil, ctx.Err()))
	})
	future.OnComplete(func(async.Result) {
		timer.Stop()
		stopContext()
	})
	return future
}

// At 在指定时间以当前时间完成 Future；ctx 取消时以 ctx.Err 完成。
func At(ctx context.Context, at time.Time) async.Future {
	return After(ctx, time.Until(at))
}

// Every 按 dur 周期持续产出当前时间，直到 ctx 取消。
func Every(ctx context.Context, dur time.Duration) async.Stream {
	if ctx == nil {
		ctx = context.Background()
	}
	if dur <= 0 {
		exception.Panicf("%w: %w: duration must be positive", ErrCore, ErrArgs)
	}
	emitter, stream := async.NewStream()
	go func() {
		defer emitter.Close()
		ticker := time.NewTicker(dur)
		defer ticker.Stop()
		for {
			select {
			case now := <-ticker.C:
				if !emitter.Emit(ctx, async.NewResult(now, nil)) {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return stream
}

// FromChan 将 ch 中的值转换为 Stream，直到 ch 关闭或 ctx 取消。
func FromChan[T any](ctx context.Context, ch <-chan T) async.Stream {
	if ctx == nil {
		ctx = context.Background()
	}
	if ch == nil {
		exception.Panicf("%w: %w: ch is nil", ErrCore, ErrArgs)
	}
	emitter, stream := async.NewStream(cap(ch))
	go func() {
		defer emitter.Close()
		for {
			select {
			case value, ok := <-ch:
				if !ok || !emitter.Emit(ctx, async.NewResult(value, nil)) {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return stream
}
