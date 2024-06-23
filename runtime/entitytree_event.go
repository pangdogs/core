//go:generate go run git.golaxy.org/core/event/eventc event
package runtime

import "git.golaxy.org/core/ec"

// EventEntityTreeAddNode 事件：新增实体树节点
// +event-gen:export=0
type EventEntityTreeAddNode interface {
	OnEntityTreeAddNode(entityTree EntityTree, parent, child ec.Entity)
}

// EventEntityTreeRemoveNode 事件：删除实体树节点
// +event-gen:export=0
type EventEntityTreeRemoveNode interface {
	OnEntityTreeRemoveNode(entityTree EntityTree, parent, child ec.Entity)
}
