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
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

// CallAsync 异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallAsync(fun generic.FuncVar0[any, async.Ret], args ...any) async.AsyncRet {
	return ctx.callee.PushCallAsync(fun, args...)
}

// CallDelegateAsync 异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegateAsync(fun generic.DelegateVar0[any, async.Ret], args ...any) async.AsyncRet {
	return ctx.callee.PushCallDelegateAsync(fun, args...)
}

// CallVoidAsync 异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallVoidAsync(fun generic.ActionVar0[any], args ...any) async.AsyncRet {
	return ctx.callee.PushCallVoidAsync(fun, args...)
}

// CallDelegateVoidAsync 异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegateVoidAsync(fun generic.DelegateVoidVar0[any], args ...any) async.AsyncRet {
	return ctx.callee.PushCallDelegateVoidAsync(fun, args...)
}
