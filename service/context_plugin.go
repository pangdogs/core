package service

import (
	"fmt"
	"github.com/pangdogs/galaxy/plugin"
	"github.com/pangdogs/galaxy/util"
)

// GetPlugin 从服务上下文上获取插件
func GetPlugin[T any](ctx Context, pluginName string) T {
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

// TryGetPlugin 尝试从服务上下文上获取插件
func TryGetPlugin[T any](ctx Context, pluginName string) (T, bool) {
	if ctx == nil {
		return util.Zero[T](), false
	}

	pluginLib := UnsafeContext(ctx).getOptions().PluginLib
	if pluginLib == nil {
		return util.Zero[T](), false
	}

	plugin, ok := plugin.GetPlugin[T](pluginLib, pluginName)
	if !ok {
		return util.Zero[T](), false
	}

	return plugin, true
}
