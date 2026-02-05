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
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// EntityManager 实体管理器接口
type EntityManager interface {
	corectx.CurrentContextProvider

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
	// EachEntities 遍历每个实体
	EachEntities(fun generic.Action1[ec.Entity])
	// ReversedRangeEntities 反向遍历所有实体
	ReversedRangeEntities(fun generic.Func1[ec.Entity, bool])
	// ReversedEachEntities 反向遍历每个实体
	ReversedEachEntities(fun generic.Action1[ec.Entity])
	// FilterEntities 过滤并获取实体
	FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity
	// ListEntities 获取所有实体
	ListEntities() []ec.Entity
	// CountEntities 获取实体数量
	CountEntities() int

	IEntityManagerEventTab
}

type _TreeNode struct {
	parent          int
	attachedIndex   int
	attachedVersion int64
	children        generic.FreeList[int]
}

type _EntityManager struct {
	ctx             Context
	entityIdIndex   map[uid.Id]int
	entityList      generic.FreeList[ec.Entity]
	entityTreeNodes map[int]*_TreeNode

	entityManagerEventTab
	entityTreeEventTab
}

// CurrentContext 获取当前上下文
func (mgr *_EntityManager) CurrentContext() iface.Cache {
	return mgr.ctx.CurrentContext()
}

// ConcurrentContext 获取多线程安全的上下文
func (mgr *_EntityManager) ConcurrentContext() iface.Cache {
	return mgr.ctx.ConcurrentContext()
}

// AddEntity 添加实体
func (mgr *_EntityManager) AddEntity(entity ec.Entity) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.State() != ec.EntityState_Birth {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityManager, entity.Id(), entity.State())
	}

	switch entity.Scope() {
	case ec.Scope_Local, ec.Scope_Global:
		break
	default:
		return fmt.Errorf("%w: invalid entity %q scope %q", ErrEntityManager, entity.Id(), entity.Scope())
	}

	mgr.initEntity(entity)

	if _, ok := mgr.entityIdIndex[entity.Id()]; ok {
		return fmt.Errorf("%w: entity %q already exists in entity-manager", ErrEntityManager, entity.Id())
	}

	if entity.Scope() == ec.Scope_Global {
		_, loaded, err := service.Current(mgr).EntityManager().GetOrAddEntity(entity)
		if err != nil {
			return fmt.Errorf("%w: entity %q add to service entity-manager failed, %w", ErrEntityManager, entity.Id(), err)
		}
		if loaded {
			return fmt.Errorf("%w: entity %q already exists in service entity-manager", ErrEntityManager, entity.Id())
		}
	}

	entitySlot := mgr.entityList.PushBack(entity)
	mgr.entityIdIndex[entity.Id()] = entitySlot.Index()

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Enter)
	ec.UnsafeEntity(entity).SetEnteredHandle(entitySlot.Index(), entitySlot.Version())
	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Freedom)

	mgr.observeEntity(entity)

	_EmitEventEntityManagerAddEntity(mgr, mgr, entity)

	return nil
}

// RemoveEntity 删除实体
func (mgr *_EntityManager) RemoveEntity(id uid.Id) {
	slotIdx, ok := mgr.entityIdIndex[id]
	if !ok {
		return
	}
	entity := mgr.entityList.Get(slotIdx).V
	entity.Destroy()
}

// GetEntity 查询实体
func (mgr *_EntityManager) GetEntity(id uid.Id) (ec.Entity, bool) {
	slotIdx, ok := mgr.entityIdIndex[id]
	if !ok {
		return nil, false
	}
	return mgr.entityList.Get(slotIdx).V, true
}

// ContainsEntity 实体是否存在
func (mgr *_EntityManager) ContainsEntity(id uid.Id) bool {
	_, ok := mgr.entityIdIndex[id]
	return ok
}

// RangeEntities 遍历所有实体
func (mgr *_EntityManager) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.Traversal(func(slot *generic.FreeSlot[ec.Entity]) bool {
		return fun.UnsafeCall(slot.V)
	})
}

// EachEntities 遍历每个实体
func (mgr *_EntityManager) EachEntities(fun generic.Action1[ec.Entity]) {
	mgr.entityList.TraversalEach(func(slot *generic.FreeSlot[ec.Entity]) {
		fun.UnsafeCall(slot.V)
	})
}

// ReversedRangeEntities 反向遍历所有实体
func (mgr *_EntityManager) ReversedRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.ReversedTraversal(func(slot *generic.FreeSlot[ec.Entity]) bool {
		return fun.UnsafeCall(slot.V)
	})
}

// ReversedEachEntities 反向遍历每个实体
func (mgr *_EntityManager) ReversedEachEntities(fun generic.Action1[ec.Entity]) {
	mgr.entityList.ReversedTraversalEach(func(slot *generic.FreeSlot[ec.Entity]) {
		fun.UnsafeCall(slot.V)
	})
}

// FilterEntities 过滤并获取实体
func (mgr *_EntityManager) FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity {
	var entities []ec.Entity

	ver := mgr.entityList.Version()
	mgr.entityList.TraversalEach(func(slot *generic.FreeSlot[ec.Entity]) {
		if slot.Version() > ver {
			return
		}
		entity := slot.V
		if fun.UnsafeCall(entity) {
			entities = append(entities, entity)
		}
	})

	return entities
}

