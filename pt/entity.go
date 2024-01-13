package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/option"
)

// EntityPT 实体原型
type EntityPT struct {
	Prototype string // 实体原型名称
	comps     []ComponentPT
}

// Construct 创建实体
func (pt EntityPT) Construct(settings ...option.Setting[ConstructEntityOptions]) ec.Entity {
	return pt.UnsafeConstruct(option.Make(Option{}.Default(), settings...))
}

// Deprecated: UnsafeConstruct 内部创建实体
func (pt EntityPT) UnsafeConstruct(options ConstructEntityOptions) ec.Entity {
	options.Prototype = pt.Prototype
	return pt.Assemble(ec.UnsafeNewEntity(options.EntityOptions), options.ComponentCtor, options.EntityCtor)
}

// Assemble 向实体安装组件
func (pt EntityPT) Assemble(entity ec.Entity, componentCtor ComponentCtor, entityCtor EntityCtor) ec.Entity {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}

	for i := range pt.comps {
		compPT := pt.comps[i]

		comp := compPT.Construct()

		if err := entity.AddComponent(compPT.Name, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}

		componentCtor.Exec(nil, entity, comp)
	}

	entityCtor.Exec(nil, entity)

	return entity
}
