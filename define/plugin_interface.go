package define

import (
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

type _PluginInterface[PLUGIN any] struct {
	name string
}

// Name 生成插件名称
func (p _PluginInterface[PLUGIN]) Name() string {
	return p.name
}

// ServiceGet 生成从服务上下文中获取插件函数
func (p _PluginInterface[PLUGIN]) ServiceGet() func(service.Context) PLUGIN {
	return func(ctx service.Context) PLUGIN {
		return service.GetPlugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeGet 生成从运行时上下文中获取插件函数
func (p _PluginInterface[PLUGIN]) RuntimeGet() func(runtime.Context) PLUGIN {
	return func(ctx runtime.Context) PLUGIN {
		return runtime.GetPlugin[PLUGIN](ctx, p.Name())
	}
}

// ServiceTryGet 生成尝试从服务上下文中获取插件函数
func (p _PluginInterface[PLUGIN]) ServiceTryGet() func(service.Context) (PLUGIN, bool) {
	return func(ctx service.Context) (PLUGIN, bool) {
		return service.TryGetPlugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeTryGet 生成尝试从运行时上下文中获取插件函数
func (p _PluginInterface[PLUGIN]) RuntimeTryGet() func(runtime.Context) (PLUGIN, bool) {
	return func(ctx runtime.Context) (PLUGIN, bool) {
		return runtime.TryGetPlugin[PLUGIN](ctx, p.Name())
	}
}

// ServicePluginInterface 服务类插件接口
type ServicePluginInterface[PLUGIN any] struct {
	Name   string
	Get    func(service.Context) PLUGIN
	TryGet func(service.Context) (PLUGIN, bool)
}

// ServicePluginInterface 生成服务类插件接口定义
func (p _PluginInterface[PLUGIN]) ServicePluginInterface() ServicePluginInterface[PLUGIN] {
	return ServicePluginInterface[PLUGIN]{
		Name:   p.Name(),
		Get:    p.ServiceGet(),
		TryGet: p.ServiceTryGet(),
	}
}

// RuntimePluginInterface 运行时类插件接口
type RuntimePluginInterface[PLUGIN any] struct {
	Name   string
	Get    func(runtime.Context) PLUGIN
	TryGet func(runtime.Context) (PLUGIN, bool)
}

// RuntimePluginInterface 生成运行时类插件接口定义
func (p _PluginInterface[PLUGIN]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN] {
	return RuntimePluginInterface[PLUGIN]{
		Name:   p.Name(),
		Get:    p.RuntimeGet(),
		TryGet: p.RuntimeTryGet(),
	}
}

// PluginInterface 插件接口
type PluginInterface[PLUGIN any] struct {
	Name          string
	RuntimeGet    func(runtime.Context) PLUGIN
	RuntimeTryGet func(runtime.Context) (PLUGIN, bool)
	ServiceGet    func(service.Context) PLUGIN
	ServiceTryGet func(service.Context) (PLUGIN, bool)
}

// PluginInterface 生成插件接口定义
func (p _PluginInterface[PLUGIN]) PluginInterface() PluginInterface[PLUGIN] {
	return PluginInterface[PLUGIN]{
		Name:          p.Name(),
		RuntimeGet:    p.RuntimeGet(),
		RuntimeTryGet: p.RuntimeTryGet(),
		ServiceGet:    p.ServiceGet(),
		ServiceTryGet: p.ServiceTryGet(),
	}
}

// DefinePluginInterface 定义插件接口，因为仅有接口没有实现，所以不能用于向插件库注册
func DefinePluginInterface[PLUGIN any]() _PluginInterface[PLUGIN] {
	return _PluginInterface[PLUGIN]{
		name: util.TypeFullName[PLUGIN](),
	}
}
