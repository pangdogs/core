//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package ec

// EventEntityDestroySelf [EmitUnExport] 事件：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
