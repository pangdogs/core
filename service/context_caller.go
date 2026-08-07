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

package service

import (
	"fmt"
	_ "unsafe"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
)

// Caller 通过全局实体 ID 将调用调度到实体所属的运行时。
// 调用不会阻塞发起方；回调在目标运行时 goroutine 中执行。目标 Runtime 启用
// AutoRecover 时，panic 会写入 Result.Error；否则会继续向其工作循环传播。
type Caller interface {
	// CallAsync 查找实体并异步执行有返回值的函数。
	CallAsync(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future

	// CallDelegateAsync 查找实体并异步执行有返回值的委托。
	CallDelegateAsync(entityId uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future

	// CallVoidAsync 查找实体并异步执行无返回值的函数。
	CallVoidAsync(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future

	// CallDelegateVoidAsync 查找实体并异步执行无返回值的委托。
	CallDelegateVoidAsync(entityId uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future
}

//go:linkname callAsync git.golaxy.org/core/runtime.callAsync
func callAsync(entity ec.ConcurrentEntity, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future

//go:linkname callDelegateAsync git.golaxy.org/core/runtime.callDelegateAsync
func callDelegateAsync(entity ec.ConcurrentEntity, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future

//go:linkname callVoidAsync git.golaxy.org/core/runtime.callVoidAsync
func callVoidAsync(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future

//go:linkname callDelegateVoidAsync git.golaxy.org/core/runtime.callDelegateVoidAsync
func callDelegateVoidAsync(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future

// CallAsync 查找实体并在其所属运行时异步执行有返回值的函数。
// 实体不存在或调用失败时，错误通过 Future 的 Result.Error 返回。
func (ctx *ContextBehavior) CallAsync(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return async.Return(async.NewFutureChan(), async.NewResult(nil, err))
	}
	return callAsync(entity, fun, args...)
}

// CallDelegateAsync 查找实体并在其所属运行时异步执行有返回值的委托。
// 实体不存在或调用失败时，错误通过 Future 的 Result.Error 返回。
func (ctx *ContextBehavior) CallDelegateAsync(entityId uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return async.Return(async.NewFutureChan(), async.NewResult(nil, err))
	}
	return callDelegateAsync(entity, fun, args...)
}

// CallVoidAsync 查找实体并在其所属运行时异步执行无返回值的函数。
// 实体不存在或调用失败时，错误通过 Future 的 Result.Error 返回。
func (ctx *ContextBehavior) CallVoidAsync(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return async.Return(async.NewFutureChan(), async.NewResult(nil, err))
	}
	return callVoidAsync(entity, fun, args...)
}

// CallDelegateVoidAsync 查找实体并在其所属运行时异步执行无返回值的委托。
// 实体不存在或调用失败时，错误通过 Future 的 Result.Error 返回。
func (ctx *ContextBehavior) CallDelegateVoidAsync(entityId uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return async.Return(async.NewFutureChan(), async.NewResult(nil, err))
	}
	return callDelegateVoidAsync(entity, fun, args...)
}

func (ctx *ContextBehavior) getEntity(id uid.Id) (ec.ConcurrentEntity, error) {
	entity, ok := ctx.entityManager.GetEntity(id)
	if !ok {
		return nil, fmt.Errorf("%w: entity not exist", ErrContext)
	}
	return entity, nil
}
