//go:generate go run github.com/pangdogs/galaxy/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE  -export_emit=false
package core

// EventECTreeAddChild 事件定义：EC树中子实体加入父实体
type EventECTreeAddChild interface {
	OnAddChild(ecTree IECTree, parent, child Entity)
}

// EventECTreeRemoveChild 事件定义：EC树中子实体离开父实体
type EventECTreeRemoveChild interface {
	OnRemoveChild(ecTree IECTree, parent, child Entity)
}
