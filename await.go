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
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"sync"
	"sync/atomic"
	_ "unsafe"
)

var (
	ErrAllFailures = fmt.Errorf("%w: all of async result failures", ErrCore)
)

//go:linkname getRuntimeContext git.golaxy.org/core/runtime.getRuntimeContext
func getRuntimeContext(provider gctx.CurrentContextProvider) runtime.Context

// Await 异步等待结果返回
func Await(provider gctx.CurrentContextProvider, asyncRet ...async.AsyncRet) AwaitDirector {
	return AwaitDirector{
		rtCtx:     getRuntimeContext(provider),
		asyncRets: asyncRet,
	}
}

// AwaitDirector 异步等待分发器
type AwaitDirector struct {
	rtCtx     runtime.Context
	asyncRets []async.AsyncRet
}

// Any 异步等待任意一个结果返回
func (ad AwaitDirector) Any(fun generic.ActionVar2[runtime.Context, async.Ret, any], va ...any) {
	if ad.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrCore))
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	var b atomic.Bool
	ctx, cancel := context.WithCancel(ad.rtCtx)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func(b *atomic.Bool, ctx context.Context, cancel context.CancelFunc,
			asyncRet async.AsyncRet, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, async.Ret, any], va []any) {

			ret := asyncRet.Wait(ctx)

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			rtCtx.CallVoid(func(va ...any) {
				rtCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, async.Ret, any])
				ret := va[2].(async.Ret)
				funVa := va[3].([]any)
				fun.Exec(rtCtx, ret, funVa...)
			}, rtCtx, fun, ret, va)
		}(&b, ctx, cancel, asyncRet, ad.rtCtx, fun, va)
	}
}

// AnyOK 异步等待任意一个结果成功返回
func (ad AwaitDirector) AnyOK(fun generic.ActionVar2[runtime.Context, async.Ret, any], va ...any) {
	if ad.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrCore))
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	var wg sync.WaitGroup
	var b atomic.Bool
	ctx, cancel := context.WithCancel(ad.rtCtx)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, b *atomic.Bool, ctx context.Context, cancel context.CancelFunc,
			asyncRet async.AsyncRet, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, async.Ret, any], va []any) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			rtCtx.CallVoid(func(va ...any) {
				rtCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, async.Ret, any])
				ret := va[2].(async.Ret)
				funVa := va[3].([]any)
				fun.Exec(rtCtx, ret, funVa...)
			}, rtCtx, fun, ret, va)
		}(&wg, &b, ctx, cancel, asyncRet, ad.rtCtx, fun, va)
	}

	go func(wg *sync.WaitGroup, b *atomic.Bool, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, async.Ret, any], va []any) {
		wg.Wait()

		if b.Load() {
			return
		}

		rtCtx.CallVoid(func(va ...any) {
			rtCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, async.Ret, any])
			funVa := va[2].([]any)
			fun.Exec(rtCtx, async.MakeRet(nil, ErrAllFailures), funVa...)
		}, rtCtx, fun, va)
	}(&wg, &b, ad.rtCtx, fun, va)
}

// All 异步等待所有结果返回
func (ad AwaitDirector) All(fun generic.ActionVar2[runtime.Context, []async.Ret, any], va ...any) {
	if ad.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrCore))
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ad.rtCtx)
	rets := make([]async.Ret, len(ad.asyncRets))

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, ret *async.Ret, asyncRet async.AsyncRet) {
			defer wg.Done()
			*ret = asyncRet.Wait(ctx)
		}(&wg, ctx, cancel, &rets[i], asyncRet)
	}

	go func(wg *sync.WaitGroup, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, []async.Ret, any], rets []async.Ret, va []any) {
		wg.Wait()

		rtCtx.CallVoid(func(va ...any) {
			rtCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, []async.Ret, any])
			rets := va[2].([]async.Ret)
			funVa := va[3].([]any)
			fun.Exec(rtCtx, rets, funVa...)
		}, rtCtx, fun, rets, va)
	}(&wg, ad.rtCtx, fun, rets, va)
}

// Pipe 异步等待管道返回
func (ad AwaitDirector) Pipe(ctx context.Context, fun generic.ActionVar2[runtime.Context, async.Ret, any], va ...any) {
	if ctx == nil {
		ctx = context.Background()
	}

	if ad.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrCore))
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func(ctx context.Context, rtCtx runtime.Context, asyncRet async.AsyncRet, fun generic.ActionVar2[runtime.Context, async.Ret, any], va []any) {
			for {
				select {
				case ret, ok := <-asyncRet:
					if !ok {
						return
					}
					rtCtx.CallVoid(func(va ...any) {
						rtCtx := va[0].(runtime.Context)
						fun := va[1].(generic.ActionVar2[runtime.Context, async.Ret, any])
						ret := va[2].(async.Ret)
						funVa := va[3].([]any)
						fun.Exec(rtCtx, ret, funVa...)
					}, rtCtx, fun, ret, va)
				case <-ctx.Done():
					return
				}
			}
		}(ctx, ad.rtCtx, asyncRet, fun, va)
	}
}
