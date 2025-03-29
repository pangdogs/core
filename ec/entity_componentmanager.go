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
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
	"slices"
)

// iComponentManager 组件管理器接口
type iComponentManager interface {
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
	// ReversedRangeComponents 反向遍历所有组件
	ReversedRangeComponents(fun generic.Func1[Component, bool])
	// FilterComponents 过滤并获取组件
	FilterComponents(fun generic.Func1[Component, bool]) []Component
	// GetComponents 获取所有组件
	GetComponents() []Component
	// CountComponents 统计所有组件数量
	CountComponents() int
	// AddComponent 添加组件，允许组件同名
	AddComponent(name string, components ...Component) error
	// RemoveComponent 使用名称删除组件，同名组件均会删除
	RemoveComponent(name string)
	// RemoveComponentById 使用组件Id删除组件（需要开启为实体组件分配唯一Id特性）
	RemoveComponentById(id uid.Id)

	getComponentNameIndex() *generic.SliceMap[string, *generic.Node[Component]]
	getComponentList() *generic.List[Component]
	removeComponentByRef(comp Component)

	IEntityComponentManagerEventTab
}

// GetComponent 使用名称查询组件，组件同名时，返回首个组件
func (entity *EntityBehavior) GetComponent(name string) Component {
	if compNode, ok := entity.getComponentNode(name); ok {
		return entity.touchComponent(compNode.V)
	}
	return nil
}

// GetComponentById 使用组件Id查询组件（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) GetComponentById(id uid.Id) Component {
	if compNode, ok := entity.getComponentNodeById(id); ok {
		return entity.touchComponent(compNode.V)
	}
	return nil
}

// GetComponentByPT 使用组件原型查询组件
func (entity *EntityBehavior) GetComponentByPT(prototype string) Component {
	if compNode, ok := entity.getComponentNodeByPT(prototype); ok {
		return entity.touchComponent(compNode.V)
	}
	return nil
}

// ContainsComponent 组件是否存在
func (entity *EntityBehavior) ContainsComponent(name string) bool {
	_, ok := entity.getComponentNode(name)
	return ok
}

// ContainsComponentById 使用组件Id检测组件是否存在（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) ContainsComponentById(id uid.Id) bool {
	_, ok := entity.getComponentNodeById(id)
	return ok
}

// ContainsComponentByPT 使用组件原型查询组件
func (entity *EntityBehavior) ContainsComponentByPT(prototype string) bool {
	_, ok := entity.getComponentNodeByPT(prototype)
	return ok
}

