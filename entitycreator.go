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
func CreateEntity(provider runtime.CurrentContextProvider, prototype string) EntityCreator {
	return EntityCreator{
		rtCtx:     runtime.Current(provider),
		prototype: prototype,
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx     runtime.Context
	prototype string
	parentId  uid.Id
	settings  []option.Setting[ec.EntityOptions]
}

// CompositeFace 设置扩展者，在扩展实体自身能力时使用
func (c EntityCreator) CompositeFace(face iface.Face[ec.Entity]) EntityCreator {
	c.settings = append(c.settings, ec.With.CompositeFace(face))
	return c
}

// Composite 设置扩展者，在扩展实体自身能力时使用
func (c EntityCreator) Composite(e ec.Entity) EntityCreator {
	c.settings = append(c.settings, ec.With.CompositeFace(iface.MakeFace(e)))
	return c
}

// Scope 设置实体的可访问作用域
func (c EntityCreator) Scope(scope ec.Scope) EntityCreator {
	c.settings = append(c.settings, ec.With.Scope(scope))
	return c
}

// PersistId 设置实体持久化Id
func (c EntityCreator) PersistId(id uid.Id) EntityCreator {
	c.settings = append(c.settings, ec.With.PersistId(id))
	return c
}

// AwakeOnFirstAccess 设置开启组件被首次访问时，检测并调用Awake()
func (c EntityCreator) AwakeOnFirstAccess(b bool) EntityCreator {
	c.settings = append(c.settings, ec.With.AwakeOnFirstAccess(b))
	return c
}

// Meta 设置Meta信息
func (c EntityCreator) Meta(m ec.Meta) EntityCreator {
	c.settings = append(c.settings, ec.With.Meta(m))
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

	entity := pt.For(service.Current(c.rtCtx), c.prototype).Construct(c.settings...)

	if c.parentId.IsNil() {
		if err := c.rtCtx.GetEntityMgr().AddEntity(entity); err != nil {
			return nil, err
		}
	} else {
		if err := c.rtCtx.GetEntityTree().AddNode(entity, c.parentId); err != nil {
			return nil, err
		}
	}

	return entity, nil
}
