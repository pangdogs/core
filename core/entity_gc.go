package core

import "github.com/pangdogs/galaxy/core/container"

type _EntityBehaviorGC struct {
	*EntityBehavior
	gcMark, gcCollected bool
}

// GC GC
func (egc *_EntityBehaviorGC) GC() {
	if !egc.gcMark {
		return
	}
	egc.gcMark = false
	egc.gcCollected = false

	egc.componentList.GC()
	egc._eventEntityDestroySelf.gc()
	egc.eventCompMgrAddComponents.gc()
	egc.eventCompMgrRemoveComponent.gc()
}

// NeedGC 是否需要GC
func (egc *_EntityBehaviorGC) NeedGC() bool {
	return egc.gcMark
}

// CollectGC 收集GC
func (egc *_EntityBehaviorGC) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	egc.gcMark = true

	if egc.runtimeCtx != nil && !egc.gcCollected {
		egc.gcCollected = true
		egc.runtimeCtx.CollectGC(egc.getGC())
	}
}

func (entity *EntityBehavior) getGC() container.GC {
	return &entity.gc
}

func (entity *EntityBehavior) getGCCollector() container.GCCollector {
	return &entity.gc
}
