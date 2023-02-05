package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util/container"
)

type _EntityInnerGC struct {
	*EntityBehavior
	gcMark, gcCollected bool
}

func (igc *_EntityInnerGC) Init(eb *EntityBehavior) {
	igc.EntityBehavior = eb
}

func (igc *_EntityInnerGC) GC() {
	if !igc.gcMark {
		return
	}
	igc.gcMark = false
	igc.gcCollected = false

	igc.componentList.GC()
	localevent.UnsafeEvent(&igc._eventEntityDestroySelf).GC()
	localevent.UnsafeEvent(&igc.eventCompMgrAddComponents).GC()
	localevent.UnsafeEvent(&igc.eventCompMgrRemoveComponent).GC()
}

func (igc *_EntityInnerGC) NeedGC() bool {
	return igc.gcMark
}

func (igc *_EntityInnerGC) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	igc.gcMark = true

	if igc.gcCollector != nil && !igc.gcCollected {
		igc.gcCollected = true
		igc.gcCollector.CollectGC(igc)
	}
}
