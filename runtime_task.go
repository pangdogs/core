package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/generic"
)

type _TaskKind int8

const (
	_TaskKind_Call _TaskKind = iota
	_TaskKind_Frame
)

type _Task struct {
	kind           _TaskKind
	fun            generic.FuncVar0[any, runtime.Ret]
	delegateFun    generic.DelegateFuncVar0[any, runtime.Ret]
	action         generic.ActionVar0[any]
	delegateAction generic.DelegateActionVar0[any]
	va             []any
	asyncRet       chan runtime.Ret
}

func (task _Task) run(autoRecover bool, reportError chan error) {
	var ret runtime.Ret
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
