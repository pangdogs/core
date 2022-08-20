package core

import (
	"errors"
	"github.com/pangdogs/core/container"
)

func (entity *EntityBehavior) GetComponent(name string) Component {
	if e, ok := entity.getComponentElement(name); ok {
		comp := Cache2IFace[Component](e.Value.Cache)
		return comp
	}

	return nil
}

func (entity *EntityBehavior) GetComponentByID(id uint64) Component {
	if e, ok := entity.getComponentElementByID(id); ok {
		comp := Cache2IFace[Component](e.Value.Cache)
		return comp
	}

	return nil
}

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

func (entity *EntityBehavior) RangeComponents(fun func(component Component) bool) {
	if fun == nil {
		return
	}

	entity.componentList.Traversal(func(e *container.Element[FaceAny]) bool {
		comp := Cache2IFace[Component](e.Value.Cache)
		return fun(comp)
	})
}

func (entity *EntityBehavior) AddComponents(name string, components []Component) error {
	for i := range components {
		if err := entity.addSingleComponent(name, components[i]); err != nil {
			return err
		}
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.IFace, components)
	return nil
}

func (entity *EntityBehavior) AddComponent(name string, component Component) error {
	if err := entity.addSingleComponent(name, component); err != nil {
		return err
	}

	emitEventCompMgrAddComponents(&entity.eventCompMgrAddComponents, entity.opts.Inheritor.IFace, []Component{component})
	return nil
}

func (entity *EntityBehavior) RemoveComponent(name string) {
	e, ok := entity.getComponentElement(name)
	if !ok {
		return
	}

	if entity.opts.EnableFastGetComponent {
		delete(entity.componentMap, name)
	}

	entity.componentList.TraversalAt(func(other *container.Element[FaceAny]) bool {
		comp := Cache2IFace[Component](other.Value.Cache)
		if comp.GetName() == name {
			other.Escape()
			emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.IFace, comp)
			return true
		}
		return false
	}, e)
}

func (entity *EntityBehavior) RemoveComponentByID(id uint64) {
	e, ok := entity.getComponentElementByID(id)
	if !ok {
		return
	}

	if entity.opts.EnableFastGetComponentByID {
		delete(entity.componentByIDMap, id)
	}

	e.Escape()
	emitEventCompMgrRemoveComponent(&entity.eventCompMgrRemoveComponent, entity.opts.Inheritor.IFace, Cache2IFace[Component](e.Value.Cache))
}

func (entity *EntityBehavior) EventCompMgrAddComponents() IEvent {
	return &entity.eventCompMgrAddComponents
}

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

		if entity.opts.EnableFastGetComponent {
			entity.componentMap[name] = e
		}

		if entity.opts.EnableFastGetComponentByID {
			entity.componentByIDMap[component.GetID()] = e
		}
	}

	entity.CollectGC(component)

	return nil
}

func (entity *EntityBehavior) getComponentElement(name string) (*container.Element[FaceAny], bool) {
	if entity.opts.EnableFastGetComponent {
		e, ok := entity.componentMap[name]
		return e, ok
	}

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

func (entity *EntityBehavior) getComponentElementByID(id uint64) (*container.Element[FaceAny], bool) {
	if entity.opts.EnableFastGetComponentByID {
		e, ok := entity.componentByIDMap[id]
		return e, ok
	}

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
