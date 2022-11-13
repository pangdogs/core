package localevent

import (
	"github.com/galaxy-kit/galaxy-go/util"
	"github.com/galaxy-kit/galaxy-go/util/container"
)

// Hook 事件绑定句柄，主要用于重新绑定或解除绑定事件，由BindEvent()或BindEventWithPriority()创建并返回，请勿自己创建
type Hook struct {
	delegateFace util.FaceAny
	priority     int32
	element      *container.Element[Hook]
	received     int
}

// Bind 重新绑定事件与订阅者
func (hook *Hook) Bind(event IEvent) {
	hook.BindWithPriority(event, 0)
}

// BindWithPriority 重新绑定事件与订阅者，可以设置优先级调整回调先后顺序，按优先级升序排列
func (hook *Hook) BindWithPriority(event IEvent, priority int32) {
	if event == nil {
		panic("nil event")
	}

	if hook.IsBound() {
		panic("already bound")
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
