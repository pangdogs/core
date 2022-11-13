package ec

import (
	"github.com/galaxy-kit/galaxy/localevent"
	"github.com/galaxy-kit/galaxy/util/container"
)

type _ComponentInnerGC struct {
	*ComponentBehavior
	gcMark, gcCollected bool
}

func (igc *_ComponentInnerGC) Init(cb *ComponentBehavior) {
	igc.ComponentBehavior = cb
}

func (igc *_ComponentInnerGC) GC() {
	if !igc.gcMark {
		return
	}
	igc.gcMark = false
	igc.gcCollected = false

	localevent.UnsafeEvent(&igc._eventComponentDestroySelf).GC()
}

func (igc *_ComponentInnerGC) NeedGC() bool {
	return igc.gcMark
}

func (igc *_ComponentInnerGC) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	igc.gcMark = true

	if igc.entity != nil && !igc.gcCollected {
		igc.gcCollected = true
		igc.entity.getInnerGCCollector().CollectGC(igc)
	}
}
