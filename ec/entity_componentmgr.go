package ec

import (
	"errors"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// _ComponentMgr 组件管理器接口
type _ComponentMgr interface {
	// GetComponent 使用名称查询组件，一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件
	GetComponent(name string) Component
	// GetComponentById 使用组件Id查询组件
	GetComponentById(id uid.Id) Component
	// GetComponents 使用名称查询所有指向的组件
	GetComponents(name string) []Component
	// RangeComponents 遍历所有组件
	RangeComponents(fun func(component Component) bool)
	// ReverseRangeComponents 反向遍历所有组件
	ReverseRangeComponents(fun func(component Component) bool)
	// CountComponents 统计所有组件数量
	CountComponents() int
	// AddComponents 使用同个名称添加多个组件，一般情况下名称指组件接口名称，也可以自定义名称
	AddComponents(name string, components []Component) error
	// AddComponent 添加单个组件，因为同个名称可以指向多个组件，所以名称指向的组件已存在时，不会返回错误
	AddComponent(name string, component Component) error
	// RemoveComponent 删除名称指向的组件，会删除同个名称指向的多个组件
	RemoveComponent(name string)
	// RemoveComponentById 使用组件Id删除组件
	RemoveComponentById(id uid.Id)
	// EventCompMgrAddComponents 事件：实体的组件管理器加入一些组件
	EventCompMgrAddComponents() localevent.IEvent
	// EventCompMgrRemoveComponent 事件：实体的组件管理器删除组件
	EventCompMgrRemoveComponent() localevent.IEvent
	// EventCompMgrFirstAccessComponent 事件：实体的组件管理器首次访问组件
	EventCompMgrFirstAccessComponent() localevent.IEvent
}

// GetComponent 使用名称查询组件，一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件
func (entity *EntityBehavior) GetComponent(name string) Component {
	if e, ok := entity.getComponentElement(name); ok {
		comp := util.Cache2Iface[Component](e.Value.Cache)

		if entity.opts.ComponentAwakeByAccess && comp.GetState() == ComponentState_Attach {
			switch entity.GetState() {
			case EntityState_Init, EntityState_Inited, EntityState_Living:
				emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.CompositeFace.Iface, comp)
			}
		}

		return comp
	}

	return nil
}

// GetComponentById 使用组件Id查询组件
func (entity *EntityBehavior) GetComponentById(id uid.Id) Component {
	if e, ok := entity.getComponentElementById(id); ok {
		comp := util.Cache2Iface[Component](e.Value.Cache)

		if entity.opts.ComponentAwakeByAccess && comp.GetState() == ComponentState_Attach {
			switch entity.GetState() {
			case EntityState_Init, EntityState_Inited, EntityState_Living:
				emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.CompositeFace.Iface, comp)
			}
		}

		return comp
	}

	return nil
}

// GetComponents 使用名称查询所有指向的组件
func (entity *EntityBehavior) GetComponents(name string) []Component {
	if e, ok := entity.getComponentElement(name); ok {
		var components []Component

		entity.componentList.TraversalAt(func(other *container.Element[util.FaceAny]) bool {
			comp := util.Cache2Iface[Component](other.Value.Cache)
			if comp.GetName() == name {
				components = append(components, comp)
				return true
			}
			return false
		}, e)

		if entity.opts.ComponentAwakeByAccess {
			for i := range components {
				if components[i].GetState() == ComponentState_Attach {
					switch entity.GetState() {
					case EntityState_Init, EntityState_Inited, EntityState_Living:
						emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.CompositeFace.Iface, components[i])
					}
				}
			}
		}

		return components
	}

	return nil
}

// RangeComponents 遍历所有组件
func (entity *EntityBehavior) RangeComponents(fun func(component Component) bool) {
	if fun == nil {
		return
	}

	entity.componentList.Traversal(func(e *container.Element[util.FaceAny]) bool {
		comp := util.Cache2Iface[Component](e.Value.Cache)

		if entity.opts.ComponentAwakeByAccess && comp.GetState() == ComponentState_Attach {
			switch entity.GetState() {
			case EntityState_Init, EntityState_Inited, EntityState_Living:
				emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.CompositeFace.Iface, comp)
			}
		}

		return fun(comp)
	})
}

// ReverseRangeComponents 反向遍历所有组件
func (entity *EntityBehavior) ReverseRangeComponents(fun func(component Component) bool) {
	if fun == nil {
		return
	}

	entity.componentList.ReverseTraversal(func(e *container.Element[util.FaceAny]) bool {
		comp := util.Cache2Iface[Component](e.Value.Cache)

		if entity.opts.ComponentAwakeByAccess && comp.GetState() == ComponentState_Attach {
			switch entity.GetState() {
			case EntityState_Init, EntityState_Inited, EntityState_Living:
				emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.CompositeFace.Iface, comp)
			}
		}

		return fun(comp)
	})
}

