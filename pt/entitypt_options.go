package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	AssignCompId func(entity ec.Entity, compPt ComponentPt) uid.Id // 设置组件Id函数
	Scope        ec.Scope                                          // 实体的可访问作用域
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
		WithEntityOption{}.AssignCompId(nil)(o)
		WithEntityOption{}.Scope(ec.Scope_Local)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithEntityOption) CompositeFace(face util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// Prototype 实体原型名称
func (WithEntityOption) Prototype(pt string) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.Prototype(pt)(&o.EntityOptions)
	}
}

// PersistId 实体持久化Id
func (WithEntityOption) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.PersistId(id)(&o.EntityOptions)
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithEntityOption) EnableComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.EnableComponentAwakeByAccess(b)(&o.EntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithEntityOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithEntityOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// AssignCompId 设置组件Id函数
func (WithEntityOption) AssignCompId(fn func(entity ec.Entity, compPt ComponentPt) uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.AssignCompId = fn
	}
}

// Scope 实体的可访问作用域
func (WithEntityOption) Scope(scope ec.Scope) EntityOption {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}
