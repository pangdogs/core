package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	CompositeFace                util.Face[Entity]                    // 扩展者，在扩展实体自身能力时使用
	Prototype                    string                               // 实体原型名称
	PersistId                    uid.Id                               // 实体持久化Id
	EnableComponentAwakeByAccess bool                                 // 开启组件被访问时，检测并调用Awake()
	FaceAnyAllocator             container.Allocator[util.FaceAny]    // 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
	HookAllocator                container.Allocator[localevent.Hook] // 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
	GCCollector                  container.GCCollector                // 自定义GC收集器，通常不传或者传入运行时上下文
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// WithEntityOption 创建实体的所有选项设置器
type WithEntityOption struct{}

// Default 默认值
func (WithEntityOption) Default() EntityOption {
	return func(o *EntityOptions) {
		WithEntityOption{}.CompositeFace(util.Face[Entity]{})(o)
		WithEntityOption{}.Prototype("")(o)
		WithEntityOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithEntityOption{}.EnableComponentAwakeByAccess(true)(o)
		WithEntityOption{}.FaceAnyAllocator(container.DefaultAllocator[util.FaceAny]())(o)
		WithEntityOption{}.HookAllocator(container.DefaultAllocator[localevent.Hook]())(o)
		WithEntityOption{}.GCCollector(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithEntityOption) CompositeFace(face util.Face[Entity]) EntityOption {
	return func(o *EntityOptions) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (WithEntityOption) Prototype(pt string) EntityOption {
	return func(o *EntityOptions) {
		o.Prototype = pt
	}
}

// PersistId 实体持久化Id
func (WithEntityOption) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.PersistId = id
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithEntityOption) EnableComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		o.EnableComponentAwakeByAccess = b
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithEntityOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithEntityOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.HookAllocator = allocator
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithEntityOption) GCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		o.GCCollector = collector
	}
}
