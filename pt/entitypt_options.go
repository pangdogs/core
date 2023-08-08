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

type (
	ComponentConstructor = func(entity ec.Entity, comp ec.Component) // 组件构造函数
	EntityConstructor    = func(entity ec.Entity)                    // 实体构造函数
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	ComponentConstructor ComponentConstructor // 组件构造函数
	EntityConstructor    EntityConstructor    // 实体构造函数
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// Default 默认值
func (WithOption) Default() EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.Default()(&o.EntityOptions)
		WithOption{}.ComponentConstructor(nil)
		WithOption{}.EntityConstructor(nil)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// PersistId 实体持久化Id
func (WithOption) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.PersistId(id)(&o.EntityOptions)
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) ComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.ComponentAwakeByAccess(b)(&o.EntityOptions)
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

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithOption) GCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		ec.WithOption{}.GCCollector(collector)(&o.EntityOptions)
	}
}

// ComponentConstructor 组件构造函数
func (WithOption) ComponentConstructor(fn ComponentConstructor) EntityOption {
	return func(o *EntityOptions) {
		o.ComponentConstructor = fn
	}
}

// EntityConstructor 实体构造函数
func (WithOption) EntityConstructor(fn EntityConstructor) EntityOption {
	return func(o *EntityOptions) {
		o.EntityConstructor = fn
	}
}
