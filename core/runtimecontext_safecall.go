package core

import (
	"fmt"
)

func (runtimeCtx *RuntimeContextBehavior) SafeCall(segment func() SafeRet) <-chan SafeRet {
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

func (runtimeCtx *RuntimeContextBehavior) SafeCallNoRet(segment func()) {
	if segment == nil {
		panic("nil segment")
	}

	runtimeCtx.callee.pushCall(segment)
}

func (runtimeCtx *RuntimeContextBehavior) setCallee(callee _Callee) {
	runtimeCtx.callee = callee
}
