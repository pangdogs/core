//go:generate go run github.com/pangdogs/galaxy/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE -export_emit=false
package core

// EventCompMgrAddComponents 事件定义：实体的组件管理器加入一些组件
type EventCompMgrAddComponents interface {
	OnCompMgrAddComponents(entity Entity, components []Component)
}

// EventCompMgrRemoveComponent 事件定义：实体的组件管理器删除组件
type EventCompMgrRemoveComponent interface {
	OnCompMgrRemoveComponent(entity Entity, component Component)
}
