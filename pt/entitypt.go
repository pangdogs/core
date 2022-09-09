package pt

import (
	"github.com/pangdogs/galaxy/core"
)

// EntityPt 实体原型
type EntityPt struct {
	Prototype string
	compPts   []ComponentPt
}

// New 创建实体
func (pt *EntityPt) New(optSetterFuncs ...core.EntityOptionSetterFunc) core.Entity {
	opts := core.EntityOptions{}
	core.EntityOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return pt.NewWithOpts(opts)
}

// NewWithOpts 创建实体并传入参数
func (pt *EntityPt) NewWithOpts(opts core.EntityOptions) core.Entity {
	opts.Prototype = pt.Prototype

	entity := core.NewEntityWithOpts(opts)
	for i := range pt.compPts {
		entity.AddComponent(pt.compPts[i].Api, pt.compPts[i].New())
	}

	return entity
}
