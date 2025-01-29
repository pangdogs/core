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
	"git.golaxy.org/core/ec/ectx"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityManager 实体管理器接口
type EntityManager interface {
	ectx.CurrentContextProvider

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

type (
	_EntityNode = *generic.Node[iface.FaceAny]

	_TreeNode struct {
		parentAt *generic.Node[iface.FaceAny]
		children *generic.List[iface.FaceAny]
	}
)

type _EntityManagerBehavior struct {
	ctx         Context
	entityIndex map[uid.Id]_EntityNode
	entityList  generic.List[iface.FaceAny]
	treeNodes   map[uid.Id]*_TreeNode

	entityManagerEventTab
	entityTreeEventTab
}

func (mgr *_EntityManagerBehavior) init(ctx Context) {
	if ctx == nil {
		exception.Panicf("%w: %w: ctx is nil", ErrEntityManager, exception.ErrArgs)
	}

	mgr.ctx = ctx
	mgr.entityIndex = map[uid.Id]_EntityNode{}
	mgr.treeNodes = map[uid.Id]*_TreeNode{}

	ctx.ActivateEvent(&mgr.entityManagerEventTab, event.EventRecursion_Allow)
	ctx.ActivateEvent(&mgr.entityTreeEventTab, event.EventRecursion_Allow)
}

func (mgr *_EntityManagerBehavior) changeRunningStatus(status RunningStatus, args ...any) {
	switch status {
	case RunningStatus_Started:
		mgr.RangeEntities(func(entity ec.Entity) bool {
			_EmitEventEntityManagerAddEntityWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity) bool {
				return entity.GetState() > ec.EntityState_Alive
			}, mgr, entity)
			return true
		})
	case RunningStatus_Terminating:
		mgr.ReversedRangeEntities(func(entity ec.Entity) bool {
			entity.DestroySelf()
			return true
		})
	case RunningStatus_Terminated:
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
	entityNode, ok := mgr.entityIndex[id]
	if !ok {
		return nil, false
	}

	if entityNode.Escaped() {
		return nil, false
	}

	return iface.Cache2Iface[ec.Entity](entityNode.V.Cache), true
}

// ContainsEntity 实体是否存在
func (mgr *_EntityManagerBehavior) ContainsEntity(id uid.Id) bool {
	_, ok := mgr.entityIndex[id]
	return ok
}

// RangeEntities 遍历所有实体
func (mgr *_EntityManagerBehavior) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.Traversal(func(entityNode *generic.Node[iface.FaceAny]) bool {
		return fun.UnsafeCall(iface.Cache2Iface[ec.Entity](entityNode.V.Cache))
	})
}

// ReversedRangeEntities 反向遍历所有实体
func (mgr *_EntityManagerBehavior) ReversedRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.ReversedTraversal(func(entityNode *generic.Node[iface.FaceAny]) bool {
		return fun.UnsafeCall(iface.Cache2Iface[ec.Entity](entityNode.V.Cache))
	})
}

// FilterEntities 过滤并获取实体
func (mgr *_EntityManagerBehavior) FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	var entities []ec.Entity

	mgr.entityList.Traversal(func(entityNode *generic.Node[iface.FaceAny]) bool {
		entity := iface.Cache2Iface[ec.Entity](entityNode.V.Cache)

		if fun.UnsafeCall(entity) {
			entities = append(entities, entity)
		}

		return true
	})

	return entities
}

// GetEntities 获取所有实体
func (mgr *_EntityManagerBehavior) GetEntities() []ec.Entity {
	entities := make([]ec.Entity, 0, mgr.entityList.Len())

	mgr.entityList.Traversal(func(entityNode *generic.Node[iface.FaceAny]) bool {
		entities = append(entities, iface.Cache2Iface[ec.Entity](entityNode.V.Cache))
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
		mgr.initComponent(entity, components[i])
	}

	_EmitEventEntityManagerEntityAddComponents(mgr, mgr, entity, components)
}

func (mgr *_EntityManagerBehavior) OnComponentManagerRemoveComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityRemoveComponent(mgr, mgr, entity, component)
}

func (mgr *_EntityManagerBehavior) OnComponentManagerFirstTouchComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityFirstTouchComponent(mgr, mgr, entity, component)
}

