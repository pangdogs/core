//go:generate go run github.com/pangdogs/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE  -export_emit=false
package core

type EventECTreeAddChild interface {
	OnAddChild(ecTree IECTree, parent, child Entity)
}

type EventECTreeRemoveChild interface {
	OnRemoveChild(ecTree IECTree, parent, child Entity)
}
