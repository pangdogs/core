//go:generate go run github.com/pangdogs/galaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package runtime

import "github.com/pangdogs/galaxy/ec"

// EventEntityMgrAddEntity [EmitUnExport] 事件定义：运行时上下文添加实体
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(runtimeCtx Context, entity ec.Entity)
}

// EventEntityMgrRemoveEntity [EmitUnExport] 事件定义：运行时上下文删除实体
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(runtimeCtx Context, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents [EmitUnExport] 事件定义：运行时上下文中的实体添加组件
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(runtimeCtx Context, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent [EmitUnExport] 事件定义：运行时上下文中的实体删除组件
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(runtimeCtx Context, entity ec.Entity, component ec.Component)
}

// eventEntityMgrNotifyECTreeRemoveEntity [EmitUnExport]
type eventEntityMgrNotifyECTreeRemoveEntity interface {
	onEntityMgrNotifyECTreeRemoveEntity(runtimeCtx Context, entity ec.Entity)
}
