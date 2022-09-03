package core

// ServiceContextOptions ...
type ServiceContextOptions struct {
	Inheritor   Face[ServiceContext]
	Prototype   string
	NodeID      int64
	ReportError chan error
	StartedCallback,
	StoppingCallback,
	StoppedCallback func(serv Service)
}

// ServiceContextOptionSetter ...
var ServiceContextOptionSetter = &_ServiceContextOptionSetter{}

type _ServiceContextOptionSetterFunc func(o *ServiceContextOptions)

type _ServiceContextOptionSetter struct{}

// Default ...
func (*_ServiceContextOptionSetter) Default() _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = Face[ServiceContext]{}
		o.Prototype = ""
		o.ReportError = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
	}
}

// Inheritor ...
func (*_ServiceContextOptionSetter) Inheritor(v Face[ServiceContext]) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = v
	}
}

// NodeID 服务分布式节点ID
func (*_ServiceContextOptionSetter) NodeID(v int64) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.NodeID = v
	}
}

// Prototype 服务（Service）原型
func (*_ServiceContextOptionSetter) Prototype(v string) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Prototype = v
	}
}

// ReportError ...
func (*_ServiceContextOptionSetter) ReportError(v chan error) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.ReportError = v
	}
}

// StartedCallback ...
func (*_ServiceContextOptionSetter) StartedCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback ...
func (*_ServiceContextOptionSetter) StoppingCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback ...
func (*_ServiceContextOptionSetter) StoppedCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StoppedCallback = v
	}
}
