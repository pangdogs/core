package golaxy

import (
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/runtime"
)

func (_runtime *RuntimeBehavior) gc() {
	runtime.UnsafeContext(_runtime.ctx).GC()
	event.UnsafeEvent(&_runtime.eventUpdate).GC()
	event.UnsafeEvent(&_runtime.eventLateUpdate).GC()

	if _runtime.opts.CustomGC != nil {
		_runtime.opts.CustomGC(_runtime.opts.CompositeFace.Iface)
	}
}
