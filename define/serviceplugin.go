package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/generic"
)

// ServicePlugin 定义服务插件
func ServicePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePluginDefinition[PLUGIN_IFACE, OPTION] {
	plug := definePlugin[PLUGIN_IFACE, OPTION](creator)

	return ServicePluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      plug.Name,
		Install:   plug.Install,
		Uninstall: plug.Uninstall,
		Using:     func(ctx service.Context) PLUGIN_IFACE { return plug.Using(ctx) },
	}
}

// ServicePluginDefinition 服务插件定义，只能在服务上下文中安装与使用
type ServicePluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[service.Context, PLUGIN_IFACE]      // 使用插件
}
