package event

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrEvent = fmt.Errorf("%w: event", internal.ErrGolaxy) // 事件错误
	ErrArgs  = internal.ErrArgs                            // 参数错误
)
