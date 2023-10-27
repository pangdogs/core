package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrEC = fmt.Errorf("%w: ec", errors.ErrGolaxy) // EC错误
)
