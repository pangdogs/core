package ec

import (
	"github.com/golaxy-kit/golaxy/util"
)

func UnsafeContextHolder(ctxHolder ContextHolder) _UnsafeContextHolder {
	return _UnsafeContextHolder{
		ContextHolder: ctxHolder,
	}
}

type _UnsafeContextHolder struct {
	ContextHolder
}

func (u _UnsafeContextHolder) GetContext() util.IfaceCache {
	return u.getContext()
}