// RangeComponents 遍历所有组件
func (entity *EntityBehavior) RangeComponents(fun generic.Func1[Component, bool]) {
	entity.components.Traversal(func(compNode *generic.Node[Component]) bool {
		comp := entity.touchComponent(compNode.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// ReversedRangeComponents 反向遍历所有组件
func (entity *EntityBehavior) ReversedRangeComponents(fun generic.Func1[Component, bool]) {
	entity.components.ReversedTraversal(func(compNode *generic.Node[Component]) bool {
		comp := entity.touchComponent(compNode.V)
		if comp == nil {
			return true
		}
		return fun.UnsafeCall(comp)
	})
}

// FilterComponents 过滤并获取组件
func (entity *EntityBehavior) FilterComponents(fun generic.Func1[Component, bool]) []Component {
	var components []Component

	entity.components.Traversal(func(compNode *generic.Node[Component]) bool {
		comp := compNode.V
		if fun.UnsafeCall(comp) {
			components = append(components, comp)
		}
		return true
	})

	for i := range components {
		if entity.touchComponent(components[i]) == nil {
			components[i] = nil
		}
	}

	components = slices.DeleteFunc(components, func(comp Component) bool {
		return comp == nil
	})

	return components
}

// GetComponents 获取所有组件
func (entity *EntityBehavior) GetComponents() []Component {
	components := make([]Component, 0, entity.components.Len())

	entity.components.Traversal(func(compNode *generic.Node[Component]) bool {
		components = append(components, compNode.V)
		return true
	})

	for i := range components {
		if entity.touchComponent(components[i]) == nil {
			components[i] = nil
		}
	}

	components = slices.DeleteFunc(components, func(comp Component) bool {
		return comp == nil
	})

	return components
}

// CountComponents 统计所有组件数量
func (entity *EntityBehavior) CountComponents() int {
	return entity.components.Len()
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

	_EmitEventComponentManagerAddComponents(entity, entity.opts.InstanceFace.Iface, components)
	return nil
}

// RemoveComponent 使用名称删除组件，同名组件均会删除
func (entity *EntityBehavior) RemoveComponent(name string) {
	compNode, ok := entity.getComponentNode(name)
	if !ok {
		return
	}

	entity.components.TraversalAt(func(compNode *generic.Node[Component]) bool {
		comp := compNode.V

		if comp.GetName() != name {
			return false
		}

		if !comp.GetRemovable() {
			return true
		}

		if comp.GetState() > ComponentState_Alive {
			return true
		}

		comp.setState(ComponentState_Detach)

		_EmitEventComponentManagerRemoveComponent(entity, entity.opts.InstanceFace.Iface, comp)

		if comp.GetState() >= ComponentState_Destroyed {
			compNode.Escape()
			entity.updateComponentNameIndex(comp.GetName())
		}

		return true
	}, compNode)
}

// RemoveComponentById 使用组件Id删除组件（需要开启为实体组件分配唯一Id特性）
func (entity *EntityBehavior) RemoveComponentById(id uid.Id) {
	compNode, ok := entity.getComponentNodeById(id)
	if !ok {
		return
	}

	comp := compNode.V

	if !comp.GetRemovable() {
		return
	}

	if comp.GetState() > ComponentState_Alive {
		return
	}

	comp.setState(ComponentState_Detach)

	_EmitEventComponentManagerRemoveComponent(entity, entity.opts.InstanceFace.Iface, comp)

	if comp.GetState() >= ComponentState_Destroyed {
		compNode.Escape()
		entity.updateComponentNameIndex(comp.GetName())
	}
}

// EventComponentManagerAddComponents 事件：实体的组件管理器添加组件
func (entity *EntityBehavior) EventComponentManagerAddComponents() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerRemoveComponent()
}

// EventComponentManagerRemoveComponent 事件：实体的组件管理器删除组件
func (entity *EntityBehavior) EventComponentManagerRemoveComponent() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerRemoveComponent()
}

// EventComponentManagerFirstTouchComponent 事件：实体的组件管理器首次访问组件
func (entity *EntityBehavior) EventComponentManagerFirstTouchComponent() event.IEvent {
	return entity.entityComponentManagerEventTab.EventComponentManagerFirstTouchComponent()
}

func (entity *EntityBehavior) getComponentNameIndex() *generic.SliceMap[string, *generic.Node[Component]] {
	return &entity.componentNameIndex
}

func (entity *EntityBehavior) getComponentList() *generic.List[Component] {
	return &entity.components
}

func (entity *EntityBehavior) removeComponentByRef(comp Component) {
	compNode, ok := entity.getComponentNodeByRef(comp)
	if !ok {
		return
	}

	if !comp.GetRemovable() {
		return
	}

	if comp.GetState() > ComponentState_Alive {
		return
	}

	comp.setState(ComponentState_Detach)

	_EmitEventComponentManagerRemoveComponent(entity, entity.opts.InstanceFace.Iface, comp)

	if comp.GetState() >= ComponentState_Destroyed {
		compNode.Escape()
		entity.updateComponentNameIndex(comp.GetName())
	}
}

func (entity *EntityBehavior) addComponent(name string, component Component) {
	component.init(name, entity.opts.InstanceFace.Iface, component)

	if at, ok := entity.getComponentNode(name); ok {
		entity.components.TraversalAt(func(compNode *generic.Node[Component]) bool {
			if compNode.V.GetName() == name {
				at = compNode
				return true
			}
			return false
		}, at)

		entity.components.InsertAfter(component, at)

	} else {
		compNode := entity.components.PushBack(component)

		if entity.opts.ComponentNameIndexing {
			entity.componentNameIndex.Add(name, compNode)
		}
	}

	component.setState(ComponentState_Attach)
}

func (entity *EntityBehavior) getComponentNode(name string) (*generic.Node[Component], bool) {
	if entity.opts.ComponentNameIndexing {
		return entity.componentNameIndex.Get(name)
	}

	var compNode *generic.Node[Component]

	entity.components.Traversal(func(node *generic.Node[Component]) bool {
		if node.V.GetName() == name {
			compNode = node
			return false
		}
		return true
	})

	return compNode, compNode != nil
}

func (entity *EntityBehavior) getComponentNodeById(id uid.Id) (*generic.Node[Component], bool) {
	var compNode *generic.Node[Component]

	entity.components.Traversal(func(node *generic.Node[Component]) bool {
		if node.V.GetId() == id {
			compNode = node
			return false
		}
		return true
	})

	return compNode, compNode != nil
}

func (entity *EntityBehavior) getComponentNodeByPT(prototype string) (*generic.Node[Component], bool) {
	var compNode *generic.Node[Component]

	entity.components.Traversal(func(node *generic.Node[Component]) bool {
		if node.V.GetBuiltin().PT.Prototype() == prototype {
			compNode = node
			return false
		}
		return true
	})

	return compNode, compNode != nil
}

func (entity *EntityBehavior) getComponentNodeByRef(comp Component) (*generic.Node[Component], bool) {
	var compNode *generic.Node[Component]

	entity.components.Traversal(func(node *generic.Node[Component]) bool {
		if node.V == comp {
			compNode = node
			return false
		}
		return true
	})

	return compNode, compNode != nil
}

func (entity *EntityBehavior) touchComponent(comp Component) Component {
	if entity.opts.ComponentAwakeOnFirstTouch && comp.GetState() == ComponentState_Attach {
		_EmitEventComponentManagerFirstTouchComponent(entity, entity.opts.InstanceFace.Iface, comp)
	}

	if comp.GetState() >= ComponentState_Destroyed {
		return nil
	}

	return comp
}

func (entity *EntityBehavior) updateComponentNameIndex(name string) {
	if !entity.opts.ComponentNameIndexing {
		return
	}

	if compNode, ok := entity.getComponentNode(name); ok {
		entity.componentNameIndex.Add(name, compNode)
	} else {
		entity.componentNameIndex.Delete(name)
	}
}
