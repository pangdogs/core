package ec

import (
	"kit.golaxy.org/golaxy/util/iface"
)

// GetComposite 获取实体的扩展者
func GetComposite[T any](entity Entity) T {
	return iface.Cache2Iface[T](entity.getOptions().CompositeFace.Cache)
}
