package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// ServicePlugin 定义服务插件
func ServicePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePluginDefinition[PLUGIN_IFACE, OPTION] {
	return _DefinePlugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.ServicePlugin(creator)
}

// ServicePluginDefinition 服务插件定义，只能在服务上下文中安装与使用
type ServicePluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[service.Context, PLUGIN_IFACE]      // 使用插件
}

// ServicePlugin 生成服务插件定义
func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) ServicePlugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePluginDefinition[PLUGIN_IFACE, OPTION] {
	return ServicePluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      d.name,
		Install:   d.install(creator),
		Uninstall: d.uninstall(),
		Using:     func(ctx service.Context) PLUGIN_IFACE { return d.using()(ctx) },
	}
}
