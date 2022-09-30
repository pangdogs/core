package runtime

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/util"
)

func EntityContext(entity ec.Entity) Context {
	if entity == nil {
		panic("nil entity")
	}

	ctx := ec.UnsafeEntity(entity).GetContext()
	if ctx == util.NilIfaceCache {
		panic("nil context")
	}

	return util.Cache2Iface[Context](ctx)
}

func ComponentContext(comp ec.Component) Context {
	if comp == nil {
		panic("nil comp")
	}

	return EntityContext(comp.GetEntity())
}
