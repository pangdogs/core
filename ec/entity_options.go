package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// WithOption 所有选项设置器
type WithOption struct{}

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	CompositeFace          util.Face[Entity]                    // 扩展者，在扩展实体自身能力时使用
	Prototype              string                               // 实体原型名称
	PersistId              uid.Id                               // 实体持久化Id
	ComponentAwakeByAccess bool                                 // 开启组件被访问时，检测并调用Awake()
	FaceAnyAllocator       container.Allocator[util.FaceAny]    // 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
	HookAllocator          container.Allocator[localevent.Hook] // 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
	GCCollector            container.GCCollector                // 自定义GC收集器，通常不传或者传入运行时上下文
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// Default 默认值
func (WithOption) Default() EntityOption {
	return func(o *EntityOptions) {
		WithOption{}.CompositeFace(util.Face[Entity]{})(o)
		WithOption{}.Prototype("")(o)
		WithOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithOption{}.ComponentAwakeByAccess(true)(o)
		WithOption{}.FaceAnyAllocator(container.DefaultAllocator[util.FaceAny]())(o)
		WithOption{}.HookAllocator(container.DefaultAllocator[localevent.Hook]())(o)
		WithOption{}.GCCollector(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[Entity]) EntityOption {
	return func(o *EntityOptions) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (WithOption) Prototype(pt string) EntityOption {
	return func(o *EntityOptions) {
		o.Prototype = pt
	}
}

// PersistId 实体持久化Id
func (WithOption) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.PersistId = id
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) ComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		o.ComponentAwakeByAccess = b
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.HookAllocator = allocator
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithOption) GCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		o.GCCollector = collector
	}
}
