//go:generate go run git.golaxy.org/core/event/eventcode gen_event
package ec

// EventEntityDestroySelf [EmitUnExport] 事件：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