// CountComponents 统计所有组件数量
func (entity *EntityBehavior) CountComponents() int {
	return entity.componentList.Len()
}

// AddComponents 使用同个名称添加多个组件，一般情况下名称指组件接口名称，也可以自定义名称
func (entity *EntityBehavior) AddComponents(name string, components []Component) error {
	for i := range components {
		if err := entity.addSingleComponent(name, components[i]); err != nil {
			return err
		}
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.CompositeFace.Iface, components)
	return nil
}

// AddComponent 添加单个组件，因为同个名称可以指向多个组件，所以名称指向的组件已存在时，不会返回错误
func (entity *EntityBehavior) AddComponent(name string, component Component) error {
	if err := entity.addSingleComponent(name, component); err != nil {
		return err
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.CompositeFace.Iface, []Component{component})
	return nil
}

// RemoveComponent 删除名称指向的组件，会删除同个名称指向的多个组件
func (entity *EntityBehavior) RemoveComponent(name string) {
	e, ok := entity.getComponentElement(name)
	if !ok {
		return
	}

	entity.componentList.TraversalAt(func(other *container.Element[util.FaceAny]) bool {
		comp := util.Cache2Iface[Component](other.Value.Cache)
		if comp.GetName() != name {
			return false
		}

		if comp.getFixed() {
			return true
		}

		other.Escape()
		comp.setState(ComponentState_Detach)

		entity.changedVersion++

		emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.CompositeFace.Iface, comp)

		return true
	}, e)
}

// RemoveComponentById 使用组件Id删除组件
func (entity *EntityBehavior) RemoveComponentById(id uid.Id) {
	e, ok := entity.getComponentElementById(id)
	if !ok {
		return
	}

	comp := util.Cache2Iface[Component](e.Value.Cache)

	if comp.getFixed() {
		return
	}

	e.Escape()
	comp.setState(ComponentState_Detach)

	entity.changedVersion++

	emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.CompositeFace.Iface, comp)
}

// EventCompMgrAddComponents 事件：实体的组件管理器加入一些组件
func (entity *EntityBehavior) EventCompMgrAddComponents() localevent.IEvent {
	return &entity.eventCompMgrAddComponents
}

// EventCompMgrRemoveComponent 事件：实体的组件管理器删除组件
func (entity *EntityBehavior) EventCompMgrRemoveComponent() localevent.IEvent {
	return &entity.eventCompMgrRemoveComponent
}

// EventCompMgrFirstAccessComponent 事件：实体的组件管理器首次访问组件
func (entity *EntityBehavior) EventCompMgrFirstAccessComponent() localevent.IEvent {
	return &entity.eventCompMgrFirstAccessComponent
}

func (entity *EntityBehavior) addSingleComponent(name string, component Component) error {
	if component == nil {
		return errors.New("nil component")
	}

	if component.GetState() != ComponentState_Birth {
		return errors.New("component state not birth is invalid")
	}

	component.init(name, entity.opts.CompositeFace.Iface, component, entity.opts.HookAllocator, entity.opts.GCCollector)

	face := util.NewFacePair[any](component, component)

	if e, ok := entity.getComponentElement(name); ok {
		entity.componentList.TraversalAt(func(other *container.Element[util.FaceAny]) bool {
			if util.Cache2Iface[Component](other.Value.Cache).GetName() == name {
				e = other
				return true
			}
			return false
		}, e)

		e = entity.componentList.InsertAfter(face, e)

	} else {
		e = entity.componentList.PushBack(face)
	}

	component.setState(ComponentState_Attach)

	entity.changedVersion++

	return nil
}

func (entity *EntityBehavior) getComponentElement(name string) (*container.Element[util.FaceAny], bool) {
	var e *container.Element[util.FaceAny]

	entity.componentList.Traversal(func(other *container.Element[util.FaceAny]) bool {
		if util.Cache2Iface[Component](other.Value.Cache).GetName() == name {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}

func (entity *EntityBehavior) getComponentElementById(id uid.Id) (*container.Element[util.FaceAny], bool) {
	var e *container.Element[util.FaceAny]

	entity.componentList.Traversal(func(other *container.Element[util.FaceAny]) bool {
		if util.Cache2Iface[Component](other.Value.Cache).GetId() == id {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}
