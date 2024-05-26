//go:generate stringer -type EntityState
package ec

// EntityState 实体状态
type EntityState int8

const (
	EntityState_Birth EntityState = iota // 出生
	EntityState_Enter                    // 进入容器
	EntityState_Awake                    // 唤醒
	EntityState_Start                    // 开始
	EntityState_Alive                    // 活跃
	EntityState_Leave                    // 离开容器
	EntityState_Shut                     // 结束
	EntityState_Death                    // 死亡
)
