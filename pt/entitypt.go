package pt

import (
	"github.com/pangdogs/galaxy/ec"
)

// EntityPt 实体原型
type EntityPt struct {
	Prototype string // 实体原型名称
	compPts   []ComponentPt
}

// New 创建实体
func (pt *EntityPt) New(optSetter ...ec.EntityOptionSetter) ec.Entity {
	opts := ec.EntityOptions{}
	ec.EntityOption.Default()(&opts)

	for i := range optSetter {
		optSetter[i](&opts)
	}

	opts.Prototype = pt.Prototype

	entity := ec.UnsafeNewEntity(opts)

	for i := range pt.compPts {
		entity.AddComponent(pt.compPts[i].Interface, pt.compPts[i].New())
	}

	return entity
}

// InstallComponents 向实体安装组件
func (pt *EntityPt) InstallComponents(entity ec.Entity) {
	if entity == nil {
		return
	}

	for i := range pt.compPts {
		entity.AddComponent(pt.compPts[i].Interface, pt.compPts[i].New())
	}
}
