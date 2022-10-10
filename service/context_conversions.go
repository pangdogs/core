package service

import "github.com/pangdogs/galaxy/util"

// Inheritor 获取服务上下文的继承者
func Inheritor[T any](ctx Context) T {
	return util.Cache2Iface[T](ctx.getOptions().Inheritor.Cache)
}
