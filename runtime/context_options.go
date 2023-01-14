package runtime

import (
	"context"
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/plugin"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/golaxy-kit/golaxy/util/container"
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	Inheritor          util.Face[Context]                // 继承者，需要扩展运行时上下文自身能力时需要使用
	Context            context.Context                   // 父Context
	AutoRecover        bool                              // 是否开启panic时自动恢复
	ReportError        chan error                        // panic时错误写入的error channel
	Name               string                            // 运行时名称
	PluginBundle       plugin.PluginBundle               // 插件包
	StartedCallback    func(runtimeCtx Context)          // 启动运行时回调函数
	StoppingCallback   func(runtimeCtx Context)          // 开始停止运行时回调函数
	StoppedCallback    func(runtimeCtx Context)          // 完全停止运行时回调函数
	FrameBeginCallback func(runtimeCtx Context)          // 帧开始时的回调函数
	FrameEndCallback   func(runtimeCtx Context)          // 帧结束时的回调函数
	FaceCache          *container.Cache[util.FaceAny]    // Face缓存，用于提高性能
	HookCache          *container.Cache[localevent.Hook] // Hook缓存，用于提高性能
}

// ContextOption 创建运行时上下文的选项设置器
type ContextOption func(o *ContextOptions)

// WithContextOption 创建运行时上下文的所有选项设置器
type WithContextOption struct{}

// Default 默认值
func (WithContextOption) Default() ContextOption {
	return func(o *ContextOptions) {
		o.Inheritor = util.Face[Context]{}
		o.Context = nil
		o.AutoRecover = false
		o.ReportError = nil
		o.Name = ""
		o.PluginBundle = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
		o.FrameBeginCallback = nil
		o.FrameEndCallback = nil
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor 继承者，需要扩展运行时上下文自身功能时需要使用
func (WithContextOption) Inheritor(v util.Face[Context]) ContextOption {
	return func(o *ContextOptions) {
		o.Inheritor = v
	}
}

// Context 父Context
func (WithContextOption) Context(v context.Context) ContextOption {
	return func(o *ContextOptions) {
		o.Context = v
	}
}

// AutoRecover 是否开启panic时自动恢复
func (WithContextOption) AutoRecover(v bool) ContextOption {
	return func(o *ContextOptions) {
		o.AutoRecover = v
	}
}

// ReportError panic时错误写入的error channel
func (WithContextOption) ReportError(v chan error) ContextOption {
	return func(o *ContextOptions) {
		o.ReportError = v
	}
}

// Name 服务名称
func (WithContextOption) Name(v string) ContextOption {
	return func(o *ContextOptions) {
		o.Name = v
	}
}

// PluginBundle 插件包
func (WithContextOption) PluginBundle(v plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = v
	}
}

// StartedCallback 启动运行时回调函数
func (WithContextOption) StartedCallback(v func(runtimeCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (WithContextOption) StoppingCallback(v func(runtimeCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (WithContextOption) StoppedCallback(v func(runtimeCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}

// FrameBeginCallback 帧更新开始时的回调函数
func (WithContextOption) FrameBeginCallback(v func(runtimeCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.FrameBeginCallback = v
	}
}

// FrameEndCallback 帧更新结束时的回调函数
func (WithContextOption) FrameEndCallback(v func(runtimeCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.FrameEndCallback = v
	}
}

// FaceCache Face缓存，用于提高性能
func (WithContextOption) FaceCache(v *container.Cache[util.FaceAny]) ContextOption {
	return func(o *ContextOptions) {
		o.FaceCache = v
	}
}

// HookCache Hook缓存，用于提高性能
func (WithContextOption) HookCache(v *container.Cache[localevent.Hook]) ContextOption {
	return func(o *ContextOptions) {
		o.HookCache = v
	}
}
