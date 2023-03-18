package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/util"
)

// PluginResolver 用于从运行时上下文或服务上下文中获取插件
type PluginResolver interface {
	// GetPlugin 获取插件
	GetPlugin(pluginName string) (util.FaceAny, bool)
}

// GetPlugin 获取插件。
//
//	@param pluginResolver 运行时上下文或服务上下文。
//	@param pluginName 插件名称。
func GetPlugin[T any](pluginResolver PluginResolver, pluginName string) T {
	if pluginResolver == nil {
		panic("nil pluginResolver")
	}

	pluginFace, ok := pluginResolver.GetPlugin(pluginName)
	if !ok {
		panic(fmt.Errorf("plugin '%s' not installed", pluginName))
	}

	return util.Cache2Iface[T](pluginFace.Cache)
}

// TryGetPlugin 尝试获取插件
//
//	@param pluginResolver 运行时上下文或服务上下文。
//	@param pluginName 插件名称。
func TryGetPlugin[T any](pluginResolver PluginResolver, pluginName string) (T, bool) {
	if pluginResolver == nil {
		return util.Zero[T](), false
	}

	pluginFace, ok := pluginResolver.GetPlugin(pluginName)
	if !ok {
		return util.Zero[T](), false
	}

	return util.Cache2Iface[T](pluginFace.Cache), true
}
