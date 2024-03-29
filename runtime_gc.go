package core

import (
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/runtime"
)

func (rt *RuntimeBehavior) gc() {
	runtime.UnsafeContext(rt.ctx).GC()
	event.UnsafeEvent(&rt.eventUpdate).GC()
	event.UnsafeEvent(&rt.eventLateUpdate).GC()
	rt.opts.CustomGC.Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), nil, rt.opts.CompositeFace.Iface)
}
