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

// GetAndCheckEntity 查询并检测实体
func (ue _UnsafeECTree) GetAndCheckEntity(entityId uid.Id) (ec.Entity, error) {
	return ue.getAndCheckEntity(entityId)
}
