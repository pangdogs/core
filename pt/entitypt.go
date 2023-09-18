package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
)

// EntityPt 实体原型
type EntityPt struct {
	Prototype string        // 实体原型名称
	compPts   []ComponentPt // 组件原型列表
}

// Construct 创建实体
func (pt *EntityPt) Construct(options ...ConstructEntityOption) ec.Entity {
	opts := ConstructEntityOptions{}
	Option{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return pt.UnsafeConstruct(opts)
}

// Deprecated: UnsafeConstruct 内部创建实体
func (pt *EntityPt) UnsafeConstruct(options ConstructEntityOptions) ec.Entity {
	options.Prototype = pt.Prototype
	return pt.Assemble(ec.UnsafeNewEntity(options.EntityOptions), options.ComponentConstructor, options.EntityConstructor)
}

// Assemble 向实体安装组件
func (pt *EntityPt) Assemble(entity ec.Entity, componentConstructor ComponentConstructor, entityConstructor EntityConstructor) ec.Entity {
	if entity == nil {
		return nil
	}

	for i := range pt.compPts {
		compPt := pt.compPts[i]

		comp := compPt.Construct()

		if err := entity.AddComponent(compPt.Name, comp); err != nil {
			panic(fmt.Errorf("%w: %w", ErrPt, err))
		}

		if componentConstructor != nil {
			componentConstructor(entity, comp)
		}
	}

	if entityConstructor != nil {
		entityConstructor(entity)
	}

	return entity
}
