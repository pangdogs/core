package service

import (
	"errors"
	"fmt"
	"github.com/pangdogs/galaxy/ec"
)

// EntityMgr 服务上下文的实体管理器
type EntityMgr interface {
	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error

	// RemoveEntity 删除实体
	RemoveEntity(id int64)
}

// GetEntity 查询实体
func (ctx *ContextBehavior) GetEntity(id int64) (ec.Entity, bool) {
	ctx.entityMapMutex.RLock()
	defer ctx.entityMapMutex.RUnlock()

	entity, ok := ctx.entityMap[id]
	return entity, ok
}

// GetOrCreateEntity 查询实体，不存在时创建实体
func (ctx *ContextBehavior) GetOrCreateEntity(id int64, creator func(id int64) ec.Entity) (actual ec.Entity, loaded bool, err error) {
	if creator == nil {
		return nil, false, errors.New("nil creator")
	}

	ctx.entityMapMutex.Lock()
	defer ctx.entityMapMutex.Unlock()

	entity, ok := ctx.entityMap[id]
	if ok {
		return entity, true, nil
	}

	if id <= 0 {
		return nil, false, errors.New("input id less equal 0 invalid")
	}

	entity = creator(id)

	if entity == nil {
		return nil, false, errors.New("creator return nil entity invalid")
	}

	if entity.GetID() != id {
		return nil, false, errors.New("creator return entity id not equal input id invalid")
	}

	ctx.entityMap[entity.GetID()] = entity

	return entity, false, nil
}

// AddEntity 添加实体
func (ctx *ContextBehavior) AddEntity(entity ec.Entity) error {
	ctx.entityMapMutex.Lock()
	defer ctx.entityMapMutex.Unlock()

	if entity == nil {
		return errors.New("nil entity")
	}

	if entity.GetID() <= 0 {
		return errors.New("entity id less equal 0 invalid")
	}

	_, ok := ctx.entityMap[entity.GetID()]
	if ok {
		return fmt.Errorf("repeated entity '%d' in this service context", entity.GetID())
	}

	ctx.entityMap[entity.GetID()] = entity

	return nil
}

// RemoveEntity 删除实体
func (ctx *ContextBehavior) RemoveEntity(id int64) {
	ctx.entityMapMutex.Lock()
	defer ctx.entityMapMutex.Unlock()

	delete(ctx.entityMap, id)
}
