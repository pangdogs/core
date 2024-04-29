//go:generate stringer -type TreeNodeState
package ec

// TreeNodeState 实体树节点状态
type TreeNodeState int8

const (
	TreeNodeState_Freedom   TreeNodeState = iota // 自由实体
	TreeNodeState_Attaching                      // 正在加入实体树
	TreeNodeState_Attached                       // 在实体树中
	TreeNodeState_Detaching                      // 正在脱离实体树
)
