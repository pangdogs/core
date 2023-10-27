package event

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrEvent = fmt.Errorf("%w: event", errors.ErrGolaxy) // 事件错误
	ErrArgs  = errors.ErrArgs                            // 参数错误
)
