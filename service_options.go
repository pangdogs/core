package golaxy

import (
	"kit.golaxy.org/golaxy/util"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	Inheritor util.Face[Service] // 继承者，需要扩展服务自身功能时需要使用
}

// ServiceOption 创建服务的选项设置器
type ServiceOption func(o *ServiceOptions)

// WithServiceOption 创建服务的所有选项设置器
type WithServiceOption struct{}

// Default 默认值
func (WithServiceOption) Default() ServiceOption {
	return func(o *ServiceOptions) {
		WithServiceOption{}.Inheritor(util.Face[Service]{})(o)
	}
}

// Inheritor 继承者，需要扩展服务自身功能时需要使用
func (WithServiceOption) Inheritor(v util.Face[Service]) ServiceOption {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}
