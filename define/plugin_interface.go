package define

import (
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
)

type _PluginInterface[PLUGIN_IFACE any] struct {
	_name string
}

func (p _PluginInterface[PLUGIN_IFACE]) name() string {
	return p._name
}

func (p _PluginInterface[PLUGIN_IFACE]) get() func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
	return func(pluginResolver plugin.PluginResolver) PLUGIN_IFACE {
		return plugin.GetPlugin[PLUGIN_IFACE](pluginResolver, p.name())
	}
}

func (p _PluginInterface[PLUGIN_IFACE]) tryGet() func(pluginResolver plugin.PluginResolver) (PLUGIN_IFACE, bool) {
	return func(pluginResolver plugin.PluginResolver) (PLUGIN_IFACE, bool) {
		return plugin.TryGetPlugin[PLUGIN_IFACE](pluginResolver, p.name())
	}
}

// ServicePluginInterface 服务插件接口
type ServicePluginInterface[PLUGIN_IFACE any] struct {
	Name   string                                     // 插件名称
	Get    func(service.Context) PLUGIN_IFACE         // 从服务上下文获取插件
	TryGet func(service.Context) (PLUGIN_IFACE, bool) // 尝试从服务上下文获取插件
}

// ServicePluginInterface 生成服务插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterface[PLUGIN_IFACE] {
	return ServicePluginInterface[PLUGIN_IFACE]{
		Name:   p.name(),
		Get:    func(ctx service.Context) PLUGIN_IFACE { return p.get()(ctx) },
		TryGet: func(ctx service.Context) (PLUGIN_IFACE, bool) { return p.tryGet()(ctx) },
	}
}

// RuntimePluginInterface 运行时插件接口
type RuntimePluginInterface[PLUGIN_IFACE any] struct {
	Name   string                                     // 插件名称
	Get    func(runtime.Context) PLUGIN_IFACE         // 从运行时上下文获取插件
	TryGet func(runtime.Context) (PLUGIN_IFACE, bool) // 尝试从运行时上下文获取插件
}

// RuntimePluginInterface 生成运行时插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN_IFACE] {
	return RuntimePluginInterface[PLUGIN_IFACE]{
		Name:   p.name(),
		Get:    func(ctx runtime.Context) PLUGIN_IFACE { return p.get()(ctx) },
		TryGet: func(ctx runtime.Context) (PLUGIN_IFACE, bool) { return p.tryGet()(ctx) },
	}
}

// PluginInterface 插件接口
type PluginInterface[PLUGIN_IFACE any] struct {
	Name   string                                           // 插件名称
	Get    func(plugin.PluginResolver) PLUGIN_IFACE         // 获取插件
	TryGet func(plugin.PluginResolver) (PLUGIN_IFACE, bool) // 尝试获取插件
}

// PluginInterface 生成插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterface[PLUGIN_IFACE] {
	return PluginInterface[PLUGIN_IFACE]{
		Name:   p.name(),
		Get:    p.get(),
		TryGet: p.tryGet(),
	}
}

// DefinePluginInterface 定义插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefinePluginInterface[PLUGIN_IFACE any]() PluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.PluginInterface()
}

// DefineRuntimePluginInterface 定义运行时插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineRuntimePluginInterface[PLUGIN_IFACE any]() RuntimePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.RuntimePluginInterface()
}

// DefineServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefineServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}.ServicePluginInterface()
}
