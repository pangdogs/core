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
	va             []any
	asyncRet       chan async.Ret
}

func (task _Task) run(autoRecover bool, reportError chan error) {
	var ret async.Ret
	var panicErr error

	if task.fun != nil {
		ret, panicErr = task.fun.Call(autoRecover, reportError, task.va...)
	} else if task.delegateFun != nil {
		ret, panicErr = task.delegateFun.Call(autoRecover, reportError, nil, task.va...)
	} else if task.action != nil {
		panicErr = task.action.Call(autoRecover, reportError, task.va...)
	} else if task.delegateAction != nil {
		panicErr = task.delegateAction.Call(autoRecover, reportError, nil, task.va...)
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
