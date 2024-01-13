package plugin

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", exception.ErrGolaxy) // 插件错误
)
