package define

import (
	"github.com/galaxy-kit/galaxy-go/runtime"
	"github.com/galaxy-kit/galaxy-go/service"
	"github.com/galaxy-kit/galaxy-go/util"
)

type _PluginInterface[PLUGIN_IFACE any] struct {
	_name string
}

func (p _PluginInterface[PLUGIN_IFACE]) name() string {
	return p._name
}

func (p _PluginInterface[PLUGIN_IFACE]) serviceGet() func(service.Context) PLUGIN_IFACE {
	return func(ctx service.Context) PLUGIN_IFACE {
		return service.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

func (p _PluginInterface[PLUGIN_IFACE]) serviceTryGet() func(service.Context) (PLUGIN_IFACE, bool) {
	return func(ctx service.Context) (PLUGIN_IFACE, bool) {
		return service.TryGetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

func (p _PluginInterface[PLUGIN_IFACE]) runtimeGet() func(runtime.Context) PLUGIN_IFACE {
	return func(ctx runtime.Context) PLUGIN_IFACE {
		return runtime.GetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

func (p _PluginInterface[PLUGIN_IFACE]) runtimeTryGet() func(runtime.Context) (PLUGIN_IFACE, bool) {
	return func(ctx runtime.Context) (PLUGIN_IFACE, bool) {
		return runtime.TryGetPlugin[PLUGIN_IFACE](ctx, p.name())
	}
}

// ServicePluginInterface 服务类插件接口
type ServicePluginInterface[PLUGIN_IFACE any] struct {
	Name   string                                     // 插件名称
	Get    func(service.Context) PLUGIN_IFACE         // 从服务上下文获取
	TryGet func(service.Context) (PLUGIN_IFACE, bool) // 从服务上下文尝试获取
}

// ServicePluginInterface 生成服务类插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) ServicePluginInterface() ServicePluginInterface[PLUGIN_IFACE] {
	return ServicePluginInterface[PLUGIN_IFACE]{
		Name:   p.name(),
		Get:    p.serviceGet(),
		TryGet: p.serviceTryGet(),
	}
}

// RuntimePluginInterface 运行时类插件接口
type RuntimePluginInterface[PLUGIN_IFACE any] struct {
	Name   string                                     // 插件名称
	Get    func(runtime.Context) PLUGIN_IFACE         // 从运行时上下文获取
	TryGet func(runtime.Context) (PLUGIN_IFACE, bool) // 从运行时上下文尝试获取
}

// RuntimePluginInterface 生成运行时类插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN_IFACE] {
	return RuntimePluginInterface[PLUGIN_IFACE]{
		Name:   p.name(),
		Get:    p.runtimeGet(),
		TryGet: p.runtimeTryGet(),
	}
}

// PluginInterface 插件接口
type PluginInterface[PLUGIN_IFACE any] struct {
	Name          string                                     // 插件名称
	RuntimeGet    func(runtime.Context) PLUGIN_IFACE         // 从运行时上下文获取
	RuntimeTryGet func(runtime.Context) (PLUGIN_IFACE, bool) // 从运行时上下文尝试获取
	ServiceGet    func(service.Context) PLUGIN_IFACE         // 从服务上下文获取
	ServiceTryGet func(service.Context) (PLUGIN_IFACE, bool) // 从服务上下文尝试获取
}

// PluginInterface 生成插件接口定义
func (p _PluginInterface[PLUGIN_IFACE]) PluginInterface() PluginInterface[PLUGIN_IFACE] {
	return PluginInterface[PLUGIN_IFACE]{
		Name:          p.name(),
		RuntimeGet:    p.runtimeGet(),
		RuntimeTryGet: p.runtimeTryGet(),
		ServiceGet:    p.serviceGet(),
		ServiceTryGet: p.serviceTryGet(),
	}
}

// DefinePluginInterface 定义插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func DefinePluginInterface[PLUGIN_IFACE any]() _PluginInterface[PLUGIN_IFACE] {
	return _PluginInterface[PLUGIN_IFACE]{
		_name: util.TypeFullName[PLUGIN_IFACE](),
	}
}
