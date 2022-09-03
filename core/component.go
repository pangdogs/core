package core

import (
	"github.com/pangdogs/galaxy/core/container"
	"reflect"
)

// Component 组件接口
type Component interface {
	_InnerGCCollector
	_InnerGC

	init(name string, entity Entity, inheritor Component, hookCache *container.Cache[Hook])

	setID(id int64)

	// GetID 获取组件（Component）全局唯一ID，线程安全
	GetID() int64

	// GetName 获取组件（Component）名称，线程安全
	GetName() string

	// GetEntity 获取组件（Component）依附的实体（Entity），非线程安全
	GetEntity() Entity

	// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
	GetRuntimeCtx() RuntimeContext

	// GetServiceCtx 获取服务上下文（Service Context），线程安全
	GetServiceCtx() ServiceContext

	setPrimary(v bool)

	getPrimary() bool

	getReflectValue() reflect.Value

	// DestroySelf 销毁自身，注意在生命周期[Awake,Start,Shut]中调用无效，非线程安全
	DestroySelf()

	eventComponentDestroySelf() IEvent
}

// ComponentSetPersistID 组件（Component）设置持久化ID，需要在组件加入实体（Entity）前设置，通常用于组件持久化
func ComponentSetPersistID(comp Component, persistID int64) {
	if persistID <= 0 {
		panic("persistID not invalid")
	}

	if comp.GetEntity() != nil {
		panic("component already added in entity")
	}

	comp.setID(persistID)
}

// ComponentBehavior 组件行为，开发组件时需要将此结构体匿名嵌入至组件结构体中
type ComponentBehavior struct {
	id                         int64
	name                       string
	entity                     Entity
	inheritor                  Component
	reflectValue               reflect.Value
	_eventComponentDestroySelf Event
	primary                    bool
	gc                         _ComponentBehaviorGC
}

func (comp *ComponentBehavior) init(name string, entity Entity, inheritor Component, hookCache *container.Cache[Hook]) {
	comp.gc.ComponentBehavior = comp
	comp.name = name
	comp.entity = entity
	comp.inheritor = inheritor
	comp._eventComponentDestroySelf.Init(false, nil, EventRecursion_Discard, hookCache, comp.getGCCollector())
}

func (comp *ComponentBehavior) setID(id int64) {
	comp.id = id
}

// GetID 获取组件（Component）全局唯一ID，线程安全
func (comp *ComponentBehavior) GetID() int64 {
	return comp.id
}

// GetName 获取组件（Component）名称，线程安全
func (comp *ComponentBehavior) GetName() string {
	return comp.name
}

// GetEntity 获取组件（Component）依附的实体（Entity），非线程安全
func (comp *ComponentBehavior) GetEntity() Entity {
	return comp.entity
}

func (comp *ComponentBehavior) setPrimary(v bool) {
	comp.primary = v
}

func (comp *ComponentBehavior) getPrimary() bool {
	return comp.primary
}

func (comp *ComponentBehavior) getReflectValue() reflect.Value {
	if comp.reflectValue.IsValid() {
		return comp.reflectValue
	}

	comp.reflectValue = reflect.ValueOf(comp.inheritor)

	return comp.reflectValue
}

// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
func (comp *ComponentBehavior) GetRuntimeCtx() RuntimeContext {
	return comp.entity.GetRuntimeCtx()
}

// GetServiceCtx 获取服务上下文（Service Context），线程安全
func (comp *ComponentBehavior) GetServiceCtx() ServiceContext {
	return comp.entity.GetServiceCtx()
}

// DestroySelf 销毁自身，注意在生命周期[Awake,Start,Shut]中调用无效，非线程安全
func (comp *ComponentBehavior) DestroySelf() {
	emitEventComponentDestroySelf(&comp._eventComponentDestroySelf, comp.inheritor)
}

func (comp *ComponentBehavior) eventComponentDestroySelf() IEvent {
	return &comp._eventComponentDestroySelf
}
