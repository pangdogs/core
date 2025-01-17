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
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
	_ "unsafe"
)

// Caller 异步调用发起者
type Caller interface {
	// CallAsync 查找实体并异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//	- 调用过程中的panic信息，均会转换为error返回。
	CallAsync(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Ret], args ...any) async.AsyncRet

	// CallDelegateAsync 查找实体并异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//	- 调用过程中的panic信息，均会转换为error返回。
	CallDelegateAsync(entityId uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Ret], args ...any) async.AsyncRet

	// CallVoidAsync 查找实体并异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//	- 调用过程中的panic信息，均会转换为error返回。
	CallVoidAsync(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.AsyncRet

	// CallDelegateVoidAsync 查找实体并异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//	- 调用过程中的panic信息，均会转换为error返回。
	CallDelegateVoidAsync(entityId uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.AsyncRet
}

//go:linkname getCaller git.golaxy.org/core/runtime.getCaller
func getCaller(provider ictx.ConcurrentContextProvider) async.Caller

func makeAsyncErr(err error) async.AsyncRet {
	asyncRet := async.MakeAsyncRet()
	asyncRet <- async.MakeRet(nil, err)
	close(asyncRet)
	return asyncRet
}

func checkEntity(entity ec.Entity) error {
	if entity.GetState() > ec.EntityState_Alive {
		return fmt.Errorf("%w: entity not alive", ErrContext)
	}
	return nil
}

// CallAsync 查找实体并异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallAsync(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Ret], args ...any) async.AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).CallAsync(func(...any) async.Ret {
		if err := checkEntity(entity); err != nil {
			return async.MakeRet(nil, err)
		}
		return fun.UnsafeCall(entity, args...)
	})
}

// CallDelegateAsync 查找实体并异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegateAsync(entityId uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Ret], args ...any) async.AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).CallAsync(func(...any) async.Ret {
		if err := checkEntity(entity); err != nil {
			return async.MakeRet(nil, err)
		}
		return fun.UnsafeCall(nil, entity, args...)
	})
}

// CallVoidAsync 查找实体并异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallVoidAsync(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).CallAsync(func(...any) async.Ret {
		if err := checkEntity(entity); err != nil {
			return async.MakeRet(nil, err)
		}
		fun.UnsafeCall(entity, args...)
		return async.VoidRet
	})
}

// CallDelegateVoidAsync 查找实体并异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegateVoidAsync(entityId uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).CallAsync(func(...any) async.Ret {
		if err := checkEntity(entity); err != nil {
			return async.MakeRet(nil, err)
		}
		fun.UnsafeCall(nil, entity, args...)
		return async.VoidRet
	})
}

func (ctx *ContextBehavior) getEntity(id uid.Id) (ec.Entity, error) {
	entity, ok := ctx.entityManager.GetEntity(id)
	if !ok {
		return nil, fmt.Errorf("%w: entity not exist", ErrContext)
	}
	return ec.UnsafeConcurrentEntity(entity).GetEntity(), nil
}
