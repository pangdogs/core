package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/generic"
	"time"
)

var (
	ErrProcessQueueClosed = fmt.Errorf("%w: process queue is closed", ErrRuntime) // 任务处理流水线关闭
	ErrProcessQueueFull   = fmt.Errorf("%w: process queue is full", ErrRuntime)   // 任务处理流水线已满
)

// PushCall 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCall(fun generic.FuncVar0[any, runtime.Ret], va ...any) runtime.AsyncRet {
	return rt.pushCallTask(_Task{
		fun: fun,
		va:  va,
	})
}

// PushCallDelegate 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallDelegate(fun generic.DelegateFuncVar0[any, runtime.Ret], va ...any) runtime.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateFun: fun,
		va:          va,
	})
}

// PushCallVoid 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoid(fun generic.ActionVar0[any], va ...any) runtime.AsyncRet {
	return rt.pushCallTask(_Task{
		action: fun,
		va:     va,
	})
}

// PushCallVoidDelegate 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
func (rt *RuntimeBehavior) PushCallVoidDelegate(fun generic.DelegateActionVar0[any], va ...any) runtime.AsyncRet {
	return rt.pushCallTask(_Task{
		delegateAction: fun,
		va:             va,
	})
}

func (rt *RuntimeBehavior) pushCallTask(task _Task) (asyncRet chan runtime.Ret) {
	task.kind = _TaskKind_Call
	task.asyncRet = concurrent.MakeAsyncRet()

	asyncRet = task.asyncRet

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			asyncRet <- runtime.MakeRet(nil, ErrProcessQueueClosed)
			close(asyncRet)
		}
	}()

	if rt.opts.ProcessQueueTimeout > 0 {
		timeoutTimer := time.NewTimer(rt.opts.ProcessQueueTimeout)
		defer timeoutTimer.Stop()

		select {
		case rt.processQueue <- task:
			return
		case <-timeoutTimer.C:
			break
		}
	} else {
		select {
		case rt.processQueue <- task:
			return
		default:
			break
		}
	}

	asyncRet <- runtime.MakeRet(nil, ErrProcessQueueFull)
	close(asyncRet)

	return
}
