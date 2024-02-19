package ec

import (
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	CompositeFace      iface.Face[Entity]                 // 扩展者，在扩展实体自身能力时使用
	Prototype          string                             // 实体原型名称
	Scope              Scope                              // 可访问作用域
	PersistId          uid.Id                             // 实体持久化Id
	AwakeOnFirstAccess bool                               // 开启组件被首次访问时，检测并调用Awake()
	Meta               Meta                               // Meta信息
	FaceAnyAllocator   container.Allocator[iface.FaceAny] // 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
	HookAllocator      container.Allocator[event.Hook]    // 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
	GCCollector        container.GCCollector              // 自定义GC收集器，通常不传或者传入运行时上下文
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		With.CompositeFace(iface.Face[Entity]{})(o)
		With.Prototype("")(o)
		With.Scope(Scope_Local)(o)
		With.PersistId(uid.Nil)(o)
		With.AwakeOnFirstAccess(true)(o)
		With.Meta(nil)(o)
		With.FaceAnyAllocator(container.DefaultAllocator[iface.FaceAny]())(o)
		With.HookAllocator(container.DefaultAllocator[event.Hook]())(o)
		With.GCCollector(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (_Option) CompositeFace(face iface.Face[Entity]) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.CompositeFace = face
	}
}

// Prototype 实体原型名称
func (_Option) Prototype(pt string) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Prototype = pt
	}
}

// Scope 可访问作用域
func (_Option) Scope(scope Scope) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}

// PersistId 实体持久化Id
func (_Option) PersistId(id uid.Id) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.PersistId = id
	}
}

// AwakeOnFirstAccess 开启组件被首次访问时，检测并调用Awake()
func (_Option) AwakeOnFirstAccess(b bool) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.AwakeOnFirstAccess = b
	}
}

// Meta Meta信息
func (_Option) Meta(m Meta) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.Meta = m
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (_Option) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrEC, exception.ErrArgs))
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (_Option) HookAllocator(allocator container.Allocator[event.Hook]) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrEC, exception.ErrArgs))
		}
		o.HookAllocator = allocator
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (_Option) GCCollector(collector container.GCCollector) option.Setting[EntityOptions] {
	return func(o *EntityOptions) {
		o.GCCollector = collector
	}
}
