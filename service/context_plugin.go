package service

import (
	"fmt"
	"github.com/pangdogs/galaxy/plugin"
)

// Plugin 从服务上下文上获取插件
func Plugin[T any](ctx Context, pluginName string) T {
	if ctx == nil {
		panic("nil ctx")
	}

	pluginLib := UnsafeContext(ctx).getOptions().PluginLib
	if pluginLib == nil {
		panic("nil pluginLib")
	}

	plugin, ok := plugin.GetPlugin[T](pluginLib, pluginName)
	if !ok {
		panic(fmt.Errorf("plugin '%s' not registered", pluginName))
	}

	return plugin
}
