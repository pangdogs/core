//go:generate go run github.com/pangdogs/galaxy/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE -export_emit=false
package core

// EventEntityMgrAddEntity 事件定义：运行时上下文（Runtime Context）添加实体（Entity）
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(runtimeCtx RuntimeContext, entity Entity)
}

// EventEntityMgrRemoveEntity 事件定义：运行时上下文（Runtime Context）删除实体（Entity）
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(runtimeCtx RuntimeContext, entity Entity)
}

// EventEntityMgrEntityAddComponents 事件定义：运行时上下文（Runtime Context）中的实体（Entity）添加组件（Component）
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(runtimeCtx RuntimeContext, entity Entity, components []Component)
}

// EventEntityMgrEntityRemoveComponent 事件定义：运行时上下文（Runtime Context）中的实体（Entity）删除组件（Component）
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(runtimeCtx RuntimeContext, entity Entity, component Component)
}

type eventEntityMgrNotifyECTreeRemoveEntity interface {
	onEntityMgrNotifyECTreeRemoveEntity(runtimeCtx RuntimeContext, entity Entity)
}
