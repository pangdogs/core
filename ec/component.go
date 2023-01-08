package ec

import (
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/golaxy-kit/golaxy/util/container"
	"reflect"
)

// Component 组件接口
type Component interface {
	_InnerGC
	_InnerGCCollector
	ContextHolder

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

	init(name string, entity Entity, inheritor Component, hookCache *container.Cache[localevent.Hook])
	setID(id ID)
	setSerialNo(sn int64)
	setFixed(v bool)
	getFixed() bool
	setState(state ComponentState)
	setReflectValue(v reflect.Value)
	getReflectValue() reflect.Value
	eventComponentDestroySelf() localevent.IEvent
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         ID
	serialNo                   int64
	name                       string
	entity                     Entity
	inheritor                  Component
	primary                    bool
	state                      ComponentState
	reflectValue               reflect.Value
	_eventComponentDestroySelf localevent.Event
	innerGC                    _ComponentInnerGC
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
		emitEventComponentDestroySelf(&comp._eventComponentDestroySelf, comp.inheritor)
	}
}

func (comp *ComponentBehavior) init(name string, entity Entity, inheritor Component, hookCache *container.Cache[localevent.Hook]) {
	comp.innerGC.Init(comp)
	comp.name = name
	comp.entity = entity
	comp.inheritor = inheritor
	comp._eventComponentDestroySelf.Init(false, nil, localevent.EventRecursion_NotEmit, hookCache, &comp.innerGC)
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
	comp.primary = v
}

func (comp *ComponentBehavior) getFixed() bool {
	return comp.primary
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

	comp.reflectValue = reflect.ValueOf(comp.inheritor)

	return comp.reflectValue
}

func (comp *ComponentBehavior) eventComponentDestroySelf() localevent.IEvent {
	return &comp._eventComponentDestroySelf
}

func (comp *ComponentBehavior) getInnerGC() container.GC {
	return &comp.innerGC
}

func (comp *ComponentBehavior) getInnerGCCollector() container.GCCollector {
	return &comp.innerGC
}
