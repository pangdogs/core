package pt

import (
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/golaxy-kit/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	AssignCompID  func(entity ec.Entity, compPt ComponentPt) ec.ID // 设置组件ID函数
	Accessibility ec.Accessibility                                 // 实体的可访问性
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// WithEntityOption 创建实体的所有选项设置器
type WithEntityOption struct {
	ec.WithEntityOption
}

// Default 默认值
func (WithEntityOption) Default() EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.Default()(&o.EntityOptions)
		WithEntityOption{}.AssignCompID(nil)(o)
		WithEntityOption{}.Accessibility(ec.Local)(o)
	}
}

// Inheritor 继承者，在扩展实体自身能力时使用
func (WithEntityOption) Inheritor(v util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.Inheritor(v)(&o.EntityOptions)
	}
}

// Prototype 实体原型名称
func (WithEntityOption) Prototype(v string) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.Prototype(v)(&o.EntityOptions)
	}
}

// PersistID 实体持久化ID
func (WithEntityOption) PersistID(v ec.ID) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.PersistID(v)(&o.EntityOptions)
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithEntityOption) EnableComponentAwakeByAccess(v bool) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.EnableComponentAwakeByAccess(v)(&o.EntityOptions)
	}
}

// FaceCache FaceCache用于提高性能，通常传入运行时上下文选项中的FaceCache
func (WithEntityOption) FaceCache(v *container.Cache[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.FaceCache(v)(&o.EntityOptions)
	}
}

// HookCache HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
func (WithEntityOption) HookCache(v *container.Cache[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.HookCache(v)(&o.EntityOptions)
	}
}

// AssignCompID 设置组件ID函数
func (WithEntityOption) AssignCompID(v func(entity ec.Entity, compPt ComponentPt) ec.ID) EntityOption {
	return func(o *EntityOptions) {
		o.AssignCompID = v
	}
}

// Accessibility 实体的可访问性
func (WithEntityOption) Accessibility(v ec.Accessibility) EntityOption {
	return func(o *EntityOptions) {
		o.Accessibility = v
	}
}
