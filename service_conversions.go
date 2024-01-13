package core

import (
	"git.golaxy.org/core/util/iface"
)

// GetServiceComposite 获取服务的扩展者
func GetServiceComposite[T any](service Service) T {
	return iface.Cache2Iface[T](service.getOptions().CompositeFace.Cache)
}
