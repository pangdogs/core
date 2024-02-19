package core

import (
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	CompositeFace iface.Face[Service] // 扩展者，在扩展服务自身功能时使用
}

type _ServiceOption struct{}

// Default 默认值
func (_ServiceOption) Default() option.Setting[ServiceOptions] {
	return func(o *ServiceOptions) {
		With.Service.CompositeFace(iface.Face[Service]{})(o)
	}
}

// CompositeFace 扩展者，在扩展服务自身功能时使用
func (_ServiceOption) CompositeFace(face iface.Face[Service]) option.Setting[ServiceOptions] {
	return func(o *ServiceOptions) {
		o.CompositeFace = face
	}
}
