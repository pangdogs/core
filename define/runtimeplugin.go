package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// RuntimePlugin 定义运行时插件
func RuntimePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) RuntimePluginDefinition[PLUGIN_IFACE, OPTION] {
	return _DefinePlugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePlugin(creator)
}

// RuntimePluginDefinition 运行时插件定义，只能在运行时上下文中安装与使用
type RuntimePluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[runtime.Context, PLUGIN_IFACE]      // 使用插件
}

// RuntimePlugin 生成运行时插件定义
func (d _DefinePlugin[PLUGIN_IFACE, OPTION]) RuntimePlugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) RuntimePluginDefinition[PLUGIN_IFACE, OPTION] {
	return RuntimePluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      d.name,
		Install:   d.install(creator),
		Uninstall: d.uninstall(),
		Using:     func(ctx runtime.Context) PLUGIN_IFACE { return d.using()(ctx) },
	}
}
