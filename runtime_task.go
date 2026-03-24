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
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

type TaskType int8

const (
	TaskType_Call TaskType = iota
	TaskType_Frame
)

type _Task struct {
	typ          TaskType
	fun          generic.FuncVar1[runtime.Context, any, async.Result]
	action       generic.ActionVar1[runtime.Context, any]
	delegate     generic.DelegateVar1[runtime.Context, any, async.Result]
	delegateVoid generic.DelegateVoidVar1[runtime.Context, any]
	args         []any
	future       async.FutureChan
	done         chan struct{}
}

func (task _Task) run(ctx runtime.Context) {
	var ret async.Result
	var panicErr error

	switch {
	case task.fun != nil:
		ret, panicErr = task.fun.Call(ctx.AutoRecover(), ctx.ReportError(), ctx, task.args...)
	case task.action != nil:
		panicErr = task.action.Call(ctx.AutoRecover(), ctx.ReportError(), ctx, task.args...)
	case task.delegate != nil:
		ret, panicErr = task.delegate.Call(ctx.AutoRecover(), ctx.ReportError(), nil, ctx, task.args...)
	case task.delegateVoid != nil:
		panicErr = task.delegateVoid.Call(ctx.AutoRecover(), ctx.ReportError(), nil, ctx, task.args...)
	}

	if panicErr != nil {
		ret.Value = nil
		ret.Error = panicErr
	}

	if !task.future.IsNil() {
		async.Return(task.future, ret)
	}

	if task.done != nil {
		select {
		case task.done <- struct{}{}:
		default:
		}
	}
}
