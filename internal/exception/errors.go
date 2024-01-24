package exception

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrGolaxy   = errors.New("golaxy")   // Golaxy框架错误
	ErrPanicked = errors.New("panicked") // panic错误
	ErrArgs     = errors.New("args")     // 参数错误
)

func PrintStackTrace(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("%w\nstack: %s\n", err, stackBuf[:n])
}
