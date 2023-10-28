package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/types"
)

type _Plugin[PLUGIN_IFACE, OPTION any] struct {
	_name string
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) name() string {
	return p._name
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) install(creator func(...OPTION) PLUGIN_IFACE) func(plugin.PluginBundle, ...OPTION) {
	return func(pluginBundle plugin.PluginBundle, options ...OPTION) {
		plugin.Install[PLUGIN_IFACE](pluginBundle, p.name(), creator(options...))
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) uninstall() func(plugin.PluginBundle) {
	return func(pluginBundle plugin.PluginBundle) {
		plugin.Uninstall(pluginBundle, p.name())
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) using() func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
	return func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
		plugin, err := plugin.Using[PLUGIN_IFACE](pluginResolver, p.name())
		if err != nil {
			panic(err)
		}
		return plugin
	}
}

// ServicePlugin 服务插件
type ServicePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                               // 插件名称
	Install   func(plugin.PluginBundle, ...OPTION) // 向插件包安装
	Uninstall func(plugin.PluginBundle)            // 从插件包卸载
	Using     func(service.Context) PLUGIN_IFACE   // 使用插件
}

// ServicePlugin 生成服务插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) ServicePlugin(creator func(...OPTION) PLUGIN_IFACE) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return ServicePlugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name(),
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     func(ctx service.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}

// RuntimePlugin 运行时插件
type RuntimePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                               // 插件名称
	Install   func(plugin.PluginBundle, ...OPTION) // 向插件包安装
	Uninstall func(plugin.PluginBundle)            // 从插件包卸载
	Using     func(runtime.Context) PLUGIN_IFACE   // 使用插件
}

// RuntimePlugin 生成运行时插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) RuntimePlugin(creator func(...OPTION) PLUGIN_IFACE) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return RuntimePlugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name(),
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     func(ctx runtime.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}

// Plugin 插件
type Plugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                   // 插件名称
	Install   func(plugin.PluginBundle, ...OPTION)     // 向插件包安装
	Uninstall func(plugin.PluginBundle)                // 从插件包卸载
	Using     func(plugin.PluginResolver) PLUGIN_IFACE // 使用插件
}

// Plugin 生成插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) Plugin(creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return Plugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name(),
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     p.using(),
	}
}

// DefinePlugin 定义通用插件，可以用于向插件包安装插件
func DefinePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.Plugin(creator)
}

// DefineRuntimePlugin 定义运行时插件，可以用于向插件包安装插件
func DefineRuntimePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePlugin(creator)
}

// DefineServicePlugin 定义服务插件，可以用于向插件包安装插件
func DefineServicePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.ServicePlugin(creator)
}
