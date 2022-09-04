package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// EntityOptions 创建实体（Entity）的所有选项
type EntityOptions struct {
	Inheritor                    Face[Entity]              // 继承者，需要拓展实体（Entity）自身功能时需要使用
	Prototype                    string                    // 实体（Entity）原型
	FaceCache                    *container.Cache[FaceAny] // FaceCache用于提高性能，通常传入运行时上下文（Runtime Context）选项中的FaceCache
	HookCache                    *container.Cache[Hook]    // HookCache用于提高性能，通常传入运行时上下文（Runtime Context）选项中的HookCache
	EnableRemovePrimaryComponent bool                      // 主要组件能否被删除
}

// EntityOptionSetter 实体（Entity）选项设置器
var EntityOptionSetter = &_EntityOptionSetter{}

type _EntityOptionSetterFunc func(o *EntityOptions)

type _EntityOptionSetter struct{}

// Default 默认值
func (*_EntityOptionSetter) Default() _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.Inheritor = Face[Entity]{}
		o.Prototype = ""
		o.FaceCache = nil
		o.HookCache = nil
		o.EnableRemovePrimaryComponent = false
	}
}

// Inheritor 继承者，需要拓展实体自身功能时需要使用
func (*_EntityOptionSetter) Inheritor(v Face[Entity]) _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.Inheritor = v
	}
}

// Prototype 实体（Entity）原型
func (*_EntityOptionSetter) Prototype(v string) _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.Prototype = v
	}
}

// FaceCache FaceCache用于提高性能，通常传入运行时上下文（Runtime Context）选项中的FaceCache
func (*_EntityOptionSetter) FaceCache(v *container.Cache[FaceAny]) _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.FaceCache = v
	}
}

// HookCache HookCache用于提高性能，通常传入运行时上下文（Runtime Context）选项中的HookCache
func (*_EntityOptionSetter) HookCache(v *container.Cache[Hook]) _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.HookCache = v
	}
}

// EnableRemovePrimaryComponent 主要组件能否被删除
func (*_EntityOptionSetter) EnableRemovePrimaryComponent(v bool) _EntityOptionSetterFunc {
	return func(o *EntityOptions) {
		o.EnableRemovePrimaryComponent = v
	}
}
