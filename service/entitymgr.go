package service

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/concurrent"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetServiceCtx 获取服务上下文
	GetServiceCtx() Context
	// GetEntity 查询实体
	GetEntity(id ec.ID) (ec.Entity, bool)
	// GetEntityWithSerialNo 查询实体，同时使用id与serialNo可以在多线程环境中准确定位实体
	GetEntityWithSerialNo(id ec.ID, serialNo int64) (ec.Entity, bool)
	// GetOrAddEntity 查询或添加实体
	GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error)
	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// GetAndRemoveEntity 查询并删除实体
	GetAndRemoveEntity(id ec.ID) (ec.Entity, bool)
	// GetAndRemoveEntityWithSerialNo 查询并删除实体，同时使用id与serialNo可以在多线程环境中准确定位实体
	GetAndRemoveEntityWithSerialNo(id ec.ID, serialNo int64) (ec.Entity, bool)
	// RemoveEntity 删除实体
	RemoveEntity(id ec.ID)
	// RemoveEntityWithSerialNo 删除实体，同时使用id与serialNo可以在多线程环境中准确定位实体
	RemoveEntityWithSerialNo(id ec.ID, serialNo int64)
}

type _EntityMgr struct {
	serviceCtx Context
	entityMap  concurrent.Map[ec.ID, ec.Entity]
}

func (entityMgr *_EntityMgr) init(serviceCtx Context) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	entityMgr.serviceCtx = serviceCtx
}

// GetServiceCtx 获取服务上下文
func (entityMgr *_EntityMgr) GetServiceCtx() Context {
	return entityMgr.serviceCtx
}

// GetEntity 查询实体
func (entityMgr *_EntityMgr) GetEntity(id ec.ID) (ec.Entity, bool) {
	entity, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	return entity, true
}

// GetEntityWithSerialNo 查询实体，同时使用id与serialNo可以在多线程环境中准确定位实体
func (entityMgr *_EntityMgr) GetEntityWithSerialNo(id ec.ID, serialNo int64) (ec.Entity, bool) {
	entity, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	if entity.GetSerialNo() != serialNo {
		return nil, false
	}

	return entity, true
}

// GetOrAddEntity 查询或添加实体
func (entityMgr *_EntityMgr) GetOrAddEntity(entity ec.Entity) (ec.Entity, bool, error) {
	if entity == nil {
		return nil, false, errors.New("nil entity")
	}

	if entity.GetID() == util.Zero[ec.ID]() {
		return nil, false, errors.New("entity id is zero invalid")
	}

	if ec.UnsafeEntity(entity).GetContext() == util.NilIfaceCache {
		return nil, false, errors.New("entity context can't be resolve")
	}

	actual, loaded := entityMgr.entityMap.LoadOrStore(entity.GetID(), entity)
	return actual, loaded, nil
}

// AddEntity 添加实体
func (entityMgr *_EntityMgr) AddEntity(entity ec.Entity) error {
	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetID() == util.Zero[ec.ID]() {
		return errors.New("entity id is zero invalid")
	}

	if ec.UnsafeEntity(entity).GetContext() == util.NilIfaceCache {
		return errors.New("entity context can't be resolve")
	}

	entityMgr.entityMap.Store(entity.GetID(), entity)

	return nil
}

// GetAndRemoveEntity 查询并删除实体
func (entityMgr *_EntityMgr) GetAndRemoveEntity(id ec.ID) (ec.Entity, bool) {
	return entityMgr.entityMap.LoadAndDelete(id)
}

// GetAndRemoveEntityWithSerialNo 查询并删除实体，同时使用id与serialNo可以在多线程环境中准确定位实体
func (entityMgr *_EntityMgr) GetAndRemoveEntityWithSerialNo(id ec.ID, serialNo int64) (ec.Entity, bool) {
	return entityMgr.entityMap.TryLoadAndDelete(id, func(entity ec.Entity) bool {
		return entity.GetSerialNo() == serialNo
	})
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgr) RemoveEntity(id ec.ID) {
	entityMgr.entityMap.Delete(id)
}

// RemoveEntityWithSerialNo 删除实体，同时使用id与serialNo可以在多线程环境中准确定位实体
func (entityMgr *_EntityMgr) RemoveEntityWithSerialNo(id ec.ID, serialNo int64) {
	entityMgr.entityMap.TryDelete(id, func(entity ec.Entity) bool {
		return entity.GetSerialNo() == serialNo
	})
}
