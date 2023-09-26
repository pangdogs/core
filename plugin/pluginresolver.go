package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/types"
)

// PluginResolver 插件解析器
type PluginResolver interface {
	// ResolvePlugin 解析插件
	ResolvePlugin(name string) (PluginInfo, bool)
}

// Fetch 获取插件。
//
//	@param pluginResolver 插件解析器。
//	@param name 插件名称。
func Fetch[T any](pluginResolver PluginResolver, name string) T {
	plugin, err := Access[T](pluginResolver, name)
	if err != nil {
		panic(err)
	}
	return plugin
}

// Access 访问插件
//
//	@param pluginResolver 插件解析器。
//	@param name 插件名称。
func Access[T any](pluginResolver PluginResolver, name string) (T, error) {
	if pluginResolver == nil {
		return types.Zero[T](), fmt.Errorf("%w: pluginResolver is nil", ErrPlugin)
	}

	pluginInfo, ok := pluginResolver.ResolvePlugin(name)
	if !ok {
		return types.Zero[T](), fmt.Errorf("%w: %q not installed", ErrPlugin, name)
	}

	if !pluginInfo.Active {
		return types.Zero[T](), fmt.Errorf("%w: %q not actived", ErrPlugin, name)
	}

	return iface.Cache2Iface[T](pluginInfo.Face.Cache), nil
}
