package core

import (
	"fmt"
	"runtime"
)

func callOuter[T any](autoRecover bool, reportError chan error, fun func() T) (ret T, exception error) {
	if fun == nil {
		return Zero[T](), nil
	}

	if autoRecover {
		defer func() {
			if info := recover(); info != nil {
				if reportError != nil {
					exception = errorAddStackTrace(info)
					select {
					case reportError <- exception:
					default:
					}
				}
			}
		}()
	}

	ret = fun()

	return
}

func callOuterNoRet(autoRecover bool, reportError chan error, fun func()) (exception error) {
	if fun == nil {
		return nil
	}

	if autoRecover {
		defer func() {
			if info := recover(); info != nil {
				if reportError != nil {
					exception = errorAddStackTrace(info)
					select {
					case reportError <- exception:
					default:
					}
				}
			}
		}()
	}

	fun()

	return
}

func errorAddStackTrace(info interface{}) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("Error: %v\nStack: %s\n", info, stackBuf[:n])
}
