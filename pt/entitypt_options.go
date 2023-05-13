package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// Options 创建实体的所有选项
type Options struct {
	ec.Options
	AssignCompId func(entity ec.Entity, compPt ComponentPt) uid.Id // 设置组件Id函数
	Scope        ec.Scope                                          // 实体的可访问作用域
}

// Option 创建实体的选项设置器
type Option func(o *Options)

// WithOption 创建实体的所有选项设置器
type WithOption struct {
	ec.WithOption
}

// Default 默认值
func (WithOption) Default() Option {
	return func(o *Options) {
		ec.WithOption{}.Default()(&o.Options)
		WithOption{}.AssignCompId(nil)(o)
		WithOption{}.Scope(ec.Scope_Local)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[ec.Entity]) Option {
	return func(o *Options) {
		ec.WithOption{}.CompositeFace(face)(&o.Options)
	}
}

// Prototype 实体原型名称
func (WithOption) Prototype(pt string) Option {
	return func(o *Options) {
		ec.WithOption{}.Prototype(pt)(&o.Options)
	}
}

// PersistId 实体持久化Id
func (WithOption) PersistId(id uid.Id) Option {
	return func(o *Options) {
		ec.WithOption{}.PersistId(id)(&o.Options)
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) EnableComponentAwakeByAccess(b bool) Option {
	return func(o *Options) {
		ec.WithOption{}.EnableComponentAwakeByAccess(b)(&o.Options)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) Option {
	return func(o *Options) {
		ec.WithOption{}.FaceAnyAllocator(allocator)(&o.Options)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) HookAllocator(allocator container.Allocator[localevent.Hook]) Option {
	return func(o *Options) {
		ec.WithOption{}.HookAllocator(allocator)(&o.Options)
	}
}

// AssignCompId 设置组件Id函数
func (WithOption) AssignCompId(fn func(entity ec.Entity, compPt ComponentPt) uid.Id) Option {
	return func(o *Options) {
		o.AssignCompId = fn
	}
}

// Scope 实体的可访问作用域
func (WithOption) Scope(scope ec.Scope) Option {
	return func(o *Options) {
		o.Scope = scope
	}
}
