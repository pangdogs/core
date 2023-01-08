package galaxy

import (
	"github.com/golaxy-kit/golaxy/util"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	Inheritor util.Face[Service] // 继承者，需要扩展服务自身功能时需要使用
}

// ServiceOption 创建服务的选项设置器
type ServiceOption func(o *ServiceOptions)

// WithServiceOption 创建服务的选项
var WithServiceOption = _WithServiceOption{}

type _WithServiceOption struct{}

// Default 默认值
func (_WithServiceOption) Default() ServiceOption {
	return func(o *ServiceOptions) {
		o.Inheritor = util.Face[Service]{}
	}
}

// Inheritor 继承者，需要扩展服务自身功能时需要使用
func (_WithServiceOption) Inheritor(v util.Face[Service]) ServiceOption {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}
