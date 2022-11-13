package galaxy

import "github.com/galaxy-kit/galaxy/util"

// GetRuntimeInheritor 获取运行时的继承者
func GetRuntimeInheritor[T any](runtime Runtime) T {
	return util.Cache2Iface[T](runtime.getOptions().Inheritor.Cache)
}
