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
	CompositeFace    util.Face[Context]       // 扩展者，需要扩展服务上下文自身能力时需要使用
	Context          context.Context          // 父Context
	AutoRecover      bool                     // 是否开启panic时自动恢复
	ReportError      chan error               // panic时错误写入的error channel
	Name             string                   // 服务名称
	PersistID        ec.ID                    // 服务持久化ID
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
		WithContextOption{}.CompositeFace(util.Face[Context]{})(o)
		WithContextOption{}.Context(nil)(o)
		WithContextOption{}.AutoRecover(false)(o)
		WithContextOption{}.ReportError(nil)(o)
		WithContextOption{}.Name("")(o)
		WithContextOption{}.PersistID(util.Zero[ec.ID]())(o)
		WithContextOption{}.GenPersistID(func() ec.ID { return ec.ID(ksuid.New()) })(o)
		WithContextOption{}.EntityLib(nil)(o)
		WithContextOption{}.PluginBundle(nil)(o)
		WithContextOption{}.StartedCallback(nil)(o)
		WithContextOption{}.StoppingCallback(nil)(o)
		WithContextOption{}.StoppedCallback(nil)(o)
	}
}

// CompositeFace 扩展者，需要扩展服务上下文自身能力时需要使用
func (WithContextOption) CompositeFace(face util.Face[Context]) ContextOption {
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

// Name 服务名称
func (WithContextOption) Name(name string) ContextOption {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistID 服务持久化ID
func (WithContextOption) PersistID(id ec.ID) ContextOption {
	return func(o *ContextOptions) {
		o.PersistID = id
	}
}

// GenPersistID 生成持久化ID的函数
func (WithContextOption) GenPersistID(fn func() ec.ID) ContextOption {
	return func(o *ContextOptions) {
		if fn == nil {
			panic("GenPersistID nil invalid")
		}
		o.GenPersistID = fn
	}
}

// EntityLib 实体原型库
func (WithContextOption) EntityLib(lib pt.EntityLib) ContextOption {
	return func(o *ContextOptions) {
		o.EntityLib = lib
	}
}

// PluginBundle 插件包
func (WithContextOption) PluginBundle(bundle plugin.PluginBundle) ContextOption {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// StartedCallback 启动运行时回调函数
func (WithContextOption) StartedCallback(fn func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StartedCallback = fn
	}
}

// StoppingCallback 开始停止运行时回调函数
func (WithContextOption) StoppingCallback(fn func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppingCallback = fn
	}
}

// StoppedCallback 完全停止运行时回调函数
func (WithContextOption) StoppedCallback(fn func(serviceCtx Context)) ContextOption {
	return func(o *ContextOptions) {
		o.StoppedCallback = fn
	}
}
