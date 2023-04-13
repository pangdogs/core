package service

import (
	"kit.golaxy.org/golaxy/util"
)

// ResolvePlugin 解析插件
func (ctx *ContextBehavior) ResolvePlugin(pluginName string) (util.FaceAny, bool) {
	pluginBundle := ctx.getOptions().PluginBundle
	if pluginBundle == nil {
		return util.Zero[util.FaceAny](), false
	}

	pluginFace, ok := pluginBundle.Get(pluginName)
	if !ok {
		return util.Zero[util.FaceAny](), false
	}

	return pluginFace, true
}
