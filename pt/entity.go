package pt

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/option"
)

// CompInfo 组件信息
type CompInfo struct {
	PT    ComponentPT // 原型
	Alias string      // 别名
}

// EntityPT 实体原型
type EntityPT struct {
	Prototype string     // 实体原型名称
	CompInfos []CompInfo // 组件信息
}

// Construct 创建实体
func (pt EntityPT) Construct(settings ...option.Setting[ec.EntityOptions]) ec.Entity {
	return pt.UnsafeConstruct(option.Make(ec.With.Default(), settings...))
}

// Deprecated: UnsafeConstruct 内部创建实体
func (pt EntityPT) UnsafeConstruct(options ec.EntityOptions) ec.Entity {
	options.Prototype = pt.Prototype
	return pt.Assemble(ec.UnsafeNewEntity(options))
}

// Assemble 向实体安装组件
func (pt EntityPT) Assemble(entity ec.Entity) ec.Entity {
	if entity == nil {
		panic(fmt.Errorf("%w: %w: entity is nil", ErrPt, exception.ErrArgs))
	}

	for i := range pt.CompInfos {
		compInfo := &pt.CompInfos[i]

		comp := compInfo.PT.Construct()

		if err := entity.AddComponent(compInfo.Alias, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}
	}

	return entity
}
