package define

import (
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/generic"
)

// RuntimePluginInterface 定义运行时插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func RuntimePluginInterface[PLUGIN_IFACE any]() RuntimePluginInterfaceDefinition[PLUGIN_IFACE] {
	plug := definePluginInterface[PLUGIN_IFACE]()

	return RuntimePluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  plug.Name,
		Using: func(ctx runtime.Context) PLUGIN_IFACE { return plug.Using(ctx) },
	}
}

// RuntimePluginInterfaceDefinition 运行时插件接口定义，只能在运行时上下文中使用
type RuntimePluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[runtime.Context, PLUGIN_IFACE] // 使用插件
}
