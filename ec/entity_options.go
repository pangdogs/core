package ec

import (
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/golaxy-kit/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	Inheritor                    util.Face[Entity]                 // 继承者，在扩展实体自身能力时使用
	Prototype                    string                            // 实体原型名称
	PersistID                    ID                                // 实体持久化ID
	EnableComponentAwakeByAccess bool                              // 开启组件被访问时，检测并调用Awake()
	FaceCache                    *container.Cache[util.FaceAny]    // FaceCache用于提高性能，通常传入运行时上下文选项中的FaceCache
	HookCache                    *container.Cache[localevent.Hook] // HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// WithEntityOption 创建实体的所有选项设置器
type WithEntityOption struct{}

// Default 默认值
func (WithEntityOption) Default() EntityOption {
	return func(o *EntityOptions) {
		o.Inheritor = util.Face[Entity]{}
		o.Prototype = ""
		o.PersistID = util.Zero[ID]()
		o.EnableComponentAwakeByAccess = true
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，在扩展实体自身能力时使用
func (WithEntityOption) Inheritor(v util.Face[Entity]) EntityOption {
	return func(o *EntityOptions) {
		o.Inheritor = v
	}
}

// Prototype 实体原型名称
func (WithEntityOption) Prototype(v string) EntityOption {
	return func(o *EntityOptions) {
		o.Prototype = v
	}
}

// PersistID 实体持久化ID
func (WithEntityOption) PersistID(v ID) EntityOption {
	return func(o *EntityOptions) {
		o.PersistID = v
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithEntityOption) EnableComponentAwakeByAccess(v bool) EntityOption {
	return func(o *EntityOptions) {
		o.EnableComponentAwakeByAccess = v
	}
}

// FaceCache FaceCache用于提高性能，通常传入运行时上下文选项中的FaceCache
func (WithEntityOption) FaceCache(v *container.Cache[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		o.FaceCache = v
	}
}

// HookCache HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
func (WithEntityOption) HookCache(v *container.Cache[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		o.HookCache = v
	}
}
