package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

func (comp *ComponentBehavior) String() string {
	return ComponentStringer.String(comp.composite)
}

var ComponentStringer internal.Stringer[Component] = DefaultComponentStringer{}

type DefaultComponentStringer struct{}

func (DefaultComponentStringer) String(comp Component) string {
	var entityInfo string
	if entity := comp.GetEntity(); entity != nil {
		entityInfo = entity.GetId().String()
	} else {
		entityInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s SerialNo:%d Name:%s Entity:%s State:%s}",
		comp.GetId(), comp.GetSerialNo(), comp.GetName(), entityInfo, comp.GetState())
}
