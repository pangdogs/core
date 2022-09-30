package galaxy

import (
	"github.com/pangdogs/galaxy/util"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	Inheritor         util.Face[Service] // 继承者，需要拓展服务自身功能时需要使用
	EnableAutoRecover bool               // 是否开启panic时自动恢复
}

// ServiceOption 创建服务的选项
var ServiceOption = &_ServiceOptionSetter{}

// ServiceOptionSetterFunc 创建服务的选项设置器
type ServiceOptionSetterFunc func(o *ServiceOptions)

type _ServiceOptionSetter struct{}

// Default 默认值
func (*_ServiceOptionSetter) Default() ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = util.Face[Service]{}
		o.EnableAutoRecover = false
	}
}

// Inheritor 继承者，需要拓展服务自身功能时需要使用
func (*_ServiceOptionSetter) Inheritor(v util.Face[Service]) ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRecover 是否开启panic时自动恢复
func (*_ServiceOptionSetter) EnableAutoRecover(v bool) ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.EnableAutoRecover = v
	}
}
