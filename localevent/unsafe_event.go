package localevent

import "github.com/pangdogs/galaxy/util"

func UnsafeEvent(v *Event) _UnsafeEvent {
	return _UnsafeEvent{
		Event: v,
	}
}

type _UnsafeEvent struct {
	*Event
}

func (ue _UnsafeEvent) NewHook(delegateFace util.FaceAny, priority int32) Hook {
	return ue.newHook(delegateFace, priority)
}

func (ue _UnsafeEvent) RemoveDelegate(delegate interface{}) {
	ue.removeDelegate(delegate)
}

func (ue _UnsafeEvent) GC() {
	ue.gc()
}
