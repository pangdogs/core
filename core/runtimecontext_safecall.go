package core

import (
	"fmt"
)

type _Callee interface {
	pushCall(segment func())
}

// _SafeCall 安全调用
type _SafeCall interface {
	// SafeCall 在运行时（Runtime）中，将代码片段插入任务流水线，安全的进行调用，返回result channel，可以选择阻塞并等待返回结果
	SafeCall(segment func() SafeRet) <-chan SafeRet

	// SafeCallNoRet 在运行时（Runtime）中，将代码片段插入任务流水线，安全的进行调用，没有返回值，无法阻塞
	SafeCallNoRet(segment func())

	setCallee(callee _Callee)
}

// SafeRet 安全调用结果
type SafeRet struct {
	Err error       // error
	Ret interface{} // 结果
}

// SafeCall 在运行时（Runtime）中，将代码片段插入任务流水线，安全的进行调用，返回result channel，可以选择阻塞并等待返回结果
func (runtimeCtx *_RuntimeContextBehavior) SafeCall(segment func() SafeRet) <-chan SafeRet {
	if segment == nil {
		panic("nil segment")
	}

	ret := make(chan SafeRet, 1)

	runtimeCtx.callee.pushCall(func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				ret <- SafeRet{Err: err}
				panic(err)
			}
		}()

		ret <- segment()
	})

	return ret
}

// SafeCallNoRet 在运行时（Runtime）中，将代码片段插入任务流水线，安全的进行调用，没有返回值，无法阻塞
func (runtimeCtx *_RuntimeContextBehavior) SafeCallNoRet(segment func()) {
	if segment == nil {
		panic("nil segment")
	}

	runtimeCtx.callee.pushCall(segment)
}

func (runtimeCtx *_RuntimeContextBehavior) setCallee(callee _Callee) {
	runtimeCtx.callee = callee
}
