package ec

import (
	"github.com/galaxy-kit/galaxy-go/localevent"
	"github.com/galaxy-kit/galaxy-go/util"
	"github.com/galaxy-kit/galaxy-go/util/container"
	"reflect"
)

// Component 组件接口
type Component interface {
	_InnerGC
	_InnerGCCollector
	ContextHolder
	init(name string, entity Entity, inheritor Component, hookCache *container.Cache[localevent.Hook])
	setID(id int64)
	// GetID 获取组件全局唯一ID
	GetID() int64
	// GetName 获取组件名称
	GetName() string
	// GetEntity 获取组件依附的实体
	GetEntity() Entity
	setPrimary(v bool)
	getPrimary() bool
	setAwoke(v bool)
	getAwoke() bool
	setStarted(v bool)
	getStarted() bool
	setReflectValue(v reflect.Value)
	getReflectValue() reflect.Value
	// DestroySelf 销毁自身，注意在生命周期[Awake,Start,Shut]中调用无效
	DestroySelf()
	eventComponentDestroySelf() localevent.IEvent
}

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         int64
	name                       string
	entity                     Entity
	inheritor                  Component
	primary                    bool
	awoke, started             bool
	reflectValue               reflect.Value
	_eventComponentDestroySelf localevent.Event
	innerGC                    _ComponentInnerGC
}

func (comp *ComponentBehavior) init(name string, entity Entity, inheritor Component, hookCache *container.Cache[localevent.Hook]) {
	comp.innerGC.Init(comp)
	comp.name = name
	comp.entity = entity
	comp.inheritor = inheritor
	comp._eventComponentDestroySelf.Init(false, nil, localevent.EventRecursion_NotEmit, hookCache, &comp.innerGC)
}

func (comp *ComponentBehavior) setID(id int64) {
	comp.id = id
}

// GetID 获取组件全局唯一ID
func (comp *ComponentBehavior) GetID() int64 {
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

func (comp *ComponentBehavior) getContext() util.IfaceCache {
	return comp.entity.getContext()
}

func (comp *ComponentBehavior) setPrimary(v bool) {
	comp.primary = v
}

func (comp *ComponentBehavior) getPrimary() bool {
	return comp.primary
}

func (comp *ComponentBehavior) setAwoke(v bool) {
	comp.awoke = v
}

func (comp *ComponentBehavior) getAwoke() bool {
	return comp.awoke
}

func (comp *ComponentBehavior) setStarted(v bool) {
	comp.started = v
}

func (comp *ComponentBehavior) getStarted() bool {
	return comp.started
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

// DestroySelf 销毁自身，注意在生命周期[Awake,Start,Shut]中调用无效
func (comp *ComponentBehavior) DestroySelf() {
	emitEventComponentDestroySelf(&comp._eventComponentDestroySelf, comp.inheritor)
	comp._eventComponentDestroySelf.Close()
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
