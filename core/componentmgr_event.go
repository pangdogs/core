//go:generate go run github.com/pangdogs/galaxy/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE -export_emit=false
package core

type EventCompMgrAddComponents[T any] interface {
	OnCompMgrAddComponents(compMgr T, components []Component)
}

type EventCompMgrRemoveComponent[T any] interface {
	OnCompMgrRemoveComponent(compMgr T, component Component)
}
