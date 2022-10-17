package ec

import (
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	Inheritor                    util.Face[Entity]                 // 继承者，在拓展实体自身能力时使用
	Prototype                    string                            // 实体原型名称
	PersistID                    int64                             // 实体持久化ID
	ComponentPersistID           func(comp Component) int64        // 组件持久化ID
	EnableRemovePrimaryComponent bool                              // 开启主要组件可以被删除，主要组件是指实体加入运行时上下文前添加的组件
	EnableComponentAwakeByAccess bool                              // 开启组件被访问时，检测并调用Awake()
	FaceCache                    *container.Cache[util.FaceAny]    // FaceCache用于提高性能，通常传入运行时上下文选项中的FaceCache
	HookCache                    *container.Cache[localevent.Hook] // HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
}

// EntityOption 创建实体的选项
var EntityOption = &_EntityOption{}

// EntityOptionSetter 创建实体的选项设置器
type EntityOptionSetter func(o *EntityOptions)

type _EntityOption struct{}

// Default 默认值
func (*_EntityOption) Default() EntityOptionSetter {
	return func(o *EntityOptions) {
		o.Inheritor = util.Face[Entity]{}
		o.Prototype = ""
		o.PersistID = 0
		o.ComponentPersistID = nil
		o.EnableRemovePrimaryComponent = false
		o.EnableComponentAwakeByAccess = true
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，在拓展实体自身能力时使用
func (*_EntityOption) Inheritor(v util.Face[Entity]) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.Inheritor = v
	}
}

// Prototype 实体原型名称
func (*_EntityOption) Prototype(v string) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.Prototype = v
	}
}

// PersistID 实体持久化ID
func (*_EntityOption) PersistID(v int64) EntityOptionSetter {
	return func(o *EntityOptions) {
		if v < 0 {
			panic("persistID less 0 invalid")
		}
		o.PersistID = v
	}
}

// ComponentPersistID 组件持久化ID
func (*_EntityOption) ComponentPersistID(v func(comp Component) int64) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.ComponentPersistID = v
	}
}

// EnableRemovePrimaryComponent 开启主要组件可以被删除，主要组件是指实体加入运行时上下文前添加的组件
func (*_EntityOption) EnableRemovePrimaryComponent(v bool) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.EnableRemovePrimaryComponent = v
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (*_EntityOption) EnableComponentAwakeByAccess(v bool) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.EnableComponentAwakeByAccess = v
	}
}

// FaceCache FaceCache用于提高性能，通常传入运行时上下文选项中的FaceCache
func (*_EntityOption) FaceCache(v *container.Cache[util.FaceAny]) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.FaceCache = v
	}
}

// HookCache HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
func (*_EntityOption) HookCache(v *container.Cache[localevent.Hook]) EntityOptionSetter {
	return func(o *EntityOptions) {
		o.HookCache = v
	}
}
