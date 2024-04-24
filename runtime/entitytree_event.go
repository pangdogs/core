//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package runtime

import "git.golaxy.org/core/ec"

// EventEntityTreeAddNode [EmitUnExport] 事件：新增实体树节点
type EventEntityTreeAddNode interface {
	OnEntityTreeAddNode(entityTree EntityTree, parent, child ec.Entity)
}

// EventEntityTreeRemoveNode [EmitUnExport] 事件：删除实体树节点
type EventEntityTreeRemoveNode interface {
	OnEntityTreeRemoveNode(entityTree EntityTree, parent, child ec.Entity)
}
