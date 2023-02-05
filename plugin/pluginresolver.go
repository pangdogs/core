package plugin

import "kit.golaxy.org/golaxy/util"

// PluginResolver 用于从运行时上下文或服务上下文中获取插件
type PluginResolver interface {
	// GetPlugin 获取插件
	GetPlugin(pluginName string) (util.FaceAny, bool)
}
