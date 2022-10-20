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

// Plugin 用于定义插件
func Plugin[PLUGIN, OPTION any]() _Plugin[PLUGIN, OPTION] {
	return _Plugin[PLUGIN, OPTION]{
		name: util.TypeFullName[PLUGIN](),
	}
}
