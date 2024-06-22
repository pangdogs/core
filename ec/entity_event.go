//go:generate go run git.golaxy.org/core/event/eventc event
package ec

// EventEntityDestroySelf [EmitUnExport] 事件：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
