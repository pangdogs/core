package localevent

import "github.com/galaxy-kit/galaxy/util"

// BindEvent 绑定事件与订阅者
func BindEvent[T any](event IEvent, delegate T) Hook {
	return BindEventWithPriority(event, delegate, 0)
}

// BindEventWithPriority 绑定事件与订阅者，可以设置优先级调整回调先后顺序，按优先级升序排列
func BindEventWithPriority[T any](event IEvent, delegate T, priority int32) Hook {
	if event == nil {
		panic("nil event")
	}
	return event.newHook(util.NewFacePair[any](delegate, delegate), priority)
}

// UnbindEvent 解绑定事件与订阅者，比使用事件绑定句柄解绑定性能差，且在同个订阅者多次绑定事件的情况下，只能从最后依次解除，无法指定解除哪一个
func UnbindEvent(event IEvent, delegate any) {
	if event == nil {
		panic("nil event")
	}
	event.removeDelegate(delegate)
}
