package pt

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

var (
	ErrPt = fmt.Errorf("%w: pt", exception.ErrCore) // 原型错误
)
