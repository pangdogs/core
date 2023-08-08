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

// WithOption 所有选项设置器
type WithOption struct{}

type (
	ServiceCtxInit = func(serviceName string, entityLib pt.EntityLib, pluginBundle plugin.PluginBundle) []service.ContextOption // 服务上下文初始化函数
	ServiceInit    = func(serviceName string) []golaxy.ServiceOption                                                            // 服务初始化函数
)

// Cmd 应用指令
type Cmd struct {
	Clause *kingpin.CmdClause        // cmd clause
	Flags  []interface{}             // cmd flags
	Run    func(flags []interface{}) // run cmd
}

// AppOptions 创建应用的所有选项
type AppOptions struct {
	Commands          []Cmd                     // 自定义应用指令
	QuitSignals       []os.Signal               // 退出信号
	ServiceCtxInitTab map[string]ServiceCtxInit // 所有服务上下文初始化函数
	ServiceInitTab    map[string]ServiceInit    // 所有服务初始化函数
}

// AppOption 创建应用的选项设置器
type AppOption func(o *AppOptions)

// Default 默认值
func (WithOption) Default() AppOption {
	return func(o *AppOptions) {
		WithOption{}.Commands(nil)(o)
		WithOption{}.QuitSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)(o)
		WithOption{}.ServiceCtxInitTab(nil)(o)
		WithOption{}.ServiceInitTab(nil)(o)
	}
}

// Commands 自定义应用指令
func (WithOption) Commands(cmds []Cmd) AppOption {
	return func(o *AppOptions) {
		o.Commands = cmds
	}
}

// QuitSignals 退出信号
func (WithOption) QuitSignals(signals ...os.Signal) AppOption {
	return func(o *AppOptions) {
		o.QuitSignals = signals
	}
}

// ServiceCtxInitTab 所有服务上下文初始化函数
func (WithOption) ServiceCtxInitTab(tab map[string]ServiceCtxInit) AppOption {
	return func(o *AppOptions) {
		o.ServiceCtxInitTab = tab
	}
}

// ServiceInitTab 所有服务初始化函数
func (WithOption) ServiceInitTab(tab map[string]ServiceInit) AppOption {
	return func(o *AppOptions) {
		o.ServiceInitTab = tab
	}
}
