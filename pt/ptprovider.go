package pt

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
)

// EntityPTProvider 实体原型提供者
type EntityPTProvider interface {
	// GetEntityLib 获取实体原型库
	GetEntityLib() EntityLib
}

// For 使用实体原型
func For(entityPTProvider EntityPTProvider, prototype string) EntityPT {
	if entityPTProvider == nil {
		panic(fmt.Errorf("%w: %w: entityPTProvider is nil", ErrPt, exception.ErrArgs))
	}

	entity, ok := entityPTProvider.GetEntityLib().Get(prototype)
	if !ok {
		panic(fmt.Errorf("%w: entity %q was not declared", ErrPt, prototype))
	}

	return entity
}
