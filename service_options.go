package golaxy

import (
	"kit.golaxy.org/golaxy/util"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	CompositeFace util.Face[Service] // 扩展者，需要扩展服务自身功能时需要使用
}

// ServiceOption 创建服务的选项设置器
type ServiceOption func(o *ServiceOptions)

// ServiceDefault 默认值
func (Option) ServiceDefault() ServiceOption {
	return func(o *ServiceOptions) {
		Option{}.ServiceCompositeFace(util.Face[Service]{})(o)
	}
}

// ServiceCompositeFace 扩展者，需要扩展服务自身功能时需要使用
func (Option) ServiceCompositeFace(face util.Face[Service]) ServiceOption {
	return func(o *ServiceOptions) {
		o.CompositeFace = face
	}
}
