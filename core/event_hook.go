package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// BindEvent ...
func BindEvent[T any](event IEvent, delegate T) Hook {
	return BindEventWithPriority(event, delegate, 0)
}

// BindEventWithPriority ...
func BindEventWithPriority[T any](event IEvent, delegate T, priority int32) Hook {
	if event == nil {
		panic("nil event")
	}
	return event.newHook(delegate, IFace2Cache(delegate), priority)
}

// UnbindEvent ...
func UnbindEvent(event IEvent, delegate interface{}) {
	if event == nil {
		panic("nil event")
	}
	event.removeDelegate(delegate)
}

// Hook ...
type Hook struct {
	delegate          interface{}
	delegateFastIFace IFaceCache
	priority          int32
	element           *container.Element[Hook]
	received          int
}

// Bind ...
func (hook *Hook) Bind(event IEvent) {
	hook.BindWithPriority(event, 0)
}

// BindWithPriority ...
func (hook *Hook) BindWithPriority(event IEvent, priority int32) {
	if event == nil {
		panic("nil event")
	}

	if hook.IsBound() {
		panic("repeated bind event invalid")
	}

	*hook = event.newHook(hook.delegate, hook.delegateFastIFace, priority)
}

// Unbind ...
func (hook *Hook) Unbind() {
	if hook.element != nil {
		hook.element.Escape()
		hook.element = nil
	}
}

// IsBound ...
func (hook *Hook) IsBound() bool {
	return hook.element != nil && !hook.element.Escaped()
}

// Delegate ...
func (hook *Hook) Delegate() IFaceCache {
	return hook.delegateFastIFace
}
