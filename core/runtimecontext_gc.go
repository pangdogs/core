package core

import "github.com/pangdogs/galaxy/core/container"

type _RuntimeContextBehaviorGC struct {
	*_RuntimeContextBehavior
}

// GC ...
func (rgc *_RuntimeContextBehaviorGC) GC() {
	for i := range rgc._RuntimeContextBehavior.gcList {
		rgc._RuntimeContextBehavior.gcList[i].GC()
	}
	rgc._RuntimeContextBehavior.gcList = rgc._RuntimeContextBehavior.gcList[:0]
}

// NeedGC ...
func (rgc *_RuntimeContextBehaviorGC) NeedGC() bool {
	return len(rgc._RuntimeContextBehavior.gcList) > 0
}

// CollectGC ...
func (runtimeCtx *_RuntimeContextBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	runtimeCtx.gcList = append(runtimeCtx.gcList, gc)
}

func (runtimeCtx *_RuntimeContextBehavior) getGC() container.GC {
	return &runtimeCtx.gc
}
