//go:generate go run github.com/pangdogs/galaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE

package ec

// EventComponentDestroySelf [EmitUnExport] 事件定义：组件销毁自身
type EventComponentDestroySelf interface {
	OnComponentDestroySelf(comp Component)
}
