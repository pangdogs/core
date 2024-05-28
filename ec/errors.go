package ec

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

var (
	ErrEC = fmt.Errorf("%w: ec", exception.ErrCore) // EC错误
)
