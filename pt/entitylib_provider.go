package pt

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

// EntityPTProvider 实体原型提供者
type EntityPTProvider interface {
	// GetEntityLib 获取实体原型库
	GetEntityLib() EntityLib
}

// For 查询实体原型
func For(provider EntityPTProvider, prototype string) EntityPT {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPt, exception.ErrArgs))
	}

	entity, ok := provider.GetEntityLib().Get(prototype)
	if !ok {
		panic(fmt.Errorf("%w: entity %q was not declared", ErrPt, prototype))
	}

	return entity
}
