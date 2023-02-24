package localevent

import (
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// EventRecursion 发生事件递归的处理方式，事件递归是指在一个事件的订阅者中再次发送这个事件
type EventRecursion int32

const (
	EventRecursion_Allow    EventRecursion = iota // 允许事件递归，但是可能会造成无限递归
	EventRecursion_Disallow                       // 不允许事件递归，发生时会panic
	EventRecursion_NotEmit                        // 不再发送事件，如果在订阅者中再次发送这个事件，那么不会再发送给任何订阅者
	EventRecursion_Discard                        // 丢弃递归的事件，如果在订阅者中再次发送这个事件，那么不会再次进入这个订阅者，但是会进入其他订阅者
	EventRecursion_Deepest                        // 深度优先处理递归事件，如果在订阅者中再次发送这个事件，那么会中断上次事件发送过程，并在本次事件发送过程中，不会再次进入这个订阅者
)

// EventRecursionLimit 事件递归次数上限
var EventRecursionLimit = 128

// IEvent 本地事件接口，非线程安全，不能用于跨线程事件通知
type IEvent interface {
	emit(fun func(delegate util.IfaceCache) bool)
	newHook(delegateFace util.FaceAny, priority int32) Hook
	removeDelegate(delegate any)
	setGCCollector(gcCollect container.GCCollector)
	gc()
}

// Event 本地事件，非线程安全，不能用于跨线程事件通知
type Event struct {
	subscribers    container.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	emitted        int
	depth          int
	opened         bool
	inited         bool
}

// Init 初始化事件
func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion, hookAllocator container.Allocator[Hook], gcCollector container.GCCollector) {
	if event.inited {
		panic("event initialized")
	}

	event.autoRecover = autoRecover
	event.reportError = reportError
	event.eventRecursion = eventRecursion
	event.subscribers.Init(hookAllocator, gcCollector)
	event.inited = true

	event.Open()
}

// Open 打开事件
func (event *Event) Open() {
	if !event.inited {
		panic("event not initialized")
	}

	event.opened = true
}

// Close 关闭事件
func (event *Event) Close() {
	event.Clean()
	event.opened = false
}

// Clean 清除全部订阅者
func (event *Event) Clean() {
	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		e.Value.Unbind()
		return true
	})
}

func (event *Event) emit(fun func(delegate util.IfaceCache) bool) {
	if fun == nil || !event.opened {
		return
	}

	if event.emitted >= EventRecursionLimit {
		panic("recursive event calls cause stack overflow")
	}

	switch event.eventRecursion {
	case EventRecursion_NotEmit:
		if event.emitted > 0 {
			return
		}
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

		if e.Value.delegateFace.IsNil() {
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

		ret, err := internal.CallOuter(event.autoRecover, event.reportError, func() bool {
			return fun(e.Value.delegateFace.Cache)
		})

		if err != nil {
			return true
		}

		return ret
	})
}

func (event *Event) newHook(delegateFace util.FaceAny, priority int32) Hook {
	if !event.opened {
		panic("event closed")
	}

	hook := Hook{
		delegateFace: delegateFace,
		priority:     priority,
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

func (event *Event) removeDelegate(delegate any) {
	event.subscribers.ReverseTraversal(func(other *container.Element[Hook]) bool {
		if other.Value.delegateFace.Iface == delegate {
			other.Escape()
			return false
		}
		return true
	})
}

func (event *Event) setGCCollector(gcCollect container.GCCollector) {
	event.subscribers.SetGCCollector(gcCollect)
}

func (event *Event) gc() {
	event.subscribers.GC()
}
