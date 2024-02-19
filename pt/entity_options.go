package pt

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

type (
	ComponentCtor = generic.DelegateAction2[ec.Entity, ec.Component] // 组件构造函数
	EntityCtor    = generic.DelegateAction1[ec.Entity]               // 实体构造函数
)

// ConstructEntityOptions 创建实体的所有选项
type ConstructEntityOptions struct {
	ec.EntityOptions
	ComponentCtor ComponentCtor // 组件构造函数
	EntityCtor    EntityCtor    // 实体构造函数
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.Default()(&o.EntityOptions)
		With.ComponentCtor(nil)
		With.EntityCtor(nil)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (_Option) CompositeFace(face iface.Face[ec.Entity]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.CompositeFace(face)(&o.EntityOptions)
	}
}

// Prototype 实体原型名称
func (_Option) Prototype(pt string) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.Prototype(pt)(&o.EntityOptions)
	}
}

// Scope 可访问作用域
func (_Option) Scope(s ec.Scope) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.Scope(s)(&o.EntityOptions)
	}
}

// PersistId 实体持久化Id
func (_Option) PersistId(id uid.Id) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.PersistId(id)(&o.EntityOptions)
	}
}

// AwakeOnFirstAccess 开启组件被首次访问时，检测并调用Awake()
func (_Option) AwakeOnFirstAccess(b bool) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.AwakeOnFirstAccess(b)(&o.EntityOptions)
	}
}

// Meta Meta信息
func (_Option) Meta(m ec.Meta) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.Meta(m)(&o.EntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (_Option) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (_Option) HookAllocator(allocator container.Allocator[event.Hook]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (_Option) GCCollector(collector container.GCCollector) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.With.GCCollector(collector)(&o.EntityOptions)
	}
}

// ComponentCtor 组件构造函数
func (_Option) ComponentCtor(ctor ComponentCtor) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		o.ComponentCtor = ctor
	}
}

// EntityCtor 实体构造函数
func (_Option) EntityCtor(ctor EntityCtor) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		o.EntityCtor = ctor
	}
}
