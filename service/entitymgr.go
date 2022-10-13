package service

import (
	"errors"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/util"
	"sync"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetServiceCtx 获取服务上下文
	GetServiceCtx() Context

	// GetEntity 查询实体
	GetEntity(id int64) (ec.Entity, bool)

	// GetOrAddEntity 查询或添加实体
	GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error)

	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error

	// GetAndRemoveEntity 查询并删除实体
	GetAndRemoveEntity(id int64) (ec.Entity, bool)

	// RemoveEntity 删除实体
	RemoveEntity(id int64)
}

type _EntityMgr struct {
	serviceCtx Context
	entityMap  sync.Map
	inited     bool
}

func (entityMgr *_EntityMgr) Init(serviceCtx Context) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	if entityMgr.inited {
		panic("repeated init entity manager")
	}

	entityMgr.serviceCtx = serviceCtx
	entityMgr.inited = true
}

func (entityMgr *_EntityMgr) GetServiceCtx() Context {
	return entityMgr.serviceCtx
}

func (entityMgr *_EntityMgr) GetEntity(id int64) (ec.Entity, bool) {
	v, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}
	return v.(ec.Entity), true
}

func (entityMgr *_EntityMgr) GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error) {
	if entity == nil {
		return nil, false, errors.New("nil entity")
	}

	if entity.GetID() == 0 {
		return nil, false, errors.New("entity id invalid")
	}

	if ec.UnsafeEntity(entity).GetContext() == util.NilIfaceCache {
		return nil, false, errors.New("entity context has not been setup")
	}

	v, loaded := entityMgr.entityMap.LoadOrStore(entity.GetID(), entity)
	if loaded {
		return v.(ec.Entity), true, nil
	}

	return entity, false, nil
}

func (entityMgr *_EntityMgr) AddEntity(entity ec.Entity) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetID() == 0 {
		return errors.New("entity id invalid")
	}

	if ec.UnsafeEntity(entity).GetContext() == util.NilIfaceCache {
		return errors.New("entity context has not been setup")
	}

	entityMgr.entityMap.Store(entity.GetID(), entity)

	return nil
}

func (entityMgr *_EntityMgr) GetAndRemoveEntity(id int64) (ec.Entity, bool) {
	v, loaded := entityMgr.entityMap.LoadAndDelete(id)
	if loaded {
		return v.(ec.Entity), true
	}
	return nil, false
}

func (entityMgr *_EntityMgr) RemoveEntity(id int64) {
	entityMgr.entityMap.Delete(id)
}
