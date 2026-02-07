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
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

// PushCallAsync 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallAsync(fun generic.FuncVar0[any, async.Ret], args ...any) async.AsyncRet {
	return rt.taskQueue.enqueueCall(fun, nil, nil, nil, args)
}

// PushCallDelegateAsync 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallDelegateAsync(fun generic.DelegateVar0[any, async.Ret], args ...any) async.AsyncRet {
	return rt.taskQueue.enqueueCall(nil, nil, fun, nil, args)
}

// PushCallVoidAsync 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoidAsync(fun generic.ActionVar0[any], args ...any) async.AsyncRet {
	return rt.taskQueue.enqueueCall(nil, fun, nil, nil, args)
}

// PushCallDelegateVoidAsync 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallDelegateVoidAsync(fun generic.DelegateVoidVar0[any], args ...any) async.AsyncRet {
	return rt.taskQueue.enqueueCall(nil, nil, nil, fun, args)
}
