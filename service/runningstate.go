//go:generate stringer -type RunningState
package service

// RunningState 运行状态
type RunningState int32

const (
	RunningState_Birth       RunningState = iota // 出生
	RunningState_Starting                        // 开始启动
	RunningState_Started                         // 已启动
	RunningState_Terminating                     // 开始停止
	RunningState_Terminated                      // 已停止
)
