package runtime

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/util/uid"
)

// Deprecated: UnsafeEntityTree 访问实体树内部方法
func UnsafeEntityTree(entityTree EntityTree) _UnsafeEntityTree {
	return _UnsafeEntityTree{
		EntityTree: entityTree,
	}
}

type _UnsafeEntityTree struct {
	EntityTree
}

// GetAndCheckEntity 查询并检测实体
func (ue _UnsafeEntityTree) GetAndCheckEntity(entityId uid.Id) (ec.Entity, error) {
	return ue.getAndCheckEntity(entityId)
}
