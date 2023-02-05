package launcher

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/service"
)

// Cmd 应用指令
type Cmd struct {
	Clause *kingpin.CmdClause        // cmd clause
	Flags  []interface{}             // cmd flags
	Run    func(flags []interface{}) // run cmd
}

// AppOptions 创建应用的所有选项
type AppOptions struct {
	SetupCommands                func() []Cmd                                                        // 设置自定义应用指令
	ServiceInstallPlugin         func(serviceName string, pluginBundle plugin.PluginBundle)          // 安装服务插件
	ServiceSetupRecover          func(serviceName string) (autoRecover bool, reportError chan error) // 设置服务panic恢复选项
	ServiceSetupStartedCallback  func(serviceName string) func(serviceCtx service.Context)           // 设置服务启动时回调
	ServiceSetupStoppingCallback func(serviceName string) func(serviceCtx service.Context)           // 设置服务开始停止时回调
	ServiceSetupStoppedCallback  func(serviceName string) func(serviceCtx service.Context)           // 设置服务完全停止时回调
}

// AppOption 创建应用的选项设置器
type AppOption func(o *AppOptions)

// WithAppOption 创建应用的所有选项设置器
type WithAppOption struct{}

// Default 默认值
func (WithAppOption) Default() AppOption {
	return func(o *AppOptions) {
		o.SetupCommands = nil
		o.ServiceInstallPlugin = nil
		o.ServiceSetupRecover = nil
		o.ServiceSetupStartedCallback = nil
		o.ServiceSetupStoppingCallback = nil
		o.ServiceSetupStoppedCallback = nil
	}
}

// SetupCustomCommands 设置自定义应用指令
func (WithAppOption) SetupCustomCommands(v func() []Cmd) AppOption {
	return func(o *AppOptions) {
		o.SetupCommands = v
	}
}

// ServiceInstallPlugin 安装服务插件
func (WithAppOption) ServiceInstallPlugin(v func(serviceName string, pluginBundle plugin.PluginBundle)) AppOption {
	return func(o *AppOptions) {
		o.ServiceInstallPlugin = v
	}
}

// ServiceSetupRecover 设置服务panic恢复选项
func (WithAppOption) ServiceSetupRecover(v func(serviceName string) (autoRecover bool, reportError chan error)) AppOption {
	return func(o *AppOptions) {
		o.ServiceSetupRecover = v
	}
}

// ServiceSetupStartedCallback 设置服务启动时回调
func (WithAppOption) ServiceSetupStartedCallback(v func(serviceName string) func(serviceCtx service.Context)) AppOption {
	return func(o *AppOptions) {
		o.ServiceSetupStartedCallback = v
	}
}

// ServiceSetupStoppingCallback 设置服务开始停止时回调
func (WithAppOption) ServiceSetupStoppingCallback(v func(serviceName string) func(serviceCtx service.Context)) AppOption {
	return func(o *AppOptions) {
		o.ServiceSetupStoppingCallback = v
	}
}

// ServiceSetupStoppedCallback 设置服务完全停止时回调
func (WithAppOption) ServiceSetupStoppedCallback(v func(serviceName string) func(serviceCtx service.Context)) AppOption {
	return func(o *AppOptions) {
		o.ServiceSetupStoppedCallback = v
	}
}
