package golaxy

import (
	"kit.golaxy.org/golaxy/util"
)

type _ServiceOption struct{}

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	CompositeFace util.Face[Service] // 扩展者，需要扩展服务自身功能时需要使用
}

// ServiceOption 创建服务的选项设置器
type ServiceOption func(o *ServiceOptions)

// Default 默认值
func (_ServiceOption) Default() ServiceOption {
	return func(o *ServiceOptions) {
		_ServiceOption{}.CompositeFace(util.Face[Service]{})(o)
	}
}

// CompositeFace 扩展者，需要扩展服务自身功能时需要使用
func (_ServiceOption) CompositeFace(face util.Face[Service]) ServiceOption {
	return func(o *ServiceOptions) {
		o.CompositeFace = face
	}
}
