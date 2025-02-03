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
	"git.golaxy.org/core/ec/ectx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"slices"
	"sync"
	"sync/atomic"
)

var (
	ErrAllAsyncRetExceeded = fmt.Errorf("%w: all of async result exceeded, %w", ErrCore, context.DeadlineExceeded)
)

// Await 异步等待结果返回
func Await(provider ectx.CurrentContextProvider, asyncRets ...async.AsyncRet) AwaitDirector {
	return AwaitDirector{
		rtCtx:     runtime.Current(provider),
		asyncRets: slices.DeleteFunc(asyncRets, func(asyncRet async.AsyncRet) bool { return asyncRet == nil }),
	}
}

// AwaitDirector 异步等待分发器
type AwaitDirector struct {
	rtCtx     runtime.Context
	asyncRets []async.AsyncRet
}

// Any 异步等待任意一个结果返回，有返回值
func (ad AwaitDirector) Any(fun generic.FuncVar2[runtime.Context, async.Ret, any, async.Ret], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	var ctx context.Context
	var cancel context.CancelFunc

	if len(ad.asyncRets) > 1 {
		ctx, cancel = context.WithCancel(ad.rtCtx)
	} else {
		ctx = ad.rtCtx
	}

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			callRet := ad.rtCtx.CallAsync(func(args ...any) async.Ret {
				return fun.UnsafeCall(ad.rtCtx, ret, args...)
			}, args...)

			async.Return(awaitRet, callRet.Wait(ad.rtCtx))

		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(awaitRet, async.MakeRet(nil, ErrAllAsyncRetExceeded))
	}()

	return awaitRet
}

// AnyVoid 异步等待任意一个结果返回，无返回值
func (ad AwaitDirector) AnyVoid(fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	var ctx context.Context
	var cancel context.CancelFunc

	if len(ad.asyncRets) > 1 {
		ctx, cancel = context.WithCancel(ad.rtCtx)
	} else {
		ctx = ad.rtCtx
	}

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			callRet := ad.rtCtx.CallVoidAsync(func(args ...any) {
				fun.UnsafeCall(ad.rtCtx, ret, args...)
			}, args...)

			async.Return(awaitRet, callRet.Wait(ad.rtCtx))

		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(awaitRet, async.MakeRet(nil, ErrAllAsyncRetExceeded))
	}()

	return awaitRet
}

// OK 异步等待任意一个结果成功返回，有返回值
func (ad AwaitDirector) OK(fun generic.FuncVar2[runtime.Context, async.Ret, any, async.Ret], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	var ctx context.Context
	var cancel context.CancelFunc

	if len(ad.asyncRets) > 1 {
		ctx, cancel = context.WithCancel(ad.rtCtx)
	} else {
		ctx = ad.rtCtx
	}

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			callRet := ad.rtCtx.CallAsync(func(args ...any) async.Ret {
				return fun.UnsafeCall(ad.rtCtx, ret, args...)
			}, args...)

			async.Return(awaitRet, callRet.Wait(ad.rtCtx))

		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(awaitRet, async.MakeRet(nil, ErrAllAsyncRetExceeded))
	}()

	return awaitRet
}

// OKVoid 异步等待任意一个结果成功返回，无返回值
func (ad AwaitDirector) OKVoid(fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	var ctx context.Context
	var cancel context.CancelFunc

	if len(ad.asyncRets) > 1 {
		ctx, cancel = context.WithCancel(ad.rtCtx)
	} else {
		ctx = ad.rtCtx
	}

	var once atomic.Bool
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !once.CompareAndSwap(false, true) {
				return
			}

			if cancel != nil {
				cancel()
			}

			callRet := ad.rtCtx.CallVoidAsync(func(args ...any) {
				fun.UnsafeCall(ad.rtCtx, ret, args...)
			}, args...)

			async.Return(awaitRet, callRet.Wait(ad.rtCtx))

		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()

		if cancel != nil {
			cancel()
		}

		if once.Load() {
			return
		}

		async.Return(awaitRet, async.MakeRet(nil, ErrAllAsyncRetExceeded))
	}()

	return awaitRet
}

// All 异步等待所有结果返回，有返回值
func (ad AwaitDirector) All(fun generic.FuncVar2[runtime.Context, []async.Ret, any, async.Ret], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	rets := make([]async.Ret, len(ad.asyncRets))
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet, ret *async.Ret) {
			defer wg.Done()
			*ret = asyncRet.Wait(ad.rtCtx)
		}(ad.asyncRets[i], &rets[i])
	}

	go func() {
		wg.Wait()

		callRet := ad.rtCtx.CallAsync(func(args ...any) async.Ret {
			return fun.UnsafeCall(ad.rtCtx, rets, args...)
		}, args...)

		async.Return(awaitRet, callRet.Wait(ad.rtCtx))
	}()

	return awaitRet
}

// AllVoid 异步等待所有结果返回，无返回值
func (ad AwaitDirector) AllVoid(fun generic.ActionVar2[runtime.Context, []async.Ret, any], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	rets := make([]async.Ret, len(ad.asyncRets))
	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet, ret *async.Ret) {
			defer wg.Done()
			*ret = asyncRet.Wait(ad.rtCtx)
		}(ad.asyncRets[i], &rets[i])
	}

	go func() {
		wg.Wait()

		callRet := ad.rtCtx.CallVoidAsync(func(args ...any) {
			fun.UnsafeCall(ad.rtCtx, rets, args...)
		}, args...)

		async.Return(awaitRet, callRet.Wait(ad.rtCtx))
	}()

	return awaitRet
}

// Transform 异步等待产出（yield）返回，并变换结果
func (ad AwaitDirector) Transform(fun generic.FuncVar2[runtime.Context, async.Ret, any, async.Ret], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		async.End(awaitRet)
		return awaitRet
	}

	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			for ret := range asyncRet {
				callRet := ad.rtCtx.CallAsync(func(args ...any) async.Ret {
					return fun.UnsafeCall(ad.rtCtx, ret, args...)
				}, args...)

				async.Yield(nil, awaitRet, callRet.Wait(ad.rtCtx))
			}
		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()
		async.End(awaitRet)
	}()

	return awaitRet
}

// Foreach 异步等待产出（yield）返回
func (ad AwaitDirector) Foreach(fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) async.AsyncRet {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	awaitRet := async.MakeAsyncRet()

	if len(ad.asyncRets) <= 0 {
		return async.Return(awaitRet, async.VoidRet)
	}

	var wg sync.WaitGroup

	for i := range ad.asyncRets {
		wg.Add(1)

		go func(asyncRet async.AsyncRet) {
			defer wg.Done()

			for ret := range asyncRet {
				callRet := ad.rtCtx.CallVoidAsync(func(args ...any) {
					fun.UnsafeCall(ad.rtCtx, ret, args...)
				}, args...)

				callRet.Wait(ad.rtCtx)
			}
		}(ad.asyncRets[i])
	}

	go func() {
		wg.Wait()
		async.Return(awaitRet, async.VoidRet)
	}()

	return awaitRet
}
