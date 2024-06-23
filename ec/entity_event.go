//go:generate go run git.golaxy.org/core/event/eventc event
package ec

// EventEntityDestroySelf 事件：实体销毁自身
// +event-gen:export=0
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
