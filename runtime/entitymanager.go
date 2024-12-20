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

package runtime

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityManager 实体管理器接口
type EntityManager interface {
	ictx.CurrentContextProvider

	// AddEntity 添加实体
	AddEntity(entity ec.Entity) error
	// RemoveEntity 删除实体
	RemoveEntity(id uid.Id)
	// GetEntity 查询实体
	GetEntity(id uid.Id) (ec.Entity, bool)
	// ContainsEntity 实体是否存在
	ContainsEntity(id uid.Id) bool
	// RangeEntities 遍历所有实体
	RangeEntities(fun generic.Func1[ec.Entity, bool])
	// ReversedRangeEntities 反向遍历所有实体
	ReversedRangeEntities(fun generic.Func1[ec.Entity, bool])
	// FilterEntities 过滤并获取实体
	FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity
	// GetEntities 获取所有实体
	GetEntities() []ec.Entity
	// CountEntities 获取实体数量
	CountEntities() int

	IEntityManagerEventTab
}

type _EntityEntry struct {
	at    *generic.Node[iface.FaceAny]
	hooks [3]event.Hook
}

type _TreeNode struct {
	parentAt *generic.Node[iface.FaceAny]
	children *generic.List[iface.FaceAny]
}

type _EntityManagerBehavior struct {
	ctx        Context
	entityIdx  map[uid.Id]*_EntityEntry
	entityList generic.List[iface.FaceAny]
	treeNodes  map[uid.Id]*_TreeNode

	entityManagerEventTab
	entityTreeEventTab
}

func (mgr *_EntityManagerBehavior) init(ctx Context) {
	if ctx == nil {
		exception.Panicf("%w: %w: ctx is nil", ErrEntityManager, exception.ErrArgs)
	}

	mgr.ctx = ctx
	mgr.entityIdx = map[uid.Id]*_EntityEntry{}
	mgr.treeNodes = map[uid.Id]*_TreeNode{}

	ctx.ActivateEvent(&mgr.entityManagerEventTab, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.entityTreeEventTab, event.EventRecursion_Allow)
}

func (mgr *_EntityManagerBehavior) changeRunningState(state RunningState) {
	switch state {
	case RunningState_Started:
		mgr.RangeEntities(func(entity ec.Entity) bool {
			_EmitEventEntityManagerAddEntity(mgr, mgr, entity)
			return true
		})
	case RunningState_Terminating:
		mgr.ReversedRangeEntities(func(entity ec.Entity) bool {
			entity.DestroySelf()
			return true
		})
	case RunningState_Terminated:
		mgr.entityManagerEventTab.Close()
		mgr.entityTreeEventTab.Close()
	}
}

// GetCurrentContext 获取当前上下文
func (mgr *_EntityManagerBehavior) GetCurrentContext() iface.Cache {
	return mgr.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (mgr *_EntityManagerBehavior) GetConcurrentContext() iface.Cache {
	return mgr.ctx.GetConcurrentContext()
}

// AddEntity 添加实体
func (mgr *_EntityManagerBehavior) AddEntity(entity ec.Entity) error {
	return mgr.addEntity(entity, uid.Nil)
}

// RemoveEntity 删除实体
func (mgr *_EntityManagerBehavior) RemoveEntity(id uid.Id) {
	mgr.removeEntity(id)
}

// GetEntity 查询实体
func (mgr *_EntityManagerBehavior) GetEntity(id uid.Id) (ec.Entity, bool) {
	entry, ok := mgr.entityIdx[id]
	if !ok {
		return nil, false
	}

	if entry.at.Escaped() {
		return nil, false
	}

	return iface.Cache2Iface[ec.Entity](entry.at.V.Cache), true
}

// ContainsEntity 实体是否存在
func (mgr *_EntityManagerBehavior) ContainsEntity(id uid.Id) bool {
	_, ok := mgr.entityIdx[id]
	return ok
}

// RangeEntities 遍历所有实体
func (mgr *_EntityManagerBehavior) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// ReversedRangeEntities 反向遍历所有实体
func (mgr *_EntityManagerBehavior) ReversedRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.ReversedTraversal(func(n *generic.Node[iface.FaceAny]) bool {
		return fun.Exec(iface.Cache2Iface[ec.Entity](n.V.Cache))
	})
}

// FilterEntities 过滤并获取实体
func (mgr *_EntityManagerBehavior) FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	var entities []ec.Entity

	mgr.entityList.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		entity := iface.Cache2Iface[ec.Entity](n.V.Cache)

		if fun.Exec(entity) {
			entities = append(entities, entity)
		}

		return true
	})

	return entities
}

// GetEntities 获取所有实体
func (mgr *_EntityManagerBehavior) GetEntities() []ec.Entity {
	entities := make([]ec.Entity, 0, mgr.entityList.Len())

	mgr.entityList.Traversal(func(n *generic.Node[iface.FaceAny]) bool {
		entities = append(entities, iface.Cache2Iface[ec.Entity](n.V.Cache))
		return true
	})

	return entities
}

// CountEntities 获取实体数量
func (mgr *_EntityManagerBehavior) CountEntities() int {
	return mgr.entityList.Len()
}

