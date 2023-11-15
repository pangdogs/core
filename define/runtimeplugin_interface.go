package define

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/types"
)

// DefineRuntimePluginInterface 定义运行时插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineRuntimePluginInterface[PLUGIN_IFACE any]() RuntimePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePluginInterface()
}

// RuntimePluginInterface 运行时插件接口，只能在运行时上下文中使用
type RuntimePluginInterface[PLUGIN_IFACE any] struct {
	Name  string                             // 插件名称
	Using func(runtime.Context) PLUGIN_IFACE // 使用插件
}

// RuntimePluginInterface 生成运行时插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN_IFACE] {
	return RuntimePluginInterface[PLUGIN_IFACE]{
		Name:  p.name,
		Using: func(ctx runtime.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}
