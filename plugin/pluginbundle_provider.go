package plugin

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// PluginProvider 插件提供者
type PluginProvider interface {
	// GetPluginBundle 获取插件包
	GetPluginBundle() PluginBundle
}

// Using 使用插件
func Using[T any](provider PluginProvider, name string) T {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}

	pluginInfo, ok := provider.GetPluginBundle().Get(name)
	if !ok {
		panic(fmt.Errorf("%w: plugin %q not installed", ErrPlugin, name))
	}

	if !pluginInfo.Active {
		panic(fmt.Errorf("%w: plugin %q not actived", ErrPlugin, name))
	}

	return iface.Cache2Iface[T](pluginInfo.Face.Cache)
}

// Install 安装插件
func Install[T any](provider PluginProvider, plugin T, name ...string) {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}
	provider.GetPluginBundle().Install(iface.MakeFaceAny(plugin), name...)
}

// Uninstall 卸载插件
func Uninstall(provider PluginProvider, name string) {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}
	provider.GetPluginBundle().Uninstall(name)
}
