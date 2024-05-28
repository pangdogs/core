package event

import (
	"git.golaxy.org/core/utils/iface"
)

// Deprecated: UnsafeEvent 访问事件内部方法
func UnsafeEvent(event IEvent) _UnsafeEvent {
	return _UnsafeEvent{
		IEvent: event,
	}
}

type _UnsafeEvent struct {
	IEvent
}

// Emit 发送事件
func (ue _UnsafeEvent) Emit(fun func(subscriber iface.Cache) bool) {
	ue.emit(fun)
}

// NewHook 创建事件钩子
func (ue _UnsafeEvent) NewHook(subscriberFace iface.FaceAny, priority int32) Hook {
	return ue.newHook(subscriberFace, priority)
}

// RemoveSubscriber 删除订阅者
func (ue _UnsafeEvent) RemoveSubscriber(subscriber any) {
	ue.removeSubscriber(subscriber)
}
