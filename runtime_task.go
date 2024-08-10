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

type _TaskType int8

const (
	_TaskType_Call _TaskType = iota
	_TaskType_Frame
)

type _Task struct {
	typ            _TaskType
	fun            generic.FuncVar0[any, async.Ret]
	delegateFun    generic.DelegateFuncVar0[any, async.Ret]
	action         generic.ActionVar0[any]
	delegateAction generic.DelegateActionVar0[any]
	args           []any
	asyncRet       chan async.Ret
}

func (task _Task) run(autoRecover bool, reportError chan error) {
	var ret async.Ret
	var panicErr error

	if task.fun != nil {
		ret, panicErr = task.fun.Call(autoRecover, reportError, task.args...)
	} else if task.delegateFun != nil {
		ret, panicErr = task.delegateFun.Call(autoRecover, reportError, nil, task.args...)
	} else if task.action != nil {
		panicErr = task.action.Call(autoRecover, reportError, task.args...)
	} else if task.delegateAction != nil {
		panicErr = task.delegateAction.Call(autoRecover, reportError, nil, task.args...)
	}

	if panicErr != nil {
		ret.Value = nil
		ret.Error = panicErr
	}

	if task.asyncRet != nil {
		task.asyncRet <- ret
		close(task.asyncRet)
	}
}
