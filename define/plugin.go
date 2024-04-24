package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// Plugin 定义通用插件
func Plugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) PluginDefinition[PLUGIN_IFACE, OPTION] {
	return _DefinePlugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.Plugin(creator)
}

// PluginDefinition 通用插件定义，在运行时上下文和服务上下文中，均可安装与使用
type PluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                             // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION]  // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]             // 从插件包卸载
	Using     generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] // 使用插件
}

type _DefinePlugin[PLUGIN_IFACE, OPTION any] struct {
	name string
}

func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) install(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) generic.ActionVar1[plugin.PluginProvider, OPTION] {
	return func(provider plugin.PluginProvider, options ...OPTION) {
		plugin.Install[PLUGIN_IFACE](provider, creator(options...), d.name)
	}
}

func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) uninstall() generic.Action1[plugin.PluginProvider] {
	return func(provider plugin.PluginProvider) {
		plugin.Uninstall(provider, d.name)
	}
}

func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) using() generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] {
	return func(provider plugin.PluginProvider) PLUGIN_IFACE {
		return plugin.Using[PLUGIN_IFACE](provider, d.name)
	}
}

// Plugin 生成通用插件定义
func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) Plugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) PluginDefinition[PLUGIN_IFACE, OPTION] {
	return PluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      d.name,
		Install:   d.install(creator),
		Uninstall: d.uninstall(),
		Using:     d.using(),
	}
}
