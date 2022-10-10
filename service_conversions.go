package galaxy

import (
	"github.com/pangdogs/galaxy/util"
)

// ServiceInheritor 获取服务的继承者
func ServiceInheritor[T any](service Service) T {
	return util.Cache2Iface[T](service.getOptions().Inheritor.Cache)
}
