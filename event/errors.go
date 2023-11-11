package event

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

var (
	ErrEvent = fmt.Errorf("%w: event", exception.ErrGolaxy) // 事件错误
	ErrArgs  = exception.ErrArgs                            // 参数错误
)
