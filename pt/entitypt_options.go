package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// WithOption 所有选项设置器
type WithOption struct{}

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	AssignCompId func(entity ec.Entity, compPt ComponentPt) uid.Id // 设置组件Id函数
	Scope        ec.Scope                                          // 实体的可访问作用域
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// Default 默认值
func (WithOption) Default() EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.Default()(&o.EntityOptions)
		WithOption{}.AssignCompId(nil)(o)
		WithOption{}.Scope(ec.Scope_Local)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// Prototype 实体原型名称
func (WithOption) Prototype(pt string) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.Prototype(pt)(&o.EntityOptions)
	}
}

// PersistId 实体持久化Id
func (WithOption) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.PersistId(id)(&o.EntityOptions)
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) EnableComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.EnableComponentAwakeByAccess(b)(&o.EntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// AssignCompId 设置组件Id函数
func (WithOption) AssignCompId(fn func(entity ec.Entity, compPt ComponentPt) uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.AssignCompId = fn
	}
}

// Scope 实体的可访问作用域
func (WithOption) Scope(scope ec.Scope) EntityOption {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}
