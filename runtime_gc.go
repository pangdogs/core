package golaxy

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/container"
)

func (_runtime *RuntimeBehavior) gc() {
	runtime.UnsafeContext(_runtime.ctx).GetInnerGC().GC()
	localevent.UnsafeEvent(&_runtime.eventUpdate).GC()
	localevent.UnsafeEvent(&_runtime.eventLateUpdate).GC()
}

// CollectGC 收集GC
func (_runtime *RuntimeBehavior) CollectGC(gc container.GC) {
}
