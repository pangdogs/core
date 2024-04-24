package define

import (
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// RuntimePluginInterface 定义运行时插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func RuntimePluginInterface[PLUGIN_IFACE any]() RuntimePluginInterfaceDefinition[PLUGIN_IFACE] {
	return _DefinePluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePluginInterface()
}

// RuntimePluginInterfaceDefinition 运行时插件接口定义，只能在运行时上下文中使用
type RuntimePluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[runtime.Context, PLUGIN_IFACE] // 使用插件
}

// RuntimePluginInterface 生成运行时插件接口定义
func (d _DefinePluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterfaceDefinition[PLUGIN_IFACE] {
	return RuntimePluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  d.name,
		Using: func(ctx runtime.Context) PLUGIN_IFACE { return d.using()(ctx) },
	}
}
