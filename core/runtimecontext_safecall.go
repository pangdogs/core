package core

import (
	"fmt"
)

// SafeCall ...
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

// SafeCallNoRet ...
func (runtimeCtx *_RuntimeContextBehavior) SafeCallNoRet(segment func()) {
	if segment == nil {
		panic("nil segment")
	}

	runtimeCtx.callee.pushCall(segment)
}

func (runtimeCtx *_RuntimeContextBehavior) setCallee(callee _Callee) {
	runtimeCtx.callee = callee
}
