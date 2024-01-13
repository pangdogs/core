package core

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

var (
	ErrGolaxy   = exception.ErrGolaxy                  // GOLAXY框架错误
	ErrPanicked = exception.ErrPanicked                // panic错误
	ErrArgs     = exception.ErrArgs                    // 参数错误
	ErrRuntime  = fmt.Errorf("%w: runtime", ErrGolaxy) // 运行时错误
	ErrService  = fmt.Errorf("%w: service", ErrGolaxy) // 服务错误
)
