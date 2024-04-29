//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package ec

// EventComponentMgrAddComponents [EmitUnExport] 事件：实体的组件管理器添加组件
type EventComponentMgrAddComponents interface {
	OnComponentMgrAddComponents(entity Entity, components []Component)
}

// EventComponentMgrRemoveComponent [EmitUnExport] 事件：实体的组件管理器删除组件
type EventComponentMgrRemoveComponent interface {
	OnComponentMgrRemoveComponent(entity Entity, component Component)
}

// EventComponentMgrFirstAccessComponent [EmitUnExport] 事件：实体的组件管理器首次访问组件
type EventComponentMgrFirstAccessComponent interface {
	OnComponentMgrFirstAccessComponent(entity Entity, component Component)
}
