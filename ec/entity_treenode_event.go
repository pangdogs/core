//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package ec

// EventTreeNodeAddChild [EmitUnExport] 事件：实体节点添加子实体
type EventTreeNodeAddChild interface {
	OnTreeNodeAddChild(parent, child Entity)
}

// EventTreeNodeRemoveChild [EmitUnExport] 事件：实体节点删除子实体
type EventTreeNodeRemoveChild interface {
	OnTreeNodeRemoveChild(parent, child Entity)
}

// EventTreeNodeEnterParent [EmitUnExport] 事件：实体加入父实体节点
type EventTreeNodeEnterParent interface {
	OnTreeNodeEnterParent(child, parent Entity)
}

// EventTreeNodeLeaveParent [EmitUnExport] 事件：实体离开父实体节点
type EventTreeNodeLeaveParent interface {
	OnTreeNodeLeaveParent(child, parent Entity)
}
