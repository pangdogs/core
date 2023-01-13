package define

import (
	"github.com/golaxy-kit/golaxy/runtime"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
)

type _PluginInterface[PLUGIN_IFACE any] struct {
	_name string
}

func (p _PluginInterface[PLUGIN_IFACE]) name() string {
	return p._name
}

func (p _PluginInterface[PLUGIN_IFACE]) serviceContext() func(service.Context) PLUGIN_IFACE {
	return func(ctx service.Context) PLUGIN_IFACE {
		return service.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

func (p _PluginInterface[PLUGIN_IFACE]) runtimeContext() func(runtime.Context) PLUGIN_IFACE {
	return func(ctx runtime.Context) PLUGIN_IFACE {
		return runtime.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

// ServicePluginInterface 服务类插件接口
type ServicePluginInterface[PLUGIN_IFACE any] struct {
	Name    string                             // 插件名称
	Context func(service.Context) PLUGIN_IFACE // 从服务上下文获取插件
}

// ServicePluginInterface 生成服务类插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterface[PLUGIN_IFACE] {
	return ServicePluginInterface[PLUGIN_IFACE]{
		Name:    p.name(),
		Context: p.serviceContext(),
	}
}

// RuntimePluginInterface 运行时类插件接口
type RuntimePluginInterface[PLUGIN_IFACE any] struct {
	Name    string                             // 插件名称
	Context func(runtime.Context) PLUGIN_IFACE // 从运行时上下文获取插件
}

// RuntimePluginInterface 生成运行时类插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN_IFACE] {
	return RuntimePluginInterface[PLUGIN_IFACE]{
		Name:    p.name(),
		Context: p.runtimeContext(),
	}
}

// PluginInterface 插件接口
type PluginInterface[PLUGIN_IFACE any] struct {
	Name           string                             // 插件名称
	RuntimeContext func(runtime.Context) PLUGIN_IFACE // 从运行时上下文获取插件
	ServiceContext func(service.Context) PLUGIN_IFACE // 从服务上下文获取插件
}

// PluginInterface 生成插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterface[PLUGIN_IFACE] {
	return PluginInterface[PLUGIN_IFACE]{
		Name:           p.name(),
		RuntimeContext: p.runtimeContext(),
		ServiceContext: p.serviceContext(),
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
