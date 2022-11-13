//go:generate go run github.com/galaxy-kit/galaxy-go/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE

package ec

// EventComponentDestroySelf [EmitUnExport] 事件定义：组件销毁自身
type EventComponentDestroySelf interface {
	OnComponentDestroySelf(comp Component)
}
