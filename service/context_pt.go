package service

import (
	"kit.golaxy.org/golaxy/pt"
)

// ResolveEntityPt 解析实体原型
func (ctx *ContextBehavior) ResolveEntityPt(prototype string) (pt.EntityPt, bool) {
	entityLib := ctx.getOptions().EntityLib
	if entityLib == nil {
		return pt.EntityPt{}, false
	}

	return entityLib.Get(prototype)
}
