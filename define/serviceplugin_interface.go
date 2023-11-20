package define

import (
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/types"
)

// DefineServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.ServicePluginInterface()
}

// ServicePluginInterface 服务插件接口，只能在服务上下文中使用
type ServicePluginInterface[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[service.Context, PLUGIN_IFACE] // 使用插件
}

// ServicePluginInterface 生成服务插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterface[PLUGIN_IFACE] {
	return ServicePluginInterface[PLUGIN_IFACE]{
		Name:  p.name,
		Using: func(ctx service.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}
