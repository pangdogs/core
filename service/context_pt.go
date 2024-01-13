package service

import (
	"git.golaxy.org/core/pt"
)

// GetEntityLib 获取实体原型库
func (ctx *ContextBehavior) GetEntityLib() pt.EntityLib {
	return ctx.opts.EntityLib
}
