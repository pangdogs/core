package service

import (
	"context"
	"github.com/pangdogs/galaxy/plugin"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/util"
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	Inheritor        util.Face[Context]       // 继承者，需要拓展服务上下文自身能力时需要使用
	Context          context.Context          // 父Context
	AutoRecover      bool                     // 是否开启panic时自动恢复
	ReportError      chan error               // panic时错误写入的error channel
	Prototype        string                   // 服务原型名称
	NodeID           int64                    // 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
	EntityLib        pt.EntityLib             // 实体原型库
	PluginLib        plugin.PluginLib         // 插件库
	StartedCallback  func(serviceCtx Context) // 启动运行时回调函数
	StoppingCallback func(serviceCtx Context) // 开始停止运行时回调函数
	StoppedCallback  func(serviceCtx Context) // 完全停止运行时回调函数
}

// ContextOptionSetter 创建服务上下文的选项设置器
type ContextOptionSetter func(o *ContextOptions)

// ContextOption 创建服务上下文的选项
var ContextOption = &_ContextOption{}

type _ContextOption struct{}

// Default 默认值
func (*_ContextOption) Default() ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = util.Face[Context]{}
		o.Context = nil
		o.AutoRecover = false
		o.ReportError = nil
		o.Prototype = ""
		o.NodeID = 0
		o.EntityLib = nil
		o.PluginLib = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
	}
}

// Inheritor 继承者，需要拓展服务上下文自身能力时需要使用
func (*_ContextOption) Inheritor(v util.Face[Context]) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = v
	}
}

// Context 父Context
func (*_ContextOption) Context(v context.Context) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Context = v
	}
}

// AutoRecover 是否开启panic时自动恢复
func (*_ContextOption) AutoRecover(v bool) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.AutoRecover = v
	}
}

// ReportError panic时错误写入的error channel
func (*_ContextOption) ReportError(v chan error) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.ReportError = v
	}
}

// Prototype 服务原型名称
func (*_ContextOption) Prototype(v string) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Prototype = v
	}
}

// NodeID 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
func (*_ContextOption) NodeID(v int64) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.NodeID = v
	}
}

// EntityLib 实体原型库
func (*_ContextOption) EntityLib(v pt.EntityLib) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.EntityLib = v
	}
}

// PluginLib 插件库
func (*_ContextOption) PluginLib(v plugin.PluginLib) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.PluginLib = v
	}
}

// StartedCallback 启动运行时回调函数
func (*_ContextOption) StartedCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (*_ContextOption) StoppingCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (*_ContextOption) StoppedCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}
