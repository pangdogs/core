package ec

import (
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

func UnsafeEntity(entity Entity) _UnsafeEntity {
	return _UnsafeEntity{
		Entity: entity,
	}
}

type _UnsafeEntity struct {
	Entity
}

func (ue _UnsafeEntity) Init(opts *EntityOptions) {
	ue.init(opts)
}

func (ue _UnsafeEntity) GetOptions() *EntityOptions {
	return ue.getOptions()
}

func (ue _UnsafeEntity) SetID(id int64) {
	ue.setID(id)
}

func (ue _UnsafeEntity) SetContext(ctx util.IfaceCache) {
	ue.setContext(ctx)
}

func (ue _UnsafeEntity) GetContext() util.IfaceCache {
	return ue.getContext()
}

func (ue _UnsafeEntity) SetGCCollector(gcCollect container.GCCollector) {
	ue.setGCCollector(gcCollect)
}

func (ue _UnsafeEntity) GetGCCollector() container.GCCollector {
	return ue.getGCCollector()
}

func (ue _UnsafeEntity) SetParent(parent Entity) {
	ue.setParent(parent)
}

func (ue _UnsafeEntity) SetInitialing(v bool) {
	ue.setInitialing(v)
}

func (ue _UnsafeEntity) GetInitialing() bool {
	return ue.getInitialing()
}

func (ue _UnsafeEntity) SetShutting(v bool) {
	ue.setShutting(v)
}

func (ue _UnsafeEntity) GetShutting() bool {
	return ue.getShutting()
}

func (ue _UnsafeEntity) EventEntityDestroySelf() localevent.IEvent {
	return ue.eventEntityDestroySelf()
}

func (ue _UnsafeEntity) GetInnerGC() container.GC {
	return ue.getInnerGC()
}

func (ue _UnsafeEntity) GetInnerGCCollector() container.GCCollector {
	return ue.getInnerGCCollector()
}
