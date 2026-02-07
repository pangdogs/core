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

type TaskType int8

const (
	TaskType_Call TaskType = iota
	TaskType_Frame
)

type _Task struct {
	typ          TaskType
	fun          generic.FuncVar0[any, async.Ret]
	action       generic.ActionVar0[any]
	delegate     generic.DelegateVar0[any, async.Ret]
	delegateVoid generic.DelegateVoidVar0[any]
	args         []any
	asyncRet     chan async.Ret
	done         chan struct{}
}

func (task _Task) run(autoRecover bool, reportError chan error) {
	var ret async.Ret
	var panicErr error

	if task.fun != nil {
		ret, panicErr = task.fun.Call(autoRecover, reportError, task.args...)
	} else if task.action != nil {
		panicErr = task.action.Call(autoRecover, reportError, task.args...)
	} else if task.delegate != nil {
		ret, panicErr = task.delegate.Call(autoRecover, reportError, nil, task.args...)
	} else if task.delegateVoid != nil {
		panicErr = task.delegateVoid.Call(autoRecover, reportError, nil, task.args...)
	}

	if panicErr != nil {
		ret.Value = nil
		ret.Error = panicErr
	}

	if task.asyncRet != nil {
		async.Return(task.asyncRet, ret)
	}

	if task.done != nil {
		select {
		case task.done <- struct{}{}:
		default:
		}
	}
}
