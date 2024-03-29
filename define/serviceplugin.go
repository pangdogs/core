package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// DefineServicePlugin 定义服务插件
func DefineServicePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.ServicePlugin(creator)
}

// ServicePlugin 服务插件，只能在服务上下文中安装与使用
type ServicePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[service.Context, PLUGIN_IFACE]      // 使用插件
}

// ServicePlugin 生成服务插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) ServicePlugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePlugin[PLUGIN_IFACE, OPTION] {
	return ServicePlugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name,
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     func(ctx service.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}
