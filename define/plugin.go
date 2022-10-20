package define

import (
	"github.com/pangdogs/galaxy/plugin"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

type _Plugin[PT, OPT any] struct {
	name string
}

// Name 生成插件名称
func (p _Plugin[PT, OPT]) Name() string {
	return p.name
}

// Register 生成插件注册函数
func (p _Plugin[PT, OPT]) Register(creator func(...OPT) PT) func(plugin.PluginLib, ...OPT) {
	return func(lib plugin.PluginLib, options ...OPT) {
		plugin.RegisterPlugin[PT](lib, p.Name(), creator(options...))
	}
}

// Deregister 生成插件取消注册函数
func (p _Plugin[PT, OPT]) Deregister() func(plugin.PluginLib) {
	return func(lib plugin.PluginLib) {
		lib.Deregister(p.Name())
	}
}

// ServiceGet 生成从服务上下文中获取插件函数
func (p _Plugin[PT, OPT]) ServiceGet() func(service.Context) PT {
	return func(ctx service.Context) PT {
		return service.Plugin[PT](ctx, p.Name())
	}
}

// RuntimeGet 生成从运行时上下文中获取插件函数
func (p _Plugin[PT, OPT]) RuntimeGet() func(runtime.Context) PT {
	return func(ctx runtime.Context) PT {
		return runtime.Plugin[PT](ctx, p.Name())
	}
}

// Plugin 用于定义插件
func Plugin[PT, OPT any]() _Plugin[PT, OPT] {
	return _Plugin[PT, OPT]{
		name: util.TypeFullName[PT](),
	}
}
