package ec

import (
	"fmt"
)

// String implements fmt.Stringer
func (comp *ComponentBehavior) String() string {
	var entityInfo string
	if entity := comp.GetEntity(); entity != nil {
		entityInfo = entity.GetId().String()
	} else {
		entityInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s Name:%s Entity:%s State:%s}",
		comp.GetId(), comp.GetName(), entityInfo, comp.GetState())
}
