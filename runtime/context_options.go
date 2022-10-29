package runtime

import (
	"context"
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/plugin"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	Inheritor          util.Face[Context]                // 继承者，需要拓展运行时上下文自身能力时需要使用
	Context            context.Context                   // 父Context
	AutoRecover        bool                              // 是否开启panic时自动恢复
	ReportError        chan error                        // panic时错误写入的error channel
	PluginLib          plugin.PluginLib                  // 插件库
	StartedCallback    func(runtimeCtx Context)          // 启动运行时回调函数
	StoppingCallback   func(runtimeCtx Context)          // 开始停止运行时回调函数
	StoppedCallback    func(runtimeCtx Context)          // 完全停止运行时回调函数
	FrameBeginCallback func(runtimeCtx Context)          // 帧开始时的回调函数
	FrameEndCallback   func(runtimeCtx Context)          // 帧结束时的回调函数
	FaceCache          *container.Cache[util.FaceAny]    // Face缓存，用于提高性能
	HookCache          *container.Cache[localevent.Hook] // Hook缓存，用于提高性能
}

// WithContextOption 创建运行时上下文的选项设置器
type WithContextOption func(o *ContextOptions)

// ContextOption 创建运行时上下文的选项
var ContextOption = &_ContextOption{}

type _ContextOption struct{}

// Default 默认值
func (*_ContextOption) Default() WithContextOption {
	return func(o *ContextOptions) {
		o.Inheritor = util.Face[Context]{}
		o.Context = nil
		o.AutoRecover = false
		o.ReportError = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，需要拓展运行时上下文自身功能时需要使用
func (*_ContextOption) Inheritor(v util.Face[Context]) WithContextOption {
	return func(o *ContextOptions) {
		o.Inheritor = v
	}
}

// Context 父Context
func (*_ContextOption) Context(v context.Context) WithContextOption {
	return func(o *ContextOptions) {
		o.Context = v
	}
}

// AutoRecover 是否开启panic时自动恢复
func (*_ContextOption) AutoRecover(v bool) WithContextOption {
	return func(o *ContextOptions) {
		o.AutoRecover = v
	}
}

// ReportError panic时错误写入的error channel
func (*_ContextOption) ReportError(v chan error) WithContextOption {
	return func(o *ContextOptions) {
		o.ReportError = v
	}
}

// PluginLib 插件库
func (*_ContextOption) PluginLib(v plugin.PluginLib) WithContextOption {
	return func(o *ContextOptions) {
		o.PluginLib = v
	}
}

// StartedCallback 启动运行时回调函数
func (*_ContextOption) StartedCallback(v func(runtimeCtx Context)) WithContextOption {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (*_ContextOption) StoppingCallback(v func(runtimeCtx Context)) WithContextOption {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (*_ContextOption) StoppedCallback(v func(runtimeCtx Context)) WithContextOption {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}

// FrameBeginCallback 帧更新开始时的回调函数
func (*_ContextOption) FrameBeginCallback(v func(runtimeCtx Context)) WithContextOption {
	return func(o *ContextOptions) {
		o.FrameBeginCallback = v
	}
}

// FrameEndCallback 帧更新结束时的回调函数
func (*_ContextOption) FrameEndCallback(v func(runtimeCtx Context)) WithContextOption {
	return func(o *ContextOptions) {
		o.FrameEndCallback = v
	}
}

// FaceCache Face缓存，用于提高性能
func (*_ContextOption) FaceCache(v *container.Cache[util.FaceAny]) WithContextOption {
	return func(o *ContextOptions) {
		o.FaceCache = v
	}
}

// HookCache Hook缓存，用于提高性能
func (*_ContextOption) HookCache(v *container.Cache[localevent.Hook]) WithContextOption {
	return func(o *ContextOptions) {
		o.HookCache = v
	}
}
