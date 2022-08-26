//go:generate go run github.com/pangdogs/galaxy/core/eventcode -decl $GOFILE -core "" -emit_package $GOPACKAGE -export_emit=false
package core

// EventEntityMgrAddEntity ...
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(runtimeCtx RuntimeContext, entity Entity)
}

// EventEntityMgrRemoveEntity ...
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(runtimeCtx RuntimeContext, entity Entity)
}

// EventEntityMgrEntityAddComponents ...
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(runtimeCtx RuntimeContext, entity Entity, components []Component)
}

// EventEntityMgrEntityRemoveComponent ...
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(runtimeCtx RuntimeContext, entity Entity, component Component)
}

type eventEntityMgrNotifyECTreeRemoveEntity interface {
	onEntityMgrNotifyECTreeRemoveEntity(runtimeCtx RuntimeContext, entity Entity)
}
