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
type ServiceCtxInitFunc = func(serviceName string, entityLib pt.EntityLib, pluginBundle plugin.PluginBundle) []service.ContextOption

// ServiceInitFunc 服务初始化函数
type ServiceInitFunc = func(serviceName string) []golaxy.ServiceOption

// Options 创建应用的所有选项
type Options struct {
	Commands          func() []Cmd                  // 自定义应用指令
	QuitSignals       []os.Signal                   // 退出信号
	ServiceCtxInitTab map[string]ServiceCtxInitFunc // 所有服务上下文初始化函数
	ServiceInitTab    map[string]ServiceInitFunc    // 所有服务初始化函数
}

// Option 创建应用的选项设置器
type Option func(o *Options)

// WithOption 创建应用的所有选项设置器
type WithOption struct{}

// Default 默认值
func (WithOption) Default() Option {
	return func(o *Options) {
		WithOption{}.Commands(nil)(o)
		WithOption{}.QuitSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)(o)
		WithOption{}.ServiceCtxInitTab(nil)(o)
		WithOption{}.ServiceInitTab(nil)(o)
	}
}

// Commands 自定义应用指令
func (WithOption) Commands(fn func() []Cmd) Option {
	return func(o *Options) {
		o.Commands = fn
	}
}

// QuitSignals 退出信号
func (WithOption) QuitSignals(signals ...os.Signal) Option {
	return func(o *Options) {
		o.QuitSignals = signals
	}
}

// ServiceCtxInitTab 所有服务上下文初始化函数
func (WithOption) ServiceCtxInitTab(tab map[string]ServiceCtxInitFunc) Option {
	return func(o *Options) {
		o.ServiceCtxInitTab = tab
	}
}

// ServiceInitTab 所有服务初始化函数
func (WithOption) ServiceInitTab(tab map[string]ServiceInitFunc) Option {
	return func(o *Options) {
		o.ServiceInitTab = tab
	}
}
