package core

import "github.com/pangdogs/galaxy/core/container"

func (runtime *_RuntimeBehavior) gc() {
	runtime.ctx.getGC().GC()
	runtime.eventUpdate.gc()
	runtime.eventLateUpdate.gc()
}

// CollectGC ...
func (runtime *_RuntimeBehavior) CollectGC(gc container.GC) {
}
