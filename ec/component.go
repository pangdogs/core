package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
	"reflect"
)

// Component 组件接口
type Component interface {
	_Component
	concurrent.CurrentContextProvider
	fmt.Stringer

	// GetId 获取组件Id
	GetId() uid.Id
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// DestroySelf 销毁自身
	DestroySelf()
}

type _Component interface {
	init(name string, entity Entity, composite Component, hookAllocator container.Allocator[event.Hook], gcCollector container.GCCollector)
	setId(id uid.Id)
	setFixed(v bool)
	getFixed() bool
	setState(state ComponentState)
	setReflectValue(v reflect.Value)
	getReflectValue() reflect.Value
	getComposite() Component
	setGCCollector(gcCollector container.GCCollector)
	eventComponentDestroySelf() event.IEvent
	cleanHooks()
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         uid.Id
	name                       string
	entity                     Entity
	composite                  Component
	fixed                      bool
	state                      ComponentState
	reflectValue               reflect.Value
	_eventComponentDestroySelf event.Event
	hooks                      []event.Hook
}

// GetId 获取组件Id
func (comp *ComponentBehavior) GetId() uid.Id {
	return comp.id
}

// GetName 获取组件名称
func (comp *ComponentBehavior) GetName() string {
	return comp.name
}

// GetEntity 获取组件依附的实体
func (comp *ComponentBehavior) GetEntity() Entity {
	return comp.entity
}

// GetState 获取组件状态
func (comp *ComponentBehavior) GetState() ComponentState {
	return comp.state
}

// DestroySelf 销毁自身
func (comp *ComponentBehavior) DestroySelf() {
	switch comp.GetState() {
	case ComponentState_Awake, ComponentState_Start, ComponentState_Living:
		emitEventComponentDestroySelf(UnsafeComponent(comp), comp.composite)
	}
}

// GetContext 获取上下文
func (comp *ComponentBehavior) GetContext() iface.Cache {
	return comp.entity.GetContext()
}

// GetCurrentContext 获取当前上下文
func (comp *ComponentBehavior) GetCurrentContext() iface.Cache {
	return comp.entity.GetCurrentContext()
}

// String implements fmt.Stringer
func (comp *ComponentBehavior) String() string {
	return fmt.Sprintf(`{"id":%q "name":%q "entity_id":%q "state":%q}`, comp.GetId(), comp.GetName(), comp.GetEntity().GetId(), comp.GetState())
}

func (comp *ComponentBehavior) init(name string, entity Entity, composite Component, hookAllocator container.Allocator[event.Hook], gcCollector container.GCCollector) {
	comp.name = name
	comp.entity = entity
	comp.composite = composite
	comp._eventComponentDestroySelf.Init(false, nil, event.EventRecursion_NotEmit, hookAllocator, gcCollector)
}

func (comp *ComponentBehavior) setId(id uid.Id) {
	comp.id = id
}

func (comp *ComponentBehavior) setFixed(v bool) {
	comp.fixed = v
}

func (comp *ComponentBehavior) getFixed() bool {
	return comp.fixed
}

func (comp *ComponentBehavior) setState(state ComponentState) {
	if state <= comp.state {
		return
	}

	comp.state = state

	switch comp.state {
	case ComponentState_Detach:
		comp._eventComponentDestroySelf.Close()
	}
}

func (comp *ComponentBehavior) setReflectValue(v reflect.Value) {
	comp.reflectValue = v
}

func (comp *ComponentBehavior) getReflectValue() reflect.Value {
	if comp.reflectValue.IsValid() {
		return comp.reflectValue
	}

	comp.reflectValue = reflect.ValueOf(comp.composite)

	return comp.reflectValue
}

func (comp *ComponentBehavior) getComposite() Component {
	return comp.composite
}

func (comp *ComponentBehavior) setGCCollector(gcCollector container.GCCollector) {
	event.UnsafeEvent(&comp._eventComponentDestroySelf).SetGCCollector(gcCollector)
}

func (comp *ComponentBehavior) eventComponentDestroySelf() event.IEvent {
	return &comp._eventComponentDestroySelf
}