// ListEntities 获取所有实体
func (mgr *_EntityManager) ListEntities() []ec.Entity {
	return mgr.entityList.ToSlice()
}

// CountEntities 获取实体数量
func (mgr *_EntityManager) CountEntities() int {
	return mgr.entityList.Len() - mgr.entityList.OrphanCount()
}

func (mgr *_EntityManager) OnEntityDestroy(entity ec.Entity) {
	mgr.onEntityDestroyIfVersion(ec.UnsafeEntity(entity).EnteredHandle())
}

func (mgr *_EntityManager) OnComponentManagerAddComponents(entity ec.Entity, components []ec.Component) {
	for i := range components {
		mgr.initComponent(entity, components[i])
	}
	_EmitEventEntityManagerEntityAddComponents(mgr, mgr, entity, components)
}

func (mgr *_EntityManager) OnComponentManagerRemoveComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityRemoveComponent(mgr, mgr, entity, component)
}

func (mgr *_EntityManager) OnComponentManagerComponentEnableChanged(entity ec.Entity, component ec.Component, enable bool) {
	_EmitEventEntityManagerEntityComponentEnableChanged(mgr, mgr, entity, component, enable)
}

func (mgr *_EntityManager) OnComponentManagerFirstTouchComponent(entity ec.Entity, component ec.Component) {
	_EmitEventEntityManagerEntityFirstTouchComponent(mgr, mgr, entity, component)
}

func (mgr *_EntityManager) init(ctx Context) {
	if ctx == nil {
		exception.Panicf("%w: %w: ctx is nil", ErrEntityManager, exception.ErrArgs)
	}

	mgr.ctx = ctx
	mgr.entityIdIndex = map[uid.Id]int{}
	mgr.entityTreeNodes = map[int]*_TreeNode{forestNodeIdx: {parent: forestNodeIdx}}

	mgr.entityManagerEventTab.SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	mgr.entityTreeEventTab.SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
}

func (mgr *_EntityManager) onContextRunningEvent(ctx Context, runningEvent RunningEvent, args ...any) {
	switch runningEvent {
	case RunningEvent_Started:
		mgr.EachEntities(func(entity ec.Entity) {
			_EmitEventEntityManagerAddEntity(mgr, mgr, entity)
		})
	case RunningEvent_Terminating:
		mgr.ReversedEachEntities(func(entity ec.Entity) {
			entity.Destroy()
		})
	case RunningEvent_Terminated:
		mgr.entityManagerEventTab.SetEnabled(false)
		mgr.entityTreeEventTab.SetEnabled(false)
	}
}

func (mgr *_EntityManager) initEntity(entity ec.Entity) {
	if entity.Id().IsNil() {
		ec.UnsafeEntity(entity).SetId(uid.New())
	}
	ec.UnsafeEntity(entity).SetContext(iface.Iface2Cache[Context](mgr.ctx))
	ec.UnsafeEntity(entity).WithContext(mgr.ctx)

	event.UnsafeEvent(entity.EventEntityDestroy()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())

	event.UnsafeEvent(entity.EventComponentManagerAddComponents()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventComponentManagerRemoveComponent()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventComponentManagerComponentEnableChanged()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventComponentManagerFirstTouchComponent()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())

	event.UnsafeEvent(entity.EventTreeNodeAddChild()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventTreeNodeRemoveChild()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventTreeNodeAttachParent()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventTreeNodeDetachParent()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(entity.EventTreeNodeMoveTo()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())

	entity.EachComponents(func(comp ec.Component) {
		mgr.initComponent(entity, comp)
	})
}

func (mgr *_EntityManager) initComponent(entity ec.Entity, comp ec.Component) {
	event.UnsafeEvent(comp.EventComponentEnableChanged()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(comp.EventComponentDestroy()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())

	if ec.UnsafeEntity(entity).Options().ComponentUniqueID {
		if comp.Id().IsNil() {
			ec.UnsafeComponent(comp).SetId(uid.New())
		}
	}
}

func (mgr *_EntityManager) observeEntity(entity ec.Entity) {
	ec.BindEventEntityDestroy(entity, mgr)

	ec.BindEventComponentManagerAddComponents(entity, mgr)
	ec.BindEventComponentManagerRemoveComponent(entity, mgr)
	ec.BindEventComponentManagerComponentEnableChanged(entity, mgr)

	if ec.UnsafeEntity(entity).Options().ComponentAwakeOnFirstTouch {
		ec.BindEventComponentManagerFirstTouchComponent(entity, mgr)
	}
}

func (mgr *_EntityManager) onEntityDestroyIfVersion(idx int, ver int64) {
	entitySlot := mgr.entityList.Get(idx)
	if !checkEntitySlot(entitySlot, ver) {
		return
	}

	entity := entitySlot.V

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Leave)

	mgr.onEntityDestroyRemoveNode(entity.Id())

	_EmitEventEntityManagerRemoveEntity(mgr, mgr, entity)

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Death)

	delete(mgr.entityIdIndex, entity.Id())
	mgr.entityList.ReleaseIfVersion(idx, ver)

	if entity.Scope() == ec.Scope_Global {
		service.Current(mgr).EntityManager().RemoveEntity(entity.Id())
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Destroyed)
}

func checkEntitySlot(slot *generic.FreeSlot[ec.Entity], ver int64) bool {
	return slot != nil && !slot.Orphaned() && !slot.Freed() && slot.Version() == ver
}
