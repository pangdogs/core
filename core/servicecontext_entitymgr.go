package core

import (
	"errors"
	"fmt"
)

// _ServiceContextEntityMgr 服务上下文（Service Context）的实体（Entity）管理器
type _ServiceContextEntityMgr interface {
	// GetEntity 查询实体，线程安全
	GetEntity(id int64) (Entity, bool)

	// GetOrCreateEntity 查询实体，不存在时创建实体，线程安全
	GetOrCreateEntity(id int64, creator func(id int64) Entity) (actual Entity, loaded bool, err error)

	// AddEntity 添加实体，线程安全
	AddEntity(entity Entity) error

	// RemoveEntity 删除实体，线程安全
	RemoveEntity(id int64)
}

// GetEntity 查询实体，线程安全
func (servCtx *_ServiceContextBehavior) GetEntity(id int64) (Entity, bool) {
	servCtx.entityMapMutex.RLock()
	defer servCtx.entityMapMutex.RUnlock()

	entity, ok := servCtx.entityMap[id]
	return entity, ok
}

// GetOrCreateEntity 查询实体，不存在时创建实体，线程安全
func (servCtx *_ServiceContextBehavior) GetOrCreateEntity(id int64, creator func(id int64) Entity) (actual Entity, loaded bool, err error) {
	if creator == nil {
		return nil, false, errors.New("nil creator")
	}

	servCtx.entityMapMutex.Lock()
	defer servCtx.entityMapMutex.Unlock()

	entity, ok := servCtx.entityMap[id]
	if ok {
		return entity, true, nil
	}

	if id <= 0 {
		return nil, false, errors.New("entity id less equal 0 invalid")
	}

	entity = creator(id)

	if entity == nil {
		return nil, false, errors.New("creator return nil entity invalid")
	}

	if entity.GetID() != id {
		return nil, false, errors.New("creator return entity id not equal input id invalid")
	}

	servCtx.entityMap[entity.GetID()] = entity

	return entity, false, nil
}

// AddEntity 添加实体，线程安全
func (servCtx *_ServiceContextBehavior) AddEntity(entity Entity) error {
	servCtx.entityMapMutex.Lock()
	defer servCtx.entityMapMutex.Unlock()

	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetID() <= 0 {
		return errors.New("entity id less equal 0 invalid")
	}

	entity, ok := servCtx.entityMap[entity.GetID()]
	if ok {
		return fmt.Errorf("repeated entity '%d' in this service context", entity.GetID())
	}

	servCtx.entityMap[entity.GetID()] = entity

	return nil
}

// RemoveEntity 删除实体，线程安全
func (servCtx *_ServiceContextBehavior) RemoveEntity(id int64) {
	servCtx.entityMapMutex.Lock()
	defer servCtx.entityMapMutex.Unlock()

	delete(servCtx.entityMap, id)
}
