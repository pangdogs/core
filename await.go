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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"sync"
	"sync/atomic"
)

var (
	ErrAllFailures = fmt.Errorf("%w: all of async result failures", ErrCore)
)

// Await 异步等待结果返回
func Await(provider ictx.CurrentContextProvider, asyncRet ...async.AsyncRet) AwaitDirector {
	return AwaitDirector{
		rtCtx:     runtime.Current(provider),
		asyncRets: asyncRet,
	}
}

// AwaitDirector 异步等待分发器
type AwaitDirector struct {
	rtCtx     runtime.Context
	asyncRets []async.AsyncRet
}

// Any 异步等待任意一个结果返回
func (ad AwaitDirector) Any(fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
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

		go func() {
			ret := asyncRet.Wait(ctx)

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			ad.rtCtx.CallVoid(func(...any) {
				fun.Exec(ad.rtCtx, ret, args...)
			})
		}()
	}
}

// AnyOK 异步等待任意一个结果成功返回
func (ad AwaitDirector) AnyOK(fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
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
		go func() {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			ad.rtCtx.CallVoid(func(...any) {
				fun.Exec(ad.rtCtx, ret, args...)
			})
		}()
	}

	go func() {
		wg.Wait()

		if b.Load() {
			return
		}

		ad.rtCtx.CallVoid(func(...any) {
			fun.Exec(ad.rtCtx, async.MakeRet(nil, ErrAllFailures), args...)
		})
	}()
}

// All 异步等待所有结果返回
func (ad AwaitDirector) All(fun generic.ActionVar2[runtime.Context, []async.Ret, any], args ...any) {
	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	var wg sync.WaitGroup
	rets := make([]async.Ret, len(ad.asyncRets))

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		wg.Add(1)
		go func(ret *async.Ret) {
			defer wg.Done()
			*ret = asyncRet.Wait(ad.rtCtx)
		}(&rets[i])
	}

	go func() {
		wg.Wait()
		ad.rtCtx.CallVoid(func(...any) {
			fun.Exec(ad.rtCtx, rets, args...)
		})
	}()
}

// Pipe 异步等待管道返回
func (ad AwaitDirector) Pipe(ctx context.Context, fun generic.ActionVar2[runtime.Context, async.Ret, any], args ...any) {
	if ctx == nil {
		ctx = context.Background()
	}

	if ad.rtCtx == nil {
		exception.Panicf("%w: rtCtx is nil", ErrCore)
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func() {
			for {
				select {
				case ret, ok := <-asyncRet:
					if !ok {
						return
					}
					ad.rtCtx.CallVoid(func(...any) {
						fun.Exec(ad.rtCtx, ret, args...)
					})
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}
