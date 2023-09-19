package pt

import (
	"fmt"
)

// EntityPtResolver 实体原型解析器
type EntityPtResolver interface {
	// ResolveEntityPt 解析实体原型
	ResolveEntityPt(prototype string) (EntityPt, bool)
}

// Fetch 获取实体原型。
//
//	@param entityPtResolver 实体原型解析器。
//	@param prototype 实体原型名称。
func Fetch(entityPtResolver EntityPtResolver, prototype string) EntityPt {
	entityPt, err := Access(entityPtResolver, prototype)
	if err != nil {
		panic(err)
	}
	return entityPt
}

// Access 尝试获取实体原型
//
//	@param entityPtResolver 实体原型解析器。
//	@param prototype 实体原型名称。
func Access(entityPtResolver EntityPtResolver, prototype string) (EntityPt, error) {
	if entityPtResolver == nil {
		return EntityPt{}, fmt.Errorf("%w: entityPtResolver is nil", ErrPt)
	}

	entityPt, ok := entityPtResolver.ResolveEntityPt(prototype)
	if !ok {
		return EntityPt{}, fmt.Errorf("%w: entity %q not registered", ErrPt, prototype)
	}

	return entityPt, nil
}
