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

// iComponentManager 组件管理器接口
type iComponentManager interface {
	iiComponentManager

	// GetComponent 使用名称查询组件，组件同名时，返回首个组件
	GetComponent(name string) Component
	// GetComponentById 使用组件Id查询组件（需要开启为实体组件分配唯一Id特性）
	GetComponentById(id uid.Id) Component
	// GetComponentByPT 使用组件原型查询组件
	GetComponentByPT(prototype string) Component
	// ContainsComponent 组件是否存在
	ContainsComponent(name string) bool
	// ContainsComponentById 使用组件Id检测组件是否存在（需要开启为实体组件分配唯一Id特性）
	ContainsComponentById(id uid.Id) bool
	// ContainsComponentByPT 使用组件原型查询组件
	ContainsComponentByPT(prototype string) bool
	// RangeComponents 遍历所有组件
	RangeComponents(fun generic.Func1[Component, bool])
	// EachComponents 遍历每个组件
	EachComponents(fun generic.Action1[Component])
	// ReversedRangeComponents 反向遍历所有组件
	ReversedRangeComponents(fun generic.Func1[Component, bool])
	// ReversedEachComponents 反向遍历每个组件
	ReversedEachComponents(fun generic.Action1[Component])
	// FilterComponents 过滤并获取组件
	FilterComponents(fun generic.Func1[Component, bool]) []Component
	// ListComponents 获取所有组件
	ListComponents() []Component
	// CountComponents 统计所有组件数量
	CountComponents() int
	// AddComponent 添加组件，允许组件同名
	AddComponent(name string, components ...Component) error
	// RemoveComponent 使用名称删除组件，同名组件均会删除
	RemoveComponent(name string)
	// RemoveComponentById 使用组件Id删除组件（需要开启为实体组件分配唯一Id特性）
	RemoveComponentById(id uid.Id)

	IEntityComponentManagerEventTab
}

type iiComponentManager interface {
	getVersion() int64
	getComponentNameIndex() *generic.SliceMap[string, int]
	getComponentList() *generic.FreeList[Component]
	onComponentEnableChangedIfVersion(idx int, ver int64)
	removeComponentIfVersion(idx int, ver int64)
}

// GetComponent 使用名称查询组件，组件同名时，返回首个组件
func (entity *EntityBehavior) GetComponent(name string) Component {
	if slot, ok := entity.getComponentSlot(name); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// GetComponentById 使用组件Id查询组件（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) GetComponentById(id uid.Id) Component {
	if slot, ok := entity.getComponentSlotById(id); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// GetComponentByPT 使用组件原型查询组件
func (entity *EntityBehavior) GetComponentByPT(prototype string) Component {
	if slot, ok := entity.getComponentSlotByPT(prototype); ok {
		return entity.touchComponent(slot.V)
	}
	return nil
}

// ContainsComponent 组件是否存在
func (entity *EntityBehavior) ContainsComponent(name string) bool {
	_, ok := entity.getComponentSlot(name)
	return ok
}

// ContainsComponentById 使用组件Id检测组件是否存在（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) ContainsComponentById(id uid.Id) bool {
	_, ok := entity.getComponentSlotById(id)
	return ok
}

// ContainsComponentByPT 使用组件原型查询组件
func (entity *EntityBehavior) ContainsComponentByPT(prototype string) bool {
	_, ok := entity.getComponentSlotByPT(prototype)
	return ok
}

// RangeComponents 遍历所有组件
func (entity *EntityBehavior) RangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.Traversal(func(slot *generic.FreeSlot[Component]) bool {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// EachComponents 遍历每个组件
func (entity *EntityBehavior) EachComponents(fun generic.Action1[Component]) {
	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return
		}
		fun.UnsafeCall(comp)
	})
}

