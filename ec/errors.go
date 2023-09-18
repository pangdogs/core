package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrEC = fmt.Errorf("%w: ec", internal.ErrGolaxy) // EC错误
)
