package ec

import (
	"github.com/golaxy-kit/golaxy/util"
)

func UnsafeContextResolver(ctxHolder ContextResolver) _UnsafeContextResolver {
	return _UnsafeContextResolver{
		ContextResolver: ctxHolder,
	}
}

type _UnsafeContextResolver struct {
	ContextResolver
}

func (u _UnsafeContextResolver) GetContext() util.IfaceCache {
	return u.getContext()
}
