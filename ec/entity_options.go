package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// Options 创建实体的所有选项
type Options struct {
	CompositeFace                util.Face[Entity]                    // 扩展者，在扩展实体自身能力时使用
	Prototype                    string                               // 实体原型名称
	PersistId                    uid.Id                               // 实体持久化Id
	EnableComponentAwakeByAccess bool                                 // 开启组件被访问时，检测并调用Awake()
	FaceAnyAllocator             container.Allocator[util.FaceAny]    // 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
	HookAllocator                container.Allocator[localevent.Hook] // 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
	GCCollector                  container.GCCollector                // 自定义GC收集器，通常不传或者传入运行时上下文
}

// Option 创建实体的选项设置器
type Option func(o *Options)

// WithOption 创建实体的所有选项设置器
type WithOption struct{}

// Default 默认值
func (WithOption) Default() Option {
	return func(o *Options) {
		WithOption{}.CompositeFace(util.Face[Entity]{})(o)
		WithOption{}.Prototype("")(o)
		WithOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithOption{}.EnableComponentAwakeByAccess(true)(o)
		WithOption{}.FaceAnyAllocator(container.DefaultAllocator[util.FaceAny]())(o)
		WithOption{}.HookAllocator(container.DefaultAllocator[localevent.Hook]())(o)
		WithOption{}.GCCollector(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[Entity]) Option {
	return func(o *Options) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (WithOption) Prototype(pt string) Option {
	return func(o *Options) {
		o.Prototype = pt
	}
}

// PersistId 实体持久化Id
func (WithOption) PersistId(id uid.Id) Option {
	return func(o *Options) {
		o.PersistId = id
	}
}

// EnableComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) EnableComponentAwakeByAccess(b bool) Option {
	return func(o *Options) {
		o.EnableComponentAwakeByAccess = b
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) Option {
	return func(o *Options) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) HookAllocator(allocator container.Allocator[localevent.Hook]) Option {
	return func(o *Options) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.HookAllocator = allocator
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithOption) GCCollector(collector container.GCCollector) Option {
	return func(o *Options) {
		o.GCCollector = collector
	}
}
