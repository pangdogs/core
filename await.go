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
	"slices"
	"sync"
	"sync/atomic"

	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

var (
	ErrAllFuturesExceeded = fmt.Errorf("%w: all futures exceeded deadline: %w", ErrCore, context.DeadlineExceeded)
)

// Await 异步等待结果返回
func Await(provider corectx.CurrentContextProvider, futures ...async.Future) AwaitDirector {
	return AwaitDirector{
		rtCtx:   runtime.Current(provider),
		futures: slices.DeleteFunc(futures, func(future async.Future) bool { return future == nil }),
	}
}

// AwaitDirector 异步等待分发器
type AwaitDirector struct {
	rtCtx   runtime.Context
	futures []async.Future
}

func (ad AwaitDirector) waitContext() (context.Context, context.CancelFunc) {
	if len(ad.futures) > 1 {
		return context.WithCancel(ad.rtCtx)
	}
	return ad.rtCtx, nil
}

// Any 异步等待任意一个结果返回，有返回值
func (ad AwaitDirector) Any(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	ctx, cancel := ad.waitContext()

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			ret := future.Wait(ctx)

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, ErrAllFuturesExceeded))
	}()

	return resultFuture.Out()
}

// AnyVoid 异步等待任意一个结果返回，无返回值
func (ad AwaitDirector) AnyVoid(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	ctx, cancel := ad.waitContext()

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			ret := future.Wait(ctx)

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, ErrAllFuturesExceeded))
	}()

	return resultFuture.Out()
}

// OK 异步等待任意一个结果成功返回，有返回值
func (ad AwaitDirector) OK(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	ctx, cancel := ad.waitContext()

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			ret := future.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, ErrAllFuturesExceeded))
	}()

	return resultFuture.Out()
}

// OKVoid 异步等待任意一个结果成功返回，无返回值
func (ad AwaitDirector) OKVoid(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	ctx, cancel := ad.waitContext()

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			ret := future.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, ErrAllFuturesExceeded))
	}()

	return resultFuture.Out()
}

// All 异步等待所有结果返回，有返回值
func (ad AwaitDirector) All(fun generic.FuncVar2[runtime.Context, []async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	rets := make([]async.Result, len(ad.futures))
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future, ret *async.Result) {
			defer wg.Done()
			*ret = future.Wait(ad.rtCtx)
		}(ad.futures[i], &rets[i])
	}

	go func() {
		wg.Wait()

		nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
			return fun.UnsafeCall(ctx, rets, args...)
		}, args...)

		async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
	}()

	return resultFuture.Out()
}

// AllVoid 异步等待所有结果返回，无返回值
func (ad AwaitDirector) AllVoid(fun generic.ActionVar2[runtime.Context, []async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	rets := make([]async.Result, len(ad.futures))
	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future, ret *async.Result) {
			defer wg.Done()
			*ret = future.Wait(ad.rtCtx)
		}(ad.futures[i], &rets[i])
	}

	go func() {
		wg.Wait()

		nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
			fun.UnsafeCall(ctx, rets, args...)
		}, args...)

		async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
	}()

	return resultFuture.Out()
}

// Transform 异步等待产出（yield）返回，并变换结果
func (ad AwaitDirector) Transform(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream(len(ad.futures))

	if len(ad.futures) <= 0 {
		return async.YieldBreak(resultFuture)
	}

	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			for ret := range future {
				nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
					return fun.UnsafeCall(ctx, ret, args...)
				}, args...)

				if !async.YieldReturn(ad.rtCtx, resultFuture, nextFuture.Wait(ad.rtCtx)) {
					return
				}
			}
		}(ad.futures[i])
	}

	go func() {
		wg.Wait()
		async.YieldBreak(resultFuture)
	}()

	return resultFuture.Out()
}

// Foreach 异步等待产出（yield）返回
func (ad AwaitDirector) Foreach(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureStream(len(ad.futures))

	if len(ad.futures) <= 0 {
		return async.YieldBreak(resultFuture)
	}

	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			for ret := range future {
				nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
					fun.UnsafeCall(ctx, ret, args...)
				}, args...)

				nextFuture.Wait(ad.rtCtx)
			}
		}(ad.futures[i])
	}

	go func() {
		wg.Wait()
		async.YieldBreak(resultFuture)
	}()

	return resultFuture.Out()
}
