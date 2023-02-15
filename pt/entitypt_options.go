package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	AssignCompID func(entity ec.Entity, compPt ComponentPt) ec.ID // 设置组件ID函数
	Scope        ec.Scope                                         // 实体的可访问作用域
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
		WithEntityOption{}.Scope(ec.Scope_Local)(o)
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
func (WithEntityOption) FaceCache(v container.Cache[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.FaceCache(v)(&o.EntityOptions)
	}
}

// HookCache HookCache用于提高性能，通常传入运行时上下文选项中的HookCache
func (WithEntityOption) HookCache(v container.Cache[localevent.Hook]) EntityOption {
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

// Scope 实体的可访问作用域
func (WithEntityOption) Scope(v ec.Scope) EntityOption {
	return func(o *EntityOptions) {
		o.Scope = v
	}
}
