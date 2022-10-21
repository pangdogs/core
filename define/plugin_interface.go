package define

import (
	"github.com/pangdogs/galaxy/ec"
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
		return service.Plugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeGet 生成从运行时上下文中获取插件函数
func (p _PluginInterface[PLUGIN]) RuntimeGet() func(runtime.Context) PLUGIN {
	return func(ctx runtime.Context) PLUGIN {
		return runtime.Plugin[PLUGIN](ctx, p.Name())
	}
}

// ServicePluginInterface 服务类插件接口
type ServicePluginInterface[PLUGIN any] struct {
	Name  string
	Get   func(service.Context) PLUGIN
	ECGet func(ec.ContextHolder) PLUGIN
}

// ServicePluginInterface 生成服务类插件接口定义
func (p _PluginInterface[PLUGIN]) ServicePluginInterface() ServicePluginInterface[PLUGIN] {
	get := p.ServiceGet()
	return ServicePluginInterface[PLUGIN]{
		Name:  p.Name(),
		Get:   get,
		ECGet: func(ctxHolder ec.ContextHolder) PLUGIN { return get(service.Get(ctxHolder)) },
	}
}

// RuntimePluginInterface 运行时类插件接口
type RuntimePluginInterface[PLUGIN any] struct {
	Name  string
	Get   func(runtime.Context) PLUGIN
	ECGet func(ec.ContextHolder) PLUGIN
}

// RuntimePluginInterface 生成运行时类插件接口定义
func (p _PluginInterface[PLUGIN]) RuntimePluginInterface() RuntimePluginInterface[PLUGIN] {
	get := p.RuntimeGet()
	return RuntimePluginInterface[PLUGIN]{
		Name:  p.Name(),
		Get:   get,
		ECGet: func(ctxHolder ec.ContextHolder) PLUGIN { return get(runtime.Get(ctxHolder)) },
	}
}

// PluginInterface 插件接口
type PluginInterface[PLUGIN any] struct {
	Name       string
	RuntimeGet func(runtime.Context) PLUGIN
	ServiceGet func(service.Context) PLUGIN
}

// PluginInterface 生成插件接口定义
func (p _PluginInterface[PLUGIN]) PluginInterface() PluginInterface[PLUGIN] {
	return PluginInterface[PLUGIN]{
		Name:       p.Name(),
		RuntimeGet: p.RuntimeGet(),
		ServiceGet: p.ServiceGet(),
	}
}

// DefinePluginInterface 定义插件接口，因为仅有接口没有实现，所以不能用于向插件库注册
func DefinePluginInterface[PLUGIN any]() _PluginInterface[PLUGIN] {
	return _PluginInterface[PLUGIN]{
		name: util.TypeFullName[PLUGIN](),
	}
}
