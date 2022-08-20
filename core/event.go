package core

import "github.com/pangdogs/galaxy/core/container"

type IEvent interface {
	Emit(fun func(delegate IFaceCache) bool)
	newHook(delegate interface{}, delegateFastIFace IFaceCache, priority int32) Hook
	removeDelegate(delegate interface{})
}

type IEventTab interface {
	EventTab(id int) IEvent
}

type IEventAssist interface {
	IEventTab
	Init(autoRecover bool, reportError chan error, hookCache *container.Cache[Hook], gcCollector container.GCCollector)
	Open()
	Close()
	Clear()
}

type EventRecursion int32

const (
	EventRecursion_Allow EventRecursion = iota
	EventRecursion_Disallow
	EventRecursion_Discard
	EventRecursion_Deepest
)

type Event struct {
	subscribers    container.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	gcCollector    container.GCCollector
	emitted        int
	depth          int
	opened         bool
}

func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion, hookCache *container.Cache[Hook], gcCollector container.GCCollector) {
	if gcCollector == nil {
		panic("nil gcCollector")
	}

	event.autoRecover = autoRecover
	event.reportError = reportError
	event.eventRecursion = eventRecursion
	event.subscribers.Init(hookCache, gcCollector)
	event.gcCollector = gcCollector
	event.opened = true
}

func (event *Event) GC() {
	event.subscribers.GC()
}

func (event *Event) NeedGC() bool {
	return event.subscribers.NeedGC()
}

func (event *Event) Emit(fun func(delegate IFaceCache) bool) {
	if fun == nil {
		return
	}

	event.emitted++
	defer func() {
		event.emitted--
	}()

	event.depth = event.emitted

	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		if !event.opened {
			return false
		}

		if e.Value.delegateFastIFace == NilIFaceCache {
			return true
		}

		switch event.eventRecursion {
		case EventRecursion_Allow:
			break
		case EventRecursion_Disallow:
			if e.Value.received > 0 {
				panic("recursive event disallowed")
			}
		case EventRecursion_Discard:
			if e.Value.received > 0 {
				return true
			}
		case EventRecursion_Deepest:
			if event.depth != event.emitted {
				return false
			}
			if e.Value.received > 0 {
				return true
			}
		}

		e.Value.received++
		defer func() {
			e.Value.received--
		}()

		ret, err := CallOuter(event.autoRecover, event.reportError, func() bool {
			return fun(e.Value.delegateFastIFace)
		})

		if err != nil {
			return true
		}

		return ret
	})
}

func (event *Event) newHook(delegate interface{}, delegateFastIFace IFaceCache, priority int32) Hook {
	if !event.opened {
		panic("event closed")
	}

	hook := Hook{
		delegate:          delegate,
		delegateFastIFace: delegateFastIFace,
		priority:          priority,
	}

	var mark *container.Element[Hook]

	event.subscribers.ReverseTraversal(func(other *container.Element[Hook]) bool {
		if hook.priority >= other.Value.priority {
			mark = other
			return false
		}
		return true
	})

	if mark != nil {
		hook.element = event.subscribers.InsertAfter(Hook{}, mark)
	} else {
		hook.element = event.subscribers.PushFront(Hook{})
	}

	hook.element.Value = hook

	return hook
}

func (event *Event) removeDelegate(delegate interface{}) {
	event.subscribers.ReverseTraversal(func(other *container.Element[Hook]) bool {
		if other.Value.delegate == delegate {
			other.Escape()
			return false
		}
		return true
	})
}

func (event *Event) Open() {
	event.subscribers.SetGCCollector(event.gcCollector)
	event.opened = true
}

func (event *Event) Close() {
	event.subscribers.SetGCCollector(nil)
	event.Clear()
	event.opened = false
}

func (event *Event) Clear() {
	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		e.Value.Unbind()
		return true
	})
}
