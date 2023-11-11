package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
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
	RangeComponents(fun generic.Func1[Component, bool])
	// ReverseRangeComponents 反向遍历所有组件
	ReverseRangeComponents(fun generic.Func1[Component, bool])
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

	iAutoEventCompMgrAddComponents        // 事件：实体的组件管理器加入一些组件
	iAutoEventCompMgrRemoveComponent      // 事件：实体的组件管理器删除组件
	iAutoEventCompMgrFirstAccessComponent // 事件：实体的组件管理器首次访问组件
}

// GetComponent 使用名称查询组件，一般情况下名称指组件接口名称，也可以自定义名称，同个名称指向多个组件时，返回首个组件
func (entity *EntityBehavior) GetComponent(name string) Component {
	if e, ok := entity.getComponentElement(name); ok {
		return entity.accessComponent(iface.Cache2Iface[Component](e.Value.Cache))
	}
	return nil
}

// GetComponentById 使用组件Id查询组件
func (entity *EntityBehavior) GetComponentById(id uid.Id) Component {
	if e, ok := entity.getComponentElementById(id); ok {
		return entity.accessComponent(iface.Cache2Iface[Component](e.Value.Cache))
	}
	return nil
}

// GetComponents 使用名称查询所有指向的组件
func (entity *EntityBehavior) GetComponents(name string) []Component {
	if e, ok := entity.getComponentElement(name); ok {
		var components []Component

		entity.componentList.TraversalAt(func(other *container.Element[iface.FaceAny]) bool {
			comp := iface.Cache2Iface[Component](other.Value.Cache)
			if comp.GetName() == name {
				components = append(components, comp)
				return true
			}
			return false
		}, e)

		for i := range components {
			if entity.accessComponent(components[i]) == nil {
				components[i] = nil
			}
		}

		for i := len(components) - 1; i >= 0; i-- {
			if components[i] == nil {
				components = append(components[:i], components[i+1:]...)
			}
		}

		return components
	}

	return nil
}

// RangeComponents 遍历所有组件
func (entity *EntityBehavior) RangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		comp := entity.accessComponent(iface.Cache2Iface[Component](e.Value.Cache))
		if comp == nil {
			return true
		}
		return fun.Exec(comp)
	})
}

// ReverseRangeComponents 反向遍历所有组件
func (entity *EntityBehavior) ReverseRangeComponents(fun generic.Func1[Component, bool]) {
	entity.componentList.ReverseTraversal(func(e *container.Element[iface.FaceAny]) bool {
		comp := entity.accessComponent(iface.Cache2Iface[Component](e.Value.Cache))
		if comp == nil {
			return true
		}
		return fun.Exec(comp)
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

	emitEventCompMgrAddComponents(entity, entity.opts.CompositeFace.Iface, components)
	return nil
}

// AddComponent 添加单个组件，因为同个名称可以指向多个组件，所以名称指向的组件已存在时，不会返回错误
func (entity *EntityBehavior) AddComponent(name string, component Component) error {
	if err := entity.addSingleComponent(name, component); err != nil {
		return err
	}

	emitEventCompMgrAddComponents(entity, entity.opts.CompositeFace.Iface, []Component{component})
	return nil
}

// RemoveComponent 删除名称指向的组件，会删除同个名称指向的多个组件
func (entity *EntityBehavior) RemoveComponent(name string) {
	e, ok := entity.getComponentElement(name)
	if !ok {
		return
	}

	entity.componentList.TraversalAt(func(other *container.Element[iface.FaceAny]) bool {
		comp := iface.Cache2Iface[Component](other.Value.Cache)
		if comp.GetName() != name {
			return false
		}

		if comp.getFixed() {
			return true
		}

		other.Escape()
		comp.setState(ComponentState_Detach)

		entity.version++

		emitEventCompMgrRemoveComponent(entity, entity.opts.CompositeFace.Iface, comp)

		return true
	}, e)
}

// RemoveComponentById 使用组件Id删除组件
func (entity *EntityBehavior) RemoveComponentById(id uid.Id) {
	e, ok := entity.getComponentElementById(id)
	if !ok {
		return
	}

	comp := iface.Cache2Iface[Component](e.Value.Cache)

	if comp.getFixed() {
		return
	}

	e.Escape()
	comp.setState(ComponentState_Detach)

	entity.version++

	emitEventCompMgrRemoveComponent(entity, entity.opts.CompositeFace.Iface, comp)
}

// EventCompMgrAddComponents 事件：实体的组件管理器加入一些组件
func (entity *EntityBehavior) EventCompMgrAddComponents() event.IEvent {
	return &entity.eventCompMgrAddComponents
}

// EventCompMgrRemoveComponent 事件：实体的组件管理器删除组件
func (entity *EntityBehavior) EventCompMgrRemoveComponent() event.IEvent {
	return &entity.eventCompMgrRemoveComponent
}

// EventCompMgrFirstAccessComponent 事件：实体的组件管理器首次访问组件
func (entity *EntityBehavior) EventCompMgrFirstAccessComponent() event.IEvent {
	return &entity.eventCompMgrFirstAccessComponent
}

func (entity *EntityBehavior) addSingleComponent(name string, component Component) error {
	if component == nil {
		return fmt.Errorf("%w: %w: component is nil", ErrEC, exception.ErrArgs)
	}

	if component.GetState() != ComponentState_Birth {
		return fmt.Errorf("%w: invalid component state %q", ErrEC, component.GetState())
	}

	component.init(name, entity.opts.CompositeFace.Iface, component, entity.opts.HookAllocator, entity.opts.GCCollector)

	face := iface.MakeFacePair[any](component, component)

	if e, ok := entity.getComponentElement(name); ok {
		entity.componentList.TraversalAt(func(other *container.Element[iface.FaceAny]) bool {
			if iface.Cache2Iface[Component](other.Value.Cache).GetName() == name {
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

	entity.version++

	return nil
}

func (entity *EntityBehavior) getComponentElement(name string) (*container.Element[iface.FaceAny], bool) {
	var e *container.Element[iface.FaceAny]

	entity.componentList.Traversal(func(other *container.Element[iface.FaceAny]) bool {
		if iface.Cache2Iface[Component](other.Value.Cache).GetName() == name {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}

func (entity *EntityBehavior) getComponentElementById(id uid.Id) (*container.Element[iface.FaceAny], bool) {
	var e *container.Element[iface.FaceAny]

	entity.componentList.Traversal(func(other *container.Element[iface.FaceAny]) bool {
		if iface.Cache2Iface[Component](other.Value.Cache).GetId() == id {
			e = other
			return false
		}
		return true
	})

	return e, e != nil
}

func (entity *EntityBehavior) accessComponent(comp Component) Component {
	if entity.opts.AwakeOnFirstAccess && comp.GetState() == ComponentState_Attach {
		switch entity.GetState() {
		case EntityState_Awake, EntityState_Start, EntityState_Living:
			emitEventCompMgrFirstAccessComponent(entity, entity.opts.CompositeFace.Iface, comp)
		}
	}

	if comp.GetState() >= ComponentState_Detach {
		return nil
	}

	return comp
}
