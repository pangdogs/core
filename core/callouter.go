package core

import (
	"fmt"
	"runtime"
)

// CallOuter 调用外部逻辑，有返回值
func CallOuter[T any](autoRecover bool, reportError chan error, fun func() T) (ret T, exception error) {
	if fun == nil {
		return Zero[T](), nil
	}

	if autoRecover {
		defer func() {
			if info := recover(); info != nil {
				if reportError != nil {
					exception = ErrorAddStackTrace(info)
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

// CallOuterNoRet 调用外部逻辑，无返回值
func CallOuterNoRet(autoRecover bool, reportError chan error, fun func()) (exception error) {
	if fun == nil {
		return nil
	}

	if autoRecover {
		defer func() {
			if info := recover(); info != nil {
				if reportError != nil {
					exception = ErrorAddStackTrace(info)
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

// ErrorAddStackTrace 获取错误堆栈信息
func ErrorAddStackTrace(info interface{}) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("Error: %v\nStack: %s\n", info, stackBuf[:n])
}
