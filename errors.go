package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrGolaxy   = internal.ErrGolaxy                   // GOLAXY框架错误
	ErrPanicked = internal.ErrPanicked                 // panic错误
	ErrArgs     = internal.ErrArgs                     // 参数类错误
	ErrRuntime  = fmt.Errorf("%w: runtime", ErrGolaxy) // 运行时错误
	ErrService  = fmt.Errorf("%w: service", ErrGolaxy) // 服务错误
)
