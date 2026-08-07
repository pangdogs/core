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
	"errors"
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
	ErrNoFutureSucceeded = fmt.Errorf("%w: no future succeeded", ErrCore) // 所有候选 Future 均失败。
)

// Await 创建与 provider 当前运行时关联的异步等待分发器。
// nil Future 会被忽略，后续回调均通过该运行时的任务队列执行。
func Await(provider corectx.CurrentContextProvider, futures ...async.Future) AwaitDirector {
	return AwaitDirector{
		rtCtx:   runtime.Current(provider),
		futures: slices.DeleteFunc(futures, func(future async.Future) bool { return future.IsNil() }),
	}
}

// AwaitDirector 组合一个或多个 Future，并把后续回调调度回所属运行时。
type AwaitDirector struct {
	rtCtx   runtime.Context
	futures []async.Future
}

func (ad AwaitDirector) singleFuture() (async.Future, bool) {
	if len(ad.futures) != 1 {
		return async.Future{}, false
	}
	return ad.futures[0], true
}

func joinAwaitErrors(errs []error) error {
	errs = slices.DeleteFunc(errs, func(err error) bool { return err == nil })
	if len(errs) <= 0 {
		return ErrNoFutureSucceeded
	}
	return fmt.Errorf("%w: %w", ErrNoFutureSucceeded, errors.Join(errs...))
}

// Any 等待首个 Future 返回结果，再在所属运行时执行 fun。
// 结果是否包含错误不影响选择。
func (ad AwaitDirector) Any(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			ret := future.Wait(ad.rtCtx)
			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
	}

	ctx, cancel := context.WithCancel(ad.rtCtx)

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

			cancel()

			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		cancel()

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors(nil)))
	}()

	return resultFuture.Out()
}

// AnyVoid 等待首个 Future 返回结果，再在所属运行时执行无返回值的 fun。
func (ad AwaitDirector) AnyVoid(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			ret := future.Wait(ad.rtCtx)
			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
	}

	ctx, cancel := context.WithCancel(ad.rtCtx)

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

			cancel()

			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i])
	}

	go func() {
		wg.Wait()

		cancel()

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors(nil)))
	}()

	return resultFuture.Out()
}

// OK 等待首个成功结果，再在所属运行时执行 fun。
// 所有 Future 均失败时返回包装 ErrNoFutureSucceeded 的错误。
func (ad AwaitDirector) OK(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			ret := future.Wait(ad.rtCtx)
			if !ret.OK() {
				async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors([]error{ret.Error})))
				return
			}

			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
	}

	ctx, cancel := context.WithCancel(ad.rtCtx)

	var once atomic.Bool
	var wg sync.WaitGroup
	errs := make([]error, len(ad.futures))

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future, errRef *error) {
			defer wg.Done()

			ret := future.Wait(ctx)
			if !ret.OK() {
				*errRef = ret.Error
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			cancel()

			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i], &errs[i])
	}

	go func() {
		wg.Wait()

		cancel()

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors(errs)))
	}()

	return resultFuture.Out()
}

// OKVoid 等待首个成功结果，再在所属运行时执行无返回值的 fun。
// 所有 Future 均失败时返回包装 ErrNoFutureSucceeded 的错误。
func (ad AwaitDirector) OKVoid(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			ret := future.Wait(ad.rtCtx)
			if !ret.OK() {
				async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors([]error{ret.Error})))
				return
			}

			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
	}

	ctx, cancel := context.WithCancel(ad.rtCtx)

	var once atomic.Bool
	var wg sync.WaitGroup
	errs := make([]error, len(ad.futures))

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future, errRef *error) {
			defer wg.Done()

			ret := future.Wait(ctx)
			if !ret.OK() {
				*errRef = ret.Error
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			cancel()

			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, ret, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))

		}(ad.futures[i], &errs[i])
	}

	go func() {
		wg.Wait()

		cancel()

		if once.Load() {
			return
		}

		async.Return(resultFuture, async.NewResult(nil, joinAwaitErrors(errs)))
	}()

	return resultFuture.Out()
}

// All 按传入顺序收集所有 Future 的结果，再在所属运行时执行 fun。
func (ad AwaitDirector) All(fun generic.FuncVar2[runtime.Context, []async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			rets := []async.Result{future.Wait(ad.rtCtx)}
			nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
				return fun.UnsafeCall(ctx, rets, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
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

// AllVoid 按传入顺序收集所有 Future 的结果，再在所属运行时执行无返回值的 fun。
func (ad AwaitDirector) AllVoid(fun generic.ActionVar2[runtime.Context, []async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			rets := []async.Result{future.Wait(ad.rtCtx)}
			nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
				fun.UnsafeCall(ctx, rets, args...)
			}, args...)

			async.Return(resultFuture, nextFuture.Wait(ad.rtCtx))
		}()
		return resultFuture.Out()
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

// Transform 合并所有 Future 的连续产出，在所属运行时逐项执行 fun，并产出转换结果。
func (ad AwaitDirector) Transform(fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan(len(ad.futures))

	if len(ad.futures) <= 0 {
		return async.YieldBreak(resultFuture)
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			defer async.YieldBreak(resultFuture)

			for ret := range future.Chan() {
				nextFuture := ad.rtCtx.CallAsync(func(ctx runtime.Context, args ...any) async.Result {
					return fun.UnsafeCall(ctx, ret, args...)
				}, args...)

				if !async.YieldReturn(ad.rtCtx, resultFuture, nextFuture.Wait(ad.rtCtx)) {
					return
				}
			}
		}()
		return resultFuture.Out()
	}

	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			for ret := range future.Chan() {
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

// Foreach 合并所有 Future 的连续产出，并在所属运行时逐项执行 fun。
func (ad AwaitDirector) Foreach(fun generic.ActionVar2[runtime.Context, async.Result, any], args ...any) async.Future {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	resultFuture := async.NewFutureChan()

	if len(ad.futures) <= 0 {
		return async.Return(resultFuture, async.NewResult(nil, nil))
	}

	if future, ok := ad.singleFuture(); ok {
		go func() {
			for ret := range future.Chan() {
				nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
					fun.UnsafeCall(ctx, ret, args...)
				}, args...)

				nextFuture.Wait(ad.rtCtx)
			}
			async.Return(resultFuture, async.NewResult(nil, nil))
		}()
		return resultFuture.Out()
	}

	var wg sync.WaitGroup

	for i := range ad.futures {
		wg.Add(1)

		go func(future async.Future) {
			defer wg.Done()

			for ret := range future.Chan() {
				nextFuture := ad.rtCtx.CallVoidAsync(func(ctx runtime.Context, args ...any) {
					fun.UnsafeCall(ctx, ret, args...)
				}, args...)

				nextFuture.Wait(ad.rtCtx)
			}
		}(ad.futures[i])
	}

	go func() {
		wg.Wait()
		async.Return(resultFuture, async.NewResult(nil, nil))
	}()

	return resultFuture.Out()
}
