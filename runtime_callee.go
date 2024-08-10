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
	"fmt"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

var (
	ErrProcessQueueClosed = fmt.Errorf("%w: process queue is closed", ErrRuntime) // 任务处理流水线关闭
	ErrProcessQueueFull   = fmt.Errorf("%w: process queue is full", ErrRuntime)   // 任务处理流水线已满
)

// PushCall 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCall(fun generic.FuncVar0[any, async.Ret], args ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		fun:  fun,
		args: args,
	})
}

// PushCallDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallDelegate(fun generic.DelegateFuncVar0[any, async.Ret], args ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateFun: fun,
		args:        args,
	})
}

// PushCallVoid 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoid(fun generic.ActionVar0[any], args ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		action: fun,
		args:   args,
	})
}

// PushCallVoidDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoidDelegate(fun generic.DelegateActionVar0[any], args ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateAction: fun,
		args:           args,
	})
}

func (rt *RuntimeBehavior) pushCallTask(task _Task) (asyncRet chan async.Ret) {
	task.typ = _TaskType_Call
	task.asyncRet = async.MakeAsyncRet()

	asyncRet = task.asyncRet

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			asyncRet <- async.MakeRet(nil, ErrProcessQueueClosed)
			close(asyncRet)
		}
	}()

	select {
	case rt.processQueue <- task:
		return
	default:
		break
	}

	asyncRet <- async.MakeRet(nil, ErrProcessQueueFull)
	close(asyncRet)

	return
}
