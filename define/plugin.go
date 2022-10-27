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

// Install 生成插件安装函数
func (p _Plugin[PLUGIN, OPTION]) Install(creator func(...OPTION) PLUGIN) func(plugin.PluginLib, ...OPTION) {
	return func(lib plugin.PluginLib, options ...OPTION) {
		plugin.InstallPlugin[PLUGIN](lib, p.Name(), creator(options...))
	}
}

// Uninstall 生成插件卸载函数
func (p _Plugin[PLUGIN, OPTION]) Uninstall() func(plugin.PluginLib) {
	return func(lib plugin.PluginLib) {
		lib.Uninstall(p.Name())
	}
}

// ServiceGet 生成从服务上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) ServiceGet() func(service.Context) PLUGIN {
	return func(ctx service.Context) PLUGIN {
		return service.GetPlugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeGet 生成从运行时上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) RuntimeGet() func(runtime.Context) PLUGIN {
	return func(ctx runtime.Context) PLUGIN {
		return runtime.GetPlugin[PLUGIN](ctx, p.Name())
	}
}

// ServiceTryGet 生成尝试从服务上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) ServiceTryGet() func(service.Context) (PLUGIN, bool) {
	return func(ctx service.Context) (PLUGIN, bool) {
		return service.TryGetPlugin[PLUGIN](ctx, p.Name())
	}
}

// RuntimeTryGet 生成尝试从运行时上下文中获取插件函数
func (p _Plugin[PLUGIN, OPTION]) RuntimeTryGet() func(runtime.Context) (PLUGIN, bool) {
	return func(ctx runtime.Context) (PLUGIN, bool) {
		return runtime.TryGetPlugin[PLUGIN](ctx, p.Name())
	}
}

// ServicePlugin 服务类插件
type ServicePlugin[PLUGIN, OPTION any] struct {
	Name      string
	Install   func(plugin.PluginLib, ...OPTION)
	Uninstall func(plugin.PluginLib)
	Get       func(service.Context) PLUGIN
	TryGet    func(service.Context) (PLUGIN, bool)
}

// ServicePlugin 生成服务类插件定义
func (p _Plugin[PLUGIN, OPTION]) ServicePlugin(creator func(...OPTION) PLUGIN) ServicePlugin[PLUGIN, OPTION] {
	return ServicePlugin[PLUGIN, OPTION]{
		Name:      p.Name(),
		Install:   p.Install(creator),
		Uninstall: p.Uninstall(),
		Get:       p.ServiceGet(),
		TryGet:    p.ServiceTryGet(),
	}
}

// RuntimePlugin 运行时类插件
type RuntimePlugin[PLUGIN, OPTION any] struct {
	Name      string
	Install   func(plugin.PluginLib, ...OPTION)
	Uninstall func(plugin.PluginLib)
	Get       func(runtime.Context) PLUGIN
	TryGet    func(runtime.Context) (PLUGIN, bool)
}

// RuntimePlugin 生成运行时类插件定义
func (p _Plugin[PLUGIN, OPTION]) RuntimePlugin(creator func(...OPTION) PLUGIN) RuntimePlugin[PLUGIN, OPTION] {
	return RuntimePlugin[PLUGIN, OPTION]{
		Name:      p.Name(),
		Install:   p.Install(creator),
		Uninstall: p.Uninstall(),
		Get:       p.RuntimeGet(),
		TryGet:    p.RuntimeTryGet(),
	}
}

// Plugin 插件
type Plugin[PLUGIN, OPTION any] struct {
	Name          string
	Install       func(plugin.PluginLib, ...OPTION)
	Uninstall     func(plugin.PluginLib)
	ServiceGet    func(service.Context) PLUGIN
	ServiceTryGet func(service.Context) (PLUGIN, bool)
	RuntimeGet    func(runtime.Context) PLUGIN
	RuntimeTryGet func(runtime.Context) (PLUGIN, bool)
}

// Plugin 生成插件定义
func (p _Plugin[PLUGIN, OPTION]) Plugin(creator func(...OPTION) PLUGIN) Plugin[PLUGIN, OPTION] {
	return Plugin[PLUGIN, OPTION]{
		Name:          p.Name(),
		Install:       p.Install(creator),
		Uninstall:     p.Uninstall(),
		ServiceGet:    p.ServiceGet(),
		ServiceTryGet: p.ServiceTryGet(),
		RuntimeGet:    p.RuntimeGet(),
		RuntimeTryGet: p.RuntimeTryGet(),
	}
}

// DefinePlugin 定义插件，可以用于向插件库安装插件
func DefinePlugin[PLUGIN, OPTION any]() _Plugin[PLUGIN, OPTION] {
	return _Plugin[PLUGIN, OPTION]{
		name: util.TypeFullName[PLUGIN](),
	}
}
