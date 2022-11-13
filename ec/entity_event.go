//go:generate go run github.com/galaxy-kit/galaxy-go/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package ec

// EventEntityDestroySelf [EmitUnExport] 事件定义：实体销毁自身
type EventEntityDestroySelf interface {
	OnEntityDestroySelf(entity Entity)
}
