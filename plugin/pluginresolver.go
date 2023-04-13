package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/util"
)

// PluginResolver 插件解析器
type PluginResolver interface {
	// ResolvePlugin 解析插件
	ResolvePlugin(pluginName string) (util.FaceAny, bool)
}

// GetPlugin 获取插件。
//
//	@param pluginResolver 插件解析器。
//	@param pluginName 插件名称。
func GetPlugin[T any](pluginResolver PluginResolver, pluginName string) T {
	if pluginResolver == nil {
		panic("nil pluginResolver")
	}

	pluginFace, ok := pluginResolver.ResolvePlugin(pluginName)
	if !ok {
		panic(fmt.Errorf("plugin '%s' not installed", pluginName))
	}

	return util.Cache2Iface[T](pluginFace.Cache)
}

// TryGetPlugin 尝试获取插件
//
//	@param pluginResolver 插件解析器。
//	@param pluginName 插件名称。
func TryGetPlugin[T any](pluginResolver PluginResolver, pluginName string) (T, bool) {
	if pluginResolver == nil {
		return util.Zero[T](), false
	}

	pluginFace, ok := pluginResolver.ResolvePlugin(pluginName)
	if !ok {
		return util.Zero[T](), false
	}

	return util.Cache2Iface[T](pluginFace.Cache), true
}
