package define

import (
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// ServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func ServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterfaceDefinition[PLUGIN_IFACE] {
	return _DefinePluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.ServicePluginInterface()
}

// ServicePluginInterfaceDefinition 服务插件接口定义，只能在服务上下文中使用
type ServicePluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[service.Context, PLUGIN_IFACE] // 使用插件
}

// ServicePluginInterface 生成服务插件接口定义
func (d _DefinePluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterfaceDefinition[PLUGIN_IFACE] {
	return ServicePluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  d.name,
		Using: func(ctx service.Context) PLUGIN_IFACE { return d.using()(ctx) },
	}
}
