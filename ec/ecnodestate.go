//go:generate stringer -type ECNodeState
package ec

// ECNodeState EC节点状态
type ECNodeState int8

const (
	ECNodeState_Detached  ECNodeState = iota // 已从EC树中脱离
	ECNodeState_Attached                     // 在EC树中
	ECNodeState_Detaching                    // 正在脱离
)
