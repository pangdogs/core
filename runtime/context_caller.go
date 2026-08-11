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

// Caller 将任务投递到 Runtime Actor 邮箱。
//
// Submit 系列返回任务执行结果；Post 系列只报告是否成功入队，不分配 Future。
// 所有回调都由 Runtime goroutine 串行执行，即使调用者已经位于同一 Runtime 中也
// 不会内联执行。
type Caller interface {
	Submit(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future
	SubmitDelegate(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future
	SubmitVoid(fun generic.ActionVar1[Context, any], args ...any) async.Future
	SubmitDelegateVoid(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
	Post(fun generic.ActionVar1[Context, any], args ...any) error
	PostDelegate(fun generic.DelegateVoidVar1[Context, any], args ...any) error
}

// Callee 接收 Caller 的任务并写入实际 Runtime 队列。
type Callee interface {
	PushSubmit(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future
	PushSubmitDelegate(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future
	PushSubmitVoid(fun generic.ActionVar1[Context, any], args ...any) async.Future
	PushSubmitDelegateVoid(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future
	PushPost(fun generic.ActionVar1[Context, any], args ...any) error
	PushPostDelegate(fun generic.DelegateVoidVar1[Context, any], args ...any) error
}

func (ctx *ContextBehavior) Submit(fun generic.FuncVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushSubmit(fun, args...)
}

func (ctx *ContextBehavior) SubmitDelegate(fun generic.DelegateVar1[Context, any, async.Result], args ...any) async.Future {
	return ctx.callee.PushSubmitDelegate(fun, args...)
}

func (ctx *ContextBehavior) SubmitVoid(fun generic.ActionVar1[Context, any], args ...any) async.Future {
	return ctx.callee.PushSubmitVoid(fun, args...)
}

func (ctx *ContextBehavior) SubmitDelegateVoid(fun generic.DelegateVoidVar1[Context, any], args ...any) async.Future {
	return ctx.callee.PushSubmitDelegateVoid(fun, args...)
}

func (ctx *ContextBehavior) Post(fun generic.ActionVar1[Context, any], args ...any) error {
	return ctx.callee.PushPost(fun, args...)
}

func (ctx *ContextBehavior) PostDelegate(fun generic.DelegateVoidVar1[Context, any], args ...any) error {
	return ctx.callee.PushPostDelegate(fun, args...)
}

func checkEntity(entity ec.Entity) error {
	if entity.State() > ec.EntityState_Alive {
		return fmt.Errorf("%w: entity is in an unexpected state %q", ErrContext, entity.State())
	}
	return nil
}

func submit(entity ec.ConcurrentEntity, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	return Concurrent(entity).Submit(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		return fun.UnsafeCall(entity, args...)
	}, args...)
}

func submitDelegate(entity ec.ConcurrentEntity, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	return Concurrent(entity).Submit(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		return fun.UnsafeCall(nil, entity, args...)
	}, args...)
}

func submitVoid(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future {
	return Concurrent(entity).Submit(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		fun.UnsafeCall(entity, args...)
		return async.NewResult(nil, nil)
	}, args...)
}

func submitDelegateVoid(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future {
	return Concurrent(entity).Submit(func(_ Context, args ...any) async.Result {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if err := checkEntity(entity); err != nil {
			return async.NewResult(nil, err)
		}
		fun.UnsafeCall(nil, entity, args...)
		return async.NewResult(nil, nil)
	}, args...)
}

func post(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) error {
	return Concurrent(entity).Post(func(_ Context, args ...any) {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if checkEntity(entity) != nil {
			return
		}
		fun.UnsafeCall(entity, args...)
	}, args...)
}

func postDelegate(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) error {
	return Concurrent(entity).Post(func(_ Context, args ...any) {
		entity := ec.UnsafeConcurrentEntity(entity).Instance()
		if checkEntity(entity) != nil {
			return
		}
		fun.UnsafeCall(nil, entity, args...)
	}, args...)
}
