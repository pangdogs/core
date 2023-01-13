package runtime

import (
	"github.com/golaxy-kit/golaxy/util"
)

// GetPlugin 获取插件
func (ctx *ContextBehavior) GetPlugin(pluginName string) (util.FaceAny, bool) {
	pluginBundle := UnsafeContext(ctx).getOptions().PluginBundle
	if pluginBundle == nil {
		return util.Zero[util.FaceAny](), false
	}

	pluginFace, ok := pluginBundle.Get(pluginName)
	if !ok {
		return util.Zero[util.FaceAny](), false
	}

	return pluginFace, true
}
