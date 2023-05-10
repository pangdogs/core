package launcher

import (
	"github.com/alecthomas/kingpin/v2"
	"kit.golaxy.org/golaxy"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/service"
	"os"
	"syscall"
)

// Cmd 应用指令
type Cmd struct {
	Clause *kingpin.CmdClause        // cmd clause
	Flags  []interface{}             // cmd flags
	Run    func(flags []interface{}) // run cmd
}

// ServiceCtxInitFunc 服务上下文初始化函数
type ServiceCtxInitFunc func(entityLib pt.EntityLib, pluginBundle plugin.PluginBundle) []service.ContextOption

// ServiceInitFunc 服务初始化函数
type ServiceInitFunc func() []golaxy.ServiceOption

// AppOptions 创建应用的所有选项
type AppOptions struct {
	Commands          func() []Cmd                  // 自定义应用指令
	QuitSignals       []os.Signal                   // 退出信号
	ServiceCtxInitTab map[string]ServiceCtxInitFunc // 所有服务上下文初始化函数
	ServiceInitTab    map[string]ServiceInitFunc    // 所有服务初始化函数
}

// AppOption 创建应用的选项设置器
type AppOption func(o *AppOptions)

// WithAppOption 创建应用的所有选项设置器
type WithAppOption struct{}

// Default 默认值
func (WithAppOption) Default() AppOption {
	return func(o *AppOptions) {
		WithAppOption{}.Commands(nil)(o)
		WithAppOption{}.QuitSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)(o)
		WithAppOption{}.ServiceCtxInitTab(nil)(o)
		WithAppOption{}.ServiceInitTab(nil)(o)
	}
}

// Commands 自定义应用指令
func (WithAppOption) Commands(fn func() []Cmd) AppOption {
	return func(o *AppOptions) {
		o.Commands = fn
	}
}

// QuitSignals 退出信号
func (WithAppOption) QuitSignals(signals ...os.Signal) AppOption {
	return func(o *AppOptions) {
		o.QuitSignals = signals
	}
}

// ServiceCtxInitTab 所有服务上下文初始化函数
func (WithAppOption) ServiceCtxInitTab(tab map[string]ServiceCtxInitFunc) AppOption {
	return func(o *AppOptions) {
		o.ServiceCtxInitTab = tab
	}
}

// ServiceInitTab 所有服务初始化函数
func (WithAppOption) ServiceInitTab(tab map[string]ServiceInitFunc) AppOption {
	return func(o *AppOptions) {
		o.ServiceInitTab = tab
	}
}
