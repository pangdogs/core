package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// Option 所有选项设置器
type Option struct{}

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	CompositeFace          iface.Face[Entity]                 // 扩展者，在扩展实体自身能力时使用
	Prototype              string                             // 实体原型名称
	PersistId              uid.Id                             // 实体持久化Id
	ComponentAwakeByAccess bool                               // 开启组件被访问时，检测并调用Awake()
	FaceAnyAllocator       container.Allocator[iface.FaceAny] // 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
	HookAllocator          container.Allocator[event.Hook]    // 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
	GCCollector            container.GCCollector              // 自定义GC收集器，通常不传或者传入运行时上下文
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// Default 默认值
func (Option) Default() EntityOption {
	return func(o *EntityOptions) {
		Option{}.CompositeFace(iface.Face[Entity]{})(o)
		Option{}.Prototype("")(o)
		Option{}.PersistId(uid.Nil)(o)
		Option{}.ComponentAwakeByAccess(true)(o)
		Option{}.FaceAnyAllocator(container.DefaultAllocator[iface.FaceAny]())(o)
		Option{}.HookAllocator(container.DefaultAllocator[event.Hook]())(o)
		Option{}.GCCollector(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (Option) CompositeFace(face iface.Face[Entity]) EntityOption {
	return func(o *EntityOptions) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (Option) Prototype(pt string) EntityOption {
	return func(o *EntityOptions) {
		o.Prototype = pt
	}
}

// PersistId 实体持久化Id
func (Option) PersistId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.PersistId = id
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (Option) ComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		o.ComponentAwakeByAccess = b
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (Option) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrEC, internal.ErrArgs))
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (Option) HookAllocator(allocator container.Allocator[event.Hook]) EntityOption {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrEC, internal.ErrArgs))
		}
		o.HookAllocator = allocator
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (Option) GCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		o.GCCollector = collector
	}
}
