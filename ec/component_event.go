//go:generate go run git.golaxy.org/core/event/eventc event

package ec

// EventComponentDestroySelf 事件：组件销毁自身
// +event-gen:export=0
type EventComponentDestroySelf interface {
	OnComponentDestroySelf(comp Component)
}
