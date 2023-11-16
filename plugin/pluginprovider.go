package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/types"
)

// PluginProvider 插件提供者
type PluginProvider interface {
	// GetPlugin 获取插件
	GetPlugin(name string) (PluginInfo, bool)
}

// Using 使用插件
//
//	@param pluginProvider 插件提供者。
//	@param name 插件名称。
func Using[T any](pluginProvider PluginProvider, name string) (T, error) {
	if pluginProvider == nil {
		return types.Zero[T](), fmt.Errorf("%w: %w: pluginProvider is nil", ErrPlugin, exception.ErrArgs)
	}

	pluginInfo, ok := pluginProvider.GetPlugin(name)
	if !ok {
		return types.Zero[T](), fmt.Errorf("%w: %q not installed", ErrPlugin, name)
	}

	if !pluginInfo.Active {
		return types.Zero[T](), fmt.Errorf("%w: %q not actived", ErrPlugin, name)
	}

	return iface.Cache2Iface[T](pluginInfo.Face.Cache), nil
}
