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

package runtime

import (
	"fmt"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

// Caller 将调用异步调度到运行时任务队列。
// 回调在运行时 goroutine 中串行执行。AutoRecover 启用时，panic 会写入 Result.Error；
// 否则 panic 会继续向运行时工作循环传播。
type Caller interface {
	// CallAsync 异步执行有返回值的函数。
	CallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future

	// CallDelegateAsync 异步执行有返回值的委托。
	CallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future

	// CallVoidAsync 异步执行无返回值的函数。
	CallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future

	// CallDelegateVoidAsync 异步执行无返回值的委托。
	CallDelegateVoidAsync(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
}

// Callee 接收异步调用并将其加入运行时任务队列。
type Callee interface {
	// PushCallAsync 将有返回值的函数加入任务队列，并返回承载调用结果的 Future。
	PushCallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future
	// PushCallDelegateAsync 将有返回值的委托加入任务队列，并返回承载调用结果的 Future。
	PushCallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future
	// PushCallVoidAsync 将无返回值的函数加入任务队列，并返回完成信号。
	PushCallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future
	// PushCallDelegateVoidAsync 将无返回值的委托加入任务队列，并返回完成信号。
	PushCallDelegateVoidAsync(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
}

// CallAsync 将有返回值的函数加入运行时任务队列。
func (ctx *ContextBehavior) CallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushCallAsync(fun, args...)
}

// CallDelegateAsync 将有返回值的委托加入运行时任务队列。
func (ctx *ContextBehavior) CallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushCallDelegateAsync(fun, args...)
}

// CallVoidAsync 将无返回值的函数加入运行时任务队列。
func (ctx *ContextBehavior) CallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future {
	return ctx.callee.PushCallVoidAsync(fun, args...)
}

// CallDelegateVoidAsync 将无返回值的委托加入运行时任务队列。
func (ctx *ContextBehavior) CallDelegateVoidAsync(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future {
	return ctx.callee.PushCallDelegateVoidAsync(fun, args...)
}

func checkEntity(entity ec.Entity) error {
	if entity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: entity is in an unexpected state %q", ErrContext, entity.State())
	}
	return nil
}

func callAsync(entity ec.ConcurrentEntity, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	return Concurrent(entity).CallAsync(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Entity()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		return fun.UnsafeCall(entity, args...)
	}, args...)
}

func callDelegateAsync(entity ec.ConcurrentEntity, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	return Concurrent(entity).CallAsync(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Entity()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		return fun.UnsafeCall(nil, entity, args...)
	}, args...)
}

func callVoidAsync(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future {
	return Concurrent(entity).CallAsync(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Entity()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		fun.UnsafeCall(entity, args...)
		return async.NewResult(nil, nil)
	}, args...)
}

func callDelegateVoidAsync(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future {
	return Concurrent(entity).CallAsync(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Entity()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		fun.UnsafeCall(nil, entity, args...)
		return async.NewResult(nil, nil)
	}, args...)
}
