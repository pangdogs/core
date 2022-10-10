package galaxy

import "github.com/pangdogs/galaxy/util"

// RuntimeInheritor 获取运行时的继承者
func RuntimeInheritor[T any](runtime Runtime) T {
	return util.Cache2Iface[T](runtime.getOptions().Inheritor.Cache)
}
