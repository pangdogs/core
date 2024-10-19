/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package service

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
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
	ctx      Context
	entities sync.Map
}

func (mgr *_EntityMgrBehavior) init(ctx Context) {
	if ctx == nil {
		exception.Panicf("%w: %w: ctx is nil", ErrEntityMgr, exception.ErrArgs)
	}

	mgr.ctx = ctx
}

// GetContext 获取服务上下文
func (mgr *_EntityMgrBehavior) GetContext() Context {
	return mgr.ctx
}

// GetEntity 查询实体
func (mgr *_EntityMgrBehavior) GetEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, ok := mgr.entities.Load(id)
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

	actual, loaded := mgr.entities.LoadOrStore(entity.GetId(), entity)
	return actual.(ec.ConcurrentEntity), loaded, nil
}

// GetAndRemoveEntity 查询并删除实体
func (mgr *_EntityMgrBehavior) GetAndRemoveEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, loaded := mgr.entities.LoadAndDelete(id)
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

	mgr.entities.Store(entity.GetId(), entity)

	return nil
}

// RemoveEntity 删除实体
func (mgr *_EntityMgrBehavior) RemoveEntity(id uid.Id) {
	mgr.entities.Delete(id)
}
