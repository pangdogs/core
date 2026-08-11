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

package ec

import (
	"fmt"
	"slices"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
)

// iComponentManager 定义实体的组件管理能力。
//
// 所有操作都应在实体所属 Runtime 的运行协程中执行。启用首次访问唤醒后，查询和
// 遍历方法可能推进组件生命周期，因而不是纯只读操作。
type iComponentManager interface {
	iiComponentManager

	// AddComponent 将 Born 状态的组件加入实体；允许同名组件。
	AddComponent(name string, components ...Component) error
	// RemoveComponent 按名称请求删除全部同名组件。
	RemoveComponent(name string)
	// RemoveComponentById 按 ID 请求删除组件；仅在启用组件唯一 ID 时有效。
	RemoveComponentById(id uid.Id)
	// RemoveComponentByPT 按原型名请求删除全部匹配组件。
	RemoveComponentByPT(prototype string)
	// GetComponent 返回首个同名组件；不存在时返回 nil。
	GetComponent(name string) Component
	// GetComponentById 按 ID 查询组件；未启用组件唯一 ID 或不存在时返回 nil。
	GetComponentById(id uid.Id) Component
	// GetComponentByPT 返回首个使用指定原型的组件；不存在时返回 nil。
	GetComponentByPT(prototype string) Component
	// GetComponents 返回全部同名组件的快照。
	GetComponents(name string) []Component
	// GetComponentsByPT 返回全部使用指定原型的组件快照。
	GetComponentsByPT(prototype string) []Component
	// RangeComponents 按加入顺序遍历组件；回调返回 false 时停止。
	RangeComponents(fun generic.Func1[Component, bool])
	// EachComponents 按加入顺序遍历全部组件。
	EachComponents(fun generic.Action1[Component])
	// ReversedRangeComponents 按加入顺序的逆序遍历组件；回调返回 false 时停止。
	ReversedRangeComponents(fun generic.Func1[Component, bool])
	// ReversedEachComponents 按加入顺序的逆序遍历全部组件。
	ReversedEachComponents(fun generic.Action1[Component])
	// FilterComponents 返回满足条件的组件快照。
	FilterComponents(fun generic.Func1[Component, bool]) []Component
	// ListComponents 返回全部组件的快照。
	ListComponents() []Component
	// CountComponents 返回当前依附于实体的组件数。
	CountComponents() int

	IEntityComponentManagerEventTab
}

type iiComponentManager interface {
	getVersion() int64
	getComponentNameIndex() *generic.SliceMap[string, int]
	getComponentList() *generic.FreeList[Component]
	onComponentEnableChangedIfVersion(idx int, ver int64)
	onComponentDestroyIfVersion(idx int, ver int64)
}

// AddComponent 将 Born 状态的组件加入实体；允许同名组件。
//
// components 为空、包含 nil 或组件不处于 Born 状态时返回错误。实体已启动时，
// Runtime 会通过添加事件同步推进新组件的生命周期。
func (entity *EntityBehavior) AddComponent(name string, components ...Component) error {
	if len(components) <= 0 {
		return fmt.Errorf("%w: %w: components is empty", ErrEC, exception.ErrArgs)
	}

	for i := range components {
		comp := components[i]

		if comp == nil {
			return fmt.Errorf("%w: %w: component is nil", ErrEC, exception.ErrArgs)
		}

		if comp.State() != ComponentState_Born {
			return fmt.Errorf("%w: invalid component state %q", ErrEC, comp.State())
		}
	}

	for i := range components {
		entity.addComponent(name, components[i])
	}

	_EmitEventComponentManagerAddComponents(entity, entity.getInstance(), components)

	return nil
}

// RemoveComponent 按名称请求删除全部可删除的同名组件。
func (entity *EntityBehavior) RemoveComponent(name string) {
	at, ok := entity.getComponentSlot(name)
	if !ok {
		return
	}

	entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
		comp := slot.V

		if comp.Name() != name {
			return false
		}

		comp.Destroy()

		return true
	}, at.Index())
}

// RemoveComponentById 按 ID 请求删除组件；未启用组件唯一 ID 时无效。
func (entity *EntityBehavior) RemoveComponentById(id uid.Id) {
	slot, ok := entity.getComponentSlotById(id)
	if !ok {
		return
	}
	comp := slot.V
	comp.Destroy()
}

// RemoveComponentByPT 按原型名请求删除全部可删除的匹配组件。
func (entity *EntityBehavior) RemoveComponentByPT(prototype string) {
	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := slot.V

		if comp.Builtin().PT.Prototype() != prototype {
			return
		}

		comp.Destroy()
	})
}

