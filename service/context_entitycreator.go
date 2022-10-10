package service

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	Prototype     string                  // 实体原型名称
	EntityOptions []ec.EntityOptionSetter // 创建实体的所有选项
}

// EntityOption 创建实体的选项
var EntityOption = &_EntityOption{}

// EntityOptionSetter 创建实体的选项设置器
type EntityOptionSetter func(o *EntityOptions)

type _EntityOption struct{}

// Default 默认值
func (*_EntityOption) Default() EntityOptionSetter {
	return func(o *EntityOptions) {
		o.Prototype = ""
	}
}

// _EntityCreator 实体创建器
type _EntityCreator interface {
	NewEntity(runtimeCtx runtime.Context, optSetter ...ec.EntityOptionSetter) (ec.Entity, error)
}

func (ctx *ContextBehavior) NewEntity(runtimeCtx runtime.Context, optSetter ...ec.EntityOptionSetter) (ec.Entity, error) {

}
