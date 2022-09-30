//go:generate go run github.com/pangdogs/galaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package ec

// EventEntityDestroySelf [EmitUnExport] 事件定义：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
