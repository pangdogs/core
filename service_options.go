package galaxy

import (
	"github.com/galaxy-kit/galaxy/util"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	Inheritor util.Face[Service] // 继承者，需要拓展服务自身功能时需要使用
}

// WithServiceOption 创建服务的选项设置器
type WithServiceOption func(o *ServiceOptions)

// ServiceOption 创建服务的选项
var ServiceOption = _ServiceOption{}

type _ServiceOption struct{}

// Default 默认值
func (_ServiceOption) Default() WithServiceOption {
	return func(o *ServiceOptions) {
		o.Inheritor = util.Face[Service]{}
	}
}

// Inheritor 继承者，需要拓展服务自身功能时需要使用
func (_ServiceOption) Inheritor(v util.Face[Service]) WithServiceOption {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}
