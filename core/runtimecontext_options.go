package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// RuntimeContextOptions 创建运行时上下文（Runtime Context）的所有选项
type RuntimeContextOptions struct {
	Inheritor       Face[RuntimeContext]      // 继承者，需要拓展运行时上下文自身功能时需要使用
	ReportError     chan error                // panic时错误写入的error channel
	StartedCallback func(runtime Runtime)     // 启动运行时回调函数
	StoppedCallback func(runtime Runtime)     // 停止运行时回调函数
	FaceCache       *container.Cache[FaceAny] // Face缓存，用于提高性能
	HookCache       *container.Cache[Hook]    // Hook缓存，用于提高性能
}

// RuntimeContextOptionSetter 运行时上下文（Runtime Context）选项设置器
var RuntimeContextOptionSetter = &_RuntimeContextOptionSetter{}

// RuntimeContextOptionSetterFunc 运行时上下文（Runtime Context）选项设置器函数
type RuntimeContextOptionSetterFunc func(o *RuntimeContextOptions)

type _RuntimeContextOptionSetter struct{}

// Default 默认值
func (*_RuntimeContextOptionSetter) Default() RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.Inheritor = Face[RuntimeContext]{}
		o.ReportError = nil
		o.StartedCallback = nil
		o.StoppedCallback = nil
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，需要拓展运行时上下文自身功能时需要使用
func (*_RuntimeContextOptionSetter) Inheritor(v Face[RuntimeContext]) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.Inheritor = v
	}
}

// ReportError panic时错误写入的error channel
func (*_RuntimeContextOptionSetter) ReportError(v chan error) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.ReportError = v
	}
}

// StartFunc 启动运行时回调函数
func (*_RuntimeContextOptionSetter) StartFunc(v func(rt Runtime)) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.StartedCallback = v
	}
}

// StopFunc 停止运行时回调函数
func (*_RuntimeContextOptionSetter) StopFunc(v func(rt Runtime)) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.StoppedCallback = v
	}
}

// FaceCache Face缓存，用于提高性能
func (*_RuntimeContextOptionSetter) FaceCache(v *container.Cache[FaceAny]) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.FaceCache = v
	}
}

// HookCache Hook缓存，用于提高性能
func (*_RuntimeContextOptionSetter) HookCache(v *container.Cache[Hook]) RuntimeContextOptionSetterFunc {
	return func(o *RuntimeContextOptions) {
		o.HookCache = v
	}
}
