package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", exception.ErrGolaxy) // 插件错误
)
