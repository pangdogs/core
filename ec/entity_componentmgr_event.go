//go:generate go run git.golaxy.org/core/event/eventc event
package ec

// EventComponentMgrAddComponents 事件：实体的组件管理器添加组件
// +event-gen:export=0
type EventComponentMgrAddComponents interface {
	OnComponentMgrAddComponents(entity Entity, components []Component)
}

// EventComponentMgrRemoveComponent 事件：实体的组件管理器删除组件
// +event-gen:export=0
type EventComponentMgrRemoveComponent interface {
	OnComponentMgrRemoveComponent(entity Entity, component Component)
}

// EventComponentMgrFirstAccessComponent 事件：实体的组件管理器首次访问组件
// +event-gen:export=0
type EventComponentMgrFirstAccessComponent interface {
	OnComponentMgrFirstAccessComponent(entity Entity, component Component)
}
