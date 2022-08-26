package core

import "github.com/pangdogs/galaxy/core/container"

// IEvent 事件接口
type IEvent interface {
	// Emit 发送事件，一般情况下是在事件生成的代码中使用
	Emit(fun func(delegate IFaceCache) bool)

	newHook(delegate interface{}, delegateFastIFace IFaceCache, priority int32) Hook

	removeDelegate(delegate interface{})
}

// IEventTab 事件表接口，我们可以把一些事件定义在同一个源码文件中，开启事件代码生成器的生成事件表选项，这样可以自动生成事件列表，无需手动声明事件
type IEventTab interface {
	// Init 初始化事件表
	Init(autoRecover bool, reportError chan error, hookCache *container.Cache[Hook], gcCollector container.GCCollector)

	// Get 获取事件
	Get(id int) IEvent

	// Open 打开事件表中所有事件
	Open()

	// Close 关闭事件表中所有事件
	Close()

	// Clean 事件表中的所有事件清除全部订阅者
	Clean()
}

// EventRecursion 发生事件递归的处理方式，事件递归是指在一个事件的订阅者中再次发送这个事件
type EventRecursion int32

const (
	EventRecursion_Allow    EventRecursion = iota // 允许事件递归，但是可能会造成无限递归
	EventRecursion_Disallow                       // 不允许事件递归，发生时会panic
	EventRecursion_NotEmit                        // 不再发送事件，在订阅者中再次发送这个事件，不会再发送给任何订阅者
	EventRecursion_Discard                        // 丢弃递归的事件，在订阅者中再次发送这个事件，不会再次进入这个订阅者，但是会进入其他订阅者
	EventRecursion_Deepest                        // 深度优先处理递归事件，在订阅者中再次发送这个事件，会中断上次事件发送过程，并在本次事件发送过程中，不会再次进入这个订阅者
)

// Event 事件
type Event struct {
	subscribers    container.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	gcCollector    container.GCCollector
	emitted        int
	depth          int
	opened         bool
	inited         bool
}

// Init ...
func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion, hookCache *container.Cache[Hook], gcCollector container.GCCollector) {
	if gcCollector == nil {
		panic("nil gcCollector")
	}

	if event.inited {
		panic("repeated init event invalid")
	}

	event.autoRecover = autoRecover
	event.reportError = reportError
	event.eventRecursion = eventRecursion
	event.subscribers.Init(hookCache, gcCollector)
	event.gcCollector = gcCollector
	event.opened = true
	event.inited = true
}

// Open ...
func (event *Event) Open() {
	event.subscribers.SetGCCollector(event.gcCollector)
	event.opened = true
}

// Close ...
func (event *Event) Close() {
	event.subscribers.SetGCCollector(nil)
	event.Clean()
	event.opened = false
}

// Clean ...
func (event *Event) Clean() {
	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		e.Value.Unbind()
		return true
	})
}

// GC ...
func (event *Event) GC() {
	event.subscribers.GC()
}

// NeedGC ...
func (event *Event) NeedGC() bool {
	return event.subscribers.NeedGC()
}

// Emit ...
func (event *Event) Emit(fun func(delegate IFaceCache) bool) {
	if fun == nil {
		return
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
