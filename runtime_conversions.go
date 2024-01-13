package core

import (
	"git.golaxy.org/core/util/iface"
)

// GetRuntimeComposite 获取运行时的扩展者
func GetRuntimeComposite[T any](runtime Runtime) T {
	return iface.Cache2Iface[T](runtime.getOptions().CompositeFace.Cache)
}
