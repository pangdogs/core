package launcher

import (
	"github.com/alecthomas/kingpin/v2"
	"kit.golaxy.org/golaxy"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/option"
	"os"
	"syscall"
)

// Option 所有选项设置器
type Option struct{}

type (
	ServiceCtxCtor = generic.DelegateFunc3[string, pt.EntityLib, plugin.PluginBundle, []option.Setting[service.ContextOptions]] // 服务上下文构造函数
	ServiceCtor    = generic.DelegateFunc1[string, []option.Setting[golaxy.ServiceOptions]]                                     // 服务构造函数
)

// Cmd 应用指令
type Cmd struct {
	Clause *kingpin.CmdClause // cmd clause
	Run    func()             // run cmd
}

// AppOptions 创建应用的所有选项
type AppOptions struct {
	Commands       []Cmd          // 应用指令
	QuitSignals    []os.Signal    // 退出信号
	ServiceCtxCtor ServiceCtxCtor // 服务上下文构造函数
	ServiceCtor    ServiceCtor    // 服务构造函数
}

// Default 默认值
func (Option) Default() option.Setting[AppOptions] {
	return func(o *AppOptions) {
		Option{}.Commands(nil)(o)
		Option{}.QuitSignals(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)(o)
		Option{}.ServiceCtxCtor(nil)(o)
		Option{}.ServiceCtor(nil)(o)
	}
}

// Commands 应用指令
func (Option) Commands(cmds []Cmd) option.Setting[AppOptions] {
	return func(o *AppOptions) {
		o.Commands = cmds
	}
}

// QuitSignals 退出信号
func (Option) QuitSignals(signals ...os.Signal) option.Setting[AppOptions] {
	return func(o *AppOptions) {
		o.QuitSignals = signals
	}
}

// ServiceCtxCtor 服务上下文构造函数
func (Option) ServiceCtxCtor(ctor ServiceCtxCtor) option.Setting[AppOptions] {
	return func(o *AppOptions) {
		o.ServiceCtxCtor = ctor
	}
}

// ServiceCtor 服务构造函数
func (Option) ServiceCtor(ctor ServiceCtor) option.Setting[AppOptions] {
	return func(o *AppOptions) {
		o.ServiceCtor = ctor
	}
}
