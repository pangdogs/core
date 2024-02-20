package core

import (
	"fmt"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/service"
)

// CreateEntityPT 创建实体原型
func CreateEntityPT(ctx service.Context, prototype string) EntityPTCreator {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrGolaxy, ErrArgs))
	}
	return EntityPTCreator{
		servCtx:   ctx,
		prototype: prototype,
	}
}

// EntityPTCreator 实体原型构建器
type EntityPTCreator struct {
	servCtx   service.Context
	prototype string
	comps     []any
}

// AddComponent 添加组件
func (c EntityPTCreator) AddComponent(comp any, alias ...string) EntityPTCreator {
	if len(alias) > 0 {
		c.comps = append(c.comps, pt.CompAlias(comp, alias[0]))
	} else {
		c.comps = append(c.comps, comp)
	}
	return c
}

// Declare 声明实体原型
func (c EntityPTCreator) Declare() {
	if c.servCtx == nil {
		panic(fmt.Errorf("%w: setting servCtx is nil", ErrGolaxy))
	}
	c.servCtx.GetEntityLib().Register(c.prototype, c.comps...)
}
