package core

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

var (
	ErrCore     = exception.ErrCore                  // 内核错误
	ErrPanicked = exception.ErrPanicked              // panic错误
	ErrArgs     = exception.ErrArgs                  // 参数错误
	ErrRuntime  = fmt.Errorf("%w: runtime", ErrCore) // 运行时错误
	ErrService  = fmt.Errorf("%w: service", ErrCore) // 服务错误
)
