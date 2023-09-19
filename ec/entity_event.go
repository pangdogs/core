//go:generate go run kit.golaxy.org/golaxy/event/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package ec

// EventEntityDestroySelf [EmitUnExport] 事件：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
