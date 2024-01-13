package event

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/iface"
)

// BindEvent 绑定事件与订阅者，可以设置优先级调整回调先后顺序（升序）
func BindEvent[T any](event IEvent, subscriber T, priority ...int32) Hook {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrEvent, exception.ErrArgs))
	}

	_priority := int32(0)
	if len(priority) > 0 {
		_priority = priority[0]
	}

	return event.newHook(iface.MakeFaceAny(subscriber), _priority)
}

// UnbindEvent 解绑定事件与订阅者，比使用事件绑定句柄解绑定性能差，且在同个订阅者多次绑定事件的情况下，只能从最后依次解除，无法指定解除哪一个
func UnbindEvent(event IEvent, subscriber any) {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrEvent, exception.ErrArgs))
	}
	event.removeSubscriber(subscriber)
}
