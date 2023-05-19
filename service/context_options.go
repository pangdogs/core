package service

import (
	"context"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
)

// Options 创建服务上下文的所有选项
type Options struct {
	CompositeFace    util.Face[Context]  // 扩展者，需要扩展服务上下文自身能力时需要使用
	Context          context.Context     // 父Context
	AutoRecover      bool                // 是否开启panic时自动恢复
	ReportError      chan error          // panic时错误写入的error channel
	Name             string              // 服务名称
	PersistId        uid.Id              // 服务持久化Id
	EntityLib        pt.EntityLib        // 实体原型库
	PluginBundle     plugin.PluginBundle // 插件包
	StartedCallback  func(ctx Context)   // 启动运行时回调函数
	StoppingCallback func(ctx Context)   // 开始停止运行时回调函数
	StoppedCallback  func(ctx Context)   // 完全停止运行时回调函数
}

// Option 创建服务上下文的选项设置器
type Option func(o *Options)

// WithOption 创建服务上下文的所有选项设置器
type WithOption struct{}

// Default 默认值
func (WithOption) Default() Option {
	return func(o *Options) {
		WithOption{}.CompositeFace(util.Face[Context]{})(o)
		WithOption{}.Context(nil)(o)
		WithOption{}.AutoRecover(false)(o)
		WithOption{}.ReportError(nil)(o)
		WithOption{}.Name("")(o)
		WithOption{}.PersistId(util.Zero[uid.Id]())(o)
		WithOption{}.EntityLib(nil)(o)
		WithOption{}.PluginBundle(nil)(o)
		WithOption{}.StartedCallback(nil)(o)
		WithOption{}.StoppingCallback(nil)(o)
		WithOption{}.StoppedCallback(nil)(o)
	}
}

// CompositeFace 扩展者，需要扩展服务上下文自身能力时需要使用
func (WithOption) CompositeFace(face util.Face[Context]) Option {
	return func(o *Options) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (WithOption) Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// AutoRecover 是否开启panic时自动恢复
func (WithOption) AutoRecover(b bool) Option {
	return func(o *Options) {
		o.AutoRecover = b
	}
}

// ReportError panic时错误写入的error channel
func (WithOption) ReportError(ch chan error) Option {
	return func(o *Options) {
		o.ReportError = ch
	}
}

// Name 服务名称
func (WithOption) Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

// PersistId 服务持久化Id
func (WithOption) PersistId(id uid.Id) Option {
	return func(o *Options) {
		o.PersistId = id
	}
}

// EntityLib 实体原型库
func (WithOption) EntityLib(lib pt.EntityLib) Option {
	return func(o *Options) {
		o.EntityLib = lib
	}
}

// PluginBundle 插件包
func (WithOption) PluginBundle(bundle plugin.PluginBundle) Option {
	return func(o *Options) {
		o.PluginBundle = bundle
	}
}

// StartedCallback 启动运行时回调函数
func (WithOption) StartedCallback(fn func(ctx Context)) Option {
	return func(o *Options) {
		o.StartedCallback = fn
	}
}

// StoppingCallback 开始停止运行时回调函数
func (WithOption) StoppingCallback(fn func(ctx Context)) Option {
	return func(o *Options) {
		o.StoppingCallback = fn
	}
}

// StoppedCallback 完全停止运行时回调函数
func (WithOption) StoppedCallback(fn func(ctx Context)) Option {
	return func(o *Options) {
		o.StoppedCallback = fn
	}
}
