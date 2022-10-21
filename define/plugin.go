package define

import (
	"github.com/pangdogs/galaxy/plugin"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

type _Plugin[PLUGIN, OPTION any] struct {
	name string
}

// Name 生成插件名称
func (p _Plugin[PLUGIN, OPTION]) Name() string {
	return p.name
}

// Register 生成插件注册函数
func (p _Plugin[PLUGIN, OPTION]) Register(creator func(...OPTION) PLUGIN) func(plugin.PluginLib, ...OPTION) {
	return func(lib plugin.PluginLib, options ...OPTION) {
		plugin.RegisterPlugin[PLUGIN](lib, p.Name(), creator(options...))
	}
}

// Deregister 生成插件取消注册函数
func (p _Plugin[PLUGIN, OPTION]) Deregister() func(plugin.PluginLib) {
	return func(lib plugin.PluginLib) {
		lib.Deregister(p.Name())
	}
}

// ServiceGet 生成从服务上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) ServiceGet() func(service.Context) PLUGIN {
	return func(ctx service.Context) PLUGIN {
		return service.Plugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeGet 生成从运行时上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) RuntimeGet() func(runtime.Context) PLUGIN {
	return func(ctx runtime.Context) PLUGIN {
		return runtime.Plugin[PLUGIN](ctx, p.Name())
	}
}

// ServicePlugin 服务类插件
type ServicePlugin[PLUGIN, OPTION any] struct {
	Name       string
	Register   func(plugin.PluginLib, ...OPTION)
	Deregister func(plugin.PluginLib)
	Get        func(service.Context) PLUGIN
}

// ServicePlugin 生成服务类插件定义
func (p _Plugin[PLUGIN, OPTION]) ServicePlugin(creator func(...OPTION) PLUGIN) ServicePlugin[PLUGIN, OPTION] {
	return ServicePlugin[PLUGIN, OPTION]{
		Name:       p.Name(),
		Register:   p.Register(creator),
		Deregister: p.Deregister(),
		Get:        p.ServiceGet(),
	}
}

// RuntimePlugin 运行时类插件
type RuntimePlugin[PLUGIN, OPTION any] struct {
	Name       string
	Register   func(plugin.PluginLib, ...OPTION)
	Deregister func(plugin.PluginLib)
	Get        func(runtime.Context) PLUGIN
}

// RuntimePlugin 生成运行时类插件定义
func (p _Plugin[PLUGIN, OPTION]) RuntimePlugin(creator func(...OPTION) PLUGIN) RuntimePlugin[PLUGIN, OPTION] {
	return RuntimePlugin[PLUGIN, OPTION]{
		Name:       p.Name(),
		Register:   p.Register(creator),
		Deregister: p.Deregister(),
		Get:        p.RuntimeGet(),
	}
}

// Plugin 插件
type Plugin[PLUGIN, OPTION any] struct {
	Name       string
	Register   func(plugin.PluginLib, ...OPTION)
	Deregister func(plugin.PluginLib)
	RuntimeGet func(runtime.Context) PLUGIN
	ServiceGet func(service.Context) PLUGIN
}

// Plugin 生成插件定义
func (p _Plugin[PLUGIN, OPTION]) Plugin(creator func(...OPTION) PLUGIN) Plugin[PLUGIN, OPTION] {
	return Plugin[PLUGIN, OPTION]{
		Name:       p.Name(),
		Register:   p.Register(creator),
		Deregister: p.Deregister(),
		RuntimeGet: p.RuntimeGet(),
		ServiceGet: p.ServiceGet(),
	}
}

// DefinePlugin 定义插件
func DefinePlugin[PLUGIN, OPTION any]() _Plugin[PLUGIN, OPTION] {
	return _Plugin[PLUGIN, OPTION]{
		name: util.TypeFullName[PLUGIN](),
	}
}
