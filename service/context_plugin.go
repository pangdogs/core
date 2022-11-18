package service

import (
	"fmt"
	"github.com/galaxy-kit/galaxy-go/plugin"
	"github.com/galaxy-kit/galaxy-go/util"
)

// GetPlugin 从服务上下文上获取插件
func GetPlugin[T any](ctx Context, pluginName string) T {
	if ctx == nil {
		panic("nil ctx")
	}

	pluginBundle := UnsafeContext(ctx).getOptions().PluginBundle
	if pluginBundle == nil {
		panic("nil pluginBundle")
	}

	plugin, ok := plugin.GetPlugin[T](pluginBundle, pluginName)
	if !ok {
		panic(fmt.Errorf("plugin '%s' not installed", pluginName))
	}

	return plugin
}

// TryGetPlugin 尝试从服务上下文上获取插件
func TryGetPlugin[T any](ctx Context, pluginName string) (T, bool) {
	if ctx == nil {
		return util.Zero[T](), false
	}

	pluginBundle := UnsafeContext(ctx).getOptions().PluginBundle
	if pluginBundle == nil {
		return util.Zero[T](), false
	}

	plugin, ok := plugin.GetPlugin[T](pluginBundle, pluginName)
	if !ok {
		return util.Zero[T](), false
	}

	return plugin, true
}
