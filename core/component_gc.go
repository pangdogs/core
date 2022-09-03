package core

import "github.com/pangdogs/galaxy/core/container"

type _ComponentBehaviorGC struct {
	*ComponentBehavior
	gcMark, gcCollected bool
}

// GC GC
func (cgc *_ComponentBehaviorGC) GC() {
	if !cgc.gcMark {
		return
	}
	cgc.gcMark = false
	cgc.gcCollected = false

	cgc._eventComponentDestroySelf.gc()
}

// NeedGC 是否需要GC
func (cgc *_ComponentBehaviorGC) NeedGC() bool {
	return cgc.gcMark
}

// CollectGC 收集GC
func (cgc *_ComponentBehaviorGC) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	cgc.gcMark = true

	if cgc.entity != nil && !cgc.gcCollected {
		cgc.gcCollected = true
		cgc.entity.getGCCollector().CollectGC(cgc.getGC())
	}
}

func (comp *ComponentBehavior) getGC() container.GC {
	return &comp.gc
}

func (comp *ComponentBehavior) getGCCollector() container.GCCollector {
	return &comp.gc
}
