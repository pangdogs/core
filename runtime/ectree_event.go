//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_event --package=$GOPACKAGE
package runtime

import "git.golaxy.org/core/ec"

// EventECTreeAddChild [EmitUnExport] 事件：EC树中子实体加入父实体
type EventECTreeAddChild interface {
	OnAddChild(ecTree ECTree, parent, child ec.Entity)
}

// EventECTreeRemoveChild [EmitUnExport] 事件：EC树中子实体离开父实体
type EventECTreeRemoveChild interface {
	OnRemoveChild(ecTree ECTree, parent, child ec.Entity)
}
