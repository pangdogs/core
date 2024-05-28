package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/generic"
)

// RuntimePlugin 定义运行时插件
func RuntimePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) RuntimePluginDefinition[PLUGIN_IFACE, OPTION] {
	plug := definePlugin[PLUGIN_IFACE, OPTION](creator)

	return RuntimePluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      plug.Name,
		Install:   plug.Install,
		Uninstall: plug.Uninstall,
		Using:     func(ctx runtime.Context) PLUGIN_IFACE { return plug.Using(ctx) },
	}
}

// RuntimePluginDefinition 运行时插件定义，只能在运行时上下文中安装与使用
type RuntimePluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[runtime.Context, PLUGIN_IFACE]      // 使用插件
}
