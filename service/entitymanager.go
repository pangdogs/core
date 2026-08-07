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
	"sync"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityManager 提供跨运行时可并发访问的全局实体索引。
// 只有 Scope_Global 实体会由运行时注册到该索引。
// 注册与注销事件在调用方 goroutine 中同步派发，因此不同 Runtime 可能并发触发回调。
type EntityManager interface {
	// GetEntity 按 ID 查询全局实体。
	GetEntity(id uid.Id) (ec.ConcurrentEntity, bool)
	// GetOrAddEntity 原子地查询或注册全局实体，并报告该 ID 是否已经存在。
	GetOrAddEntity(entity ec.ConcurrentEntity) (ec.ConcurrentEntity, bool, error)
	// RemoveEntity 按 ID 注销全局实体；实体不存在时不执行任何操作。
	RemoveEntity(id uid.Id)
}

type _EntityManager struct {
	ctx      Context
	entities sync.Map
}

// GetEntity 按 ID 查询全局实体。
func (mgr *_EntityManager) GetEntity(id uid.Id) (ec.ConcurrentEntity, bool) {
	v, ok := mgr.entities.Load(id)
	if !ok {
		return nil, false
	}

	return v.(ec.ConcurrentEntity), true
}

// GetOrAddEntity 原子地查询或注册全局实体，并报告该 ID 是否已经存在。
func (mgr *_EntityManager) GetOrAddEntity(entity ec.ConcurrentEntity) (ec.ConcurrentEntity, bool, error) {
	if entity == nil {
		return nil, false, fmt.Errorf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.Id().IsNil() {
		return nil, false, fmt.Errorf("%w: entity id is nil", ErrEntityManager)
	}

	if entity.ConcurrentContext() == iface.NilCache {
		return nil, false, fmt.Errorf("%w: entity context is nil", ErrEntityManager)
	}

	select {
	case <-entity.Done():
		return nil, false, fmt.Errorf("%w: entity is death", ErrEntityManager)
	default:
	}

	actual, loaded := mgr.entities.LoadOrStore(entity.Id(), entity)
	if !loaded {
		mgr.ctx.emitEventRunningEvent(RunningEvent_EntityRegistered, entity)
	}

	return actual.(ec.ConcurrentEntity), loaded, nil
}

// RemoveEntity 按 ID 注销全局实体；实体不存在时不执行任何操作。
func (mgr *_EntityManager) RemoveEntity(id uid.Id) {
	entity, loaded := mgr.entities.LoadAndDelete(id)
	if !loaded {
		return
	}
	mgr.ctx.emitEventRunningEvent(RunningEvent_EntityDeregistered, entity)
}

func (mgr *_EntityManager) init(ctx Context) {
	if ctx == nil {
		exception.Panicf("%w: %w: ctx is nil", ErrEntityManager, exception.ErrArgs)
	}

	mgr.ctx = ctx
}
