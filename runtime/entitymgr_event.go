//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package runtime

import "git.golaxy.org/core/ec"

// EventEntityMgrAddEntity [EmitUnExport] 事件：实体管理器添加实体
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrRemoveEntity [EmitUnExport] 事件：实体管理器删除实体
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents [EmitUnExport] 事件：实体管理器中的实体添加组件
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(entityMgr EntityMgr, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent [EmitUnExport] 事件：实体管理器中的实体删除组件
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}

// EventEntityMgrEntityFirstAccessComponent [EmitUnExport] 事件：实体管理器中的实体首次访问组件
type EventEntityMgrEntityFirstAccessComponent interface {
	OnEntityMgrEntityFirstAccessComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}
