package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

var (
	ErrEC = fmt.Errorf("%w: ec", exception.ErrGolaxy) // EC错误
)
