package pt

import (
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/runtime"
)

// EntityOptions 创建实体的所有选项
type EntityOptions struct {
	ec.EntityOptions
	GenCompID     func(entity ec.Entity, compPt ComponentPt) ec.ID // 生成组件ID函数
	Accessibility runtime.Accessibility                            // 实体的可访问性
}

// EntityOption 创建实体的选项设置器
type EntityOption func(o *EntityOptions)

// WithEntityOption 创建实体的所有选项设置器
type WithEntityOption struct {
	ec.WithEntityOption
}

// Default 默认值
func (w WithEntityOption) Default() EntityOption {
	return func(o *EntityOptions) {
		w.WithEntityOption.Default()(&o.EntityOptions)
		o.GenCompID = nil
		o.Accessibility = runtime.Local
	}
}

// GenCompID 生成组件ID函数
func (WithEntityOption) GenCompID(v func(entity ec.Entity, compPt ComponentPt) ec.ID) EntityOption {
	return func(o *EntityOptions) {
		o.GenCompID = v
	}
}

// Accessibility 实体的可访问性
func (WithEntityOption) Accessibility(v runtime.Accessibility) EntityOption {
	return func(o *EntityOptions) {
		o.Accessibility = v
	}
}

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
	return pt.Assemble(ec.UnsafeNewEntity(options.EntityOptions), options.GenCompID)
}

// Assemble 向实体安装组件
func (pt *EntityPt) Assemble(entity ec.Entity, GenCompID func(entity ec.Entity, compPt ComponentPt) ec.ID) ec.Entity {
	if entity == nil {
		return nil
	}

	for i := range pt.compPts {
		var id ec.ID

		if GenCompID != nil {
			id = GenCompID(entity, pt.compPts[i])
		}

		entity.AddComponent(pt.compPts[i].Name, pt.compPts[i].Construct(id))
	}

	return entity
}
