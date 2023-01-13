package define

import (
	"github.com/golaxy-kit/golaxy/plugin"
	"github.com/golaxy-kit/golaxy/runtime"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
)

type _Plugin[PLUGIN_IFACE, OPTION any] struct {
	_name string
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) name() string {
	return p._name
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) installTo(creator func(...OPTION) PLUGIN_IFACE) func(plugin.PluginBundle, ...OPTION) {
	return func(pluginBundle plugin.PluginBundle, options ...OPTION) {
		plugin.InstallPlugin[PLUGIN_IFACE](pluginBundle, p.name(), creator(options...))
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) uninstallFrom() func(plugin.PluginBundle) {
	return func(pluginBundle plugin.PluginBundle) {
		pluginBundle.Uninstall(p.name())
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) serviceContext() func(service.Context) PLUGIN_IFACE {
	return func(ctx service.Context) PLUGIN_IFACE {
		return service.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) runtimeContext() func(runtime.Context) PLUGIN_IFACE {
	return func(ctx runtime.Context) PLUGIN_IFACE {
		return runtime.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

// ServicePlugin 服务类插件
type ServicePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name          string                               // 插件名称
	InstallTo     func(plugin.PluginBundle, ...OPTION) // 向插件包安装
	UninstallFrom func(plugin.PluginBundle)            // 从插件包卸载
	Context       func(service.Context) PLUGIN_IFACE   // 从服务上下文获取插件
}

// ServicePlugin 生成服务类插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) ServicePlugin(creator func(...OPTION) PLUGIN_IFACE) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return ServicePlugin[PLUGIN_IFACE, OPTION]{
		Name:          p.name(),
		InstallTo:     p.installTo(creator),
		UninstallFrom: p.uninstallFrom(),
		Context:       p.serviceContext(),
	}
}

// RuntimePlugin 运行时类插件
type RuntimePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name          string                               // 插件名称
	InstallTo     func(plugin.PluginBundle, ...OPTION) // 向插件包安装
	UninstallFrom func(plugin.PluginBundle)            // 从插件包卸载
	Context       func(runtime.Context) PLUGIN_IFACE   // 从运行时上下文获取插件
}

// RuntimePlugin 生成运行时类插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) RuntimePlugin(creator func(...OPTION) PLUGIN_IFACE) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return RuntimePlugin[PLUGIN_IFACE, OPTION]{
		Name:          p.name(),
		InstallTo:     p.installTo(creator),
		UninstallFrom: p.uninstallFrom(),
		Context:       p.runtimeContext(),
	}
}

// Plugin 插件
type Plugin[PLUGIN_IFACE, OPTION any] struct {
	Name           string                               // 插件名称
	InstallTo      func(plugin.PluginBundle, ...OPTION) // 向插件包安装
	UninstallFrom  func(plugin.PluginBundle)            // 从插件包卸载
	ServiceContext func(service.Context) PLUGIN_IFACE   // 从服务上下文获取插件
	RuntimeContext func(runtime.Context) PLUGIN_IFACE   // 从运行时上下文获取插件
}

// Plugin 生成插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) Plugin(creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return Plugin[PLUGIN_IFACE, OPTION]{
		Name:           p.name(),
		InstallTo:      p.installTo(creator),
		UninstallFrom:  p.uninstallFrom(),
		ServiceContext: p.serviceContext(),
		RuntimeContext: p.runtimeContext(),
	}
}

// DefinePlugin 定义插件，可以用于向插件包安装插件
func DefinePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.Plugin(creator)
}

// DefineRuntimePlugin 定义运行时插件，可以用于向插件包安装插件
func DefineRuntimePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.RuntimePlugin(creator)
}

// DefineServicePlugin 定义插件，可以用于向插件包安装插件
func DefineServicePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.ServicePlugin(creator)
}
