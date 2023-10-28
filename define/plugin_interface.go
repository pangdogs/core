package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/types"
)

type _PluginInterface[PLUGIN_IFACE any] struct {
	_name string
}

func (p _PluginInterface[PLUGIN_IFACE]) name() string {
	return p._name
}

func (p _PluginInterface[PLUGIN_IFACE]) using() func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
	return func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
		plugin, err := plugin.Using[PLUGIN_IFACE](pluginResolver, p.name())
		if err != nil {
			panic(err)
		}
		return plugin
	}
}

// ServicePluginInterface 服务插件接口
type ServicePluginInterface[PLUGIN_IFACE any] struct {
	Name  string                             // 插件名称
	Using func(service.Context) PLUGIN_IFACE // 使用插件
}

// ServicePluginInterface 生成服务插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterface[PLUGIN_IFACE] {
	return ServicePluginInterface[PLUGIN_IFACE]{
		Name:  p.name(),
		Using: func(ctx service.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}

// RuntimePluginInterface 运行时插件接口
type RuntimePluginInterface[PLUGIN_IFACE any] struct {
	Name  string                             // 插件名称
	Using func(runtime.Context) PLUGIN_IFACE // 使用插件
}

// RuntimePluginInterface 生成运行时插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN_IFACE] {
	return RuntimePluginInterface[PLUGIN_IFACE]{
		Name:  p.name(),
		Using: func(ctx runtime.Context) PLUGIN_IFACE { return p.using()(ctx) },
	}
}

// PluginInterface 通用插件接口
type PluginInterface[PLUGIN_IFACE any] struct {
	Name  string                                   // 插件名称
	Using func(plugin.PluginResolver) PLUGIN_IFACE // 使用插件
}

// PluginInterface 生成通用插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterface[PLUGIN_IFACE] {
	return PluginInterface[PLUGIN_IFACE]{
		Name:  p.name(),
		Using: p.using(),
	}
}

// DefinePluginInterface 定义通用插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefinePluginInterface[PLUGIN_IFACE any]() PluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.PluginInterface()
}

// DefineRuntimePluginInterface 定义运行时插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineRuntimePluginInterface[PLUGIN_IFACE any]() RuntimePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.RuntimePluginInterface()
}

// DefineServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: types.FullName[PLUGIN_IFACE](),
	}.ServicePluginInterface()
}
