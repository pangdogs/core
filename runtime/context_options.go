package runtime

import (
	"context"
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	Inheritor         util.Face[Context]                // 继承者，需要拓展运行时上下文自身能力时需要使用
	EnableAutoRecover bool                              // 是否开启panic时自动恢复
	ReportError       chan error                        // panic时错误写入的error channel
	ParentContext     context.Context                   // 父Context
	StartedCallback   func(runtimeCtx Context)          // 启动运行时回调函数
	StoppingCallback  func(runtimeCtx Context)          // 开始停止运行时回调函数
	StoppedCallback   func(runtimeCtx Context)          // 完全停止运行时回调函数
	FaceCache         *container.Cache[util.FaceAny]    // Face缓存，用于提高性能
	HookCache         *container.Cache[localevent.Hook] // Hook缓存，用于提高性能
}

// ContextOption 创建运行时上下文的选项
var ContextOption = &_ContextOptionSetter{}

// ContextOptionSetter 创建运行时上下文的选项设置器
type ContextOptionSetter func(o *ContextOptions)

type _ContextOptionSetter struct{}

// Default 默认值
func (*_ContextOptionSetter) Default() ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = util.Face[Context]{}
		o.EnableAutoRecover = false
		o.ReportError = nil
		o.ParentContext = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，需要拓展运行时上下文自身功能时需要使用
func (*_ContextOptionSetter) Inheritor(v util.Face[Context]) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRecover 是否开启panic时自动恢复
func (*_ContextOptionSetter) EnableAutoRecover(v bool) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.EnableAutoRecover = v
	}
}

// ReportError panic时错误写入的error channel
func (*_ContextOptionSetter) ReportError(v chan error) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.ReportError = v
	}
}

// ParentContext 父Context
func (*_ContextOptionSetter) ParentContext(v context.Context) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.ParentContext = v
	}
}

// StartedCallback 启动运行时回调函数
func (*_ContextOptionSetter) StartedCallback(v func(runtimeCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (*_ContextOptionSetter) StoppingCallback(v func(runtimeCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (*_ContextOptionSetter) StoppedCallback(v func(runtimeCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}

// FaceCache Face缓存，用于提高性能
func (*_ContextOptionSetter) FaceCache(v *container.Cache[util.FaceAny]) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.FaceCache = v
	}
}

// HookCache Hook缓存，用于提高性能
func (*_ContextOptionSetter) HookCache(v *container.Cache[localevent.Hook]) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.HookCache = v
	}
}
