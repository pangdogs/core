package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/types"
)

// DefinePluginInterface 定义通用插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefinePluginInterface[PLUGIN_IFACE any]() PluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.PluginInterface()
}

// PluginInterface 通用插件接口，在运行时上下文和服务上下文中，均可使用
type PluginInterface[PLUGIN_IFACE any] struct {
	Name  string                                             // 插件名称
	Using generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] // 使用插件
}

type _PluginInterface[PLUGIN_IFACE any] struct {
	name string
}

func (p _PluginInterface[PLUGIN_IFACE]) using() generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] {
	return func(pluginProvider plugin.PluginProvider) PLUGIN_IFACE {
		return plugin.Using[PLUGIN_IFACE](pluginProvider, p.name)
	}
}

// PluginInterface 生成通用插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterface[PLUGIN_IFACE] {
	return PluginInterface[PLUGIN_IFACE]{
		Name:  p.name,
		Using: p.using(),
	}
}
