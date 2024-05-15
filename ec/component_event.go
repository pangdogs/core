//go:generate go run git.golaxy.org/core/event/eventcode gen_event

package ec

// EventComponentDestroySelf [EmitUnExport] 事件：组件销毁自身
type EventComponentDestroySelf interface {
	OnComponentDestroySelf(comp Component)
}
