package core

type ComponentQuery interface {
	GetComponent(name string) Component
	GetComponentByID(id uint64) Component
	GetComponents(name string) []Component
	RangeComponents(fun func(component Component) bool)
}

type ComponentMgr interface {
	ComponentQuery
	AddComponents(name string, components []Component) error
	AddComponent(name string, component Component) error
	RemoveComponent(name string)
	RemoveComponentByID(id uint64)
}

type ComponentMgrEvents interface {
	EventCompMgrAddComponents() IEvent
	EventCompMgrRemoveComponent() IEvent
}
