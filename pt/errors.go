package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrPt = fmt.Errorf("%w: Pt", errors.ErrGolaxy) // 原型错误
)
