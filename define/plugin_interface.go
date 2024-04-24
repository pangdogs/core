package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// PluginInterface 定义通用插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func PluginInterface[PLUGIN_IFACE any]() PluginInterfaceDefinition[PLUGIN_IFACE] {
	return _DefinePluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.PluginInterface()
}

// PluginInterfaceDefinition 通用插件接口定义，在运行时上下文和服务上下文中，均可使用
type PluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                             // 插件名称
	Using generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] // 使用插件
}

type _DefinePluginInterface[PLUGIN_IFACE any] struct {
	name string
}

func (d _DefinePluginInterface[PLUGIN_IFACE]) using() generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] {
	return func(provider plugin.PluginProvider) PLUGIN_IFACE {
		return plugin.Using[PLUGIN_IFACE](provider, d.name)
	}
}

// PluginInterface 生成通用插件接口定义
func (d _DefinePluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterfaceDefinition[PLUGIN_IFACE] {
	return PluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  d.name,
		Using: d.using(),
	}
}
