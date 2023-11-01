package internal

import (
	"fmt"
	"kit.golaxy.org/golaxy/util/types"
	"runtime"
)

// Call 调用外部代码，有返回值
func Call[T any](autoRecover bool, reportError chan error, fun func() T) (ret T, panicErr error) {
	if fun == nil {
		return types.Zero[T](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	return fun(), nil
}

// CallVoid 调用外部代码，没有返回值
func CallVoid(autoRecover bool, reportError chan error, fun func()) (panicErr error) {
	if fun == nil {
		return nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	fun()

	return
}

func PrintStackTrace(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("%w\nstack: %s\n", err, stackBuf[:n])
}
