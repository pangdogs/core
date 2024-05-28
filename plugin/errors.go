package plugin

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", exception.ErrCore) // 插件错误
)
