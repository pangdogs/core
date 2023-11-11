package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

// EntityPtResolver 实体原型解析器
type EntityPtResolver interface {
	// ResolveEntityPt 解析实体原型
	ResolveEntityPt(prototype string) (EntityPt, bool)
}

// Using 使用实体原型
//
//	@param entityPtResolver 实体原型解析器。
//	@param prototype 实体原型名称。
func Using(entityPtResolver EntityPtResolver, prototype string) (EntityPt, error) {
	if entityPtResolver == nil {
		return EntityPt{}, fmt.Errorf("%w: %w: entityPtResolver is nil", ErrPt, exception.ErrArgs)
	}

	entityPt, ok := entityPtResolver.ResolveEntityPt(prototype)
	if !ok {
		return EntityPt{}, fmt.Errorf("%w: entity %q not registered", ErrPt, prototype)
	}

	return entityPt, nil
}
