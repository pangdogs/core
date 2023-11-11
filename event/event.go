package event

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
)

// EventRecursion 发生事件递归时的处理方式，事件递归是指事件发送过程中，在订阅者接收并处理事件的逻辑中，再次发送这个事件
type EventRecursion int32

const (
	EventRecursion_Allow    EventRecursion = iota // 允许事件递归，但是逻辑有误时，会造成无限递归
	EventRecursion_Disallow                       // 不允许事件递归，发生无限递归时会panic
	EventRecursion_NotEmit                        // 不再发送事件，如果在订阅者逻辑中再次发送这个事件，那么不会再发送给任何订阅者
	EventRecursion_Discard                        // 丢弃递归的事件，如果在订阅者逻辑中再次发送这个事件，那么不会再次发送给这个订阅者，但是会发送给其他订阅者
	EventRecursion_Deepest                        // 深度优先处理递归事件，如果在订阅者逻辑中再次发送这个事件，那么会中断上一次事件发送过程，并在本次事件发送过程中，不会再次发送给这个订阅者
)

var (
	// EventRecursionLimit 事件递归次数上限，超过此上限会panic
	EventRecursionLimit = int32(128)
)

// IEvent 本地事件接口，非线程安全，不能用于跨线程事件通知
type IEvent interface {
	emit(fun generic.Func1[iface.Cache, bool])
	newHook(delegateFace iface.FaceAny, priority int32) Hook
	removeDelegate(delegate any)
	setGCCollector(gcCollector container.GCCollector)
	gc()
}

// Event 本地事件，非线程安全，不能用于跨线程事件通知
type Event struct {
	subscribers    container.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	emitted        int32
	emitDepth      int32
	emitBatch      int32
	opened         bool
	inited         bool
}

// Init 初始化事件
func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion, hookAllocator container.Allocator[Hook], gcCollector container.GCCollector) {
	if event.inited {
		panic(fmt.Errorf("%w: event is already initialized", ErrEvent))
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
		panic(fmt.Errorf("%w: event not initialized", ErrEvent))
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

func (event *Event) emit(fun generic.Func1[iface.Cache, bool]) {
	if !event.opened {
		return
	}

	if event.emitted >= EventRecursionLimit {
		panic(fmt.Errorf("%w: recursive event calls(%d) cause stack overflow", ErrEvent, event.emitted))
	}

	switch event.eventRecursion {
	case EventRecursion_NotEmit:
		if event.emitted > 0 {
			return
		}
	}

	event.emitted++
	defer func() { event.emitted-- }()
	event.emitDepth = event.emitted
	event.emitBatch++

	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		if !event.opened {
			return false
		}

		if e.Value.delegateFace.IsNil() || e.Value.createdBatch == event.emitBatch {
			return true
		}

		switch event.eventRecursion {
		case EventRecursion_Allow:
			break
		case EventRecursion_Disallow:
			if e.Value.received > 0 {
				panic(fmt.Errorf("%w: recursive event disallowed", ErrEvent))
			}
		case EventRecursion_Discard:
			if e.Value.received > 0 {
				return true
			}
		case EventRecursion_Deepest:
			if event.emitDepth != event.emitted {
				return false
			}
			if e.Value.received > 0 {
				return true
			}
		}

		e.Value.received++
		defer func() { e.Value.received-- }()

		ret, panicErr := fun.Call(event.autoRecover, event.reportError, e.Value.delegateFace.Cache)
		if panicErr != nil {
			return true
		}

		return ret
	})
}

func (event *Event) newHook(delegateFace iface.FaceAny, priority int32) Hook {
	if !event.opened {
		panic(fmt.Errorf("%w: event closed", ErrEvent))
	}

	if delegateFace.IsNil() {
		panic(fmt.Errorf("%w: %w: delegateFace is nil", ErrEvent, exception.ErrArgs))
	}

	hook := Hook{
		delegateFace: delegateFace,
		createdBatch: event.emitBatch,
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

func (event *Event) setGCCollector(gcCollector container.GCCollector) {
	event.subscribers.SetGCCollector(gcCollector)
}

func (event *Event) gc() {
	event.subscribers.GC()
}
