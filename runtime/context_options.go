package runtime

import (
	"context"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	CompositeFace      util.Face[Context]                   // 扩展者，需要扩展运行时上下文自身能力时需要使用
	Context            context.Context                      // 父Context
	AutoRecover        bool                                 // 是否开启panic时自动恢复
	ReportError        chan error                           // panic时错误写入的error channel
	Name               string                               // 运行时名称
	PersistId          uid.Id                               // 运行时持久化Id
	PluginBundle       plugin.PluginBundle                  // 插件包
	StartedCallback    func(ctx Context)                    // 启动运行时回调函数
	StoppingCallback   func(ctx Context)                    // 开始停止运行时回调函数
	StoppedCallback    func(ctx Context)                    // 完全停止运行时回调函数
	FrameBeginCallback func(ctx Context)                    // 帧开始时的回调函数
	FrameEndCallback   func(ctx Context)                    // 帧结束时的回调函数
	FaceAnyAllocator   container.Allocator[util.FaceAny]    // 自定义FaceAny内存分配器，用于提高性能
	HookAllocator      container.Allocator[localevent.Hook] // 自定义Hook内存分配器，用于提高性能
}

// ContextOption 创建运行时上下文的选项设置器
type ContextOption func(o *ContextOptions)

// WithContextOption 创建运行时上下文的所有选项设置器
type WithContextOption struct{}

// Default 默认值
func (WithContextOption) Default() ContextOption {
	return func(o *ContextOptions) {
		WithContextOption{}.Composite(util.Face[Context]{})(o)
		WithContextOption{}.Context(nil)(o)
		WithContextOption{}.AutoRecover(false)(o)
		WithContextOption{}.ReportError(nil)(o)
		WithContextOption{}.Name("")(o)
		WithContextOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithContextOption{}.PluginBundle(nil)(o)
		WithContextOption{}.StartedCallback(nil)(o)
		WithContextOption{}.StoppingCallback(nil)(o)
		WithContextOption{}.StoppedCallback(nil)(o)
		WithContextOption{}.FrameBeginCallback(nil)(o)
		WithContextOption{}.FrameEndCallback(nil)(o)
		WithContextOption{}.FaceAnyAllocator(container.DefaultAllocator[util.FaceAny]())(o)
		WithContextOption{}.HookAllocator(container.DefaultAllocator[localevent.Hook]())(o)
	}
}

// Composite 扩展者，需要扩展运行时上下文自身功能时需要使用
func (WithContextOption) Composite(face util.Face[Context]) ContextOption {
	return func(o *ContextOptions) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (WithContextOption) Context(ctx context.Context) ContextOption {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// AutoRecover 是否开启panic时自动恢复
func (WithContextOption) AutoRecover(b bool) ContextOption {
	return func(o *ContextOptions) {
		o.AutoRecover = b
	}
}

// ReportError panic时错误写入的error channel
func (WithContextOption) ReportError(ch chan error) ContextOption {
	return func(o *ContextOptions) {
		o.ReportError = ch
	}
}

// Name 运行时名称
func (WithContextOption) Name(name string) ContextOption {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 运行时持久化Id
func (WithContextOption) PersistId(id uid.Id) ContextOption {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// PluginBundle 插件包
func (WithContextOption) PluginBundle(bundle plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// StartedCallback 启动运行时回调函数
func (WithContextOption) StartedCallback(fn func(ctx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StartedCallback = fn
	}
}

// StoppingCallback 开始停止运行时回调函数
func (WithContextOption) StoppingCallback(fn func(ctx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppingCallback = fn
	}
}

// StoppedCallback 完全停止运行时回调函数
func (WithContextOption) StoppedCallback(fn func(ctx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppedCallback = fn
	}
}

// FrameBeginCallback 帧更新开始时的回调函数
func (WithContextOption) FrameBeginCallback(fn func(ctx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.FrameBeginCallback = fn
	}
}

// FrameEndCallback 帧更新结束时的回调函数
func (WithContextOption) FrameEndCallback(fn func(ctx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.FrameEndCallback = fn
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能
func (WithContextOption) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) ContextOption {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能
func (WithContextOption) HookAllocator(allocator container.Allocator[localevent.Hook]) ContextOption {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.HookAllocator = allocator
	}
}
