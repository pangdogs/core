//go:generate stringer -type EntityState
package ec

// EntityState 实体状态
type EntityState int8

const (
	EntityState_Birth  EntityState = iota // 出生
	EntityState_Entry                     // 进入
	EntityState_Init                      // 初始化
	EntityState_Start                     // 开始
	EntityState_Living                    // 活跃
	EntityState_Leave                     // 离开
	EntityState_Shut                      // 结束
	EntityState_Death                     // 死亡
)
