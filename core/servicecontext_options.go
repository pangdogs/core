package core

// ServiceContextOptions 创建服务上下文（Service Context）的所有选项
type ServiceContextOptions struct {
	Inheritor        Face[ServiceContext] // 继承者，需要拓展服务上下文（Service Context）自身功能时需要使用
	Prototype        string               // 服务（Service）原型
	NodeID           int64                // 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
	ReportError      chan error           // panic时错误写入的error channel
	StartedCallback  func(serv Service)   // 启动运行时回调函数
	StoppingCallback func(serv Service)   // 开始停止运行时回调函数
	StoppedCallback  func(serv Service)   // 完全停止运行时回调函数
}

// ServiceContextOptionSetter 服务上下文（Service Context）选项设置器
var ServiceContextOptionSetter = &_ServiceContextOptionSetter{}

type _ServiceContextOptionSetterFunc func(o *ServiceContextOptions)

type _ServiceContextOptionSetter struct{}

// Default 默认值
func (*_ServiceContextOptionSetter) Default() _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = Face[ServiceContext]{}
		o.Prototype = ""
		o.NodeID = 0
		o.ReportError = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
	}
}

// Inheritor 继承者，需要拓展服务上下文（Service Context）自身功能时需要使用
func (*_ServiceContextOptionSetter) Inheritor(v Face[ServiceContext]) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = v
	}
}

// Prototype 服务（Service）原型
func (*_ServiceContextOptionSetter) Prototype(v string) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.Prototype = v
	}
}

// NodeID 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
func (*_ServiceContextOptionSetter) NodeID(v int64) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.NodeID = v
	}
}

// ReportError panic时错误写入的error channel
func (*_ServiceContextOptionSetter) ReportError(v chan error) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.ReportError = v
	}
}

// StartedCallback 启动运行时回调函数
func (*_ServiceContextOptionSetter) StartedCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (*_ServiceContextOptionSetter) StoppingCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (*_ServiceContextOptionSetter) StoppedCallback(v func(serv Service)) _ServiceContextOptionSetterFunc {
	return func(o *ServiceContextOptions) {
		o.StoppedCallback = v
	}
}
