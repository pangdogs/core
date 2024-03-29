package core

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

// CreateEntity 创建实体
func CreateEntity(ctxProvider runtime.CurrentContextProvider) EntityCreator {
	return EntityCreator{
		rtCtx:   runtime.Current(ctxProvider),
		options: option.Make(pt.With.Default()),
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx    runtime.Context
	parentId uid.Id
	options  pt.ConstructEntityOptions
}

// CompositeFace 设置扩展者，在扩展实体自身能力时使用
func (c EntityCreator) CompositeFace(face iface.Face[ec.Entity]) EntityCreator {
	c.options = option.Append(c.options, pt.With.CompositeFace(face))
	return c
}

// Prototype 设置实体原型名称
func (c EntityCreator) Prototype(prototype string) EntityCreator {
	c.options = option.Append(c.options, pt.With.Prototype(prototype))
	return c
}

// Scope 设置实体的可访问作用域
func (c EntityCreator) Scope(scope ec.Scope) EntityCreator {
	c.options = option.Append(c.options, pt.With.Scope(scope))
	return c
}

// PersistId 设置实体持久化Id
func (c EntityCreator) PersistId(id uid.Id) EntityCreator {
	c.options = option.Append(c.options, pt.With.PersistId(id))
	return c
}

// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
func (c EntityCreator) AwakeOnFirstAccess(b bool) EntityCreator {
	c.options = option.Append(c.options, pt.With.AwakeOnFirstAccess(b))
	return c
}

// Meta 设置Meta信息
func (c EntityCreator) Meta(m ec.Meta) EntityCreator {
	c.options = option.Append(c.options, pt.With.Meta(m))
	return c
}

// FaceAnyAllocator 设置自定义FaceAny内存分配器，用于提高性能，通常传入运行时上下文中的FaceAnyAllocator
func (c EntityCreator) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) EntityCreator {
	c.options = option.Append(c.options, pt.With.FaceAnyAllocator(allocator))
	return c
}

// HookAllocator 设置自定义Hook内存分配器，用于提高性能，通常传入运行时上下文中的HookAllocator
func (c EntityCreator) HookAllocator(allocator container.Allocator[event.Hook]) EntityCreator {
	c.options = option.Append(c.options, pt.With.HookAllocator(allocator))
	return c
}

// GCCollector 设置自定义GC收集器，通常不传或者传入运行时上下文
func (c EntityCreator) GCCollector(collector container.GCCollector) EntityCreator {
	c.options = option.Append(c.options, pt.With.GCCollector(collector))
	return c
}

// ComponentCtor 设置组件构造函数
func (c EntityCreator) ComponentCtor(ctors pt.ComponentCtor) EntityCreator {
	c.options = option.Append(c.options, pt.With.ComponentCtor(ctors))
	return c
}

// EntityCtor 设置实体构造函数
func (c EntityCreator) EntityCtor(ctors pt.EntityCtor) EntityCreator {
	c.options = option.Append(c.options, pt.With.EntityCtor(ctors))
	return c
}

// ParentId 设置父实体Id
func (c EntityCreator) ParentId(id uid.Id) EntityCreator {
	c.parentId = id
	return c
}

// Spawn 创建实体
func (c EntityCreator) Spawn() (ec.Entity, error) {
	if c.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrGolaxy))
	}

	if !c.parentId.IsNil() {
		_, err := runtime.UnsafeECTree(c.rtCtx.GetECTree()).GetAndCheckEntity(c.parentId)
		if err != nil {
			return nil, err
		}
	}

	entity := pt.Using(service.Current(c.rtCtx), c.options.Prototype).UnsafeConstruct(c.options)

	if err := c.rtCtx.GetEntityMgr().AddEntity(entity); err != nil {
		return nil, err
	}

	if !c.parentId.IsNil() {
		if err := c.rtCtx.GetECTree().AddChild(c.options.PersistId, entity.GetId()); err != nil {
			entity.DestroySelf()
			return nil, err
		}
	}

	return entity, nil
}
