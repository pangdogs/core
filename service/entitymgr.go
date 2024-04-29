package service

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/uid"
	"sync"
)

// EntityMgr 实体管理器接口
type EntityMgr interface {
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

type _EntityMgrBehavior struct {
	ctx       Context
	entityIdx sync.Map
}

func (mgr *_EntityMgrBehavior) init(ctx Context) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrEntityMgr, exception.ErrArgs))
	}

	mgr.ctx = ctx
}

// GetContext 获取服务上下文
func (mgr *_EntityMgrBehavior) GetContext() Context {
	return mgr.ctx
}

// GetEntity 查询实体
func (mgr *_EntityMgrBehavior) GetEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, ok := mgr.entityIdx.Load(id)
	if !ok {
		return nil, false
	}

	return v.(ec.ConcurrentEntity), true
}

// GetOrAddEntity 查询或添加实体
func (mgr *_EntityMgrBehavior) GetOrAddEntity(entity ec.ConcurrentEntity) (ec.ConcurrentEntity, bool, error) {
	if entity == nil {
		return nil, false, fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetId().IsNil() {
		return nil, false, fmt.Errorf("%w: entity id is nil", ErrEntityMgr)
	}

	if entity.GetConcurrentContext() == iface.NilCache {
		return nil, false, fmt.Errorf("%w: entity context is nil", ErrEntityMgr)
	}

	actual, loaded := mgr.entityIdx.LoadOrStore(entity.GetId(), entity)
	return actual.(ec.ConcurrentEntity), loaded, nil
}

// GetAndRemoveEntity 查询并删除实体
func (mgr *_EntityMgrBehavior) GetAndRemoveEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, loaded := mgr.entityIdx.LoadAndDelete(id)
	if !loaded {
		return nil, false
	}
	return v.(ec.ConcurrentEntity), true
}

// AddEntity 添加实体
func (mgr *_EntityMgrBehavior) AddEntity(entity ec.ConcurrentEntity) error {
	if entity == nil {
		return fmt.Errorf("%w: %w: entity is nil", ErrEntityMgr, exception.ErrArgs)
	}

	if entity.GetId().IsNil() {
		return fmt.Errorf("%w: entity id is nil", ErrEntityMgr)
	}

	if entity.GetConcurrentContext() == iface.NilCache {
		return fmt.Errorf("%w: entity context is nil", ErrEntityMgr)
	}

	mgr.entityIdx.Store(entity.GetId(), entity)

	return nil
}

// RemoveEntity 删除实体
func (mgr *_EntityMgrBehavior) RemoveEntity(id uid.Id) {
	mgr.entityIdx.Delete(id)
}
