package event

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// Bind 绑定事件与订阅者，可以设置优先级调整回调先后顺序（升序）
func Bind[T any](event IEvent, subscriber T, priority ...int32) Hook {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrEvent, exception.ErrArgs))
	}

	_priority := int32(0)
	if len(priority) > 0 {
		_priority = priority[0]
	}

	return event.newHook(iface.MakeFaceAny(subscriber), _priority)
}

// Unbind 解绑定事件与订阅者，在同个订阅者多次绑定事件的情况下，会以逆序依次解除，正常情况下应该使用事件钩子（Hook）解绑定，不应该使用该函数
func Unbind(event IEvent, subscriber any) {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrEvent, exception.ErrArgs))
	}
	event.removeSubscriber(subscriber)
}
