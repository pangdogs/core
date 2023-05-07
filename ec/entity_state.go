//go:generate stringer -type EntityState
package ec

// EntityState 实体状态
type EntityState int8

const (
	EntityState_Birth  EntityState = iota // 出生
	EntityState_Entry                     // 进入容器
	EntityState_Init                      // 初始化
	EntityState_Inited                    // 已初始化
	EntityState_Living                    // 活跃
	EntityState_Leave                     // 离开容器
	EntityState_Shut                      // 结束
	EntityState_Death                     // 死亡
)
