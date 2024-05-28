package define

import (
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// PluginInterface 定义通用插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func PluginInterface[PLUGIN_IFACE any]() PluginInterfaceDefinition[PLUGIN_IFACE] {
	return definePluginInterface[PLUGIN_IFACE]()
}

// PluginInterfaceDefinition 通用插件接口定义，在运行时上下文和服务上下文中，均可使用
type PluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                             // 插件名称
	Using generic.Func1[plugin.PluginProvider, PLUGIN_IFACE] // 使用插件
}

func definePluginInterface[PLUGIN_IFACE any]() PluginInterfaceDefinition[PLUGIN_IFACE] {
	name := types.FullNameT[PLUGIN_IFACE]()

	return PluginInterfaceDefinition[PLUGIN_IFACE]{
		Name: name,
		Using: func(provider plugin.PluginProvider) PLUGIN_IFACE {
			return plugin.Using[PLUGIN_IFACE](provider, name)
		},
	}
}
