package core

// ServiceOptions 创建服务（Service）的所有选项
type ServiceOptions struct {
	Inheritor         Face[Service]
	EnableAutoRecover bool
}

// ServiceOptionSetter ...
var ServiceOptionSetter = &_ServiceOptionSetter{}

type _ServiceOptionSetterFunc func(o *ServiceOptions)

type _ServiceOptionSetter struct{}

// Default ...
func (*_ServiceOptionSetter) Default() _ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = Face[Service]{}
		o.EnableAutoRecover = false
	}
}

// Inheritor ...
func (*_ServiceOptionSetter) Inheritor(v Face[Service]) _ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRecover ...
func (*_ServiceOptionSetter) EnableAutoRecover(v bool) _ServiceOptionSetterFunc {
	return func(o *ServiceOptions) {
		o.EnableAutoRecover = v
	}
}
