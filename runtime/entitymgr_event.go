//go:generate go run kit.golaxy.org/golaxy/event/eventcode --decl_file=$GOFILE gen_emit --package=$GOPACKAGE --default_auto=true
package runtime

import "kit.golaxy.org/golaxy/ec"

// EventEntityMgrAddEntity [EmitUnExport] 事件：实体管理器添加实体
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(entityMgr IEntityMgr, entity ec.Entity)
}

// EventEntityMgrRemovingEntity [EmitUnExport] 事件：实体管理器开始删除实体
type EventEntityMgrRemovingEntity interface {
	OnEntityMgrRemovingEntity(entityMgr IEntityMgr, entity ec.Entity)
}

// EventEntityMgrRemoveEntity [EmitUnExport] 事件：实体管理器删除实体
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(entityMgr IEntityMgr, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents [EmitUnExport] 事件：实体管理器中的实体添加组件
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(entityMgr IEntityMgr, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent [EmitUnExport] 事件：实体管理器中的实体删除组件
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(entityMgr IEntityMgr, entity ec.Entity, component ec.Component)
}

// EventEntityMgrEntityFirstAccessComponent [EmitUnExport] 事件：实体管理器中的实体首次访问组件
type EventEntityMgrEntityFirstAccessComponent interface {
	OnEntityMgrEntityFirstAccessComponent(entityMgr IEntityMgr, entity ec.Entity, component ec.Component)
}
