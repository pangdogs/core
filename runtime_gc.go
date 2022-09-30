package galaxy

import (
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/util/container"
)

func (_runtime *RuntimeBehavior) gc() {
	runtime.UnsafeContext(_runtime.ctx).GetInnerGC().GC()
	localevent.UnsafeEvent(&_runtime.eventUpdate).GC()
	localevent.UnsafeEvent(&_runtime.eventLateUpdate).GC()
}

// CollectGC 收集GC
func (_runtime *RuntimeBehavior) CollectGC(gc container.GC) {
}
