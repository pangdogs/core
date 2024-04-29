package define

import (
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
)

// ServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func ServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterfaceDefinition[PLUGIN_IFACE] {
	plug := definePluginInterface[PLUGIN_IFACE]()

	return ServicePluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  plug.Name,
		Using: func(ctx service.Context) PLUGIN_IFACE { return plug.Using(ctx) },
	}
}

// ServicePluginInterfaceDefinition 服务插件接口定义，只能在服务上下文中使用
type ServicePluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[service.Context, PLUGIN_IFACE] // 使用插件
}
