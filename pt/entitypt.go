package pt

import (
	"github.com/golaxy-kit/golaxy/ec"
)

// EntityPt 实体原型
type EntityPt struct {
	Prototype string // 实体原型名称
	compPts   []ComponentPt
}

// New 创建实体
func (pt *EntityPt) New(options ...ec.EntityOption) ec.Entity {
	opts := ec.EntityOptions{}
	ec.WithEntityOption.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return pt.UnsafeNew(opts)
}

// UnsafeNew 不安全的创建实体，需要自己初始化所有选项
func (pt *EntityPt) UnsafeNew(options ec.EntityOptions) ec.Entity {
	options.Prototype = pt.Prototype
	return pt.AddTo(ec.UnsafeNewEntity(options))
}

// AddTo 向实体安装组件
func (pt *EntityPt) AddTo(entity ec.Entity) ec.Entity {
	if entity == nil {
		return nil
	}

	for i := range pt.compPts {
		entity.AddComponent(pt.compPts[i].Name, pt.compPts[i].New())
	}

	return entity
}
