package golaxy

import (
	"github.com/golaxy-kit/golaxy/util"
)

// GetServiceInheritor 获取服务的继承者
func GetServiceInheritor[T any](service Service) T {
	return util.Cache2Iface[T](service.getOptions().Inheritor.Cache)
}
