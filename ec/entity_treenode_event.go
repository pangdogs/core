//go:generate go run git.golaxy.org/core/event/eventc event
package ec

// EventTreeNodeAddChild 事件：实体节点添加子实体
// +event-gen:export=0
type EventTreeNodeAddChild interface {
	OnTreeNodeAddChild(parent, child Entity)
}

// EventTreeNodeRemoveChild 事件：实体节点删除子实体
// +event-gen:export=0
type EventTreeNodeRemoveChild interface {
	OnTreeNodeRemoveChild(parent, child Entity)
}

// EventTreeNodeEnterParent 事件：实体加入父实体节点
// +event-gen:export=0
type EventTreeNodeEnterParent interface {
	OnTreeNodeEnterParent(child, parent Entity)
}

// EventTreeNodeLeaveParent 事件：实体离开父实体节点
// +event-gen:export=0
type EventTreeNodeLeaveParent interface {
	OnTreeNodeLeaveParent(child, parent Entity)
}
