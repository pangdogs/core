//go:generate go run github.com/pangdogs/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE -export_emit=false
package core

type eventComponentDestroySelf interface {
	onComponentDestroySelf(comp Component)
}
