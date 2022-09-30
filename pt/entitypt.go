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
	entity := ec.NewEntity(append(optSetter, ec.EntityOption.Prototype(pt.Prototype))...)

	for i := range pt.compPts {
		entity.AddComponent(pt.compPts[i].Interface, pt.compPts[i].New())
	}

	return entity
}
