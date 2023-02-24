package golaxy

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/runtime"
)

func (_runtime *RuntimeBehavior) gc() {
	runtime.UnsafeContext(_runtime.ctx).GC()
	localevent.UnsafeEvent(&_runtime.eventUpdate).GC()
	localevent.UnsafeEvent(&_runtime.eventLateUpdate).GC()
}
