package pt

import (
	"github.com/golaxy-kit/golaxy/ec"
)

// EntityPt 实体原型
type EntityPt struct {
	Prototype string // 实体原型名称
	compPts   []ComponentPt
}

// Construct 创建实体
func (pt *EntityPt) Construct(options ...EntityOption) ec.Entity {
	opts := EntityOptions{}
	WithEntityOption{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return pt.UnsafeConstruct(opts)
}

// UnsafeConstruct 不安全的创建实体，需要自己初始化所有选项
func (pt *EntityPt) UnsafeConstruct(options EntityOptions) ec.Entity {
	options.Prototype = pt.Prototype
	return pt.Assemble(ec.UnsafeNewEntity(options.EntityOptions), options.AssignCompID)
}

// Assemble 向实体安装组件
func (pt *EntityPt) Assemble(entity ec.Entity, assignCompID func(entity ec.Entity, compPt ComponentPt) ec.ID) ec.Entity {
	if entity == nil {
		return nil
	}

	for i := range pt.compPts {
		var id ec.ID

		if assignCompID != nil {
			id = assignCompID(entity, pt.compPts[i])
		}

		entity.AddComponent(pt.compPts[i].Name, pt.compPts[i].Construct(id))
	}

	return entity
}
