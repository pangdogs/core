package pt

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

var (
	ErrPt = fmt.Errorf("%w: pt", exception.ErrCore) // 原型错误
)
