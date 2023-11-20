package service

import (
	"kit.golaxy.org/golaxy/pt"
)

// GetEntityLib 获取实体原型库
func (ctx *ContextBehavior) GetEntityLib() pt.EntityLib {
	return ctx.opts.EntityLib
}
