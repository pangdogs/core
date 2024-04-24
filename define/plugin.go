package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// DefinePlugin 定义通用插件
func DefinePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) Plugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.Plugin(creator)
}

// Plugin 通用插件，在运行时上下文和服务上下文中，均可安装与使用
type Plugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                             // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION]  // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]             // 从插件包卸载
	Using     generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] // 使用插件
}

type _Plugin[PLUGIN_IFACE, OPTION any] struct {
	name string
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) install(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) generic.ActionVar1[plugin.PluginProvider, OPTION] {
	return func(provider plugin.PluginProvider, options ...OPTION) {
		plugin.Install[PLUGIN_IFACE](provider, creator(options...), p.name)
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) uninstall() generic.Action1[plugin.PluginProvider] {
	return func(provider plugin.PluginProvider) {
		plugin.Uninstall(provider, p.name)
	}
}

func (p _Plugin[PLUGIN_IFACE, OPTION]) using() generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] {
	return func(provider plugin.PluginProvider) PLUGIN_IFACE {
		return plugin.Using[PLUGIN_IFACE](provider, p.name)
	}
}

// Plugin 生成通用插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) Plugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) Plugin[PLUGIN_IFACE, OPTION] {
	return Plugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name,
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     p.using(),
	}
}
