package runtime

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/util/uid"
)

// Deprecated: UnsafeECTree 访问EC树内部方法
func UnsafeECTree(ecTree IECTree) _UnsafeECTree {
	return _UnsafeECTree{
		IECTree: ecTree,
	}
}

type _UnsafeECTree struct {
	IECTree
}

// FetchEntity 访问实体
func (ue _UnsafeECTree) FetchEntity(entityId uid.Id) (ec.Entity, error) {
	return ue.fetchEntity(entityId)
}
