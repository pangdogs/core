package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
	"reflect"
)

// Component 组件接口
type Component interface {
	_Component
	ContextResolver

	// GetID 获取组件ID
	GetID() ID
	// GetSerialNo 获取序列号
	GetSerialNo() int64
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	// GetState 获取组件状态
	GetState() ComponentState
	// DestroySelf 销毁自身
	DestroySelf()
	// String 字符串化
	String() string
}

type _Component interface {
	init(name string, entity Entity, composite Component, hookAllocator container.Allocator[localevent.Hook], gcCollector container.GCCollector)
	setID(id ID)
	setSerialNo(sn int64)
	setFixed(v bool)
	getFixed() bool
	setState(state ComponentState)
	setReflectValue(v reflect.Value)
	getReflectValue() reflect.Value
	getComposite() Component
	setGCCollector(gcCollector container.GCCollector)
	eventComponentDestroySelf() localevent.IEvent
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         ID
	serialNo                   int64
	name                       string
	entity                     Entity
	composite                  Component
	fixed                      bool
	state                      ComponentState
	reflectValue               reflect.Value
	_eventComponentDestroySelf localevent.Event
}

// GetID 获取组件ID
func (comp *ComponentBehavior) GetID() ID {
	return comp.id
}

// GetSerialNo 获取序列号
func (comp *ComponentBehavior) GetSerialNo() int64 {
	return comp.serialNo
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
		emitEventComponentDestroySelf(&comp._eventComponentDestroySelf, comp.composite)
	}
}

// String 字符串化
func (comp *ComponentBehavior) String() string {
	var entityID ID
	if entity := comp.GetEntity(); entity != nil {
		entityID = entity.GetID()
	}

	return fmt.Sprintf("[ID:%s SerialNo:%d Name:%s Entity:%s State:%s]",
		comp.GetID(),
		comp.GetSerialNo(),
		comp.GetName(),
		entityID,
		comp.GetState())
}

func (comp *ComponentBehavior) init(name string, entity Entity, composite Component, hookAllocator container.Allocator[localevent.Hook], gcCollector container.GCCollector) {
	comp.name = name
	comp.entity = entity
	comp.composite = composite
	comp._eventComponentDestroySelf.Init(false, nil, localevent.EventRecursion_NotEmit, hookAllocator, gcCollector)
}

func (comp *ComponentBehavior) setID(id ID) {
	comp.id = id
}

func (comp *ComponentBehavior) setSerialNo(sn int64) {
	comp.serialNo = sn
}

func (comp *ComponentBehavior) getContext() util.IfaceCache {
	return comp.entity.getContext()
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
	localevent.UnsafeEvent(&comp._eventComponentDestroySelf).SetGCCollector(gcCollector)
}

func (comp *ComponentBehavior) eventComponentDestroySelf() localevent.IEvent {
	return &comp._eventComponentDestroySelf
}
