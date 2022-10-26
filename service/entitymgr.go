package service

import (
	"errors"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/concurrent"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetServiceCtx 获取服务上下文
	GetServiceCtx() Context
	// GetEntity 查询实体
	GetEntity(id int64) (ec.Entity, bool)
	// GetEntityWithSerialNo 使用ID与序列号查询实体
	GetEntityWithSerialNo(id, serialNo int64) (ec.Entity, bool)
	// GetOrAddEntity 查询或添加实体
	GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error)
	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// GetAndRemoveEntity 查询并删除实体
	GetAndRemoveEntity(id int64) (ec.Entity, bool)
	// GetAndRemoveEntityWithSerialNo 使用ID与序列号查询并删除实体
	GetAndRemoveEntityWithSerialNo(id, serialNo int64) (ec.Entity, bool)
	// RemoveEntity 删除实体
	RemoveEntity(id int64)
	// RemoveEntityWithSerialNo 使用ID与序列号删除实体
	RemoveEntityWithSerialNo(id, serialNo int64)
}

type _EntityMgr struct {
	serviceCtx Context
	entityMap  concurrent.Map[int64, ec.Entity]
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
	entity, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	return entity, true
}

func (entityMgr *_EntityMgr) GetEntityWithSerialNo(id, serialNo int64) (ec.Entity, bool) {
	entity, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	if entity.GetSerialNo() != serialNo {
		return nil, false
	}

	return entity, true
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

	actual, loaded := entityMgr.entityMap.LoadOrStore(entity.GetID(), entity)
	return actual, loaded, nil
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
	return entityMgr.entityMap.LoadAndDelete(id)
}

func (entityMgr *_EntityMgr) GetAndRemoveEntityWithSerialNo(id, serialNo int64) (ec.Entity, bool) {
	return entityMgr.entityMap.TryLoadAndDelete(id, func(entity ec.Entity) bool {
		return entity.GetSerialNo() == serialNo
	})
}

func (entityMgr *_EntityMgr) RemoveEntity(id int64) {
	entityMgr.entityMap.Delete(id)
}

func (entityMgr *_EntityMgr) RemoveEntityWithSerialNo(id, serialNo int64) {
	entityMgr.entityMap.TryDelete(id, func(entity ec.Entity) bool {
		return entity.GetSerialNo() == serialNo
	})
}
