//go:generate stringer -type TreeNodeState
package ec

// TreeNodeState 实体树节点状态
type TreeNodeState int8

const (
	TreeNodeState_Detached  TreeNodeState = iota // 已从实体树中脱离
	TreeNodeState_Attached                       // 在实体树中
	TreeNodeState_Detaching                      // 正在脱离
)
