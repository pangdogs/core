package plugin

import (
	"fmt"
	"github.com/golaxy-kit/golaxy/util"
)

// InstallPlugin 安装插件。
//
//	@param pluginBundle 插件包。
//	@param pluginName 插件名称。
//	@param plugin 插件。
func InstallPlugin[T any](pluginBundle PluginBundle, pluginName string, plugin T) {
	if pluginBundle == nil {
		panic("nil pluginBundle")
	}
	pluginBundle.Install(pluginName, util.NewFacePair[any](plugin, plugin))
}

// UninstallPlugin 卸载插件。
//
//	@param pluginBundle 插件包。
func UninstallPlugin(pluginBundle PluginBundle, pluginName string) {
	if pluginBundle == nil {
		panic("nil pluginBundle")
	}
	pluginBundle.Uninstall(pluginName)
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
		return nil, false
	}

	pluginFace, ok := pluginResolver.GetPlugin(pluginName)
	if !ok {
		return nil, false
	}

	return util.Cache2Iface[T](pluginFace.Cache), true
}
