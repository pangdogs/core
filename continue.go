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
	"sync/atomic"

	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
)

// ContinueOn 在 future 完成后，把 fun 作为新任务投递到 provider 所属 Runtime。
//
// 它在订阅、入队和执行前检查异步 Scope；队列关闭、容量不足、Scope 关闭和回调
// panic 都通过返回 Future 报告。回调始终在 Runtime goroutine 中串行执行。
func ContinueOn(
	provider corectx.ConcurrentContextProvider,
	future async.Future,
	fun generic.FuncVar2[runtime.Context, async.Result, any, async.Result],
	args ...any,
) async.Future {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrCore, ErrArgs)
	}
	if future.IsNil() {
		return async.Rejected(async.ErrNoCandidates)
	}
	if fun == nil {
		exception.Panicf("%w: %w: continuation is nil", ErrCore, ErrArgs)
	}

	rt := runtime.Concurrent(provider)
	scope := continuationScope(provider, rt)
	promise, next := async.NewPromise(rt.ExecutorID())
	var executing atomic.Bool
	if scope == nil || scope.Err() != nil {
		promise.Resolve(scopeClosedResult(scope))
		return next
	}

	subscription := future.OnComplete(func(ret async.Result) {
		if scope.Err() != nil {
			promise.Resolve(scopeClosedResult(scope))
			return
		}

		submitted := rt.Submit(func(ctx runtime.Context, _ ...any) async.Result {
			executing.Store(true)
			if scope.Err() != nil {
				return scopeClosedResult(scope)
			}
			return fun.UnsafeCall(ctx, ret, args...)
		})
		submitted.OnComplete(func(ret async.Result) { promise.Resolve(ret) })
	})

	stopScope := context.AfterFunc(scope.Context(), func() {
		subscription.Cancel()
		if !executing.Load() {
			promise.Resolve(scopeClosedResult(scope))
		}
	})
	next.OnComplete(func(async.Result) { stopScope() })
	return next
}

// ContinueOnDelegate 是 ContinueOn 的 Delegate 版本。
func ContinueOnDelegate(
	provider corectx.ConcurrentContextProvider,
	future async.Future,
	fun generic.DelegateVar2[runtime.Context, async.Result, any, async.Result],
	args ...any,
) async.Future {
	return ContinueOn(provider, future, func(ctx runtime.Context, ret async.Result, args ...any) async.Result {
		return fun.UnsafeCall(nil, ctx, ret, args...)
	}, args...)
}

// ContinueOnVoid 在 Runtime 中执行无业务返回值续体，并返回可等待错误的 Future。
func ContinueOnVoid(
	provider corectx.ConcurrentContextProvider,
	future async.Future,
	fun generic.ActionVar2[runtime.Context, async.Result, any],
	args ...any,
) async.Future {
	return ContinueOn(provider, future, func(ctx runtime.Context, ret async.Result, args ...any) async.Result {
		fun.UnsafeCall(ctx, ret, args...)
		return async.NewResult(nil, nil)
	}, args...)
}

// ContinueOnDelegateVoid 是 ContinueOnVoid 的 DelegateVoid 版本。
func ContinueOnDelegateVoid(
	provider corectx.ConcurrentContextProvider,
	future async.Future,
	fun generic.DelegateVoidVar2[runtime.Context, async.Result, any],
	args ...any,
) async.Future {
	return ContinueOn(provider, future, func(ctx runtime.Context, ret async.Result, args ...any) async.Result {
		fun.UnsafeCall(nil, ctx, ret, args...)
		return async.NewResult(nil, nil)
	}, args...)
}

func continuationScope(provider corectx.ConcurrentContextProvider, rt runtime.ConcurrentContext) *async.Scope {
	if scoped, ok := provider.(corectx.AsyncScopeProvider); ok {
		if scope := scoped.AsyncScope(); scope != nil {
			return scope
		}
	}
	return rt.AsyncScope()
}

func scopeClosedResult(scope *async.Scope) async.Result {
	if scope == nil || scope.Err() == nil {
		return async.NewResult(nil, async.ErrScopeClosed)
	}
	return async.NewResult(nil, fmt.Errorf("%w: %w", async.ErrScopeClosed, scope.Err()))
}
