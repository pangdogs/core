package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/util/types"
)

// DefinePlugin 定义通用插件
func DefinePlugin[PLUGIN_IFACE, OPTION any](creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.Plugin(creator)
}

// Plugin 通用插件，在运行时上下文和服务上下文中，均可安装与使用
type Plugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                   // 插件名称
	Install   func(plugin.PluginBundle, ...OPTION)     // 向插件包安装
	Uninstall func(plugin.PluginBundle)                // 从插件包卸载
	Using     func(plugin.PluginResolver) PLUGIN_IFACE // 使用插件
}

type _Plugin[PLUGIN_IFACE, OPTION any] struct {
	name string
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) install(creator func(...OPTION) PLUGIN_IFACE) func(plugin.PluginBundle, ...OPTION) {
	return func(pluginBundle plugin.PluginBundle, options ...OPTION) {
		plugin.Install[PLUGIN_IFACE](pluginBundle, p.name, creator(options...))
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) uninstall() func(plugin.PluginBundle) {
	return func(pluginBundle plugin.PluginBundle) {
		plugin.Uninstall(pluginBundle, p.name)
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) using() func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
	return func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
		plugin, err := plugin.Using[PLUGIN_IFACE](pluginResolver, p.name)
		if err != nil {
			panic(err)
		}
		return plugin
	}
}

// Plugin 生成通用插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) Plugin(creator func(...OPTION) PLUGIN_IFACE) Plugin[PLUGIN_IFACE, OPTION] {
	return Plugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name,
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     p.using(),
	}
}
