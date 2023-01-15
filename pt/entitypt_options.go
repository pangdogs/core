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
func (WithEntityOption) Default() EntityOption {
	return func(o *EntityOptions) {
		ec.WithEntityOption{}.Default()(&o.EntityOptions)
		WithEntityOption{}.GenCompID(nil)(o)
		WithEntityOption{}.Accessibility(runtime.Local)(o)
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
