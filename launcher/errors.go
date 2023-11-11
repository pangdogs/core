package launcher

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

var (
	ErrApp = fmt.Errorf("%w: app", exception.ErrGolaxy)
)
