package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util"
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

func (uc _UnsafeComponent) Init(name string, entity Entity, inheritor Component, hookAllocator container.Allocator[localevent.Hook], gcCollector container.GCCollector) {
	uc.init(name, entity, inheritor, hookAllocator, gcCollector)
}

func (uc _UnsafeComponent) SetID(id ID) {
	uc.setID(id)
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

func (uc _UnsafeComponent) GetContext() util.IfaceCache {
	return uc.getContext()
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

func (uc _UnsafeComponent) GetInheritor() Component {
	return uc.getInheritor()
}

func (uc _UnsafeComponent) SetGCCollector(gcCollector container.GCCollector) {
	uc.setGCCollector(gcCollector)
}

func (uc _UnsafeComponent) EventComponentDestroySelf() localevent.IEvent {
	return uc.eventComponentDestroySelf()
}