// ReversedRangeComponents 反向遍历所有组件
func (entity *EntityBehavior) ReversedRangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.ReversedTraversal(func(slot *generic.FreeSlot[Component]) bool {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// ReversedEachComponents 反向遍历每个组件
func (entity *EntityBehavior) ReversedEachComponents(fun generic.Action1[Component]) {
	entity.componentList.ReversedTraversalEach(func(slot *generic.FreeSlot[Component]) {
		comp := entity.touchComponent(slot.V)
		if comp == nil {
			return
		}
		fun.UnsafeCall(comp)
	})
}

// FilterComponents 过滤并获取组件
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

// ListComponents 获取所有组件
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

// CountComponents 统计所有组件数量
func (entity *EntityBehavior) CountComponents() int {
	return entity.componentList.Len() - entity.componentList.OrphanCount()
}

// AddComponent 添加组件，允许组件同名
func (entity *EntityBehavior) AddComponent(name string, components ...Component) error {
	if len(components) <= 0 {
		return fmt.Errorf("%w: %w: components is empty", ErrEC, exception.ErrArgs)
	}

	for i := range components {
		comp := components[i]

		if comp == nil {
			return fmt.Errorf("%w: %w: component is nil", ErrEC, exception.ErrArgs)
		}

		if comp.GetState() != ComponentState_Birth {
			return fmt.Errorf("%w: invalid component state %q", ErrEC, comp.GetState())
		}
	}

	for i := range components {
		entity.addComponent(name, components[i])
	}

	_EmitEventComponentManagerAddComponents(entity, entity.getInstance(), components)

	return nil
}

// RemoveComponent 使用名称删除组件，同名组件均会删除
func (entity *EntityBehavior) RemoveComponent(name string) {
	at, ok := entity.getComponentSlot(name)
	if !ok {
		return
	}

	entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
		comp := slot.V

		if comp.GetName() != name {
			return false
		}

		comp.Destroy()

		return true
	}, at.Index())
}

// RemoveComponentById 使用组件Id删除组件（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) RemoveComponentById(id uid.Id) {
	slot, ok := entity.getComponentSlotById(id)
	if !ok {
		return
	}
	comp := slot.V
	comp.Destroy()
}

// EventComponentManagerAddComponents 事件：实体的组件管理器添加组件
func (entity *EntityBehavior) EventComponentManagerAddComponents() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerAddComponents()
}

// EventComponentManagerRemoveComponent 事件：实体的组件管理器删除组件
func (entity *EntityBehavior) EventComponentManagerRemoveComponent() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerRemoveComponent()
}

// EventComponentManagerComponentEnableChanged 事件：实体组件管理器中的组件启用状态改变
func (entity *EntityBehavior) EventComponentManagerComponentEnableChanged() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerComponentEnableChanged()
}

// EventComponentManagerFirstTouchComponent 事件：实体的组件管理器首次访问组件
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

	_EmitEventComponentManagerComponentEnableChanged(entity, entity.getInstance(), comp, comp.GetEnable())
}

func (entity *EntityBehavior) removeComponentIfVersion(idx int, ver int64) {
	compSlot := entity.componentList.Get(idx)
	if !checkComponentSlot(compSlot, ver) {
		return
	}

	comp := compSlot.V

	if !comp.GetRemovable() {
		return
	}

	comp.setState(ComponentState_Detach)

	_EmitEventComponentManagerRemoveComponent(entity, entity.getInstance(), comp)

	comp.setState(ComponentState_Death)

	nameIdx, ok := entity.componentNameIndex.Get(comp.GetName())
	if ok && nameIdx == idx {
		var nextSlot *generic.FreeSlot[Component]

		entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
			if slot == compSlot {
				return true
			}
			if slot.V.GetName() == comp.GetName() {
				nextSlot = slot
			}
			return false
		}, idx)

		if nextSlot != nil {
			entity.componentNameIndex.Add(nextSlot.V.GetName(), nextSlot.Index())
		} else {
			entity.componentNameIndex.Delete(comp.GetName())
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
	var compSlot *generic.FreeSlot[Component]

	entity.componentList.Traversal(func(slot *generic.FreeSlot[Component]) bool {
		if slot.V.GetId() == id {
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
		if slot.V.GetBuiltin().PT.Prototype() == prototype {
			compSlot = slot
			return false
		}
		return true
	})

	return compSlot, compSlot != nil
}

func (entity *EntityBehavior) addComponent(name string, component Component) {
	component.init(name, entity.getInstance(), component)

	var compSlot *generic.FreeSlot[Component]

	if at, ok := entity.getComponentSlot(name); ok {
		entity.componentList.TraversalAt(func(slot *generic.FreeSlot[Component]) bool {
			if slot.V.GetName() == name {
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

	component.setState(ComponentState_Attach)
	component.setAttachedHandle(compSlot.Index(), compSlot.Version())
}

func (entity *EntityBehavior) touchComponent(comp Component) Component {
	if entity.options.ComponentAwakeOnFirstTouch && comp.GetState() == ComponentState_Attach {
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
