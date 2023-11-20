package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
)

// EntityPTProvider 实体原型提供者
type EntityPTProvider interface {
	// GetEntityLib 获取实体原型库
	GetEntityLib() EntityLib
}

// Using 使用实体原型
func Using(entityPTProvider EntityPTProvider, prototype string) EntityPT {
	if entityPTProvider == nil {
		panic(fmt.Errorf("%w: %w: entityPTProvider is nil", ErrPt, exception.ErrArgs))
	}

	entity, ok := entityPTProvider.GetEntityLib().Get(prototype)
	if !ok {
		panic(fmt.Errorf("%w: entity %q not registered", ErrPt, prototype))
	}

	return entity
}
