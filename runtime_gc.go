package galaxy

import (
	"github.com/galaxy-kit/galaxy-go/localevent"
	"github.com/galaxy-kit/galaxy-go/runtime"
	"github.com/galaxy-kit/galaxy-go/util/container"
)

func (_runtime *RuntimeBehavior) gc() {
	runtime.UnsafeContext(_runtime.ctx).GetInnerGC().GC()
	localevent.UnsafeEvent(&_runtime.eventUpdate).GC()
	localevent.UnsafeEvent(&_runtime.eventLateUpdate).GC()
}

// CollectGC 收集GC
func (_runtime *RuntimeBehavior) CollectGC(gc container.GC) {
}
