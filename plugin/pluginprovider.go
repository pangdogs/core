package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/iface"
)

// PluginProvider 插件提供者
type PluginProvider interface {
	// GetPluginBundle 获取插件包
	GetPluginBundle() PluginBundle
}

// Using 使用插件
func Using[T any](pluginProvider PluginProvider, name string) T {
	if pluginProvider == nil {
		panic(fmt.Errorf("%w: %w: pluginProvider is nil", ErrPlugin, exception.ErrArgs))
	}

	pluginInfo, ok := pluginProvider.GetPluginBundle().Get(name)
	if !ok {
		panic(fmt.Errorf("%w: plugin %q not installed", ErrPlugin, name))
	}

	if !pluginInfo.Active {
		panic(fmt.Errorf("%w: plugin %q not actived", ErrPlugin, name))
	}

	return iface.Cache2Iface[T](pluginInfo.Face.Cache)
}

// Install 安装插件
func Install[T any](pluginProvider PluginProvider, plugin T, name ...string) {
	if pluginProvider == nil {
		panic(fmt.Errorf("%w: %w: pluginProvider is nil", ErrPlugin, exception.ErrArgs))
	}
	pluginProvider.GetPluginBundle().Install(iface.MakeFaceAny(plugin), name...)
}

// Uninstall 卸载插件
func Uninstall(pluginProvider PluginProvider, name string) {
	if pluginProvider == nil {
		panic(fmt.Errorf("%w: %w: pluginProvider is nil", ErrPlugin, exception.ErrArgs))
	}
	pluginProvider.GetPluginBundle().Uninstall(name)
}
