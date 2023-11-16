package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

// EntityPtProvider 实体原型提供者
type EntityPtProvider interface {
	// GetEntityPt 获取实体原型
	GetEntityPt(prototype string) (EntityPt, bool)
}

// Using 使用实体原型
//
//	@param entityPtProvider 实体原型提供者。
//	@param prototype 实体原型名称。
func Using(entityPtProvider EntityPtProvider, prototype string) (EntityPt, error) {
	if entityPtProvider == nil {
		return EntityPt{}, fmt.Errorf("%w: %w: entityPtProvider is nil", ErrPt, exception.ErrArgs)
	}

	entityPt, ok := entityPtProvider.GetEntityPt(prototype)
	if !ok {
		return EntityPt{}, fmt.Errorf("%w: entity %q not registered", ErrPt, prototype)
	}

	return entityPt, nil
}
