package plugin

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", exception.ErrCore) // 插件错误
)
