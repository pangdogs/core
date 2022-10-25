package logger

import "github.com/pangdogs/galaxy/define"

// Plugin 定义本插件接口
var Plugin = define.DefinePluginInterface[Logger]().ServicePluginInterface()
