package core

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

// CreateEntity 创建实体
func CreateEntity(provider runtime.CurrentContextProvider) EntityCreator {
	return EntityCreator{
		rtCtx:   runtime.Current(provider),
		options: option.Make(ec.With.Default()),
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx    runtime.Context
	parentId uid.Id
	options  ec.EntityOptions
}

// CompositeFace 设置扩展者，在扩展实体自身能力时使用
func (c EntityCreator) CompositeFace(face iface.Face[ec.Entity]) EntityCreator {
	c.options = option.Append(c.options, ec.With.CompositeFace(face))
	return c
}

// Prototype 设置实体原型名称
func (c EntityCreator) Prototype(prototype string) EntityCreator {
	c.options = option.Append(c.options, ec.With.Prototype(prototype))
	return c
}

// Scope 设置实体的可访问作用域
func (c EntityCreator) Scope(scope ec.Scope) EntityCreator {
	c.options = option.Append(c.options, ec.With.Scope(scope))
	return c
}

// PersistId 设置实体持久化Id
func (c EntityCreator) PersistId(id uid.Id) EntityCreator {
	c.options = option.Append(c.options, ec.With.PersistId(id))
	return c
}

// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
func (c EntityCreator) AwakeOnFirstAccess(b bool) EntityCreator {
	c.options = option.Append(c.options, ec.With.AwakeOnFirstAccess(b))
	return c
}

// Meta 设置Meta信息
func (c EntityCreator) Meta(m ec.Meta) EntityCreator {
	c.options = option.Append(c.options, ec.With.Meta(m))
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
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrCore))
	}

	if !c.parentId.IsNil() {
		_, err := runtime.UnsafeEntityTree(c.rtCtx.GetEntityTree()).GetAndCheckEntity(c.parentId)
		if err != nil {
			return nil, err
		}
	}

	entity := pt.For(service.Current(c.rtCtx), c.options.Prototype).UnsafeConstruct(c.options)

	if err := c.rtCtx.GetEntityMgr().AddEntity(entity); err != nil {
		return nil, err
	}

	if !c.parentId.IsNil() {
		if err := c.rtCtx.GetEntityTree().AddNode(c.parentId, entity.GetId()); err != nil {
			entity.DestroySelf()
			return nil, err
		}
	}

	return entity, nil
}
