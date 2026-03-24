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

// Caller 异步调用发起者
type Caller interface {
	// CallAsync 异步调用函数，有返回值。不会阻塞当前线程，会返回Future。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future

	// CallDelegateAsync 异步调用委托，有返回值。不会阻塞当前线程，会返回Future。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future

	// CallVoidAsync 异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回Future。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future

	// CallDelegateVoidAsync 异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回Future。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallDelegateVoidAsync(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
}

// Callee 异步调用接受者
type Callee interface {
	// PushCallAsync 将调用函数压入接受者的任务处理流水线，返回Future。
	PushCallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future
	// PushCallDelegateAsync 将调用委托压入接受者的任务处理流水线，返回Future。
	PushCallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future
	// PushCallVoidAsync 将调用函数压入接受者的任务处理流水线，返回Future。
	PushCallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future
	// PushCallDelegateVoidAsync 将调用委托压入接受者的任务处理流水线，返回Future。
	PushCallDelegateVoidAsync(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
}

// CallAsync 异步调用函数，有返回值。不会阻塞当前线程，会返回Future。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallAsync(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushCallAsync(fun, args...)
}

// CallDelegateAsync 异步调用委托，有返回值。不会阻塞当前线程，会返回Future。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegateAsync(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushCallDelegateAsync(fun, args...)
}

// CallVoidAsync 异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回Future。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallVoidAsync(fun generic.ActionVar1[Context, any], args ...any) async.Future {
	return ctx.callee.PushCallVoidAsync(fun, args...)
}

// CallDelegateVoidAsync 异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回Future。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
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