func (mgr *_EntityManagerBehavior) addEntity(entity ec.Entity, parentId uid.Id) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.GetState() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityManager, entity.GetId(), entity.GetState())
	}

	switch entity.GetScope() {
	case ec.Scope_Local, ec.Scope_Global:
		break
	default:
		return fmt.Errorf("%w: invalid entity %q scope %q", ErrEntityManager, entity.GetId(), entity.GetScope())
	}

	mgr.initEntity(entity)

	parent, err := mgr.fetchParent(entity, parentId)
	if err != nil {
		return err
	}

	if _, ok := mgr.entityIndex[entity.GetId()]; ok {
		return fmt.Errorf("%w: entity %q already exists in entity-manager", ErrEntityManager, entity.GetId())
	}

	if parent != nil {
		if _, ok := mgr.treeNodes[entity.GetId()]; ok {
			return fmt.Errorf("%w: entity %q already exists in entity-tree", ErrEntityManager, entity.GetId())
		}
	}

	if entity.GetScope() == ec.Scope_Global {
		_, loaded, err := service.Current(mgr).GetEntityManager().GetOrAddEntity(entity)
		if err != nil {
			return fmt.Errorf("%w: entity %q add to service entity-manager failed, %w", ErrEntityManager, entity.GetId(), err)
		}
		if loaded {
			return fmt.Errorf("%w: entity %q already exists in service entity-manager", ErrEntityManager, entity.GetId())
		}
	}

	mgr.entityIndex[entity.GetId()] = mgr.entityList.PushBack(iface.MakeFaceAny(entity))

	mgr.observeEntity(entity)

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Enter)

	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Freedom)
	ec.UnsafeEntity(entity).SetTreeNodeParent(nil)

	if parent != nil {
		if err := mgr.appendToParentNode(entity, parent); err != nil {
			exception.Panicf("%w: entity %q append to parent %q failed, %w", ErrEntityManager, entity.GetId(), parent.GetId(), err)
		}
	}

	_EmitEventEntityManagerAddEntityWithInterrupt(mgr, func(entityManager EntityManager, entity ec.Entity) bool {
		return entity.GetState() > ec.EntityState_Alive
	}, mgr, entity)

	if parent != nil {
		if err := mgr.attachToParentNode(entity, parent); err != nil {
			return fmt.Errorf("%w: entity %q attach to parent %q failed, %w", ErrEntityManager, entity.GetId(), parent.GetId(), err)
		}
	}

	return nil
}

func (mgr *_EntityManagerBehavior) removeEntity(id uid.Id) {
	entityNode, ok := mgr.entityIndex[id]
	if !ok {
		return
	}

	entity := iface.Cache2Iface[ec.Entity](entityNode.V.Cache)

	if entity.GetState() > ec.EntityState_Alive {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Leave)

	if entity.GetTreeNodeState() != ec.TreeNodeState_Freedom {
		ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Detaching)
	}

	mgr.ReversedRangeChildren(entity.GetId(), func(child ec.Entity) bool {
		child.DestroySelf()
		return true
	})

	mgr.detachFromParentNode(entity)

	_EmitEventEntityManagerRemoveEntity(mgr, mgr, entity)

	mgr.removeFromParentNode(entity)

	delete(mgr.entityIndex, id)
	entityNode.Escape()

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
		return nil, fmt.Errorf("%w: parent %q not exist", ErrEntityManager, parentId)
	}

	if parent.GetState() > ec.EntityState_Alive {
		return nil, fmt.Errorf("%w: invalid parent %q state %q", ErrEntityManager, parent.GetId(), parent.GetState())
	}

	if parent.GetId() == entity.GetId() {
		return nil, fmt.Errorf("%w: parent and child %q can't be the same", ErrEntityManager, parent.GetId())
	}

	return parent, nil
}

func (mgr *_EntityManagerBehavior) initEntity(entity ec.Entity) {
	if entity.GetId().IsNil() {
		ec.UnsafeEntity(entity).SetId(uid.New())
	}
	ec.UnsafeEntity(entity).SetContext(iface.Iface2Cache[Context](mgr.ctx))
	ec.UnsafeEntity(entity).WithContext(mgr.ctx)

	entity.RangeComponents(func(comp ec.Component) bool {
		mgr.initComponent(entity, comp)
		return true
	})
}

func (mgr *_EntityManagerBehavior) initComponent(entity ec.Entity, comp ec.Component) {
	if ec.UnsafeEntity(entity).GetOptions().ComponentUniqueID {
		if comp.GetId().IsNil() {
			ec.UnsafeComponent(comp).SetId(uid.New())
		}
	} else {
		ec.UnsafeComponent(comp).SetId(entity.GetId())
	}
	ec.UnsafeComponent(comp).WithContext(entity)
}

func (mgr *_EntityManagerBehavior) observeEntity(entity ec.Entity) {
	ec.BindEventComponentManagerAddComponents(entity, mgr)
	ec.BindEventComponentManagerRemoveComponent(entity, mgr)

	if ec.UnsafeEntity(entity).GetOptions().ComponentAwakeOnFirstTouch {
		ec.BindEventComponentManagerFirstTouchComponent(entity, mgr)
	}
}
