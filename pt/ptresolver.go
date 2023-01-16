package pt

// PtResolver 用于从服务上下文中获取实体原型
type PtResolver interface {
	// GetEntityPt 获取实体原型
	GetEntityPt(prototype string) (EntityPt, bool)
}
