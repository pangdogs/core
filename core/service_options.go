package core

// ServiceOptions 创建服务（Service）的所有选项
type ServiceOptions struct {
	Inheritor         Face[Service] // 继承者，需要拓展服务（Service）自身功能时需要使用
	EnableAutoRecover bool          // 是否开启panic时自动恢复
}

// ServiceOptionSetter 服务（Service）选项设置器
var ServiceOptionSetter = &_ServiceOptionSetter{}

// ServiceOptionSetterFunc 服务（Service）选项设置器函数
type ServiceOptionSetterFunc func(o *ServiceOptions)

type _ServiceOptionSetter struct{}

// Default 默认值
func (*_ServiceOptionSetter) Default() ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = Face[Service]{}
		o.EnableAutoRecover = false
	}
}

// Inheritor 继承者，需要拓展服务（Service）自身功能时需要使用
func (*_ServiceOptionSetter) Inheritor(v Face[Service]) ServiceOptionSetterFunc {
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
