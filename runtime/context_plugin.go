package runtime

import (
	"git.golaxy.org/core/plugin"
)

// GetPluginBundle 获取插件包
func (ctx *ContextBehavior) GetPluginBundle() plugin.PluginBundle {
	return ctx.opts.PluginBundle
}
