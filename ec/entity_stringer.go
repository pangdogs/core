package ec

import (
	"fmt"
)

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	var parentInfo string
	if parent, ok := entity.GetParent(); ok {
		parentInfo = parent.GetId().String()
	} else {
		parentInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s Prototype:%s Parent:%s State:%s}",
		entity.GetId(), entity.GetPrototype(), parentInfo, entity.GetState())
}
