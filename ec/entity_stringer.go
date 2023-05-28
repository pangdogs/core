package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

func (entity *EntityBehavior) String() string {
	return EntityStringer.String(entity.opts.CompositeFace.Iface)
}

var EntityStringer internal.Stringer[Entity] = DefaultEntityStringer{}

type DefaultEntityStringer struct{}

func (DefaultEntityStringer) String(entity Entity) string {
	var parentInfo string
	if parent, ok := entity.GetParent(); ok {
		parentInfo = parent.GetId().String()
	} else {
		parentInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s SerialNo:%d Prototype:%s Parent:%s State:%s}",
		entity.GetId(), entity.GetSerialNo(), entity.GetPrototype(), parentInfo, entity.GetState())
}
