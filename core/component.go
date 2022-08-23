package core

import (
	"github.com/pangdogs/galaxy/core/container"
	"reflect"
)

type Component interface {
	container.GC
	container.GCCollector
	init(name string, entity Entity, inheritor Component, hookCache *container.Cache[Hook])
	setID(id uint64)
	GetID() uint64
	GetName() string
	GetEntity() Entity
	GetRuntimeCtx() RuntimeContext
	GetServiceCtx() ServiceContext
	setPrimary(v bool)
	getPrimary() bool
	getReflectValue() reflect.Value
	DestroySelf()
	eventComponentDestroySelf() IEvent
}

type ComponentBehavior struct {
	id                         uint64
	name                       string
	entity                     Entity
	inheritor                  Component
	primary                    bool
	reflectValue               reflect.Value
	_eventComponentDestroySelf Event
	gcMark, gcCollected        bool
}

func (comp *ComponentBehavior) GC() {
	if !comp.gcMark {
		return
	}
	comp.gcMark = false
	comp.gcCollected = false

	comp._eventComponentDestroySelf.GC()
}

func (comp *ComponentBehavior) NeedGC() bool {
	return comp.gcMark
}

func (comp *ComponentBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	comp.gcMark = true

	if comp.entity != nil && !comp.gcCollected {
		comp.gcCollected = true
		comp.entity.CollectGC(comp.inheritor)
	}
}

func (comp *ComponentBehavior) init(name string, entity Entity, inheritor Component, hookCache *container.Cache[Hook]) {
	comp.name = name
	comp.entity = entity
	comp.inheritor = inheritor
	comp._eventComponentDestroySelf.Init(false, nil, EventRecursion_Discard, hookCache, comp.inheritor)
}

func (comp *ComponentBehavior) setID(id uint64) {
	comp.id = id
}

func (comp *ComponentBehavior) GetID() uint64 {
	return comp.id
}

func (comp *ComponentBehavior) GetName() string {
	return comp.name
}

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

func (comp *ComponentBehavior) GetRuntimeCtx() RuntimeContext {
	return comp.entity.GetRuntimeCtx()
}

func (comp *ComponentBehavior) GetServiceCtx() ServiceContext {
	return comp.entity.GetServiceCtx()
}

func (comp *ComponentBehavior) DestroySelf() {
	emitEventComponentDestroySelf(&comp._eventComponentDestroySelf, comp.inheritor)
}

func (comp *ComponentBehavior) eventComponentDestroySelf() IEvent {
	return &comp._eventComponentDestroySelf
}