func (mgr *_EntityManagerBehavior) OnComponentManagerAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		comp := components[i]

		if ec.UnsafeEntity(entity).GetOptions().ComponentUniqueID {
			if comp.GetId().IsNil() {
				ec.UnsafeComponent(comp).SetId(uid.New())
			}
		} else {
			ec.UnsafeComponent(comp).SetId(entity.GetId())
		}

		ec.UnsafeComponent(comp).WithContext(mgr.ctx)
	}

	_EmitEventEntityManagerEntityAddComponentsWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity, components []ec.Component) bool {
		return entity.GetState() > ec.EntityState_Alive
	}, mgr, entity, components)
}

func (mgr *_EntityManagerBehavior) OnComponentManagerRemoveComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityRemoveComponentWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity, component ec.Component) bool {
		return entity.GetState() > ec.EntityState_Alive
	}, mgr, entity, component)
}

func (mgr *_EntityManagerBehavior) OnComponentManagerFirstTouchComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityFirstTouchComponentWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity, component ec.Component) bool {
		return entity.GetState() > ec.EntityState_Alive
	}, mgr, entity, component)
}

func (mgr *_EntityManagerBehavior) addEntity(entity ec.Entity, parentId uid.Id) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	parent, err := mgr.fetchParent(entity, parentId)
	if err != nil {
		return err
	}

	switch entity.GetScope() {
	case ec.Scope_Local, ec.Scope_Global:
	default:
		return fmt.Errorf("%w: %w: invalid scope", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetState() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity state %q", ErrEntityManager, entity.GetState())
	}

	if entity.GetId().IsNil() {
		ec.UnsafeEntity(entity).SetId(uid.New())
	}
	ec.UnsafeEntity(entity).SetContext(iface.Iface2Cache[Context](mgr.ctx))
	ec.UnsafeEntity(entity).WithContext(mgr.ctx)

	entity.RangeComponents(func(comp ec.Component) bool {
		if ec.UnsafeEntity(entity).GetOptions().ComponentUniqueID {
			if comp.GetId().IsNil() {
				ec.UnsafeComponent(comp).SetId(uid.New())
			}
		} else {
			ec.UnsafeComponent(comp).SetId(entity.GetId())
		}
		ec.UnsafeComponent(comp).WithContext(mgr.ctx)
		return true
	})

	if mgr.ContainsEntity(entity.GetId()) {
		return fmt.Errorf("%w: entity already exists in entity-mgr", ErrEntityManager)
	}

	if parent != nil {
		if _, ok := mgr.treeNodes[entity.GetId()]; ok {
			return fmt.Errorf("%w: entity already exists in entity-tree", ErrEntityTree)
		}
	}

	if entity.GetScope() == ec.Scope_Global {
		_, loaded, err := service.Current(mgr).GetEntityManager().GetOrAddEntity(entity)
		if err != nil {
			return err
		}
		if loaded {
			return fmt.Errorf("%w: entity already exists in service entity-mgr", ErrEntityManager)
		}
	}

	entry := &_EntityEntry{
		at: mgr.entityList.PushBack(iface.MakeFaceAny(entity)),
		hooks: [3]event.Hook{
			ec.BindEventComponentManagerAddComponents(entity, mgr),
			ec.BindEventComponentManagerRemoveComponent(entity, mgr),
		},
	}
	if ec.UnsafeEntity(entity).GetOptions().ComponentAwakeOnFirstTouch {
		entry.hooks[2] = ec.BindEventComponentManagerFirstTouchComponent(entity, mgr)
	}
	mgr.entityIdx[entity.GetId()] = entry

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Enter)

	if parent != nil {
		mgr.addToParentNode(entity, parent)
	}

	_EmitEventEntityManagerAddEntityWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity) bool {
		return entity.GetState() > ec.EntityState_Alive
	}, mgr, entity)

	if parent != nil {
		mgr.attachParentNode(entity, parent)
	}

	return nil
}

func (mgr *_EntityManagerBehavior) removeEntity(id uid.Id) {
	entry, ok := mgr.entityIdx[id]
	if !ok {
		return
	}

	entity := iface.Cache2Iface[ec.Entity](entry.at.V.Cache)

	if entity.GetState() > ec.EntityState_Alive {
		return
	}
	ec.UnsafeEntity(entity).SetState(ec.EntityState_Leave)

	if entity.GetTreeNodeState() == ec.TreeNodeState_Attached {
		ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)
	}

	mgr.ReversedRangeChildren(entity.GetId(), func(child ec.Entity) bool {
		child.DestroySelf()
		return true
	})

	mgr.detachParentNode(entity)

	_EmitEventEntityManagerRemoveEntity(mgr, mgr, entity)

	mgr.removeFromParentNode(entity)

	delete(mgr.entityIdx, id)
	entry.at.Escape()
	event.Clean(entry.hooks[:])

	if entity.GetScope() == ec.Scope_Global {
		service.Current(mgr).GetEntityManager().RemoveEntity(entity.GetId())
	}
}

func (mgr *_EntityManagerBehavior) fetchParent(entity ec.Entity, parentId uid.Id) (ec.Entity, error) {
	if parentId.IsNil() {
		return nil, nil
	}

	parent, ok := mgr.GetEntity(parentId)
	if !ok {
		return nil, fmt.Errorf("%w: parent not exist", ErrEntityManager)
	}

	if parent.GetState() > ec.EntityState_Alive {
		return nil, fmt.Errorf("%w: invalid parent state %q", ErrEntityManager, parent.GetState())
	}

	if parent.GetId() == entity.GetId() {
		return nil, fmt.Errorf("%w: parent and child cannot be the same", ErrEntityManager)
	}

	return parent, nil
}
