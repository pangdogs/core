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

// CallAsync 将 fun 提交到 provider 所属运行时执行，并返回承载调用结果的 Future。
func CallAsync(provider corectx.ConcurrentContextProvider, fun generic.FuncVar1[runtime.Context, any, async.Result], args ...any) async.Future {
	return runtime.Concurrent(provider).CallAsync(fun, args...)
}

// CallVoidAsync 将无返回值的 fun 提交到 provider 所属运行时执行，并返回完成信号。
func CallVoidAsync(provider corectx.ConcurrentContextProvider, fun generic.ActionVar1[runtime.Context, any], args ...any) async.Future {
	return runtime.Concurrent(provider).CallVoidAsync(fun, args...)
}

// GoAsync 在新的 goroutine 中执行 fun，并返回承载调用结果的 Future。
// nil 上下文会替换为 context.Background；panic 会转换为 Result.Error。
func GoAsync(ctx context.Context, fun generic.FuncVar1[context.Context, any, async.Result], args ...any) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	future := async.NewFutureChan()

	go func() {
		ret, panicErr := fun.SafeCall(ctx, args...)
		if panicErr != nil {
			ret.Error = panicErr
		}
		async.Return(future, ret)
	}()

	return future.Out()
}

// GoVoidAsync 在新的 goroutine 中执行无返回值的 fun，并返回完成信号。
// nil 上下文会替换为 context.Background；panic 会转换为 Result.Error。
func GoVoidAsync(ctx context.Context, fun generic.ActionVar1[context.Context, any], args ...any) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	future := async.NewFutureChan()

	go func() {
		async.Return(future, async.NewResult(nil, fun.SafeCall(ctx, args...)))
	}()

	return future.Out()
}

// TimeAfterAsync 在 dur 后产出一次当前时间并结束。
// 上下文取消时不再产出结果并直接结束。
func TimeAfterAsync(ctx context.Context, dur time.Duration) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	future := async.NewFutureChan()

	go func() {
		timer := time.NewTimer(dur)
		defer timer.Stop()

		select {
		case t := <-timer.C:
			async.YieldReturn(ctx, future, async.NewResult(t, nil))
		case <-ctx.Done():
		}

		async.YieldBreak(future)
	}()

	return future.Out()
}

// TimeAtAsync 在 at 到达时产出一次当前时间并结束。
// 上下文取消时不再产出结果并直接结束。
func TimeAtAsync(ctx context.Context, at time.Time) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	future := async.NewFutureChan()

	go func() {
		timer := time.NewTimer(time.Until(at))
		defer timer.Stop()

		select {
		case t := <-timer.C:
			async.YieldReturn(ctx, future, async.NewResult(t, nil))
		case <-ctx.Done():
		}

		async.YieldBreak(future)
	}()

	return future.Out()
}

// TimeTickAsync 按 dur 周期持续产出当前时间，直到上下文取消。
// dur 小于等于 0 时 time.NewTicker 会 panic；消费端停止读取时，生产协程可能阻塞到
// 上下文取消。
func TimeTickAsync(ctx context.Context, dur time.Duration) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	future := async.NewFutureChan()

	go func() {
		tick := time.NewTicker(dur)
		defer tick.Stop()

	loop:
		for {
			select {
			case t := <-tick.C:
				if !async.YieldReturn(ctx, future, async.NewResult(t, nil)) {
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}

		async.YieldBreak(future)
	}()

	return future.Out()
}

// ReadChanAsync 将 ch 中的每个值转换为 Future 的连续产出。
// ch 关闭或上下文取消时结束；ch 为 nil 时会 panic。消费端停止读取时，生产协程可能
// 阻塞到上下文取消。
func ReadChanAsync[T any](ctx context.Context, ch <-chan T) async.Future {
	if ctx == nil {
		ctx = context.Background()
	}

	if ch == nil {
		exception.Panicf("%w: %w: ch is nil", ErrCore, ErrArgs)
	}

	future := async.NewFutureChan(cap(ch))

	go func() {
	loop:
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					break loop
				}
				if !async.YieldReturn(ctx, future, async.NewResult(v, nil)) {
					break loop
				}
			case <-ctx.Done():
				break loop
			}
		}
		async.YieldBreak(future)
	}()

	return future.Out()
}