// GetComponent 返回首个同名组件；不存在时返回 nil。
func (entity *EntityBehavior) GetComponent(name string) Component {
	if slot, ok := entity.getComponentSlot(name); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// GetComponentById 按 ID 查询组件；未启用组件唯一 ID 或不存在时返回 nil。
func (entity *EntityBehavior) GetComponentById(id uid.Id) Component {
	if slot, ok := entity.getComponentSlotById(id); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// GetComponentByPT 返回首个使用指定原型的组件；不存在时返回 nil。
func (entity *EntityBehavior) GetComponentByPT(prototype string) Component {
	if slot, ok := entity.getComponentSlotByPT(prototype); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// GetComponents 返回全部同名组件的快照。
func (entity *EntityBehavior) GetComponents(name string) []Component {
	at, ok := entity.getComponentSlot(name)
	if !ok {
		return nil
	}

	var components []Component

	entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
		comp := slot.V

		if comp.Name() != name {
			return false
		}

		comp = entity.touchComponent(comp)
		if comp == nil {
			return true
		}

		components = append(components, comp)

		return true
	}, at.Index())

	return components
}

// GetComponentsByPT 返回全部使用指定原型的组件快照。
func (entity *EntityBehavior) GetComponentsByPT(prototype string) []Component {
	var components []Component

	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := slot.V

		if comp.Builtin().PT.Prototype() != prototype {
			return
		}

		comp = entity.touchComponent(comp)
		if comp == nil {
			return
		}

		components = append(components, comp)
	})

	return components
}

// RangeComponents 按加入顺序遍历组件；回调返回 false 时停止。
func (entity *EntityBehavior) RangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.Traversal(func(slot *generic.FreeSlot[Component]) bool {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// EachComponents 按加入顺序遍历全部组件。
func (entity *EntityBehavior) EachComponents(fun generic.Action1[Component]) {
	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return
		}
		fun.UnsafeCall(comp)
	})
}

// ReversedRangeComponents 按加入顺序的逆序遍历组件；回调返回 false 时停止。
func (entity *EntityBehavior) ReversedRangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.ReversedTraversal(func(slot *generic.FreeSlot[Component]) bool {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// ReversedEachComponents 按加入顺序的逆序遍历全部组件。
func (entity *EntityBehavior) ReversedEachComponents(fun generic.Action1[Component]) {
	entity.componentList.ReversedTraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return
		}
		fun.UnsafeCall(comp)
	})
}

// FilterComponents 返回满足条件的组件快照，并对返回项执行首次访问处理。
func (entity *EntityBehavior) FilterComponents(fun generic.Func1[Component, bool]) []Component {
	var components []Component

	ver := entity.componentList.Version()
	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		if slot.Version() > ver {
			return
		}
		comp := slot.V
		if fun.UnsafeCall(comp) {
			components = append(components, comp)
		}
	})

	for i := range components {
		entity.touchComponent(components[i])
	}

	components = slices.DeleteFunc(components, func(comp Component) bool {
		idx, ver := comp.getAttachedHandle()
		slot := entity.componentList.Get(idx)
		return !checkComponentSlot(slot, ver)
	})

	return components
}

// ListComponents 返回全部组件的快照，并对返回项执行首次访问处理。
func (entity *EntityBehavior) ListComponents() []Component {
	components := entity.componentList.ToSlice()

	for i := range components {
		entity.touchComponent(components[i])
	}

	components = slices.DeleteFunc(components, func(comp Component) bool {
		idx, ver := comp.getAttachedHandle()
		slot := entity.componentList.Get(idx)
		return !checkComponentSlot(slot, ver)
	})

	return components
}

// CountComponents 返回当前依附于实体的组件数。
func (entity *EntityBehavior) CountComponents() int {
	return entity.componentList.Len() - entity.componentList.OrphanCount()
}

// EventComponentManagerAddComponents 返回组件批量添加事件。
func (entity *EntityBehavior) EventComponentManagerAddComponents() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerAddComponents()
}

// EventComponentManagerRemoveComponent 返回组件移除事件。
func (entity *EntityBehavior) EventComponentManagerRemoveComponent() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerRemoveComponent()
}

// EventComponentManagerComponentEnableChanged 返回组件启用状态变更事件。
func (entity *EntityBehavior) EventComponentManagerComponentEnableChanged() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerComponentEnableChanged()
}

