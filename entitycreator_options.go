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
		WithOption{}.EntityParentId(uid.Nil)(o)
		WithOption{}.EntityScope(ec.Scope_Local)(o)
	}
}

// EntityCompositeFace 扩展者，在扩展实体自身能力时使用
func (WithOption) EntityCompositeFace(face util.Face[ec.Entity]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.CompositeFace(face)(&o.EntityOptions)
	}
}

// EntityComponentAwakeByAccess 开启组件被访问时，检测并调用Awake()
func (WithOption) EntityComponentAwakeByAccess(b bool) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.ComponentAwakeByAccess(b)(&o.EntityOptions)
	}
}

// EntityFaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (WithOption) EntityFaceAnyAllocator(allocator container.Allocator[util.FaceAny]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.FaceAnyAllocator(allocator)(&o.EntityOptions)
	}
}

// EntityHookAllocator 自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (WithOption) EntityHookAllocator(allocator container.Allocator[localevent.Hook]) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.HookAllocator(allocator)(&o.EntityOptions)
	}
}

// EntityGCCollector 自定义GC收集器，通常不传或者传入运行时上下文
func (WithOption) EntityGCCollector(collector container.GCCollector) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.GCCollector(collector)(&o.EntityOptions)
	}
}

// ComponentConstructor 组件构造函数
func (WithOption) ComponentConstructor(fn pt.ComponentConstructor) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.ComponentConstructor(fn)(&o.EntityOptions)
	}
}

// EntityConstructor 实体构造函数
func (WithOption) EntityConstructor(fn pt.EntityConstructor) EntityOption {
	return func(o *EntityOptions) {
		pt.WithOption{}.EntityConstructor(fn)(&o.EntityOptions)
	}
}

// EntityParentId 父实体Id
func (WithOption) EntityParentId(id uid.Id) EntityOption {
	return func(o *EntityOptions) {
		o.ParentID = id
	}
}

// EntityScope 实体的可访问作用域
func (WithOption) EntityScope(scope ec.Scope) EntityOption {
	return func(o *EntityOptions) {
		o.Scope = scope
	}
}
