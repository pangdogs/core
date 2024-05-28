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
func (rt *RuntimeBehavior) PushCall(fun generic.FuncVar0[any, async.Ret], va ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		fun: fun,
		va:  va,
	})
}

// PushCallDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallDelegate(fun generic.DelegateFuncVar0[any, async.Ret], va ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateFun: fun,
		va:          va,
	})
}

// PushCallVoid 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoid(fun generic.ActionVar0[any], va ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		action: fun,
		va:     va,
	})
}

// PushCallVoidDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoidDelegate(fun generic.DelegateActionVar0[any], va ...any) async.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateAction: fun,
		va:             va,
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
