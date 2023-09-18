package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", internal.ErrGolaxy) // 插件错误
)
