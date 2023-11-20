package event

import (
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
)

// Deprecated: UnsafeEvent 访问本地事件内部方法
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

// NewHook 创建事件绑定句柄
func (ue _UnsafeEvent) NewHook(subscriberFace iface.FaceAny, priority int32) Hook {
	return ue.newHook(subscriberFace, priority)
}

// RemoveSubscriber 删除订阅者
func (ue _UnsafeEvent) RemoveSubscriber(subscriber any) {
	ue.removeSubscriber(subscriber)
}

// SetGCCollector 设置GC收集器
func (ue _UnsafeEvent) SetGCCollector(gcCollector container.GCCollector) {
	ue.setGCCollector(gcCollector)
}

// GC GC
func (ue _UnsafeEvent) GC() {
	ue.gc()
}
