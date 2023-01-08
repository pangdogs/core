//go:generate go run github.com/golaxy-kit/golaxy/localevent/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE
package runtime

import "github.com/golaxy-kit/golaxy/ec"

// EventEntityMgrAddEntity [EmitUnExport] 事件定义：实体管理器中添加实体
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(entityMgr IEntityMgr, entity ec.Entity)
}

// EventEntityMgrRemoveEntity [EmitUnExport] 事件定义：实体管理器中删除实体
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(entityMgr IEntityMgr, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents [EmitUnExport] 事件定义：实体管理器中的实体添加组件
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(entityMgr IEntityMgr, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent [EmitUnExport] 事件定义：实体管理器中的实体删除组件
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(entityMgr IEntityMgr, entity ec.Entity, component ec.Component)
}

// EventEntityMgrEntityFirstAccessComponent [EmitUnExport] 事件定义：实体管理器中的实体首次访问组件
type EventEntityMgrEntityFirstAccessComponent interface {
	OnEntityMgrEntityFirstAccessComponent(entityMgr IEntityMgr, entity ec.Entity, component ec.Component)
}

// eventEntityMgrNotifyECTreeRemoveEntity [EmitUnExport]
type eventEntityMgrNotifyECTreeRemoveEntity interface {
	onEntityMgrNotifyECTreeRemoveEntity(entityMgr IEntityMgr, entity ec.Entity)
}
