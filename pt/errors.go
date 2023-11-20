package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

var (
	ErrPt = fmt.Errorf("%w: pt", exception.ErrGolaxy) // 原型错误
)
