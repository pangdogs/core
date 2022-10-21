package registry

import "github.com/pangdogs/galaxy/define"

// Plugin 定义本插件接口
var Plugin = define.DefinePluginInterface[Registry]().ServicePluginInterface()
