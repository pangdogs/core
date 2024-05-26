//go:generate stringer -type ComponentState
package ec

// ComponentState 组件状态
type ComponentState int8

const (
	ComponentState_Birth  ComponentState = iota // 出生
	ComponentState_Attach                       // 附着
	ComponentState_Awake                        // 唤醒
	ComponentState_Start                        // 开始
	ComponentState_Alive                        // 活跃
	ComponentState_Detach                       // 脱离
	ComponentState_Shut                         // 结束
	ComponentState_Death                        // 死亡
)
