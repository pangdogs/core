package runtime

import "github.com/pangdogs/galaxy/util"

func Plugin[T any](ctx Context, pluginName string) T {
	if ctx == nil {
		panic("nil ctx")
	}

	pluginLib := UnsafeContext(ctx).getOptions().PluginLib
	if pluginLib == nil {
		panic("nil pluginLib")
	}

	return util.Cache2Iface[T](pluginLib.Get(pluginName))
}