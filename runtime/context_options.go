//go:generate stringer -type RunningState
package runtime

import (
	"context"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// Option 所有选项设置器
type Option struct{}

// RunningState 运行状态
type RunningState int32

const (
	RunningState_Birth                RunningState = iota // 出生
	RunningState_Starting                                 // 开始启动
	RunningState_Started                                  // 已启动
	RunningState_FrameLoopBegin                           // 帧循环开始
	RunningState_FrameUpdateBegin                         // 帧更新开始
	RunningState_FrameUpdateEnd                           // 帧更新结束
	RunningState_FrameLoopEnd                             // 帧循环结束
	RunningState_AsyncProcessingBegin                     // 异步调用处理开始
	RunningState_AsyncProcessingEnd                       // 异步调用处理结束
	RunningState_Terminating                              // 开始停止
	RunningState_Terminated                               // 已停止
)

type (
	RunningHandler = func(ctx Context, state RunningState) // 运行状态变化处理器
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	CompositeFace    util.Face[Context]                   // 扩展者，需要扩展运行时上下文自身能力时需要使用
	Context          context.Context                      // 父Context
	AutoRecover      bool                                 // 是否开启panic时自动恢复
	ReportError      chan error                           // panic时错误写入的error channel
	Name             string                               // 运行时名称
	PersistId        uid.Id                               // 运行时持久化Id
	PluginBundle     plugin.PluginBundle                  // 插件包
	RunningHandler   RunningHandler                       // 运行状态变化处理器
	FaceAnyAllocator container.Allocator[util.FaceAny]    // 自定义FaceAny内存分配器，用于提高性能
	HookAllocator    container.Allocator[localevent.Hook] // 自定义Hook内存分配器，用于提高性能
}

// ContextOption 创建运行时上下文的选项设置器
type ContextOption func(o *ContextOptions)

// Default 默认值
func (Option) Default() ContextOption {
	return func(o *ContextOptions) {
		Option{}.Composite(util.Face[Context]{})(o)
		Option{}.Context(nil)(o)
		Option{}.AutoRecover(false)(o)
		Option{}.ReportError(nil)(o)
		Option{}.Name("")(o)
		Option{}.PersistId(util.Zero[uid.Id]())(o)
		Option{}.PluginBundle(nil)(o)
		Option{}.RunningHandler(nil)(o)
		Option{}.FaceAnyAllocator(container.DefaultAllocator[util.FaceAny]())(o)
		Option{}.HookAllocator(container.DefaultAllocator[localevent.Hook]())(o)
	}
}

// Composite 扩展者，需要扩展运行时上下文自身功能时需要使用
func (Option) Composite(face util.Face[Context]) ContextOption {
	return func(o *ContextOptions) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (Option) Context(ctx context.Context) ContextOption {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// AutoRecover 是否开启panic时自动恢复
func (Option) AutoRecover(b bool) ContextOption {
	return func(o *ContextOptions) {
		o.AutoRecover = b
	}
}

// ReportError panic时错误写入的error channel
func (Option) ReportError(ch chan error) ContextOption {
	return func(o *ContextOptions) {
		o.ReportError = ch
	}
}

// Name 运行时名称
func (Option) Name(name string) ContextOption {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 运行时持久化Id
func (Option) PersistId(id uid.Id) ContextOption {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// PluginBundle 插件包
func (Option) PluginBundle(bundle plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// RunningHandler 运行状态变化处理器
func (Option) RunningHandler(fn RunningHandler) ContextOption {
	return func(o *ContextOptions) {
		o.RunningHandler = fn
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能
func (Option) FaceAnyAllocator(allocator container.Allocator[util.FaceAny]) ContextOption {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能
func (Option) HookAllocator(allocator container.Allocator[localevent.Hook]) ContextOption {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic("nil allocator")
		}
		o.HookAllocator = allocator
	}
}
