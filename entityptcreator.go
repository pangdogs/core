package core

import (
	"fmt"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/service"
)

// CreateEntityPT 创建实体原型
func CreateEntityPT(ctx service.Context) EntityPTCreator {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrCore, ErrArgs))
	}
	return EntityPTCreator{
		servCtx: ctx,
	}
}

// EntityPTCreator 实体原型构建器
type EntityPTCreator struct {
	servCtx   service.Context
	prototype string
	comps     []any
}

// Prototype 设置实体原型名称
func (c EntityPTCreator) Prototype(prototype string) EntityPTCreator {
	c.prototype = prototype
	return c
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
		panic(fmt.Errorf("%w: setting servCtx is nil", ErrCore))
	}
	c.servCtx.GetEntityLib().Register(c.prototype, c.comps...)
}
