package ec

import (
	"errors"
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// _ComponentMgr 组件管理器接口
type _ComponentMgr interface {
	// GetComponent 使用名称查询组件，一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件
	GetComponent(name string) Component
	// GetComponentByID 使用组件唯一ID查询组件
	GetComponentByID(id int64) Component
	// GetComponents 使用名称查询所有指向的组件
	GetComponents(name string) []Component
	// RangeComponents 遍历所有组件
	RangeComponents(fun func(component Component) bool)
	// AddComponents 使用同个名称添加多个组件，一般情况下名称指组件接口名称，也可以自定义名称
	AddComponents(name string, components []Component) error
	// AddComponent 添加单个组件，因为同个名称可以指向多个组件，所以名称指向的组件已存在时，不会返回错误
	AddComponent(name string, component Component) error
	// RemoveComponent 删除名称指向的组件，会删除同个名称指向的多个组件
	RemoveComponent(name string)
	// RemoveComponentByID 使用组件唯一ID删除组件
	RemoveComponentByID(id int64)
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

		if entity.opts.EnableComponentAwakeByAccess && !comp.getAwoke() {
			emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.Inheritor.Iface, comp)
		}

		return comp
	}

	return nil
}

// GetComponentByID 使用组件唯一ID查询组件
func (entity *EntityBehavior) GetComponentByID(id int64) Component {
	if e, ok := entity.getComponentElementByID(id); ok {
		comp := util.Cache2Iface[Component](e.Value.Cache)

		if entity.opts.EnableComponentAwakeByAccess && !comp.getAwoke() {
			emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.Inheritor.Iface, comp)
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

		if entity.opts.EnableComponentAwakeByAccess {
			for i := range components {
				if !components[i].getAwoke() {
					emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.Inheritor.Iface, components[i])
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

		if entity.opts.EnableComponentAwakeByAccess && !comp.getAwoke() {
			emitEventCompMgrFirstAccessComponent(&entity.eventCompMgrFirstAccessComponent, entity.opts.Inheritor.Iface, comp)
		}

		return fun(comp)
	})
}

// AddComponents 使用同个名称添加多个组件，一般情况下名称指组件接口名称，也可以自定义名称
func (entity *EntityBehavior) AddComponents(name string, components []Component) error {
	for i := range components {
		if err := entity.addSingleComponent(name, components[i]); err != nil {
			return err
		}
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.Iface, components)
	return nil
}

// AddComponent 添加单个组件，因为同个名称可以指向多个组件，所以名称指向的组件已存在时，不会返回错误
func (entity *EntityBehavior) AddComponent(name string, component Component) error {
	if err := entity.addSingleComponent(name, component); err != nil {
		return err
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.Iface, []Component{component})
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

		if !entity.opts.EnableRemovePrimaryComponent {
			if comp.getPrimary() {
				return true
			}
		}

		other.Escape()
		emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.Iface, comp)

		return true
	}, e)
}

// RemoveComponentByID 使用组件唯一ID删除组件
func (entity *EntityBehavior) RemoveComponentByID(id int64) {
	e, ok := entity.getComponentElementByID(id)
	if !ok {
		return
	}

	comp := util.Cache2Iface[Component](e.Value.Cache)

	if !entity.opts.EnableRemovePrimaryComponent {
		if comp.getPrimary() {
			return
		}
	}

	e.Escape()
	emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.Iface, comp)
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

	if component.GetEntity() != nil {
		return errors.New("component already added in entity")
	}

	component.init(name, entity.opts.Inheritor.Iface, component, entity.opts.HookCache)

	if entity.opts.ComponentPersistID != nil {
		compID := entity.opts.ComponentPersistID(component)
		if compID < 0 {
			panic("component persistID less 0 invalid")
		}
		component.setID(compID)
	}

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

	entity.getInnerGCCollector().CollectGC(component.getInnerGC())

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

func (entity *EntityBehavior) getComponentElementByID(id int64) (*container.Element[util.FaceAny], bool) {
	var e *container.Element[util.FaceAny]

	entity.componentList.Traversal(func(other *container.Element[util.FaceAny]) bool {
		if util.Cache2Iface[Component](other.Value.Cache).GetID() == id {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}
