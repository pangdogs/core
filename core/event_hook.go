package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// BindEvent 绑定事件与订阅者，非线程安全
func BindEvent[T any](event IEvent, delegate T) Hook {
	return BindEventWithPriority(event, delegate, 0)
}

// BindEventWithPriority 绑定事件与订阅者，可以设置优先级调整回调先后顺序，按优先级升序排列，非线程安全
func BindEventWithPriority[T any](event IEvent, delegate T, priority int32) Hook {
	if event == nil {
		panic("nil event")
	}
	return event.newHook(FaceAny{
		IFace: delegate,
		Cache: IFace2Cache(delegate),
	}, priority)
}

// UnbindEvent 解绑定事件与订阅者，比使用事件绑定句柄（Hook）解绑定性能差，且在同个订阅者多次绑定事件的情况下，只能从最后依次解除，无法指定解除哪一个，非线程安全
func UnbindEvent(event IEvent, delegate interface{}) {
	if event == nil {
		panic("nil event")
	}
	event.removeDelegate(delegate)
}

// Hook 事件绑定句柄，主要用于重新绑定或解除绑定事件，由BindEvent()或BindEventWithPriority()产生，请勿手工创建
type Hook struct {
	delegateFace FaceAny
	priority     int32
	element      *container.Element[Hook]
	received     int
}

// Bind 重新绑定事件与订阅者，非线程安全
func (hook *Hook) Bind(event IEvent) {
	hook.BindWithPriority(event, 0)
}

// BindWithPriority 重新绑定事件与订阅者，可以设置优先级调整回调先后顺序，按优先级升序排列，非线程安全
func (hook *Hook) BindWithPriority(event IEvent, priority int32) {
	if event == nil {
		panic("nil event")
	}

	if hook.IsBound() {
		panic("repeated bind event invalid")
	}

	*hook = event.newHook(hook.delegateFace, priority)
}

// Unbind 解绑定事件与订阅者
func (hook *Hook) Unbind() {
	if hook.element != nil {
		hook.element.Escape()
		hook.element = nil
	}
}

// IsBound 是否已绑定事件
func (hook *Hook) IsBound() bool {
	return hook.element != nil && !hook.element.Escaped()
}