// EventComponentManagerFirstTouchComponent 返回组件首次访问事件。
func (entity *EntityBehavior) EventComponentManagerFirstTouchComponent() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerFirstTouchComponent()
}

func (entity *EntityBehavior) getVersion() int64 {
	return entity.componentList.Version()
}

func (entity *EntityBehavior) getComponentNameIndex() *generic.SliceMap[string, int] {
	return &entity.componentNameIndex
}

func (entity *EntityBehavior) getComponentList() *generic.FreeList[Component] {
	return &entity.componentList
}

func (entity *EntityBehavior) onComponentEnableChangedIfVersion(idx int, ver int64) {
	slot := entity.componentList.Get(idx)
	if !checkComponentSlot(slot, ver) {
		return
	}

	comp := slot.V

	_EmitEventComponentManagerComponentEnableChanged(entity, entity.getInstance(), comp, comp.Enabled())
}

func (entity *EntityBehavior) onComponentDestroyIfVersion(idx int, ver int64) {
	compSlot := entity.componentList.Get(idx)
	if !checkComponentSlot(compSlot, ver) {
		return
	}

	comp := compSlot.V

	if !comp.Removable() {
		return
	}

	comp.setState(ComponentState_Detaching)

	_EmitEventComponentManagerRemoveComponent(entity, entity.getInstance(), comp)

	comp.setState(ComponentState_Dead)

	nameIdx, ok := entity.componentNameIndex.Get(comp.Name())
	if ok && nameIdx == idx {
		var nextSlot *generic.FreeSlot[Component]

		entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
			if slot == compSlot {
				return true
			}
			if slot.V.Name() == comp.Name() {
				nextSlot = slot
			}
			return false
		}, idx)

		if nextSlot != nil {
			entity.componentNameIndex.Add(nextSlot.V.Name(), nextSlot.Index())
		} else {
			entity.componentNameIndex.Delete(comp.Name())
		}
	}

	entity.componentList.ReleaseIfVersion(idx, ver)

	comp.setState(ComponentState_Destroyed)
}

func (entity *EntityBehavior) getComponentSlot(name string) (*generic.FreeSlot[Component], bool) {
	slotIdx, ok := entity.componentNameIndex.Get(name)
	if !ok {
		return nil, false
	}
	return entity.componentList.Get(slotIdx), true
}

func (entity *EntityBehavior) getComponentSlotById(id uid.Id) (*generic.FreeSlot[Component], bool) {
	if !entity.options.ComponentUniqueID {
		return nil, false
	}

	var compSlot *generic.FreeSlot[Component]

	entity.componentList.Traversal(func(slot *generic.FreeSlot[Component]) bool {
		if slot.V.Id() == id {
			compSlot = slot
			return false
		}
		return true
	})

	return compSlot, compSlot != nil
}

func (entity *EntityBehavior) getComponentSlotByPT(prototype string) (*generic.FreeSlot[Component], bool) {
	var compSlot *generic.FreeSlot[Component]

	entity.componentList.Traversal(func(slot *generic.FreeSlot[Component]) bool {
		if slot.V.Builtin().PT.Prototype() == prototype {
			compSlot = slot
			return false
		}
		return true
	})

	return compSlot, compSlot != nil
}

func (entity *EntityBehavior) addComponent(name string, component Component) {
	component.init(name, entity.getInstance(), component)
	if entity.asyncScope != nil {
		component.setContext(entity.getInstance())
	}

	var compSlot *generic.FreeSlot[Component]

	if at, ok := entity.getComponentSlot(name); ok {
		entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
			if slot.V.Name() == name {
				at = slot
				return true
			}
			return false
		}, at.Index())

		compSlot = entity.componentList.InsertAfter(component, at.Index())

	} else {
		compSlot = entity.componentList.PushBack(component)
		entity.componentNameIndex.Add(name, compSlot.Index())
	}

	component.setState(ComponentState_Attached)
	component.setAttachedHandle(compSlot.Index(), compSlot.Version())
}

func (entity *EntityBehavior) touchComponent(comp Component) Component {
	if entity.options.ComponentAwakeOnFirstTouch && comp.State() == ComponentState_Attached {
		_EmitEventComponentManagerFirstTouchComponent(entity, entity.getInstance(), comp)
	}

	idx, ver := comp.getAttachedHandle()
	slot := entity.componentList.Get(idx)
	if !checkComponentSlot(slot, ver) {
		return nil
	}

	return comp
}

func checkComponentSlot(slot *generic.FreeSlot[Component], ver int64) bool {
	return slot != nil && !slot.Orphaned() && !slot.Freed() && slot.Version() == ver
}
