package runtime

import (
	"git.golaxy.org/core/util/iface"
)

// GetComposite 获取运行时上下文的扩展者
func GetComposite[T any](ctx Context) T {
	return iface.Cache2Iface[T](ctx.getOptions().CompositeFace.Cache)
}
