package core

import (
	"git.golaxy.org/core/runtime"
)

func (rt *RuntimeBehavior) gc() {
	runtime.UnsafeContext(rt.ctx).GC()
	rt.opts.CustomGC.Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError(), nil, rt.opts.CompositeFace.Iface)
}
