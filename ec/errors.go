package ec

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

var (
	ErrEC = fmt.Errorf("%w: ec", exception.ErrGolaxy) // EC错误
)
