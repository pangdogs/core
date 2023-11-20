package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/types"
)

// DefineRuntimePlugin 定义运行时插件
func DefineRuntimePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return _Plugin[PLUGIN_IFACE, OPTION]{
		name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePlugin(creator)
}

// RuntimePlugin 运行时插件，只能在运行时上下文中安装与使用
type RuntimePlugin[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                            // 插件名称
	Install   generic.ActionVar1[plugin.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[plugin.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[runtime.Context, PLUGIN_IFACE]      // 使用插件
}

// RuntimePlugin 生成运行时插件定义
func (p _Plugin[PLUGIN_IFACE, OPTION]) RuntimePlugin(creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) RuntimePlugin[PLUGIN_IFACE, OPTION] {
	return RuntimePlugin[PLUGIN_IFACE, OPTION]{
		Name:      p.name,
		Install:   p.install(creator),
		Uninstall: p.uninstall(),
		Using:     func(ctx runtime.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}
