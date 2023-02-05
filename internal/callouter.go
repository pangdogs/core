package internal

import (
	"fmt"
	"kit.golaxy.org/golaxy/util"
	"runtime"
)

// CallOuter 调用外部代码，有返回值
func CallOuter[T any](autoRecover bool, reportError chan error, fun func() T) (ret T, exception error) {
	if fun == nil {
		return util.Zero[T](), nil
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

// CallOuterNoRet 调用外部代码，没有返回值
func CallOuterNoRet(autoRecover bool, reportError chan error, fun func()) (exception error) {
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

func errorAddStackTrace(info any) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("Error: %v\nStack: %s\n", info, stackBuf[:n])
}
