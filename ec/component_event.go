//go:generate go run github.com/golaxy-kit/golaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE

package ec

// EventComponentDestroySelf [EmitUnExport] 事件：组件销毁自身
type EventComponentDestroySelf interface {
	OnComponentDestroySelf(comp Component)
}
