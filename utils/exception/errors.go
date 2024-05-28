package exception

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrCore     = errors.New("core")     // 内核错误
	ErrPanicked = errors.New("panicked") // panic错误
	ErrArgs     = errors.New("args")     // 参数错误
)

func PrintStackTrace(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("error => %w\nstack => %s\n", err, stackBuf[:n])
}
