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

// Option 所有选项设置器
type Option struct{}

type (
	ServiceCtxHandler = func(serviceName string, entityLib pt.EntityLib, pluginBundle plugin.PluginBundle) []service.ContextOption // 服务上下文初始化处理器
	ServiceHandler    = func(serviceName string) []golaxy.ServiceOption                                                            // 服务初始化处理器
)

// Cmd 应用指令
type Cmd struct {
	Clause *kingpin.CmdClause // cmd clause
	Run    func()             // run cmd
}

// AppOptions 创建应用的所有选项
type AppOptions struct {
	Commands           []Cmd                        // 应用指令
	QuitSignals        []os.Signal                  // 退出信号
	ServiceCtxHandlers map[string]ServiceCtxHandler // 服务上下文初始化处理器
	ServiceHandlers    map[string]ServiceHandler    // 服务初始化处理器
}

// AppOption 创建应用的选项设置器
type AppOption func(o *AppOptions)

// Default 默认值
func (Option) Default() AppOption {
	return func(o *AppOptions) {
		Option{}.Commands(nil)(o)
		Option{}.QuitSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)(o)
		Option{}.ServiceCtxHandlers(nil)(o)
		Option{}.ServiceHandlers(nil)(o)
	}
}

// Commands 应用指令
func (Option) Commands(cmds []Cmd) AppOption {
	return func(o *AppOptions) {
		o.Commands = cmds
	}
}

// QuitSignals 退出信号
func (Option) QuitSignals(signals ...os.Signal) AppOption {
	return func(o *AppOptions) {
		o.QuitSignals = signals
	}
}

// ServiceCtxHandlers 服务上下文初始化处理器
func (Option) ServiceCtxHandlers(handlers map[string]ServiceCtxHandler) AppOption {
	return func(o *AppOptions) {
		o.ServiceCtxHandlers = handlers
	}
}

// ServiceHandlers 服务初始化处理器
func (Option) ServiceHandlers(handlers map[string]ServiceHandler) AppOption {
	return func(o *AppOptions) {
		o.ServiceHandlers = handlers
	}
}
