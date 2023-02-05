package ec

import (
	"kit.golaxy.org/golaxy/util"
)

func UnsafeContextResolver(ctxResolver ContextResolver) _UnsafeContextResolver {
	return _UnsafeContextResolver{
		ContextResolver: ctxResolver,
	}
}

type _UnsafeContextResolver struct {
	ContextResolver
}

func (u _UnsafeContextResolver) GetContext() util.IfaceCache {
	return u.getContext()
}
