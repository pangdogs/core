//go:generate stringer -type RunningState
package runtime

// RunningState 运行状态
type RunningState int32

const (
	RunningState_Birth            RunningState = iota // 出生
	RunningState_Starting                             // 开始启动
	RunningState_Started                              // 已启动
	RunningState_FrameLoopBegin                       // 帧循环开始
	RunningState_FrameUpdateBegin                     // 帧更新开始
	RunningState_FrameUpdateEnd                       // 帧更新结束
	RunningState_FrameLoopEnd                         // 帧循环结束
	RunningState_RunCallBegin                         // Call开始执行
	RunningState_RunCallEnd                           // Call结束执行
	RunningState_RunGCBegin                           // GC开始执行
	RunningState_RunGCEnd                             // GC结束执行
	RunningState_Terminating                          // 开始停止
	RunningState_Terminated                           // 已停止
)
