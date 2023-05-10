package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util/container"
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

func (uc _UnsafeComponent) Init(name string, entity Entity, composite Component, hookAllocator container.Allocator[localevent.Hook], gcCollector container.GCCollector) {
	uc.init(name, entity, composite, hookAllocator, gcCollector)
}

func (uc _UnsafeComponent) SetId(id uid.Id) {
	uc.setId(id)
}

func (uc _UnsafeComponent) SetSerialNo(sn int64) {
	uc.setSerialNo(sn)
}

func (uc _UnsafeComponent) SetFixed(v bool) {
	uc.setFixed(v)
}

func (uc _UnsafeComponent) GetFixed() bool {
	return uc.getFixed()
}

func (uc _UnsafeComponent) SetState(state ComponentState) {
	uc.setState(state)
}

func (uc _UnsafeComponent) SetReflectValue(v reflect.Value) {
	uc.setReflectValue(v)
}

func (uc _UnsafeComponent) GetReflectValue() reflect.Value {
	return uc.getReflectValue()
}

func (uc _UnsafeComponent) GetComposite() Component {
	return uc.getComposite()
}

func (uc _UnsafeComponent) SetGCCollector(gcCollector container.GCCollector) {
	uc.setGCCollector(gcCollector)
}

func (uc _UnsafeComponent) EventComponentDestroySelf() localevent.IEvent {
	return uc.eventComponentDestroySelf()
}
