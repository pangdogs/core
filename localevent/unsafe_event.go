package localevent

import "github.com/galaxy-kit/galaxy/util"

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

func (ue _UnsafeEvent) GC() {
	ue.gc()
}
