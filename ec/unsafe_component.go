package ec

import (
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util/container"
	"reflect"
)

func UnsafeComponent(comp Component) _UnsafeComponent {
	return _UnsafeComponent{
		Component: comp,
	}
}

type _UnsafeComponent struct {
	Component
}

func (uc _UnsafeComponent) Init(name string, entity Entity, inheritor Component, hookCache *container.Cache[localevent.Hook]) {
	uc.init(name, entity, inheritor, hookCache)
}

func (uc _UnsafeComponent) SetID(id int64) {
	uc.setID(id)
}

func (uc _UnsafeComponent) SetPrimary(v bool) {
	uc.setPrimary(v)
}

func (uc _UnsafeComponent) GetPrimary() bool {
	return uc.getPrimary()
}

func (uc _UnsafeComponent) SetAwoke(v bool) {
	uc.setAwoke(v)
}

func (uc _UnsafeComponent) GetAwoke() bool {
	return uc.getAwoke()
}

func (uc _UnsafeComponent) SetStarted(v bool) {
	uc.setStarted(v)
}

func (uc _UnsafeComponent) GetStarted() bool {
	return uc.getStarted()
}

func (uc _UnsafeComponent) SetReflectValue(v reflect.Value) {
	uc.setReflectValue(v)
}

func (uc _UnsafeComponent) GetReflectValue() reflect.Value {
	return uc.getReflectValue()
}

func (uc _UnsafeComponent) EventComponentDestroySelf() localevent.IEvent {
	return uc.eventComponentDestroySelf()
}

func (uc _UnsafeComponent) GetInnerGC() container.GC {
	return uc.getInnerGC()
}

func (uc _UnsafeComponent) GetInnerGCCollector() container.GCCollector {
	return uc.getInnerGCCollector()
}
