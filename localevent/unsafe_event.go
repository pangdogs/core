package localevent

import (
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

func UnsafeEvent(v IEvent) _UnsafeEvent {
	return _UnsafeEvent{
		IEvent: v,
	}
}

type _UnsafeEvent struct {
	IEvent
}

func (ue _UnsafeEvent) Emit(fun func(delegate util.IfaceCache) bool) {
	ue.emit(fun)
}

func (ue _UnsafeEvent) NewHook(delegateFace util.FaceAny, priority int32) Hook {
	return ue.newHook(delegateFace, priority)
}

func (ue _UnsafeEvent) RemoveDelegate(delegate any) {
	ue.removeDelegate(delegate)
}

func (ue _UnsafeEvent) SetGCCollector(gcCollect container.GCCollector) {
	ue.setGCCollector(gcCollect)
}

func (ue _UnsafeEvent) GC() {
	ue.gc()
}
