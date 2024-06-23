//go:generate go run git.golaxy.org/core/event/eventc event
package runtime

import "git.golaxy.org/core/ec"

// EventEntityMgrAddEntity 事件：实体管理器添加实体
// +event-gen:export=0
type EventEntityMgrAddEntity interface {
	OnEntityMgrAddEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrRemoveEntity 事件：实体管理器删除实体
// +event-gen:export=0
type EventEntityMgrRemoveEntity interface {
	OnEntityMgrRemoveEntity(entityMgr EntityMgr, entity ec.Entity)
}

// EventEntityMgrEntityAddComponents 事件：实体管理器中的实体添加组件
// +event-gen:export=0
type EventEntityMgrEntityAddComponents interface {
	OnEntityMgrEntityAddComponents(entityMgr EntityMgr, entity ec.Entity, components []ec.Component)
}

// EventEntityMgrEntityRemoveComponent 事件：实体管理器中的实体删除组件
// +event-gen:export=0
type EventEntityMgrEntityRemoveComponent interface {
	OnEntityMgrEntityRemoveComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}

// EventEntityMgrEntityFirstAccessComponent 事件：实体管理器中的实体首次访问组件
// +event-gen:export=0
type EventEntityMgrEntityFirstAccessComponent interface {
	OnEntityMgrEntityFirstAccessComponent(entityMgr EntityMgr, entity ec.Entity, component ec.Component)
}
