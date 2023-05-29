package golaxy

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	pt.EntityOptions
	ParentID uid.Id   // 父实体Id
	Scope    ec.Scope // 实体的可访问作用域
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// EntityDefault 默认值
func (WithOption) EntityDefault() EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.Default()(&o.EntityOptions)
	}
}

// CompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) CompositeFace(face util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// ComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) ComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.ComponentAwakeByAccess(b)(&o.EntityOptions)
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) HookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// GCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithOption) GCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.GCCollector(collector)(&o.EntityOptions)
	}
}

// ParentId 父实体Id
func (WithOption) ParentId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.ParentID = id
	}
}

// AssignCompId 设置组件Id函数
func (WithOption) AssignCompId(fn func(entity ec.Entity, compPt pt.ComponentPt) uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.AssignCompId = fn
	}
}

// Scope 实体的可访问作用域
func (WithOption) Scope(scope ec.Scope) EntityOption {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}
