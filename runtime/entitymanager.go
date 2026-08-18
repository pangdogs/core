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

// EntityManager 管理当前运行时拥有的实体及其加入顺序。
// 该接口不提供并发保护，应在所属运行时 goroutine 中使用。
type EntityManager interface {
	corectx.CurrentContextProvider

	// AddEntity 接管 Born 状态的实体；运行时已启动时会同步推进其生命周期。
	AddEntity(entity ec.Entity) error
	// RemoveEntity 按 ID 请求销毁实体；实体不存在时不执行任何操作。
	RemoveEntity(id uid.ID)
	// GetEntity 按 ID 查询本地实体。
	GetEntity(id uid.ID) (ec.Entity, bool)
	// RangeEntities 按加入顺序遍历实体，回调返回 false 时停止。
	RangeEntities(fun generic.Func1[ec.Entity, bool])
	// EachEntities 按加入顺序遍历全部实体。
	EachEntities(fun generic.Action1[ec.Entity])
	// ReversedRangeEntities 按加入顺序逆向遍历实体，回调返回 false 时停止。
	ReversedRangeEntities(fun generic.Func1[ec.Entity, bool])
	// ReversedEachEntities 按加入顺序逆向遍历全部实体。
	ReversedEachEntities(fun generic.Action1[ec.Entity])
	// FilterEntities 按加入顺序返回符合条件的实体。
	FilterEntities(fun generic.Func1[ec.Entity, bool]) []ec.Entity
	// ListEntities 按加入顺序返回实体切片副本。
	ListEntities() []ec.Entity
	// CountEntities 返回当前实体数。
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
	entityIDIndex   map[uid.ID]int
	entityList      generic.FreeList[ec.Entity]
	entityTreeNodes map[int]*_TreeNode

	entityManagerEventTab
	entityTreeEventTab
}

// CurrentContextCache 返回所属运行时的当前上下文接口缓存。
func (mgr *_EntityManager) CurrentContextCache() iface.Cache {
	return mgr.ctx.CurrentContextCache()
}

// ConcurrentContextCache 返回所属运行时的并发上下文接口缓存。
func (mgr *_EntityManager) ConcurrentContextCache() iface.Cache {
	return mgr.ctx.ConcurrentContextCache()
}

// AddEntity 接管 Born 状态的实体并同步触发其加入事件。
func (mgr *_EntityManager) AddEntity(entity ec.Entity) error {
	if entity == nil {
		exception.Panicf("%w: %w: entity is nil", ErrEntityManager, exception.ErrArgs)
	}

	if entity.State() != ec.EntityState_Born {
		return fmt.Errorf("%w: invalid entity %q state %q", ErrEntityManager, entity.ID(), entity.State())
	}

	switch entity.Scope() {
	case ec.Scope_Local, ec.Scope_Global:
		break
	default:
		return fmt.Errorf("%w: invalid entity %q scope %q", ErrEntityManager, entity.ID(), entity.Scope())
	}

	mgr.initEntity(entity)

	if _, ok := mgr.entityIDIndex[entity.ID()]; ok {
		entity.AsyncScope().Close()
		return fmt.Errorf("%w: entity %q already exists in entity-manager", ErrEntityManager, entity.ID())
	}

	if entity.Scope() == ec.Scope_Global {
		_, loaded, err := service.Current(mgr).EntityManager().GetOrAddEntity(entity)
		if err != nil {
			entity.AsyncScope().Close()
			return fmt.Errorf("%w: entity %q add to service entity-manager failed, %w", ErrEntityManager, entity.ID(), err)
		}
		if loaded {
			entity.AsyncScope().Close()
			return fmt.Errorf("%w: entity %q already exists in service entity-manager", ErrEntityManager, entity.ID())
		}
	}

	entitySlot := mgr.entityList.PushBack(entity)
	mgr.entityIDIndex[entity.ID()] = entitySlot.Index()

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Entered)
	ec.UnsafeEntity(entity).SetEnteredHandle(entitySlot.Index(), entitySlot.Version())
	ec.UnsafeEntity(entity).SetTreeNodeState(ec.TreeNodeState_Free)

	mgr.observeEntity(entity)

	_EmitEventEntityManagerAddEntity(mgr, mgr, entity)

	return nil
}

