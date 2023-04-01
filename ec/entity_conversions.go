package ec

import (
	"kit.golaxy.org/golaxy/util"
)

// GetComposite 获取实体的扩展者
func GetComposite[T any](entity Entity) T {
	return util.Cache2Iface[T](entity.getOptions().CompositeFace.Cache)
}
