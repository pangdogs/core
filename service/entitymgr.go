package service

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
	"sync"
)

// IEntityMgr 实体管理器接口
type IEntityMgr interface {
	// GetContext 获取服务上下文
	GetContext() Context
	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.ConcurrentEntity, bool)
	// GetOrAddEntity 查询或添加实体
	GetOrAddEntity(entity ec.ConcurrentEntity) (ec.ConcurrentEntity, bool, error)
	// AddEntity 添加实体
	AddEntity(entity ec.ConcurrentEntity) error
	// GetAndRemoveEntity 查询并删除实体
	GetAndRemoveEntity(id uid.Id) (ec.ConcurrentEntity, bool)
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)
}

type _EntityMgr struct {
	ctx       Context
	entityMap sync.Map
}

func (entityMgr *_EntityMgr) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityMgr, exception.ErrArgs))
	}

	entityMgr.ctx = ctx
}

// GetContext 获取服务上下文
func (entityMgr *_EntityMgr) GetContext() Context {
	return entityMgr.ctx
}

// GetEntity 查询实体
func (entityMgr *_EntityMgr) GetEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, ok := entityMgr.entityMap.Load(id)
	if !ok {
		return nil, false
	}

	return v.(ec.ConcurrentEntity), true
}

// GetOrAddEntity 查询或添加实体
func (entityMgr *_EntityMgr) GetOrAddEntity(entity ec.ConcurrentEntity) (ec.ConcurrentEntity, bool, error) {
	if entity == nil {
		return nil, false, fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetId().IsNil() {
		return nil, false, fmt.Errorf("%w: entity id is nil", ErrEntityMgr)
	}

	if entity.GetConcurrentContext() == iface.NilCache {
		return nil, false, fmt.Errorf("%w: entity context can't be resolve", ErrEntityMgr)
	}

	actual, loaded := entityMgr.entityMap.LoadOrStore(entity.GetId(), entity)
	return actual.(ec.ConcurrentEntity), loaded, nil
}

// AddEntity 添加实体
func (entityMgr *_EntityMgr) AddEntity(entity ec.ConcurrentEntity) error {
	if entity == nil {
		return fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetId().IsNil() {
		return fmt.Errorf("%w: entity id is nil", ErrEntityMgr)
	}

	if entity.GetConcurrentContext() == iface.NilCache {
		return fmt.Errorf("%w: entity context can't be resolve", ErrEntityMgr)
	}

	entityMgr.entityMap.Store(entity.GetId(), entity)

	return nil
}

// GetAndRemoveEntity 查询并删除实体
func (entityMgr *_EntityMgr) GetAndRemoveEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, loaded := entityMgr.entityMap.LoadAndDelete(id)
	if !loaded {
		return nil, false
	}
	return v.(ec.ConcurrentEntity), true
}

// RemoveEntity 删除实体
func (entityMgr *_EntityMgr) RemoveEntity(id uid.Id) {
	entityMgr.entityMap.Delete(id)
}
