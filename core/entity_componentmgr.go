package core

import (
	"errors"
	"github.com/pangdogs/galaxy/core/container"
)

// _EntityComponentMgr 实体（Entity）的组件（Component）管理器接口
type _EntityComponentMgr interface {
	// GetComponent 使用名称查询组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件，非线程安全
	GetComponent(name string) Component

	// GetComponentByID 使用组件（Component）运行时ID查询组件，非线程安全
	GetComponentByID(id int64) Component

	// GetComponents 使用名称查询所有指向的组件（Component），非线程安全
	GetComponents(name string) []Component

	// RangeComponents 遍历所有组件，非线程安全
	RangeComponents(fun func(component Component) bool)

	// AddComponents 使用同个名称添加多个组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，非线程安全
	AddComponents(name string, components []Component) error

	// AddComponent 添加单个组件（Component），因为同个名称可以指向多个组件，所有名称指向的组件已存在时，不会返回错误，非线程安全
	AddComponent(name string, component Component) error

	// RemoveComponent 删除名称指向的组件（Component），会删除同个名称指向的多个组件，非线程安全
	RemoveComponent(name string)

	// RemoveComponentByID 使用组件（Component）运行时ID删除组件，非线程安全
	RemoveComponentByID(id int64)

	// EventCompMgrAddComponents 事件：实体的组件管理器加入一些组件
	EventCompMgrAddComponents() IEvent

	// EventCompMgrRemoveComponent 事件：实体的组件管理器删除组件
	EventCompMgrRemoveComponent() IEvent
}

// GetComponent 使用名称查询组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件，非线程安全
func (entity *EntityBehavior) GetComponent(name string) Component {
	if e, ok := entity.getComponentElement(name); ok {
		comp := Cache2IFace[Component](e.Value.Cache)
		return comp
	}

	return nil
}

// GetComponentByID 使用组件（Component）运行时ID查询组件，非线程安全
func (entity *EntityBehavior) GetComponentByID(id int64) Component {
	if e, ok := entity.getComponentElementByID(id); ok {
		comp := Cache2IFace[Component](e.Value.Cache)
		return comp
	}

	return nil
}

// GetComponents 使用名称查询所有指向的组件（Component），非线程安全
func (entity *EntityBehavior) GetComponents(name string) []Component {
	if e, ok := entity.getComponentElement(name); ok {
		var components []Component

		entity.componentList.TraversalAt(func(other *container.Element[FaceAny]) bool {
			comp := Cache2IFace[Component](other.Value.Cache)
			if comp.GetName() == name {
				components = append(components, comp)
				return true
			}
			return false
		}, e)

		return components
	}

	return nil
}

// RangeComponents 遍历所有组件，非线程安全
func (entity *EntityBehavior) RangeComponents(fun func(component Component) bool) {
	if fun == nil {
		return
	}

	entity.componentList.Traversal(func(e *container.Element[FaceAny]) bool {
		comp := Cache2IFace[Component](e.Value.Cache)
		return fun(comp)
	})
}

// AddComponents 使用同个名称添加多个组件（Component），一般情况下名称指组件接口名称，也可以自定义名称，非线程安全
func (entity *EntityBehavior) AddComponents(name string, components []Component) error {
	for i := range components {
		if err := entity.addSingleComponent(name, components[i]); err != nil {
			return err
		}
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.IFace, components)
	return nil
}

// AddComponent 添加单个组件（Component），因为同个名称可以指向多个组件，所有名称指向的组件已存在时，不会返回错误，非线程安全
func (entity *EntityBehavior) AddComponent(name string, component Component) error {
	if err := entity.addSingleComponent(name, component); err != nil {
		return err
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.IFace, []Component{component})
	return nil
}

// RemoveComponent 删除名称指向的组件（Component），会删除同个名称指向的多个组件，非线程安全
func (entity *EntityBehavior) RemoveComponent(name string) {
	e, ok := entity.getComponentElement(name)
	if !ok {
		return
	}

	entity.componentList.TraversalAt(func(other *container.Element[FaceAny]) bool {
		comp := Cache2IFace[Component](other.Value.Cache)
		if comp.GetName() != name {
			return false
		}

		if !entity.opts.EnableRemovePrimaryComponent {
			if comp.getPrimary() {
				return true
			}
		}

		other.Escape()
		emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.IFace, comp)

		return true
	}, e)
}

// RemoveComponentByID 使用组件（Component）运行时ID删除组件，非线程安全
func (entity *EntityBehavior) RemoveComponentByID(id int64) {
	e, ok := entity.getComponentElementByID(id)
	if !ok {
		return
	}

	if !entity.opts.EnableRemovePrimaryComponent {
		comp := Cache2IFace[Component](e.Value.Cache)
		if comp.getPrimary() {
			return
		}
	}

	e.Escape()
	emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.IFace, Cache2IFace[Component](e.Value.Cache))
}

// EventCompMgrAddComponents 事件：实体的组件管理器加入一些组件
func (entity *EntityBehavior) EventCompMgrAddComponents() IEvent {
	return &entity.eventCompMgrAddComponents
}

// EventCompMgrRemoveComponent 事件：实体的组件管理器删除组件
func (entity *EntityBehavior) EventCompMgrRemoveComponent() IEvent {
	return &entity.eventCompMgrRemoveComponent
}

func (entity *EntityBehavior) addSingleComponent(name string, component Component) error {
	if component == nil {
		return errors.New("nil component")
	}

	if component.GetEntity() != nil {
		return errors.New("component already added in entity")
	}

	component.init(name, entity.opts.Inheritor.IFace, component, entity.opts.HookCache)

	face := FaceAny{
		IFace: component,
		Cache: IFace2Cache(component),
	}

	if e, ok := entity.getComponentElement(name); ok {
		entity.componentList.TraversalAt(func(other *container.Element[FaceAny]) bool {
			if Cache2IFace[Component](other.Value.Cache).GetName() == name {
				e = other
				return true
			}
			return false
		}, e)

		e = entity.componentList.InsertAfter(face, e)

	} else {
		e = entity.componentList.PushBack(face)
	}

	entity.getGCCollector().CollectGC(component.getGC())

	return nil
}

func (entity *EntityBehavior) getComponentElement(name string) (*container.Element[FaceAny], bool) {
	var e *container.Element[FaceAny]

	entity.componentList.Traversal(func(other *container.Element[FaceAny]) bool {
		if Cache2IFace[Component](other.Value.Cache).GetName() == name {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}

func (entity *EntityBehavior) getComponentElementByID(id int64) (*container.Element[FaceAny], bool) {
	var e *container.Element[FaceAny]

	entity.componentList.Traversal(func(other *container.Element[FaceAny]) bool {
		if Cache2IFace[Component](other.Value.Cache).GetID() == id {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}
