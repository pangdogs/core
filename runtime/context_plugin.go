package runtime

import (
	"kit.golaxy.org/golaxy/plugin"
)

// ResolvePlugin 解析插件
func (ctx *ContextBehavior) ResolvePlugin(name string) (plugin.PluginInfo, bool) {
	pluginBundle := ctx.getOptions().PluginBundle
	if pluginBundle == nil {
		return plugin.PluginInfo{}, false
	}

	pluginInfo, ok := pluginBundle.Get(name)
	if !ok {
		return plugin.PluginInfo{}, false
	}

	return pluginInfo, true
}
