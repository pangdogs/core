package core

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

type _EntityCreatorOption struct{}

// EntityCreatorOptions 实体构建器的所有选项
type EntityCreatorOptions struct {
	pt.ConstructEntityOptions
	ParentID uid.Id // 父实体Id
}

// Default 默认值
func (_EntityCreatorOption) Default() option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.Default()(&o.ConstructEntityOptions)
		_EntityCreatorOption{}.ParentId(uid.Nil)(o)
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

// Scope 实体的可访问作用域
func (_EntityCreatorOption) Scope(scope ec.Scope) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.Scope(scope)
	}
}

// PersistId 实体持久化Id
func (_EntityCreatorOption) PersistId(id uid.Id) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.PersistId(id)(&o.ConstructEntityOptions)
	}
}

// AwakeOnFirstAccess 开启组件被首次访问时，检测并调用Awake()
func (_EntityCreatorOption) AwakeOnFirstAccess(b bool) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.AwakeOnFirstAccess(b)(&o.ConstructEntityOptions)
	}
}

// Meta Meta信息
func (_EntityCreatorOption) Meta(m ec.Meta) option.Setting[EntityCreatorOptions] {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.Meta(m)(&o.ConstructEntityOptions)
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
