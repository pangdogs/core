package service

import (
	"context"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

type (
	RunningHandler = generic.DelegateAction2[Context, RunningState] // 运行状态变化处理器
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	CompositeFace  iface.Face[Context] // 扩展者，在扩展服务上下文自身能力时使用
	Context        context.Context     // 父Context
	AutoRecover    bool                // 是否开启panic时自动恢复
	ReportError    chan error          // panic时错误写入的error channel
	Name           string              // 服务名称
	PersistId      uid.Id              // 服务持久化Id
	EntityLib      pt.EntityLib        // 实体原型库
	PluginBundle   plugin.PluginBundle // 插件包
	RunningHandler RunningHandler      // 运行状态变化处理器
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		With.CompositeFace(iface.Face[Context]{})(o)
		With.Context(nil)(o)
		With.PanicHandling(false, nil)(o)
		With.Name("")(o)
		With.PersistId(uid.Nil)(o)
		With.EntityLib(pt.DefaultEntityLib())(o)
		With.PluginBundle(plugin.NewPluginBundle())(o)
		With.RunningHandler(nil)(o)
	}
}

// CompositeFace 扩展者，在扩展服务上下文自身能力时使用
func (_Option) CompositeFace(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.CompositeFace = face
	}
}

// Context 父Context
func (_Option) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// PanicHandling panic时的处理方式
func (_Option) PanicHandling(autoRecover bool, reportError chan error) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AutoRecover = autoRecover
		o.ReportError = reportError
	}
}

// Name 服务名称
func (_Option) Name(name string) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 服务持久化Id
func (_Option) PersistId(id uid.Id) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// EntityLib 实体原型库
func (_Option) EntityLib(lib pt.EntityLib) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.EntityLib = lib
	}
}

// PluginBundle 插件包
func (_Option) PluginBundle(bundle plugin.PluginBundle) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PluginBundle = bundle
	}
}

// RunningHandler 运行状态变化处理器
func (_Option) RunningHandler(handler RunningHandler) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.RunningHandler = handler
	}
}