// RemoveEntity 按 ID 请求销毁实体；实体不存在时不执行任何操作。
func (mgr *_EntityManager) RemoveEntity(id uid.ID) {
	slotIdx, ok := mgr.entityIDIndex[id]
	if !ok {
		return
	}
	entity := mgr.entityList.Get(slotIdx).V
	entity.Destroy()
}

// GetEntity 按 ID 查询本地实体。
func (mgr *_EntityManager) GetEntity(id uid.ID) (ec.Entity, bool) {
	slotIdx, ok := mgr.entityIDIndex[id]
	if !ok {
		return nil, false
	}
	return mgr.entityList.Get(slotIdx).V, true
}

// RangeEntities 按加入顺序遍历实体，回调返回 false 时停止。
func (mgr *_EntityManager) RangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.Traversal(func(slot *generic.FreeSlot[ec.Entity]) bool {
		return fun.UnsafeCall(slot.V)
	})
}

// EachEntities 按加入顺序遍历全部实体。
func (mgr *_EntityManager) EachEntities(fun generic.Action1[ec.Entity]) {
	mgr.entityList.TraversalEach(func(slot *generic.FreeSlot[ec.Entity]) {
		fun.UnsafeCall(slot.V)
	})
}

// ReversedRangeEntities 按加入顺序逆向遍历实体，回调返回 false 时停止。
func (mgr *_EntityManager) ReversedRangeEntities(fun generic.Func1[ec.Entity, bool]) {
	mgr.entityList.ReversedTraversal(func(slot *generic.FreeSlot[ec.Entity]) bool {
		return fun.UnsafeCall(slot.V)
	})
}

// ReversedEachEntities 按加入顺序逆向遍历全部实体。
func (mgr *_EntityManager) ReversedEachEntities(fun generic.Action1[ec.Entity]) {
	mgr.entityList.ReversedTraversalEach(func(slot *generic.FreeSlot[ec.Entity]) {
		fun.UnsafeCall(slot.V)
	})
}

// FilterEntities 按加入顺序返回符合条件的实体。
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

// ListEntities 按加入顺序返回实体切片副本。
func (mgr *_EntityManager) ListEntities() []ec.Entity {
	return mgr.entityList.ToSlice()
}

// CountEntities 返回当前实体数。
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
	mgr.entityIDIndex = map[uid.ID]int{}
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
	if entity.ID().IsNil() {
		ec.UnsafeEntity(entity).SetID(uid.New())
	}
	ec.UnsafeEntity(entity).SetContext(mgr.ctx)

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

	ec.UnsafeEntity(entity).ComponentList().TraversalEach(func(slot *generic.FreeSlot[ec.Component]) {
		comp := slot.V
		mgr.initComponent(entity, comp)
	})
}

func (mgr *_EntityManager) initComponent(entity ec.Entity, comp ec.Component) {
	event.UnsafeEvent(comp.EventComponentEnableChanged()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())
	event.UnsafeEvent(comp.EventComponentDestroy()).Ctrl().SetPanicHandling(mgr.ctx.AutoRecover(), mgr.ctx.ReportError())

	if ec.UnsafeEntity(entity).Options().ComponentUniqueID {
		if comp.ID().IsNil() {
			ec.UnsafeComponent(comp).SetID(uid.New())
		}
	} else {
		ec.UnsafeComponent(comp).SetID(entity.ID())
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

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Leaving)

	mgr.onEntityDestroyRemoveNode(entity.ID())

	_EmitEventEntityManagerRemoveEntity(mgr, mgr, entity)

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Dead)

	delete(mgr.entityIDIndex, entity.ID())
	mgr.entityList.ReleaseIfVersion(idx, ver)

	if entity.Scope() == ec.Scope_Global {
		service.Current(mgr).EntityManager().RemoveEntity(entity.ID())
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Destroyed)
}

func checkEntitySlot(slot *generic.FreeSlot[ec.Entity], ver int64) bool {
	return slot != nil && !slot.Orphaned() && !slot.Freed() && slot.Version() == ver
}
