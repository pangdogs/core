package service

import (
	"github.com/golaxy-kit/golaxy/pt"
)

// GetEntityPt 获取实体原型
func (ctx *ContextBehavior) GetEntityPt(prototype string) (pt.EntityPt, bool) {
	entityLib := ctx.getOptions().EntityLib
	if entityLib == nil {
		return pt.EntityPt{}, false
	}

	return entityLib.Get(prototype)
}
