package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrGolaxy   = errors.ErrGolaxy                     // GOLAXY框架错误
	ErrPanicked = errors.ErrPanicked                   // panic错误
	ErrArgs     = errors.ErrArgs                       // 参数错误
	ErrRuntime  = fmt.Errorf("%w: runtime", ErrGolaxy) // 运行时错误
	ErrService  = fmt.Errorf("%w: service", ErrGolaxy) // 服务错误
)
