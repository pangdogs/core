package service

import (
	"context"
	"github.com/pangdogs/galaxy/util"
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	Inheritor        util.Face[Context]       // 继承者，需要拓展服务上下文自身能力时需要使用
	Prototype        string                   // 服务原型名称
	NodeID           int64                    // 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
	AutoRecover      bool                     // 是否开启panic时自动恢复
	ReportError      chan error               // panic时错误写入的error channel
	ParentContext    context.Context          // 父Context
	StartedCallback  func(serviceCtx Context) // 启动运行时回调函数
	StoppingCallback func(serviceCtx Context) // 开始停止运行时回调函数
	StoppedCallback  func(serviceCtx Context) // 完全停止运行时回调函数
}

// ContextOption 创建服务上下文的选项
var ContextOption = &_ContextOptionSetter{}

// ContextOptionSetter 创建服务上下文的选项设置器
type ContextOptionSetter func(o *ContextOptions)

type _ContextOptionSetter struct{}

// Default 默认值
func (*_ContextOptionSetter) Default() ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = util.Face[Context]{}
		o.Prototype = ""
		o.NodeID = 0
		o.AutoRecover = false
		o.ReportError = nil
		o.ParentContext = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
	}
}

// Inheritor 继承者，需要拓展服务上下文自身能力时需要使用
func (*_ContextOptionSetter) Inheritor(v util.Face[Context]) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Inheritor = v
	}
}

// Prototype 服务原型名称
func (*_ContextOptionSetter) Prototype(v string) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.Prototype = v
	}
}

// NodeID 服务分布式节点ID，主要用于snowflake算法生成唯一ID，需要全局唯一
func (*_ContextOptionSetter) NodeID(v int64) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.NodeID = v
	}
}

// AutoRecover 是否开启panic时自动恢复
func (*_ContextOptionSetter) AutoRecover(v bool) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.AutoRecover = v
	}
}

// ReportError panic时错误写入的error channel
func (*_ContextOptionSetter) ReportError(v chan error) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.ReportError = v
	}
}

// ParentContext 父Context
func (*_ContextOptionSetter) ParentContext(v context.Context) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.ParentContext = v
	}
}

// StartedCallback 启动运行时回调函数
func (*_ContextOptionSetter) StartedCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StartedCallback = v
	}
}

// StoppingCallback 开始停止运行时回调函数
func (*_ContextOptionSetter) StoppingCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppingCallback = v
	}
}

// StoppedCallback 完全停止运行时回调函数
func (*_ContextOptionSetter) StoppedCallback(v func(serviceCtx Context)) ContextOptionSetter {
	return func(o *ContextOptions) {
		o.StoppedCallback = v
	}
}
