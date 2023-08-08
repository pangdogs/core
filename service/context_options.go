package service

import (
	"context"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
)

// WithOption 所有选项设置器
type WithOption struct{}

type (
	Callback = func(ctx Context) // 回调函数
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	CompositeFace util.Face[Context]  // 扩展者，需要扩展服务上下文自身能力时需要使用
	Context       context.Context     // 父Context
	AutoRecover   bool                // 是否开启panic时自动恢复
	ReportError   chan error          // panic时错误写入的error channel
	Name          string              // 服务名称
	PersistId     uid.Id              // 服务持久化Id
	EntityLib     pt.EntityLib        // 实体原型库
	PluginBundle  plugin.PluginBundle // 插件包
	StartedCb     Callback            // 启动运行时回调函数
	StoppingCb    Callback            // 开始停止运行时回调函数
	StoppedCb     Callback            // 完全停止运行时回调函数
}

// ContextOption 创建服务上下文的选项设置器
type ContextOption func(o *ContextOptions)

// Default 默认值
func (WithOption) Default() ContextOption {
	return func(o *ContextOptions) {
		WithOption{}.CompositeFace(util.Face[Context]{})(o)
		WithOption{}.Context(nil)(o)
		WithOption{}.AutoRecover(false)(o)
		WithOption{}.ReportError(nil)(o)
		WithOption{}.Name("")(o)
		WithOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithOption{}.EntityLib(nil)(o)
		WithOption{}.PluginBundle(nil)(o)
		WithOption{}.StartedCb(nil)(o)
		WithOption{}.StoppingCb(nil)(o)
		WithOption{}.StoppedCb(nil)(o)
	}
}

// CompositeFace 扩展者，需要扩展服务上下文自身能力时需要使用
func (WithOption) CompositeFace(face util.Face[Context]) ContextOption {
	return func(o *ContextOptions) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (WithOption) Context(ctx context.Context) ContextOption {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// AutoRecover 是否开启panic时自动恢复
func (WithOption) AutoRecover(b bool) ContextOption {
	return func(o *ContextOptions) {
		o.AutoRecover = b
	}
}

// ReportError panic时错误写入的error channel
func (WithOption) ReportError(ch chan error) ContextOption {
	return func(o *ContextOptions) {
		o.ReportError = ch
	}
}

// Name 服务名称
func (WithOption) Name(name string) ContextOption {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 服务持久化Id
func (WithOption) PersistId(id uid.Id) ContextOption {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// EntityLib 实体原型库
func (WithOption) EntityLib(lib pt.EntityLib) ContextOption {
	return func(o *ContextOptions) {
		o.EntityLib = lib
	}
}

// PluginBundle 插件包
func (WithOption) PluginBundle(bundle plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// StartedCb 启动运行时回调函数
func (WithOption) StartedCb(fn Callback) ContextOption {
	return func(o *ContextOptions) {
		o.StartedCb = fn
	}
}

// StoppingCb 开始停止运行时回调函数
func (WithOption) StoppingCb(fn Callback) ContextOption {
	return func(o *ContextOptions) {
		o.StoppingCb = fn
	}
}

// StoppedCb 完全停止运行时回调函数
func (WithOption) StoppedCb(fn Callback) ContextOption {
	return func(o *ContextOptions) {
		o.StoppedCb = fn
	}
}
