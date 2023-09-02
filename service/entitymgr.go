package service

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"sync"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetContext 获取服务上下文
	GetContext() Context
	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.Entity, bool)
	// GetOrAddEntity 查询或添加实体
	GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error)
	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// GetAndRemoveEntity 查询并删除实体
	GetAndRemoveEntity(id uid.Id) (ec.Entity, bool)
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)
}

type _EntityMgr struct {
	ctx       Context
	entityMap sync.Map
}

func (entityMgr *_EntityMgr) init(ctx Context) {
	if ctx == nil {
		panic("nil ctx")
	}

	entityMgr.ctx = ctx
}

// GetContext 获取服务上下文
func (entityMgr *_EntityMgr) GetContext() Context {
	return entityMgr.ctx
}

// GetEntity 查询实体
func (entityMgr *_EntityMgr) GetEntity(id uid.Id) (ec.Entity, bool) {
	v, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	return v.(ec.Entity), true
}

// GetOrAddEntity 查询或添加实体
func (entityMgr *_EntityMgr) GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error) {
	if entity == nil {
		return nil, false, errors.New("nil entity")
	}

	if entity.GetId().IsNil() {
		return nil, false, errors.New("entity id equal zero is invalid")
	}

	if entity.ResolveContext() == util.NilIfaceCache {
		return nil, false, errors.New("entity context can't be resolve")
	}

	actual, loaded := entityMgr.entityMap.LoadOrStore(entity.GetId(), entity)
	return actual.(ec.Entity), loaded, nil
}

// AddEntity 添加实体
func (entityMgr *_EntityMgr) AddEntity(entity ec.Entity) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetId().IsNil() {
		return errors.New("entity id equal zero is invalid")
	}

	if entity.ResolveContext() == util.NilIfaceCache {
		return errors.New("entity context can't be resolve")
	}

	entityMgr.entityMap.Store(entity.GetId(), entity)

	return nil
}

// GetAndRemoveEntity 查询并删除实体
func (entityMgr *_EntityMgr) GetAndRemoveEntity(id uid.Id) (ec.Entity, bool) {
	v, loaded := entityMgr.entityMap.LoadAndDelete(id)
	if !loaded {
		return nil, false
	}
	return v.(ec.Entity), true
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgr) RemoveEntity(id uid.Id) {
	entityMgr.entityMap.Delete(id)
}
