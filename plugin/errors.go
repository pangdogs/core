package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrPlugin = fmt.Errorf("%w: plugin", errors.ErrGolaxy) // 插件错误
)
