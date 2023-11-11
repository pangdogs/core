package golaxy

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"kit.golaxy.org/golaxy/util/uid"
)

type _EntityCreatorOption struct{}

// EntityCreatorOptions 实体构建器的所有选项
type EntityCreatorOptions struct {
	pt.ConstructEntityOptions
	ParentID uid.Id   // 父实体Id
	Scope    ec.Scope // 实体的可访问作用域
}

// Default 默认值
func (_EntityCreatorOption) Default() option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.Default()(&o.ConstructEntityOptions)
		_EntityCreatorOption{}.ParentId(uid.Nil)(o)
		_EntityCreatorOption{}.Scope(ec.Scope_Local)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (_EntityCreatorOption) CompositeFace(face iface.Face[ec.Entity]) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.CompositeFace(face)(&o.ConstructEntityOptions)
	}
}

// Prototype 实体原型名称
func (_EntityCreatorOption) Prototype(pt string) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		o.Prototype = pt
	}
}

// PersistId 实体持久化Id
func (_EntityCreatorOption) PersistId(id uid.Id) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		ec.Option{}.PersistId(id)(&o.EntityOptions)
	}
}

// AwakeOnFirstAccess 开启组件被首次访问时，检测并调用Awake()
func (_EntityCreatorOption) AwakeOnFirstAccess(b bool) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.AwakeOnFirstAccess(b)(&o.ConstructEntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (_EntityCreatorOption) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.FaceAnyAllocator(allocator)(&o.ConstructEntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (_EntityCreatorOption) HookAllocator(allocator container.Allocator[event.Hook]) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.HookAllocator(allocator)(&o.ConstructEntityOptions)
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (_EntityCreatorOption) GCCollector(collector container.GCCollector) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.GCCollector(collector)(&o.ConstructEntityOptions)
	}
}

// ComponentCtor 组件构造函数
func (_EntityCreatorOption) ComponentCtor(ctors pt.ComponentCtor) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.ComponentCtor(ctors)(&o.ConstructEntityOptions)
	}
}

// EntityCtor 实体构造函数
func (_EntityCreatorOption) EntityCtor(ctors pt.EntityCtor) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.EntityCtor(ctors)(&o.ConstructEntityOptions)
	}
}

// ParentId 父实体Id
func (_EntityCreatorOption) ParentId(id uid.Id) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		o.ParentID = id
	}
}

// Scope 实体的可访问作用域
func (_EntityCreatorOption) Scope(scope ec.Scope) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		o.Scope = scope
	}
}
