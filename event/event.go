package event

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
)

// EventRecursion 发生事件递归时的处理方式（事件递归：事件发送过程中，在订阅者的逻辑中，再次发送这个事件）
type EventRecursion int32

const (
	EventRecursion_Allow    EventRecursion = iota // 允许事件递归，可能会无限递归
	EventRecursion_Disallow                       // 不允许事件递归，递归时会panic
	EventRecursion_Discard                        // 丢弃递归的事件，不会再发送给任何订阅者
	EventRecursion_Truncate                       // 截断递归的事件，不会再发送给当前订阅者，但是会发送给其他订阅者
	EventRecursion_Deepest                        // 深度优先处理递归的事件，会中断当前事件发送过程，并在新的事件发送过程中，不会再次发送给这个订阅者
)

var (
	// EventRecursionLimit 事件递归次数上限，超过此上限会panic
	EventRecursionLimit = int32(128)
)

// IEvent 本地事件接口
/*
定义事件：
	1.按以下格式编写一个接口，即完成事件的定义：
	type Event{事件名} interface {
		On{事件名}({参数列表})
	}

	2.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventcode gen_event

定义事件的选项（添加到定义事件的注释里）：
	1.发送事件的辅助代码的可见性
		[EmitExport]：不可见
		[EmitUnExport]：可见

	2.是否生成简化绑定事件的辅助代码
		[EmitAuto]：生成
		[EmitManual]：不生成

使用事件：
	1.事件一般作为组件的成员，在组件Awake时初始化，组件Dispose时关闭，示例如下：
	type Comp struct {
		ec.ComponentBehavior
		event{事件名} event.Event
	}
	func (c *Comp) Awake() {
		runtime.Current(c).ActivateEvent(&c.event{事件名}, event.EventRecursion_Discard)
	}
	func (c *Comp) Dispose() {
		c.event{事件名}.Close()
	}

订阅事件：
	1.在组件的成员函数，编写以下代码：
	func (c *Comp) On{事件名}({参数列表}) {
		...
	}

	2.在需要订阅事件时，编写以下代码：
	func (c *Comp) MethodXXX() {
		{事件定义包名}.Bind{事件名}({发布者}, c)
	}

	3.如果订阅者生命周期小于发布者，那么需要记录hook并且在Dispose时解除绑定，示例如下：
	type Comp struct {
		ec.ComponentBehavior
		hook event.Hook
	}
	func (c *Comp) MethodXXX() {
		c.hook = {事件定义包名}.Bind{事件名}({发布者}, c)
	}
	func (c *Comp) Dispose() {
		c.hook.Unbind()
	}

	4.如果不想写代码记录hook，可以使用ComponentBehavior的AutoHooks()来记录hook，在组件销毁时会自动解除绑定
*/
type IEvent interface {
	emit(fun generic.Func1[iface.Cache, bool])
	newHook(subscriberFace iface.FaceAny, priority int32) Hook
	removeSubscriber(subscriber any)
}

// IEventCtrl 事件控制接口
type IEventCtrl interface {
	// Init 初始化事件
	Init(autoRecover bool, reportError chan error, recursion EventRecursion)
	// Open 打开事件
	Open()
	// Close 关闭事件
	Close()
	// Clean 清除全部订阅者
	Clean()
}

// Event 本地事件
type Event struct {
	subscribers    container.List[Hook]
	autoRecover    bool
	reportError    chan error
	eventRecursion EventRecursion
	emitted        int32
	emitDepth      int32
	opened         bool
	inited         bool
}

// Init 初始化事件
func (event *Event) Init(autoRecover bool, reportError chan error, eventRecursion EventRecursion) {
	if event.inited {
		panic(fmt.Errorf("%w: event is already initialized", ErrEvent))
	}

	event.autoRecover = autoRecover
	event.reportError = reportError
	event.eventRecursion = eventRecursion
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
	case EventRecursion_Discard:
		if event.emitted > 0 {
			return
		}
	}

	event.emitted++
	defer func() { event.emitted-- }()
	event.emitDepth = event.emitted
	ver := event.subscribers.Version()

	event.subscribers.Traversal(func(e *container.Element[Hook]) bool {
		if !event.opened {
			return false
		}

		if e.Value.subscriberFace.IsNil() || e.Version() > ver {
			return true
		}

		switch event.eventRecursion {
		case EventRecursion_Allow:
			break
		case EventRecursion_Disallow:
			if e.Value.received > 0 {
				panic(fmt.Errorf("%w: recursive event disallowed", ErrEvent))
			}
		case EventRecursion_Truncate:
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

		ret, panicErr := fun.Call(event.autoRecover, event.reportError, e.Value.subscriberFace.Cache)
		if panicErr != nil {
			return true
		}

		return ret
	})
}

func (event *Event) newHook(subscriberFace iface.FaceAny, priority int32) Hook {
	if !event.opened {
		panic(fmt.Errorf("%w: event closed", ErrEvent))
	}

	if subscriberFace.IsNil() {
		panic(fmt.Errorf("%w: %w: subscriberFace is nil", ErrEvent, exception.ErrArgs))
	}

	hook := Hook{
		subscriberFace: subscriberFace,
		priority:       priority,
	}

	var at *container.Element[Hook]

	event.subscribers.ReversedTraversal(func(other *container.Element[Hook]) bool {
		if hook.priority >= other.Value.priority {
			at = other
			return false
		}
		return true
	})

	if at != nil {
		hook.element = event.subscribers.InsertAfter(Hook{}, at)
	} else {
		hook.element = event.subscribers.PushFront(Hook{})
	}

	hook.element.Value = hook

	return hook
}

func (event *Event) removeSubscriber(subscriber any) {
	event.subscribers.ReversedTraversal(func(other *container.Element[Hook]) bool {
		if other.Value.subscriberFace.Iface == subscriber {
			other.Escape()
			return false
		}
		return true
	})
}
