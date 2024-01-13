package runtime

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/util/uid"
)

// Deprecated: UnsafeECTree 访问EC树内部方法
func UnsafeECTree(ecTree ECTree) _UnsafeECTree {
	return _UnsafeECTree{
		ECTree: ecTree,
	}
}

type _UnsafeECTree struct {
	ECTree
}

// FetchEntity 访问实体
func (ue _UnsafeECTree) FetchEntity(entityId uid.Id) (ec.Entity, error) {
	return ue.fetchEntity(entityId)
}
