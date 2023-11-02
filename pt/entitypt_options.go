package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"kit.golaxy.org/golaxy/util/uid"
)

// Option 所有选项设置器
type Option struct{}

type (
	ComponentCtor = generic.Action2[ec.Entity, ec.Component] // 组件构造函数
	EntityCtor    = generic.Action1[ec.Entity]               // 实体构造函数
)

// ConstructEntityOptions 创建实体的所有选项
type ConstructEntityOptions struct {
	ec.EntityOptions
	ComponentCtors []ComponentCtor // 组件构造函数
	EntityCtors    []EntityCtor    // 实体构造函数
}

// Default 默认值
func (Option) Default() option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.Default()(&o.EntityOptions)
		Option{}.ComponentCtors(nil)
		Option{}.EntityCtors(nil)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (Option) CompositeFace(face iface.Face[ec.Entity]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// PersistId 实体持久化Id
func (Option) PersistId(id uid.Id) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.PersistId(id)(&o.EntityOptions)
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (Option) ComponentAwakeByAccess(b bool) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.ComponentAwakeByAccess(b)(&o.EntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (Option) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (Option) HookAllocator(allocator container.Allocator[event.Hook]) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (Option) GCCollector(collector container.GCCollector) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		ec.Option{}.GCCollector(collector)(&o.EntityOptions)
	}
}

// ComponentCtors 组件构造函数
func (Option) ComponentCtors(ctors []ComponentCtor) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		o.ComponentCtors = ctors
	}
}

// EntityCtors 实体构造函数
func (Option) EntityCtors(ctors []EntityCtor) option.Setting[ConstructEntityOptions] {
	return func(o *ConstructEntityOptions) {
		o.EntityCtors = ctors
	}
}
