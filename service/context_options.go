package service

import (
	"context"
	"github.com/segmentio/ksuid"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util"
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	Inheritor        util.Face[Context]       // 继承者，需要扩展服务上下文自身能力时需要使用
	Context          context.Context          // 父Context
	AutoRecover      bool                     // 是否开启panic时自动恢复
	ReportError      chan error               // panic时错误写入的error channel
	Name             string                   // 服务名称
	GenPersistID     func() ec.ID             // 生成持久化ID的函数
	EntityLib        pt.EntityLib             // 实体原型库
	PluginBundle     plugin.PluginBundle      // 插件包
	StartedCallback  func(serviceCtx Context) // 启动运行时回调函数
	StoppingCallback func(serviceCtx Context) // 开始停止运行时回调函数
	StoppedCallback  func(serviceCtx Context) // 完全停止运行时回调函数
}

// ContextOption 创建服务上下文的选项设置器
type ContextOption func(o *ContextOptions)

// WithContextOption 创建服务上下文的所有选项设置器
type WithContextOption struct{}

// Default 默认值
func (WithContextOption) Default() ContextOption {
	return func(o *ContextOptions) {
		WithContextOption{}.Inheritor(util.Face[Context]{})(o)
		WithContextOption{}.Context(nil)(o)
		WithContextOption{}.AutoRecover(false)(o)
		WithContextOption{}.ReportError(nil)(o)
		WithContextOption{}.Name("")(o)
		WithContextOption{}.GenPersistID(func() ec.ID { return ec.ID(ksuid.New()) })(o)
		WithContextOption{}.EntityLib(nil)(o)
		WithContextOption{}.PluginBundle(nil)(o)
		WithContextOption{}.StartedCallback(nil)(o)
		WithContextOption{}.StoppingCallback(nil)(o)
		WithContextOption{}.StoppedCallback(nil)(o)
	}
}

// Inheritor 继承者，需要扩展服务上下文自身能力时需要使用
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

// GenPersistID 生成持久化ID的函数
func (WithContextOption) GenPersistID(v func() ec.ID) ContextOption {
	return func(o *ContextOptions) {
		if v == nil {
			panic("GenPersistID nil invalid")
		}
		o.GenPersistID = v
	}
}

// EntityLib 实体原型库
func (WithContextOption) EntityLib(v pt.EntityLib) ContextOption {
	return func(o *ContextOptions) {
		o.EntityLib = v
	}
}

// PluginBundle 插件包
func (WithContextOption) PluginBundle(v plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = v
	}
}

// StartedCallback 启动运行时回调函数
func (WithContextOption) StartedCallback(v func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (WithContextOption) StoppingCallback(v func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (WithContextOption) StoppedCallback(v func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}
