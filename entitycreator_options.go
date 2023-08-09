package golaxy

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

type _EntityCreatorOption struct{}

// EntityCreatorOptions 实体构建器的所有选项
type EntityCreatorOptions struct {
	pt.ConstructEntityOptions
	ParentID uid.Id   // 父实体Id
	Scope    ec.Scope // 实体的可访问作用域
}

// EntityCreatorOption 实体构建器的选项设置器
type EntityCreatorOption func(o *EntityCreatorOptions)

// Default 默认值
func (_EntityCreatorOption) Default() EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.Default()(&o.ConstructEntityOptions)
		_EntityCreatorOption{}.ParentId(uid.Nil)(o)
		_EntityCreatorOption{}.Scope(ec.Scope_Local)(o)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (_EntityCreatorOption) CompositeFace(face util.Face[ec.Entity]) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.CompositeFace(face)(&o.ConstructEntityOptions)
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (_EntityCreatorOption) ComponentAwakeByAccess(b bool) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.ComponentAwakeByAccess(b)(&o.ConstructEntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (_EntityCreatorOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.FaceAnyAllocator(allocator)(&o.ConstructEntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (_EntityCreatorOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.HookAllocator(allocator)(&o.ConstructEntityOptions)
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (_EntityCreatorOption) GCCollector(collector container.GCCollector) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.GCCollector(collector)(&o.ConstructEntityOptions)
	}
}

// ComponentConstructor 组件构造函数
func (_EntityCreatorOption) ComponentConstructor(fn pt.ComponentConstructor) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.ComponentConstructor(fn)(&o.ConstructEntityOptions)
	}
}

// EntityConstructor 实体构造函数
func (_EntityCreatorOption) EntityConstructor(fn pt.EntityConstructor) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		pt.Option{}.EntityConstructor(fn)(&o.ConstructEntityOptions)
	}
}

// ParentId 父实体Id
func (_EntityCreatorOption) ParentId(id uid.Id) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		o.ParentID = id
	}
}

// Scope 实体的可访问作用域
func (_EntityCreatorOption) Scope(scope ec.Scope) EntityCreatorOption {
	return func(o *EntityCreatorOptions) {
		o.Scope = scope
	}
}
