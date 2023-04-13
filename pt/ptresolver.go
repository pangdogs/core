package pt

import (
	"fmt"
)

// EntityPtResolver 实体原型解析器
type EntityPtResolver interface {
	// ResolveEntityPt 解析实体原型
	ResolveEntityPt(prototype string) (EntityPt, bool)
}

// GetEntityPt 获取实体原型。
//
//	@param entityPtResolver 实体原型解析器。
//	@param prototype 实体原型名称。
func GetEntityPt(entityPtResolver EntityPtResolver, prototype string) EntityPt {
	if entityPtResolver == nil {
		panic("nil entityPtResolver")
	}

	entityPt, ok := entityPtResolver.ResolveEntityPt(prototype)
	if !ok {
		panic(fmt.Errorf("entity prototype '%s' not registered", prototype))
	}

	return entityPt
}

// TryGetEntityPt 尝试获取实体原型
//
//	@param entityPtResolver 实体原型解析器。
//	@param prototype 实体原型名称。
func TryGetEntityPt(entityPtResolver EntityPtResolver, prototype string) (EntityPt, bool) {
	if entityPtResolver == nil {
		return EntityPt{}, false
	}

	return entityPtResolver.ResolveEntityPt(prototype)
}
