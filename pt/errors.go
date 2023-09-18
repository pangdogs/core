package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrPt = fmt.Errorf("%w: Pt", internal.ErrGolaxy) // 原型错误
)
