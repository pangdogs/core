package ec

import (
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// Component 组件接口
type Component interface {
	iComponent
	gctx.CurrentContextProvider
	fmt.Stringer

	// GetId 获取组件Id
	GetId() uid.Id
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// DestroySelf 销毁自身
	DestroySelf()
}

type iComponent interface {
	init(name string, entity Entity, composite Component)
	setId(id uid.Id)
	setFixed(v bool)
	getFixed() bool
	setState(state ComponentState)
	setReflected(v reflect.Value)
	getComposite() Component
	eventComponentDestroySelf() event.IEvent
	cleanManagedHooks()
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         uid.Id
	name                       string
	entity                     Entity
	composite                  Component
	fixed                      bool
	state                      ComponentState
	reflected                  reflect.Value
	_eventComponentDestroySelf event.Event
	managedHooks               []event.Hook
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

// GetReflected 获取反射值
func (comp *ComponentBehavior) GetReflected() reflect.Value {
	if comp.reflected.IsValid() {
		return comp.reflected
	}
	comp.reflected = reflect.ValueOf(comp.composite)
	return comp.reflected
}

// DestroySelf 销毁自身
func (comp *ComponentBehavior) DestroySelf() {
	switch comp.GetState() {
	case ComponentState_Awake, ComponentState_Start, ComponentState_Alive:
		_EmitEventComponentDestroySelf(UnsafeComponent(comp), comp.composite)
	}
}

// GetCurrentContext 获取当前上下文
func (comp *ComponentBehavior) GetCurrentContext() iface.Cache {
	return comp.entity.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (comp *ComponentBehavior) GetConcurrentContext() iface.Cache {
	return comp.entity.GetConcurrentContext()
}

// String implements fmt.Stringer
func (comp *ComponentBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q, "entity_id":%q}`, comp.GetId(), comp.GetName(), comp.GetEntity().GetId())
}

func (comp *ComponentBehavior) init(name string, entity Entity, composite Component) {
	comp.name = name
	comp.entity = entity
	comp.composite = composite
	comp._eventComponentDestroySelf.Init(false, nil, event.EventRecursion_Discard)
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

func (comp *ComponentBehavior) setReflected(v reflect.Value) {
	comp.reflected = v
}

func (comp *ComponentBehavior) getComposite() Component {
	return comp.composite
}

func (comp *ComponentBehavior) eventComponentDestroySelf() event.IEvent {
	return &comp._eventComponentDestroySelf
}
