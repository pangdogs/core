//go:generate stringer -type RunningState
package runtime

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"kit.golaxy.org/golaxy/util/uid"
)

type _ContextOption struct{}

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
	RunningHandler = generic.Action2[Context, RunningState] // 运行状态变化处理器
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	CompositeFace    iface.Face[Context]                // 扩展者，需要扩展运行时上下文自身能力时需要使用
	Context          context.Context                    // 父Context
	AutoRecover      bool                               // 是否开启panic时自动恢复
	ReportError      chan error                         // panic时错误写入的error channel
	Name             string                             // 运行时名称
	PersistId        uid.Id                             // 运行时持久化Id
	PluginBundle     plugin.PluginBundle                // 插件包
	RunningHandler   RunningHandler                     // 运行状态变化处理器
	FaceAnyAllocator container.Allocator[iface.FaceAny] // 自定义FaceAny内存分配器，用于提高性能
	HookAllocator    container.Allocator[event.Hook]    // 自定义Hook内存分配器，用于提高性能
}

// Default 默认值
func (_ContextOption) Default() option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		_ContextOption{}.Composite(iface.Face[Context]{})(o)
		_ContextOption{}.Context(nil)(o)
		_ContextOption{}.AutoRecover(false)(o)
		_ContextOption{}.ReportError(nil)(o)
		_ContextOption{}.Name("")(o)
		_ContextOption{}.PersistId(uid.Nil)(o)
		_ContextOption{}.PluginBundle(nil)(o)
		_ContextOption{}.RunningHandler(nil)(o)
		_ContextOption{}.FaceAnyAllocator(container.DefaultAllocator[iface.FaceAny]())(o)
		_ContextOption{}.HookAllocator(container.DefaultAllocator[event.Hook]())(o)
	}
}

// Composite 扩展者，需要扩展运行时上下文自身功能时需要使用
func (_ContextOption) Composite(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (_ContextOption) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// AutoRecover 是否开启panic时自动恢复
func (_ContextOption) AutoRecover(b bool) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AutoRecover = b
	}
}

// ReportError panic时错误写入的error channel
func (_ContextOption) ReportError(ch chan error) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.ReportError = ch
	}
}

// Name 运行时名称
func (_ContextOption) Name(name string) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 运行时持久化Id
func (_ContextOption) PersistId(id uid.Id) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// PluginBundle 插件包
func (_ContextOption) PluginBundle(bundle plugin.PluginBundle) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// RunningHandler 运行状态变化处理器
func (_ContextOption) RunningHandler(handler RunningHandler) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.RunningHandler = handler
	}
}

// FaceAnyAllocator 自定义FaceAny内存分配器，用于提高性能
func (_ContextOption) FaceAnyAllocator(allocator container.Allocator[iface.FaceAny]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrContext, internal.ErrArgs))
		}
		o.FaceAnyAllocator = allocator
	}
}

// HookAllocator 自定义Hook内存分配器，用于提高性能
func (_ContextOption) HookAllocator(allocator container.Allocator[event.Hook]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		if allocator == nil {
			panic(fmt.Errorf("%w: %w: allocator is nil", ErrContext, internal.ErrArgs))
		}
		o.HookAllocator = allocator
	}
}
