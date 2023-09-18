package golaxy

import (
	"kit.golaxy.org/golaxy/util/iface"
)

type _ServiceOption struct{}

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	CompositeFace iface.Face[Service] // 扩展者，需要扩展服务自身功能时需要使用
}

// ServiceOption 创建服务的选项设置器
type ServiceOption func(o *ServiceOptions)

// Default 默认值
func (_ServiceOption) Default() ServiceOption {
	return func(o *ServiceOptions) {
		_ServiceOption{}.CompositeFace(iface.Face[Service]{})(o)
	}
}

// CompositeFace 扩展者，需要扩展服务自身功能时需要使用
func (_ServiceOption) CompositeFace(face iface.Face[Service]) ServiceOption {
	return func(o *ServiceOptions) {
		o.CompositeFace = face
	}
}
